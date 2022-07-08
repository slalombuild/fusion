//go:build integration

package terraform_test

import (
	"bytes"
	"context"
	"log"
	"testing"
	"time"

	"github.com/SlalomBuild/fusion/templates/aws"
	"github.com/SlalomBuild/fusion/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraform_Graph(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	tf, err := terraform.New(ctx, t.TempDir())
	if err != nil {
		t.Error(err)
	}

	// Build a new lambda function
	lambda := aws.NewLambdaFunction("foo", "index.js", "handler", "GO1.X")
	err = tf.RenderTemplate(lambda)
	if err != nil {
		t.Error(err)
	}

	w := &bytes.Buffer{}
	err = tf.Graph(w)
	if assert.NoError(t, err) {
		assert.NotNil(t, w)
	}
}

func ExampleTerraform_Validate() {
	// Create a new terraform instance
	tf, err := terraform.New(context.Background(), ".")
	if err != nil {
		log.Fatal(err)
	}

	// Build a new lambda function
	lambda := aws.NewLambdaFunction("foo", "index.js", "handler", "GO1.X")

	// Render the template to the terraform temp dir
	err = tf.RenderTemplate(lambda)
	if err != nil {
		log.Fatal(err)
	}

	// Validate terraform in temp dir
	err = tf.Validate()
	if err != nil {
		log.Fatal(err)
	}
}
