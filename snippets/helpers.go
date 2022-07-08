package snippets

import (
	"path/filepath"
	"strings"
)

// removeDuplicateStringSlices seeks out duplicates from a slice of slices
// and removes them by their index.
func removeDuplicateStringSlices(matrix [][]string) [][]string {
	newMatrix := [][]string{}
	for _, v := range matrix {
		if !matrixContainsSlice(newMatrix, v) {
			newMatrix = append(newMatrix, v)
		}
	}
	return newMatrix
}

// matrixContainsSlices confirms that the provided matrix
// contains the provided nested slice.
func matrixContainsSlice(matrix [][]string, slice []string) bool {
	for _, v := range matrix {
		if stringSlicesEqual(slice, v) {
			return true
		}
	}
	return false
}

// stringSlicesEqual compares two slices of strings
// to confirm their equality.
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// nameFromPath returns a valid Snippet name
// from a provided .tmpl filepath
func nameFromPath(path string) (string, bool) {
	if filepath.Ext(path) == ".tmpl" {
		ext := filepath.Ext(path)
		base := filepath.Base(path)
		return strings.ReplaceAll(base, ext, ""), true
	}

	return "", false
}
