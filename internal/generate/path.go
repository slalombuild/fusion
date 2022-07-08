package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/ettle/strcase"
)

var (
	ErrWriteFile  = "failed to write file %s"
	ErrReadFile   = "failed to read file"
	ErrCreateDirs = "failed to recursively create directories"
)

type Destination int

const (
	DESTINATION_COMMAND Destination = iota
	DESTINATION_TEMPLATE
	DESTINATION_TEMPLATE_DATA
)

func OutputPath(destination Destination, provider, resource string) string {

	var path string
	provider = strings.ToLower(provider)
	resource = strings.ToLower(resource)

	switch destination {
	case DESTINATION_COMMAND:
		folder := fmt.Sprintf("internal/commands/%scmd/", provider)
		file := strcase.ToSnake(fmt.Sprintf("cmd_%s_new_%s", provider, resource))
		path = filepath.Join(folder, file) + ".go"

	case DESTINATION_TEMPLATE:
		folder := fmt.Sprintf("templates/%s", provider)
		file := strcase.ToSnake(fmt.Sprintf("%s_%s", provider, resource))
		path = filepath.Join(folder, file) + ".tmpl"

	case DESTINATION_TEMPLATE_DATA:
		folder := fmt.Sprintf("templates/%s", provider)
		file := strcase.ToSnake(fmt.Sprintf("%s_%s", provider, resource))
		path = filepath.Join(folder, file) + ".go"
	}

	return path
}

// Save saves a file and directory path if
// it does not exist
func Save(path string, content []byte) error {
	err := createDirIfNotExist(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, content, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, ErrWriteFile, path)
	}

	return nil
}

// createDirIfNotExist creates a directory tree
// recursively if the filepath is not found
func createDirIfNotExist(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		path = filepath.Dir(path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, ErrCreateDirs)
		}
	}

	return nil
}

// CommandName properly formats the name of the go command
// from the provided resource
func CommandName(name string) string {
	return "New" + strcase.ToGoPascal(name) + "Cmd"
}
