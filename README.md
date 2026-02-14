# jsonpretty

Pretty-print and validate JSON from stdin or files.

## Usage

```
jsonpretty [options] [file...]
```

## Options

- `-v, --validate` — Validate only, don't print
- `-c, --compact` — Compact/minimize output
- `-t, --tab` — Use tabs for indentation
- `-h, --help` — Show help

## Examples

```bash
echo '{"name":"Clara"}' | jsonpretty
jsonpretty data.json
jsonpretty -v config.json
cat api.json | jsonpretty -c
```

## Install

```
go install github.com/clarabennett2626/jsonpretty@latest
```

## Contributing

Pull requests welcome!
