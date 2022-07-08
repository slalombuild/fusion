package snippets

import (
	"fmt"
	"regexp"
	"strings"
)

// replaceBlockNames takes terraform resource and data
// blocks and templatizes their references so
// they can be easily renamed
func (v VSCode) replaceBlockNames(text string, varCount int) (string, int) {
	r := regexp.MustCompile(`(resource) \"(\w+)\" \"(\w+)\"`)
	matches := r.FindAllStringSubmatch(text, -1)
	matches = removeDuplicateStringSlices(matches)
	for _, v := range matches {
		oldString := fmt.Sprintf("%s \"%s\" \"%s\"", v[1], v[2], v[3])
		newString := fmt.Sprintf("%s \"%s\" \"${%v:%s}\"", v[1], v[2], varCount, v[3])
		text = strings.ReplaceAll(text, oldString, newString)
		oldString = fmt.Sprintf("%s.%s", v[2], v[3])
		newString = fmt.Sprintf("%s.${%v:%s}", v[2], varCount, v[3])
		text = strings.ReplaceAll(text, oldString, newString)
		varCount += 1
	}
	return text, varCount
}

// replaceDirectInsertions finds Go Template blocks
// and replaces them with VSCode snippet direct
// insertions.
func (v VSCode) replaceDirectInsertions(text string, varCount int) (string, int) {
	varGroup := 2
	r := regexp.MustCompile(`{{ ([^i][^f]\w*)? ?\.(\w+) }}`)
	matches := r.FindAllStringSubmatch(text, -1)
	matches = removeDuplicateStringSlices(matches)
	for _, v := range matches {
		varName := v[varGroup]
		if v[1] == "quote" {
			text = strings.ReplaceAll(text, v[0], fmt.Sprintf(`"${%v:%s}"`, varCount, varName))
		} else {
			text = strings.ReplaceAll(text, v[0], fmt.Sprintf("${%v:%s}", varCount, varName))
		}
		varCount += 1
	}
	return text, varCount
}
