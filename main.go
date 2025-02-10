package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

type Repo struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Commit string `json:"commit,omitempty"`
	Tag    string `json:"tag,omitempty"`
	Branch string `json:"branch,omitempty"`
}

type Config struct {
	GitLabURL string `json:"gitlab_url"`
	Token     string `json:"token"`
	RepoFile  string `json:"repo_file"`
}

var rootCmd = &cobra.Command{
	Use:   "gitlab-cloner",
	Short: "CLI tool to clone GitLab repositories",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Usage: gitlab-cloner <config.json>")
		}
		configFile := args[0]
		config, err := loadConfig(configFile)
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
		client, err := gitlab.NewClient(config.Token, gitlab.WithBaseURL(config.GitLabURL))
		if err != nil {
			log.Fatalf("Failed to create GitLab client: %v", err)
		}
		repos, err := loadRepoList(config.RepoFile)
		if err != nil {
			log.Fatalf("Failed to load repositories: %v", err)
		}
		for _, repo := range repos {
			if err := cloneOrUpdateRepo(repo, config, client); err != nil {
				log.Errorf("Failed to process repo %s: %v", repo.Name, err)
			}
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func loadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func loadRepoList(filename string) ([]Repo, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var repos []Repo
	if err := json.Unmarshal(data, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func cloneOrUpdateRepo(repo Repo, config *Config, client *gitlab.Client) error {
	dir := filepath.Join("./", repo.Name)
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		log.Infof("Repo %s already exists, skipping...", repo.Name)
		return nil
	}
	cloneURL := fmt.Sprintf("https://oauth2:%s@%s/%s.git", config.Token, strings.TrimPrefix(repo.URL, "https://"))
	log.Infof("Cloning %s...", repo.Name)
	cmd := exec.Command("git", "clone", cloneURL, dir)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone failed: %w", err)
	}
	return checkoutVersion(dir, repo)
}

func checkoutVersion(dir string, repo Repo) error {
	var ref string
	if repo.Commit != "" {
		ref = repo.Commit
	} else if repo.Tag != "" {
		ref = repo.Tag
	} else if repo.Branch != "" {
		ref = repo.Branch
	} else {
		return nil
	}
	log.Infof("Checking out %s in repo %s...", ref, repo.Name)
	cmd := exec.Command("git", "checkout", ref)
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git checkout failed: %w", err)
	}
	return nil
}
