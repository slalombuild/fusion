package aws

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/slalombuild/fusion/templates"

	"github.com/stretchr/testify/assert"
)

func TestNewLambdaFunction(t *testing.T) {
	type args struct {
		name     string
		filename string
		handler  string
		runtime  string
	}
	tests := []struct {
		name string
		args args
		want *LambdaFunction
	}{
		{
			name: "Create new lambda data",
			args: args{name: "my_lambda", filename: "function.zip", handler: "index.handler", runtime: "nodejs14.x"},
			want: &LambdaFunction{
				Name:     "my_lambda",
				Filename: "function.zip",
				Handler:  "index.handler",
				Runtime:  "nodejs14.x",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewLambdaFunction(tt.args.name, tt.args.filename, tt.args.handler, tt.args.runtime))
		})
	}
}

func TestLambdaFunction_Render(t *testing.T) {
	got := bytes.NewBuffer(nil)
	want := templates.FormatHCL(`
resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "example_function"  {
  filename      = "index.js"
  function_name = "example-function"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "handler.index"
  source_code_hash = filebase64sha256("index.js")

  runtime = "node"

  environment {
    variables = {}
  }
}`)

	lambda := &LambdaFunction{
		Name:     "example-function",
		Runtime:  "node",
		Handler:  "handler.index",
		Filename: "index.js",
	}

	err := lambda.Render(got, true)
	if assert.NoError(t, err) {
		assert.Equal(t, got.String(), want)
	}
}

func ExampleLambdaFunction_Render() {
	lambda := &LambdaFunction{
		Name:     "example-function",
		Runtime:  "node",
		Handler:  "handler.index",
		Filename: "index.js",
	}

	err := lambda.Render(os.Stdout, true)
	if err != nil {
		fmt.Println("failed to render lambda", err)
		os.Exit(1)
	}
}
