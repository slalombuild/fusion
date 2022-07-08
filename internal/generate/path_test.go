package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Create lambda command name", args: args{"lambda"}, want: "NewLambdaCmd"},
		{name: "Create security group command name", args: args{"security_group"}, want: "NewSecurityGroupCmd"},
		{name: "Create eks cluster command name", args: args{"eks.cluster"}, want: "NewEksClusterCmd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommandName(tt.args.name); got != tt.want {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func Test_createDirIfNotExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Create temp folder if not present", args: args{t.TempDir()}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := createDirIfNotExist(tt.args.path)
			if assert.NoError(t, err) {
				assert.DirExists(t, tt.args.path)
			}
		})
	}
}

func TestOutputPath(t *testing.T) {
	type args struct {
		destination Destination
		provider    string
		resource    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Create new lambda command path",
			args: args{DESTINATION_COMMAND, "AWS", "lambda_function"},
			want: "internal/commands/awscmd/cmd_aws_new_lambda_function.go",
		},
		{
			name: "Create new lambda template path",
			args: args{DESTINATION_TEMPLATE, "AWS", "lambda_function"},
			want: "templates/aws/aws_lambda_function.tmpl",
		},
		{
			name: "Create new lambda template data path",
			args: args{DESTINATION_TEMPLATE_DATA, "AWS", "lambda_function"},
			want: "templates/aws/aws_lambda_function.go",
		},
		{
			name: "Create new lambda template data path and format string case",
			args: args{DESTINATION_TEMPLATE_DATA, "AWS", "lambda-function"},
			want: "templates/aws/aws_lambda_function.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OutputPath(tt.args.destination, tt.args.provider, tt.args.resource)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSave(t *testing.T) {
	type args struct {
		path    string
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Save go file if not exist", args: args{path: t.TempDir() + "test.go"}, wantErr: false},
		{name: "Save nested template file if not exist", args: args{path: t.TempDir() + "/test/" + "test.tmpl"}, wantErr: false},
		{name: "Save heavily nested template file", args: args{path: t.TempDir() + "templates/aws/lambda/aws_lambda_function.tmpl"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Save(tt.args.path, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
