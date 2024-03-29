<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# azurecmd

```go
import "github.com/slalombuild/fusion/internal/commands/azurecmd"
```

### Package azurecmd contains all commands for the aws section of the fusion cli

Commands

AWS commands can be executed with fusion new azure

## Index

- [Variables](<#variables>)
- [type Azure](<#type-azure>)
- [type Globals](<#type-globals>)
- [type NewFunctionCommand](<#type-newfunctioncommand>)
  - [func (cmd *NewFunctionCommand) Run(ctx *commands.Context) error](<#func-newfunctioncommand-run>)
- [type NewVnetCommand](<#type-newvnetcommand>)
  - [func (cmd *NewVnetCommand) Run(ctx *commands.Context) error](<#func-newvnetcommand-run>)


## Variables

```go
var (
    ErrGenVnet   = "failed to generate terraform for vnet"
    ErrHighlight = "error highlighting terraform"
)
```

## type [Azure](<https://github.com/slalombuild/fusion/blob/main/internal/commands/azurecmd/cmd_azure.go#L13-L16>)

```go
type Azure struct {
    Func NewFunctionCommand `cmd:"" help:"Create new Azure function"`
    Vnet NewVnetCommand     `cmd:"" help:"Create new Azure vnet"`
}
```

## type [Globals](<https://github.com/slalombuild/fusion/blob/main/internal/commands/azurecmd/cmd_azure.go#L9-L11>)

```go
type Globals struct {
    Region string `enum:"us-west-1,us-west-2,us-east-1" default:"us-east-1"`
}
```

## type [NewFunctionCommand](<https://github.com/slalombuild/fusion/blob/main/internal/commands/azurecmd/cmd_azure_new_function.go#L12-L15>)

```go
type NewFunctionCommand struct {
    Globals
    Data *azure.AzureFunction `embed:""`
}
```

### func \(\*NewFunctionCommand\) [Run](<https://github.com/slalombuild/fusion/blob/main/internal/commands/azurecmd/cmd_azure_new_function.go#L17>)

```go
func (cmd *NewFunctionCommand) Run(ctx *commands.Context) error
```

## type [NewVnetCommand](<https://github.com/slalombuild/fusion/blob/main/internal/commands/azurecmd/cmd_azure_new_vnet.go#L17-L19>)

NewVnetCmd creates a new vnet

```go
type NewVnetCommand struct {
    *azure.Vnet `embed:""`
}
```

### func \(\*NewVnetCommand\) [Run](<https://github.com/slalombuild/fusion/blob/main/internal/commands/azurecmd/cmd_azure_new_vnet.go#L21>)

```go
func (cmd *NewVnetCommand) Run(ctx *commands.Context) error
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
