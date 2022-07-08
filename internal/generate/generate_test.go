package generate

import (
	"testing"

	"github.com/dave/jennifer/jen"
	assert "github.com/stretchr/testify/assert"
)

func Test_parseJenniferTypeFromString(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want *jen.Statement
	}{
		{name: "Parse jen.String from string", args: args{"string"}, want: jen.String()},
		{name: "Parse jen.Int from number", args: args{"number"}, want: jen.Int()},
		{name: "Parse jen.Int from int", args: args{"int"}, want: jen.Int()},
		{name: "Parse jen.Index.String() from list", args: args{"list"}, want: jen.Index().String()},
		{name: "Parse jen.Index.String() from array", args: args{"array"}, want: jen.Index().String()},
		{name: "Parse jen.Index.String() from slice", args: args{"slice"}, want: jen.Index().String()},
		{name: "Parse jen.Index.Int() from list:int", args: args{"list:int"}, want: jen.Index().Int()},
		{name: "Parse jen.Index.Int() from array:int", args: args{"array:int"}, want: jen.Index().Int()},
		{name: "Parse jen.Index.Int() from slice:int", args: args{"slice:int"}, want: jen.Index().Int()},
		{name: "Parse jen.Map(jen.String()).String() from object", args: args{"object"}, want: jen.Map(jen.String()).String()},
		{name: "Parse jen.Map(jen.String()).Int() from map:int", args: args{"map:int"}, want: jen.Map(jen.String()).Int()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseJenniferTypeFromString(tt.args.s))
		})
	}
}
