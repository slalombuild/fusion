// Package gcpcmd contains all commands for the gcp section of the fusion cli
//
// Commands
//
// GCP commands can be executed with
// 	fusion new gcp
package gcpcmd

type Globals struct {
	Region string `enum:"us-central1" default:"us-central1"`
}

type GCP struct {
	LoadBalancer NewLoadBalancerCommand `cmd:"" aliases:"lb" help:"Create new GCP load balancer"`
}
