package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	var input io.Reader = os.Stdin
	validate := false
	compact := false
	indent := "  "

	args := os.Args[1:]
	var files []string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-v", "--validate":
			validate = true
		case "-c", "--compact":
			compact = true
		case "-t", "--tab":
			indent = "\t"
		case "-h", "--help":
			fmt.Println("Usage: jsonpretty [options] [file...]")
			fmt.Println()
			fmt.Println("Pretty-print or validate JSON from files or stdin.")
			fmt.Println()
			fmt.Println("Options:")
			fmt.Println("  -v, --validate   Validate only, don't print")
			fmt.Println("  -c, --compact    Compact output (minimize)")
			fmt.Println("  -t, --tab        Use tabs for indentation")
			fmt.Println("  -h, --help       Show this help")
			os.Exit(0)
		default:
			files = append(files, args[i])
		}
	}

	process := func(r io.Reader, name string) error {
		data, err := io.ReadAll(r)
		if err != nil {
			return fmt.Errorf("reading %s: %w", name, err)
		}

		var v interface{}
		if err := json.Unmarshal(data, &v); err != nil {
			return fmt.Errorf("%s: invalid JSON: %w", name, err)
		}

		if validate {
			fmt.Printf("%s: valid JSON\n", name)
			return nil
		}

		var out []byte
		if compact {
			out, err = json.Marshal(v)
		} else {
			out, err = json.MarshalIndent(v, "", indent)
		}
		if err != nil {
			return err
		}
		fmt.Println(string(out))
		return nil
	}

	if len(files) == 0 {
		if err := process(input, "stdin"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	exitCode := 0
	for _, f := range files {
		fh, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			exitCode = 1
			continue
		}
		if err := process(fh, f); err != nil {
			fmt.Fprintln(os.Stderr, err)
			exitCode = 1
		}
		fh.Close()
	}
	os.Exit(exitCode)
}
