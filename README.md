# gua-cli (Github User Activity CLI)

A command-line application written in Go that retrieves and displays user activity data from GitHub.

## Table of Contents

- Installation
- Usage/Examples
- License 
## Installation

**Install Go**
- Ensure Go is installed on your system. You can download and install it from the official [Golang website](https://go.dev/dl/).
- After installation, ensure Go is added to your $PATH. Verify the installation by running:
```bash
$ go version
```

**Clone the repository**
- Clone the repository and enter the project directory
```bash
$ git clone https://github.com/mhaatha/gua-cli.git
$ cd gua-cli
```

**Download Dependencies**
```bash
$ go mod download
$ go mod tidy
```

**Setup Environment Variables**
- Create `.env` file in project root directory that provide your Github Personal Access Token (PAT).
- .env file contains:
```bash
PAT="<Your Personal Access Token>"
```

## Usage/Examples

```bash
$ go run main.go username <github-username>
```


## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/mhaatha/gua-cli/blob/main/LICENSE) file for details.
