# dirgrep

`dirgrep` is a CLI tool to perform grep operations directory-wise.

Written in Go, it leverages the language built-in concurrency feature to execute grep operations concurrently on all files within a directory (even recursively!).

## Usage

```text
dirgrep is a simple and intuitive CLI tool that can perform grep operations within a specific diectory (recursively or not). Powered by concurrent Go, with love.

Usage:
  dirgrep [flags]

Flags:
  -c, --context int        The context to add to the matches (number of charachters). Defaults to 0 if not used
  -d, --directory string   The directory to search for the pattern in. Defaults to the current working directory if not specified. (default ".")
  -h, --help               Show the help message and exit.
  -p, --pattern string     Pattern to search for within the given directory. Required.
  -r, --recursive          Whether or not to search for files to grep recursively. Defaults to false if not used
  -s, --skip strings       One or more sub-directories to skip. Can be used multiple times, can be used with comma-separated values. Defaults to an empty list.
```

### More docs to come!

