# Report Generator

![CI](https://github.com/Qyroxen/Report-Generator/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Report-Generator/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Report-Generator?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Report-Generator)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Report-Generator)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Report-Generator?style=social)](https://github.com/Qyroxen/Report-Generator/stargazers)

## What is it?

Report Generator is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Report-Generator.git
cd Report-Generator
go build -o reportgenerator .

# Run
./reportgenerator --help
```

## CLI Usage

```bash
# Basic usage
./reportgenerator

# With flags
./reportgenerator --verbose --output json

# Get help
./reportgenerator --help
```

## Examples

```bash
# Example 1
./reportgenerator example1

# Example 2
./reportgenerator example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o reportgenerator .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Report-Generator/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Report-Generator?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Report-Generator/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Report-Generator?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Report-Generator/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Report-Generator" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Report-Generator/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Report-Generator" alt="Pull Requests">
  </a>
</p>
