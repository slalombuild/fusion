# 🧬 fusion

[![Go Reference](https://pkg.go.dev/badge/github.com/slalombuild/fusion.svg)](https://pkg.go.dev/github.com/slalombuild/fusion)
![Latest Release](https://img.shields.io/github/v/release/slalombuild/fusion?label=latest%20release)
[![Go Report Card](https://goreportcard.com/badge/github.com/slalombuild/fusion)](https://goreportcard.com/report/github.com/slalombuild/fusion)

Generate secure by default cloud infrastructure configuration with Go and Terraform. 

## Install 📥

Install the fusion cli

### Go

If you have Go setup on your system, you can install fusion with `go install`

```shell
go install github.com/slalombuild/fusion/cmd/fusion@latest
```

### Homebrew

```shell
brew tap slalombuild/fusion
brew install fusion

# Optionally install the fusionctl dev tool
brew install fusionctl
```

### Scoop

```shell
scoop bucket add fusion https://github.com/slalombuild/fusion.git
scoop install fusion
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

## Contributions

Is there a feature that you'd like to see implemented in Fusion? Feel free to open a [Feature Request](https://github.com/slalombuild/fusion/issues/new?assignees=&labels=enhancement&template=feature_request.yml&title=%28short+issue+description%29) issue to let us know what you'd like to see! 
We encourage submitting [Pull Requests](https://github.com/slalombuild/fusion/pulls) directly to add terraform resources to the library

For detailed development instructions, view our [DEVELOPMENT](.github/DEVELOPMENT.md) and [CONTRIBUTING](.github/CONTRIBUTING.md) guidelines.
