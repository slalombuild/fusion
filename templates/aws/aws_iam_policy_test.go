package aws

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIamPolicy_Render(t *testing.T) {
	got := bytes.NewBuffer(nil)
	policy := &IamPolicy{
		Name:        "example-policy",
		Description: "A policy description",
		Path:        "/",
		PolicyJSON:  struct{}{},
	}

	err := policy.Render(got, true)
	assert.NoError(t, err)
}

func ExampleIamPolicy_Render() {
	policy := &IamPolicy{
		Name:        "example-policy",
		Description: "A policy description",
		PolicyJSON: strings.TrimSpace(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Effect": "Allow",
					"Action": [
						"ec2:AttachVolume",
						"ec2:DetachVolume"
					],
					"Resource": "arn:aws:ec2:*:*:instance/*",
					"Condition": {
						"StringEquals": {"aws:ResourceTag/Department": "Development"}
					}
				},
				{
					"Effect": "Allow",
					"Action": [
						"ec2:AttachVolume",
						"ec2:DetachVolume"
					],
					"Resource": "arn:aws:ec2:*:*:volume/*",
					"Condition": {
						"StringEquals": {"aws:ResourceTag/VolumeUser": "${aws:username}"}
					}
				}
			]
		}`),
	}

	err := policy.Render(os.Stdout, true)
	if err != nil {
		fmt.Println("failed to render policy", err)
		os.Exit(1)
	}
}
