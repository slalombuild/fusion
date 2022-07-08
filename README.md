# 🧬 fusion

[![Go Reference](https://pkg.go.dev/badge/github.com/SlalomBuild/fusion.svg)](https://pkg.go.dev/github.com/SlalomBuild/fusion)
![Latest Release](https://img.shields.io/github/v/release/SlalomBuild/fusion?label=latest%20release)
[![Go Report Card](https://goreportcard.com/badge/github.com/SlalomBuild/fusion)](https://goreportcard.com/report/github.com/SlalomBuild/fusion)

Generate secure by default cloud infrastructure configuration with Go and Terraform. 

## Install 📥

Install the fusion cli

### Go

If you have Go setup on your system, you can install fusion with `go install`

```shell
go install github.com/SlalomBuild/fusion/cmd/fusion@latest
```

### Docker Image

```
docker run --rm -it ghcr.io/slalombuild/fusion:latest --help
```

## Usage ⚡️

Getting started with fusion is as simple as naming the type of cloud resource you want and allow fusion to generate the terraform.

See available commands with `--help`

```
fusion —help

Usage: fusion <command>

Generate secure by default cloud infrastructure configuration

Flags:
  -h, --help        Show context-sensitive help.
  -v, --verbose     Enable verbose logging
  -n, --no-color    Disable colorful output ($NO_COLOR)

Commands:
  new    Create new cloud resources with Terraform

Run "fusion <command> --help" for more information on a command.
```

For more in-depth examples of creating cloud resources with fusion, view the [Example](./_example) folder.

## Snippets

Snippets are available in all supported IDEs with the pattern `fsn-<provider>_<resource>`

### VSCode 

```bash
# Install fusion vscode snippets into default snippets 
# directory
fusion gen snippets -e vscode -i
```

<details>
<summary>VSCode not installed in the default directory?</summary>
<br>
You will need to output a json file with `fusion gen snippets -e vscode -o filename.json` and place it and `package.json` from the repository's snippets directory within `.../.vscode/extensions/fusion-snippets`, creating directories if needed. Restart your IDE to make them available.
<br><br>
</details>

### Intellij

```bash
# 1. Generate snippets
fusion gen snippets -e intellij -o filename.xml

# 2. Find your IDE's configuration directory

# 3. Create a directory within that called `templates` if it does not already exist, and drop the xml file in there. Then, restart your IDE to make them available.
```

*Intellij users must check [this page](https://intellij-support.jetbrains.com/hc/en-us/articles/206544519-Directories-used-by-the-IDE-to-store-settings-caches-plugins-and-logs) to find the Configuration directory that pertains to your IDE version and operating system.*

## Development

For detailed development instructions, view our [DEVELOPMENT.md](./github/DEVELOPMENT.md) and [CONTRIBUTING.md](.github/CONTRIBUTING.md).