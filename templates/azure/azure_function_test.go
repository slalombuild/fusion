package azure

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAzureFunction(t *testing.T) {
	tests := []struct {
		name string
		want *AzureFunction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewAzureFunction())
		})
	}
}
