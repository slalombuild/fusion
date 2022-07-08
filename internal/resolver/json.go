package resolver

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
)

// JSONFileMapper implements kong.MapperValue to decode a JSON file into
// a struct field.
//
//    var cli struct {
//      Profile Profile 	`type:"jsonfile"`
//		Policy	interface{} `type:"jsonfile"`
//    }
//
//    func main() {
//      kong.Parse(&cli, kong.NamedMapper("jsonfile", JSONFileMapper))
//    }
var JSONFileMapper = kong.MapperFunc(decodeJSONFile)

func decodeJSONFile(ctx *kong.DecodeContext, target reflect.Value) error {
	var fname string
	if err := ctx.Scan.PopValueInto("filename", &fname); err != nil {
		return errors.Wrap(err, "failed to pop json into kong context")
	}
	f, err := os.Open(fname)
	if err != nil {
		return errors.Wrap(err, "failed to open json file")
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(target.Addr().Interface())
	return errors.Wrap(err, "failed to decode json to target interface")
}

func JSON(r io.Reader) (kong.Resolver, error) {
	values := map[string]interface{}{}
	err := json.NewDecoder(r).Decode(&values)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode values from json")
	}

	return kong.ResolverFunc(func(context *kong.Context, parent *kong.Path, flag *kong.Flag) (interface{}, error) {
		name := flag.Name
		raw, ok := values[name]
		if ok {
			return raw, nil
		}
		raw = values
		for _, part := range strings.Split(name, ".") {
			if values, ok := raw.(map[string]interface{}); ok {
				raw, ok = values[part]
				if !ok {
					return nil, nil
				}
			} else {
				return nil, nil
			}
		}
		return raw, nil

	}), nil
}
