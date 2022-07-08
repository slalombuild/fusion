package terraform

import (
	"context"
	"os/exec"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/pkg/errors"
)

// execPath checks for a local terraform installation
// and installs terraform if not present
func execPath(ctx context.Context) (string, error) {
	tfpath, err := exec.LookPath("terraform")
	if err == nil {
		return tfpath, nil
	}

	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	tfpath, err = installer.Install(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to install terraform")
	}

	return tfpath, nil
}
