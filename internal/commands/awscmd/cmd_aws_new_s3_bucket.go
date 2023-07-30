package awscmd

import (
	commands "github.com/slalombuild/fusion/internal/commands"
	aws "github.com/slalombuild/fusion/templates/aws"
)

// NewS3BucketCmd creates a new s3_bucket
type NewS3BucketCmd struct {
	Globals
	*aws.S3Bucket `embed:""`
}

func (cmd *NewS3BucketCmd) Run(ctx *commands.Context) error {
	return cmd.Render(ctx.Output, ctx.NoColor.Bool())
}
