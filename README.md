# jsonpretty 🎨

Pretty-print and validate JSON from stdin or files.

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/clarabennettdev/jsonpretty)](https://goreportcard.com/report/github.com/clarabennettdev/jsonpretty)

Pipe in messy JSON, get clean output.

## Install

```bash
go install github.com/clarabennettdev/jsonpretty@latest
```

## Usage

```bash
curl -s https://api.example.com/data | jsonpretty
jsonpretty config.json
jsonpretty -v response.json    # validate only
jsonpretty -c data.json        # compact
```

## Options

| Flag | Description |
|------|-------------|
| `-v, --validate` | Validate only, don't print |
| `-c, --compact` | Compact/minimize output |
| `-t, --tab` | Use tabs for indentation |

## License

MIT
