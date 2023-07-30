//go:build integration

package terraform_test

import (
	"context"
	"testing"
	"time"

	"github.com/slalombuild/fusion/templates"
	"github.com/slalombuild/fusion/templates/aws"
	"github.com/slalombuild/fusion/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraform_Validate(t *testing.T) {
	tests := []struct {
		name     string
		resource templates.Renderer
	}{
		{name: "Validate lambda", resource: aws.NewLambdaFunction("foo", "index.js", "handler", "GO1.X")},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	tf, err := terraform.New(ctx, t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if assert.NoError(t, tf.RenderTemplate(tt.resource)) {
				assert.NoError(t, tf.Validate())
			}
		})
	}
}
