package azure

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVnet(t *testing.T) {
	type args struct {
		resourcegroupname  string
		location           string
		virtualnetworkname string
	}
	tests := []struct {
		name string
		args args
		want *Vnet
	}{
		{
			name: "Create new vnet data",
			args: args{resourcegroupname: "resource_group_name", location: "centralus", virtualnetworkname: "vnet_name"},
			want: &Vnet{
				ResourceGroupName:  "resource_group_name",
				Location:           "centralus",
				VirtualNetworkName: "vnet_name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewVnet(tt.args.resourcegroupname, tt.args.location, tt.args.virtualnetworkname))
		})
	}
}
