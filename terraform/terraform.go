package terraform

import (
	"context"
	"io"
	"os"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/pkg/errors"
	"github.com/slalombuild/fusion/templates"
)

const (
	tempDirPattern = "*fusion_temp"

	ErrTerraformInit     = "failed to perform terraform init"
	ErrTerraformValidate = "failed to perform terraform validate"
	ErrTerraformGraph    = "failed to perform terraform graph"
)

type Terraform struct {
	ctx context.Context

	tempDir string
	exec    *tfexec.Terraform
}

// New creates a new instance of the Terraform library with
// error logging, initialization, and a terraform executable fallback
// built-in.
func New(ctx context.Context, workingDir string) (*Terraform, error) {
	tfpath, err := execPath(ctx)
	if err != nil {
		return nil, err
	}

	tempDir, err := os.MkdirTemp(workingDir, tempDirPattern)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create temp directory %s", tempDir)
	}

	tf, err := tfexec.NewTerraform(tempDir, tfpath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new terraform executor")
	}

	err = tf.Init(ctx)
	if err != nil {
		return nil, errors.Wrap(err, ErrTerraformInit)
	}

	return &Terraform{
		ctx:     ctx,
		tempDir: tempDir,
		exec:    tf,
	}, nil
}

// RenderTemplate renders the terraform template to the
// temp dir and cleans up after
func (t *Terraform) RenderTemplate(r templates.Renderer) error {
	// Render the template to a temp file
	f, err := os.CreateTemp(t.tempDir, "*.tf")
	if err != nil {
		return errors.Wrapf(err, "failed to create temp terraform file in %s", t.tempDir)
	}

	err = r.Render(f, true)
	if err != nil {
		return errors.Wrap(err, "failed to render template")
	}

	return nil
}

// Validate validates terraform in the working directory.
func (t *Terraform) Validate() error {
	_, err := t.exec.Validate(t.ctx)
	if err != nil {
		return errors.Wrap(err, ErrTerraformValidate)
	}
	return nil
}

// Graph initalizes and converts valid terraform into a GraphViz visualization.
func (t *Terraform) Graph(w io.Writer) error {
	err := t.exec.Init(t.ctx)
	if err != nil {
		return errors.Wrap(err, ErrTerraformInit)
	}

	graph, err := t.exec.Graph(t.ctx)
	if err != nil {
		return errors.Wrap(err, ErrTerraformGraph)
	}

	_, err = w.Write([]byte(graph))
	if err != nil {
		return errors.Wrap(err, "failed to write graph to output")
	}

	return nil
}

// Cleanup cleans the temporary directories created during
// terraform execution
func (t *Terraform) Cleanup() error {
	err := os.RemoveAll(t.tempDir)
	return errors.Wrapf(err, "failed to delete tempdir %s", t.tempDir)
}
