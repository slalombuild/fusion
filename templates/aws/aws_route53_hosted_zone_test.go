package aws

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoute53_Render(t *testing.T) {
	got := bytes.NewBuffer(nil)

	route53 := &Route53HostedZone{
		Zone:   "example.com",
		Target: "123.123.123.123",
	}

	err := route53.Render(got, true)
	assert.NoError(t, err)
}
