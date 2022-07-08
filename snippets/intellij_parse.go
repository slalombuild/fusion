package snippets

import (
	"fmt"
	"regexp"
	"strings"
)

// replaceBlockNames takes terraform resource and data
// blocks and templatizes their references so
// they can be easily renamed
func (i IntelliJ) replaceBlockNames(text string, varCount int, vars []string) (string, int, []string) {
	r := regexp.MustCompile(`(resource) \"(\w+)\" \"(\w+)\"`)
	matches := r.FindAllStringSubmatch(text, -1)
	matches = removeDuplicateStringSlices(matches)
	for _, v := range matches {
		// variablize resource definitions
		oldString := fmt.Sprintf("%s \"%s\" \"%s\"", v[1], v[2], v[3])
		newString := fmt.Sprintf("%s \"%s\" \"$%s_%v$\"", v[1], v[2], v[3], varCount)
		text = strings.ReplaceAll(text, oldString, newString)
		// variablize references to resources
		oldString = fmt.Sprintf("%s.%s", v[2], v[3])
		newString = fmt.Sprintf("%s.$%s_%v$", v[2], v[3], varCount)
		text = strings.ReplaceAll(text, oldString, newString)
		varCount += 1
		vars = append(vars, v[3])
	}
	return text, varCount, vars
}

// replaceDirectInsertions finds Go Template blocks
// and replaces them with VSCode snippet direct
// insertions.
func (i IntelliJ) replaceDirectInsertions(text string, varCount int, vars []string) (string, int, []string) {
	varGroup := 2
	r := regexp.MustCompile(`{{ ([^i][^f]\w*)? ?\.(\w+) }}`)
	matches := r.FindAllStringSubmatch(text, -1)
	matches = removeDuplicateStringSlices(matches)
	for _, v := range matches {
		varName := v[varGroup]
		if v[1] == "quote" {
			text = strings.ReplaceAll(text, v[0], fmt.Sprintf(`"$%s_%v$"`, varName, varCount))
		} else {
			text = strings.ReplaceAll(text, v[0], fmt.Sprintf("$%s_%v$", varName, varCount))
		}
		varCount += 1
		vars = append(vars, v[2])
	}
	return text, varCount, vars
}
