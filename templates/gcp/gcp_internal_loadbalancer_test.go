package gcp

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalLoadBalancer(t *testing.T) {
	type args struct {
		forwardingrulename string
		region             string
		backendservicename string
		healthcheckname    string
		networkname        string
		subnetname         string
	}
	tests := []struct {
		name string
		args args
		want *InternalLoadBalancer
	}{
		{
			name: "Create new loadbalancer data",
			args: args{forwardingrulename: "internal_loadbalancer", region: "us-central1", backendservicename: "backend_service", healthcheckname: "health_check", networkname: "network", subnetname: "subnet"},
			want: &InternalLoadBalancer{
				ForwardingRuleName: "internal_loadbalancer",
				Region:             "us-central1",
				BackendServiceName: "backend_service",
				HealthCheckName:    "health_check",
				NetworkName:        "network",
				SubnetName:         "subnet",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewInternalLoadBalancer(tt.args.forwardingrulename, tt.args.region, tt.args.backendservicename, tt.args.healthcheckname, tt.args.networkname, tt.args.subnetname))
		})
	}
}
