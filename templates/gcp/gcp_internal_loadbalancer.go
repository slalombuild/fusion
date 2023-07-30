package gcp

import (
	_ "embed"
	"io"

	"github.com/pkg/errors"
	"github.com/slalombuild/fusion/templates"
)

//go:embed gcp_internal_loadbalancer.tmpl
var TEMPLATE_GCP_INTERNAL_LOADBALANCER string

// InternalLoadBalancer is the template data object used to create
// a gcp internal loadbalancer
type InternalLoadBalancer struct {
	ForwardingRuleName string `help:"" default:"internal-loadbalancer"`
	Region             string `help:"" default:"us-central1"`
	BackendServiceName string `help:"" default:"backend_service"`
	HealthCheckName    string `help:"" default:"health_check"`
	NetworkName        string `help:"" default:"network"`
	SubnetName         string `help:"" default:"subnet"`
}

// NewInternalLoadBalancer creates a new internal load balancer
func NewInternalLoadBalancer(forwardingrulename, region, backendservicename, healthcheckname, networkname, subnetname string) *InternalLoadBalancer {
	return &InternalLoadBalancer{
		ForwardingRuleName: forwardingrulename,
		Region:             region,
		BackendServiceName: backendservicename,
		HealthCheckName:    healthcheckname,
		NetworkName:        networkname,
		SubnetName:         subnetname,
	}
}

// Render generates the Terraform code for the InternalLoadBalancer
func (resource *InternalLoadBalancer) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_GCP_INTERNAL_LOADBALANCER, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}

	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}

	return nil
}
