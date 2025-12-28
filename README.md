# dirgrep

`dirgrep` is a CLI tool to perform grep operations directory-wise.

Written in Go, it leverages the language built-in concurrency feature to execute grep operations concurrently on all files within a directory (even recursively!).

## Install

In order to install **dirgrep** there are three ways:

1. Using `go`: if you already have `go` 1.23+ installed in your environment, installing **dirgrep** is effortless

```bash
go install github.com/AstraBert/dirgrep
```

2. Using `npm`:

```bash
npm install @cle-does-things/dirgrep
```

3. Downloading the executable from the [releases page](https://github.com/AstraBert/dirgrep/releases): you can download it directly from the GitHub repository or, if you do not want to leave your terminal, you can use `curl`:

```bash
curl -L -o dirgrep https://github.com/AstraBert/dirgrep/releases/download/<version>/dirgrep_<version>_<OS>_<processor>.tar.gz ## e.g. https://github.com/AstraBert/dirgrep/releases/download/0.1.1/dirgrep_0.1.1_darwin_amd64.tar.gz

# make sure the downloaded binary is executable (not needed for Windows)
chmod +x dirgrep
```

In this last case, be careful to specify your OS (supported: linux, windows, macos) and your processor type (supported: amd, arm).

## Usage

```bash
dirgrep [command] [flags]
```

**Commands**

- `mcp`: Start an MCP server over stdio transport

**Flags**

- `-c`, `--context int`  
  Number of characters of context to include around matches. Defaults to `0`.

- `-d`, `--directory string`  
  Directory to search for the pattern. Defaults to the current working directory (`"."`).

- `-h`, `--help`  
  Show the help message and exit.

- `-p`, `--pattern string`  
  Pattern to search for within the given directory. **Required.**

- `-r`, `--recursive`  
  Search files recursively. Defaults to `false`.

- `-s`, `--skip strings`  
  One or more sub-directories to skip. Can be specified multiple times or as comma-separated values. Defaults to an empty list.

- `-x`, `--no-pretty`
  Deactivate pretty-printing for the matches to the console.

**Examples**

```bash
# search for 'package main' in the current directory, excluding .git and .gitub
dirgrep --pattern 'package main' --skip .git --skip .github --recursive
# search in a specific directory non-recursively
dirgrep --pattern 'root' --directory cmd/
# add a context of 200 charachters around the match
dirgrep --pattern '202\d' --context 200
# deactivate pretty-printing
dirgrep --pattern '202\d' --context 100 --no-pretty
# start MCP server
dirgrep mcp
```

## Benchmark

> _Python 3.9+ is required for running the benchmark_

You can run the [benchmark for `dirgrep`](./benchmark/) (VS other tools) using:

```bash
cd benchmark
bash run.sh
```

This will:

- Create 1 million files (approx. 4GB) under the `benchmark/files` directory, containing three random lines each.
- Start `./dirgrep` search for the pattern `A password forgot itself at dawn.` (one of the random lines), redirecting the standard output to `benchmark_dirgrep.txt`
- Start `grep` search recursively for the same pattern, redirecting the standard output to `benchmark_grep.txt`
- Start `ripgrep` search recursively for the same pattern, redirecting the standard output to `benchmark_ripgrep.txt`
- Once the program is finished, you will have the time output for it. 

In the latest run, `dirgrep` VS other tools performed as follows:

| Metric     | dirgrep     | grep        | ripgrep     | Description                          |
|------------|-------------|-------------|-------------|--------------------------------------|
| user       | 35.58s      | 1.53s       | 2.06s       | CPU time in user mode (program code) |
| system     | 454.60s     | 27.67s      | 38.44s      | CPU time in system mode (kernel ops) |
| cpu        | 1166%       | 17%         | 312%        | CPU utilization across all cores     |
| **total**  | **42.037s** | **166.34s** | **12.940s** | Actual elapsed wall-clock time       |

## Contributing

We welcome contributions! Please read our [Contributing Guide](./CONTRIBUTING.md) to get started.

## License

This project is licensed under the [MIT License](./LICENSE)

