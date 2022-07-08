// Package azurecmd contains all commands for the aws section of the fusion cli
//
// Commands
//
// AWS commands can be executed with
// 	fusion new azure
package azurecmd

type Globals struct {
	Region string `enum:"us-west-1,us-west-2,us-east-1" default:"us-east-1"`
}

type Azure struct {
	Func NewFunctionCommand `cmd:"" help:"Create new Azure function"`
	Vnet NewVnetCommand     `cmd:"" help:"Create new Azure vnet"`
}
