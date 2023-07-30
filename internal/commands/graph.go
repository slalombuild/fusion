package commands

import (
	"bytes"
	"context"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/pkg/errors"
	"github.com/slalombuild/fusion/templates"
	"github.com/slalombuild/fusion/terraform"
)

func Graph(ctx *Context, r templates.Renderer) error {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriter(os.Stderr), spinner.WithColor("fgCyan"), spinner.WithSuffix(" Generating graphviz for terraform..."))
	s.Start()

	tf, err := terraform.New(context.Background(), ".")
	if err != nil {
		return err
	}

	err = tf.RenderTemplate(r)
	if err != nil {
		return err
	}

	w := &bytes.Buffer{}
	err = tf.Graph(w)
	if err != nil {
		return err
	}
	s.Stop()

	_, err = ctx.Output.Write(w.Bytes())
	if err != nil {
		return errors.Wrap(err, "failed to write graphviz to output")
	}

	return tf.Cleanup()
}
