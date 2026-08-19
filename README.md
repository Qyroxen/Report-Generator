# Report Generator

Generate reports from data - CSV, JSON, and database sources.

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![CI](https://github.com/Qyroxen/report-generator/actions/workflows/ci.yml/badge.svg)](https://github.com/Qyroxen/report-generator/actions/workflows/ci.yml)

> Generate reports from data - CSV, JSON, and database sources.

## What is it?

Report Generator is a command-line tool built with Go that helps developers generate reports from data - csv, json, and database sources. It's designed to be fast, reliable, and easy to use.

## Why?

Every developer needs report generator — but existing tools are either too complex, too slow, or require cloud dependencies. We built Report Generator to be:
- **Fast** — Written in Go for maximum performance
- **Offline** — No cloud dependencies, your data stays on your machine
- **Simple** — Clean CLI interface with sensible defaults
- **Extensible** — Easy to customize and integrate into your workflow

## Features

- **Multiple data sources** — Multiple data sources
- **Template-based generation** — Template-based generation
- **Chart and graph support** — Chart and graph support
- **Export to PDF/HTML** — Export to PDF/HTML
- **Scheduled reports** — Scheduled reports
- **CLI interface** — CLI interface

## Quick Start

### Prerequisites

- Go 1.23 or later

### Install

```bash
# Install with go install
go install github.com/Qyroxen/report-generator@latest

# Or build from source
git clone https://github.com/Qyroxen/report-generator.git
cd report-generator
go build -o report-generator .
```

### Usage

```bash
# Basic usage
.report-generator --help

# Example
./report-generator create --data data.csv --template report.html
```

## Output

```
Report Generator v1.0.0

Scanning...

✓ Analysis complete
✓ Results ready

{
  "status": "success",
  "results": [...]
}
```

## Configuration

Create a `.config.yaml` file in your project root:

```yaml
# Configuration options
verbose: true
output: json
timeout: 30s
```

## CLI Flags

```
report generator [command]

Flags:
  --path string      Target path (default ".")
  --format string    Output format: json, text (default "text")
  --verbose          Enable verbose output
  --config string    Config file path
  --output string    Output file path
```

## Examples

### Basic Example

```bash
.report-generator --path ./src
```

### Advanced Example

```bash
.report-generator --path ./src --format json --output report.json --verbose
```

### CI/CD Integration

```yaml
# .github/workflows/ci.yml
- name: Run Report Generator
  run: |
    go install github.com/Qyroxen/report-generator@latest
    report-generator --path . --format json --output report.json
```

## Documentation

- [Getting Started](docs/getting-started.md)
- [Configuration](docs/configuration.md)
- [API Reference](docs/api-reference.md)
- [Examples](examples/)
- [Contributing](CONTRIBUTING.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Qyroxen** - [GitHub](https://github.com/Qyroxen)

---

**Found this useful?** Give it a ⭐ on GitHub!
