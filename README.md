# GitLab Repo Cloner

## Overview
`gitlab-cloner` is a command-line tool that automates the process of cloning GitLab repositories. It supports authentication using a GitLab token and allows cloning repositories based on commit, tag, or branch priority.

## Features
- Authenticate with GitLab using a personal access token.
- Clone repositories specified in a JSON configuration file.
- Checkout repositories based on priority: **commit > tag > branch**.
- Skip downloading repositories that already exist.
- Uses `cobra` for CLI command processing.
- Logs activity using `logrus` for better debugging and monitoring.

## Installation
### Prerequisites
- Go (Golang) installed (v1.18+ recommended)
- Git installed and available in the system's PATH

### Build from source
```sh
# Clone the repository
git clone https://your-repo-url/gitlab-cloner.git
cd gitlab-cloner

# Build the binary
go build -o gitlab-cloner main.go
```

## Usage
### Configuration File
Create a `config.json` file with the following structure:
```json
{
  "gitlab_url": "https://gitlab.com",
  "token": "your_gitlab_token",
  "repo_file": "repo.json"
}
```

Create a `repo.json` file specifying repositories to clone:
```json
[
  {
    "name": "my-repo",
    "url": "https://gitlab.com/group/my-repo",
    "commit": "abc123"
  },
  {
    "name": "another-repo",
    "url": "https://gitlab.com/group/another-repo",
    "tag": "v1.0"
  }
]
```

### Running the CLI
```sh
./gitlab-cloner config.json
```
This will:
1. Authenticate with GitLab.
2. Clone repositories specified in `repo.json`.
3. Checkout the specified **commit**, **tag**, or **branch** in order of priority.

## Logging
This tool uses `logrus` for logging, which provides structured logs for easier debugging.

## Contributing
Feel free to submit issues or pull requests if you have improvements or feature requests.

## License
This project is licensed under the MIT License.

# GitLab Repo Cloner

## Overview
`gitlab-cloner` is a command-line tool that automates the process of cloning GitLab repositories. It supports authentication using a GitLab token and allows cloning repositories based on commit, tag, or branch priority.

## Features
- Authenticate with GitLab using a personal access token.
- Clone repositories specified in a JSON configuration file.
- Checkout repositories based on priority: **commit > tag > branch**.
- Skip downloading repositories that already exist.
- Uses `cobra` for CLI command processing.
- Logs activity using `logrus` for better debugging and monitoring.

## Installation
### Prerequisites
- Go (Golang) installed (v1.18+ recommended)
- Git installed and available in the system's PATH

### Build from source
```sh
# Clone the repository
git clone https://your-repo-url/gitlab-cloner.git
cd gitlab-cloner

# Build the binary
go build -o gitlab-cloner main.go
```

## Usage
### Configuration File
Create a `config.json` file with the following structure:
```json
{
  "gitlab_url": "https://gitlab.com",
  "token": "your_gitlab_token",
  "repo_file": "repo.json"
}
```

Create a `repo.json` file specifying repositories to clone:
```json
[
  {
    "name": "my-repo",
    "url": "https://gitlab.com/group/my-repo",
    "commit": "abc123"
  },
  {
    "name": "another-repo",
    "url": "https://gitlab.com/group/another-repo",
    "tag": "v1.0"
  }
]
```

### Running the CLI
```sh
./gitlab-cloner config.json
```
This will:
1. Authenticate with GitLab.
2. Clone repositories specified in `repo.json`.
3. Checkout the specified **commit**, **tag**, or **branch** in order of priority.

## Logging
This tool uses `logrus` for logging, which provides structured logs for easier debugging.

## Contributing
Feel free to submit issues or pull requests if you have improvements or feature requests.

## License
This project is licensed under the MIT License.

