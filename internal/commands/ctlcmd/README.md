<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# ctlcmd

```go
import "github.com/slalombuild/fusion/internal/commands/ctlcmd"
```

Package ctlcmd contains all commands for the fusionctl cli

## Index

- [Variables](<#variables>)
- [type Answers](<#type-answers>)
- [type Ctl](<#type-ctl>)
- [type NewResourceCmd](<#type-newresourcecmd>)
  - [func (cmd *NewResourceCmd) Run(ctx *commands.Context) error](<#func-newresourcecmd-run>)
- [type NewStackCmd](<#type-newstackcmd>)
  - [func (cmd *NewStackCmd) Run(ctx *commands.Context) error](<#func-newstackcmd-run>)


## Variables

```go
var (
    ErrPrompt = "failed to ask survey prompt"
)
```

## type [Answers](<https://github.com/slalombuild/fusion/blob/main/internal/commands/ctlcmd/cmd_ctl_new_resource.go#L27-L31>)

```go
type Answers struct {
    Resource string `arg:"" help:"Terraform resource name (e.g. loadbalancer, security_group)"`
    Provider string `help:"Cloud provider name (e.g. aws, gcp, azure)" enum:"aws,gcp,azure," default:"" short:"p"`
    Import   string `help:"Terraform file to import" short:"i" type:"existingfile"`
}
```

## type [Ctl](<https://github.com/slalombuild/fusion/blob/main/internal/commands/ctlcmd/cmd_ctl.go#L8-L11>)

```go
type Ctl struct {
    Resource NewResourceCmd `cmd:"" help:"Create new fusion resources from existing terraform"`
    Stack    NewStackCmd    `cmd:"" help:"Create new fusion stacks from existing terraform" hidden:"true"`
}
```

## type [NewResourceCmd](<https://github.com/slalombuild/fusion/blob/main/internal/commands/ctlcmd/cmd_ctl_new_resource.go#L20-L25>)

```go
type NewResourceCmd struct {
    Questions []*survey.Question `kong:"-"`
    Answers   Answers            `embed:""`
    Fields    map[string]string  `help:"Map of fields and their types to be used in template" short:"f" default:"name=string;description=string"`
    Save      bool               `help:"Save output" default:"false"`
}
```

### func \(\*NewResourceCmd\) [Run](<https://github.com/slalombuild/fusion/blob/main/internal/commands/ctlcmd/cmd_ctl_new_resource.go#L33>)

```go
func (cmd *NewResourceCmd) Run(ctx *commands.Context) error
```

## type [NewStackCmd](<https://github.com/slalombuild/fusion/blob/main/internal/commands/ctlcmd/cmd_ctl_new_stack.go#L8-L10>)

```go
type NewStackCmd struct {
    Questions []*survey.Question `kong:"-"`
}
```

### func \(\*NewStackCmd\) [Run](<https://github.com/slalombuild/fusion/blob/main/internal/commands/ctlcmd/cmd_ctl_new_stack.go#L12>)

```go
func (cmd *NewStackCmd) Run(ctx *commands.Context) error
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
