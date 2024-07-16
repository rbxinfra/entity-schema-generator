package templates

import (
	"regexp"
	"strings"
)

const requiredWhitespaceMarker = "{{{REQUIRED_WHITESPACE}}}"

var whitespaceRegex = regexp.MustCompile(`^\s*$`)

// returns a constant that post processing will interpret as required whitespace whereas all other empty lines will be removed
func getRequiredWhitespaceMarker() string {
	return requiredWhitespaceMarker
}

// PostProcess removes empty lines and lines that only contain whitespace from the text.
func postProcess(text string) string {
	lines := strings.Split(text, "\n")
	newLines := []string{}
	for _, line := range lines {
		// If the line is completely empty or contains only spaces, skip it
		if len(line) == 0 || whitespaceRegex.MatchString(line) {
			continue
		}

		// If the line contains requiredWhitespaceMarker, replace it with a single empty line
		if strings.Contains(line, requiredWhitespaceMarker) {
			newLines = append(newLines, "")
			continue
		}

		// Add the line to the new lines
		newLines = append(newLines, line)
	}

	return strings.Join(newLines, "\n")
}
