package parser

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	confLine = regexp.MustCompile("^\\s*?(\\w[\\w\\d_]+?)\\s*?=\\s*?([^\\s]+?)\\s*?$")
)

const (
	commentStartsWith = "#"
)

// ParseEnvVars parses string 'data' to map[string]string.
func ParseEnvVars(data string) map[string]string {
	lines := strings.Split(data, "\n")
	parsed := make(map[string]string)

	for _, line := range lines {
		if strings.Index(line, commentStartsWith) == 0 {
			continue
		}
		matches := confLine.FindStringSubmatch(line)
		if len(matches) == 3 {
			key := matches[1]
			value := maybeParseQuotes(matches[2])
			parsed[key] = value
		}
	}

	return parsed
}

func maybeParseQuotes(str string) string {
	if unquoted, err := strconv.Unquote(str); err == nil {
		return unquoted
	}

	return str
}
