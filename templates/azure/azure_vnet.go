package azure

import (
	_ "embed"
	"io"

	"github.com/SlalomBuild/fusion/templates"
	"github.com/pkg/errors"
)

//go:embed azure_vnet.tmpl
var TEMPLATE_AZURE_VNET string

type Vnet struct {
	ResourceGroupName  string `default:"resource_group_name" help:""`
	Location           string `default:"centralus" help:""`
	VirtualNetworkName string `default:"vnet_name" help:""`
}

func NewVnet(resourcegroupname, location, virtualnetworkname string) *Vnet {
	return &Vnet{
		ResourceGroupName:  resourcegroupname,
		Location:           location,
		VirtualNetworkName: virtualnetworkname,
	}
}

func (resource *Vnet) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AZURE_VNET, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}

	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}

	return nil
}
