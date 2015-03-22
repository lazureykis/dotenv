package dotenv

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	confLine = regexp.MustCompile("^\\s*?(\\w[\\w\\d_]+?)\\s*?=\\s*?([^\\s]+?)\\s*?$")
)

const (
	defaultPath       = ".env"
	commentStartsWith = "#"
)

// Go loads environment variables from '.env' file if that file exists.
func Go() {
	// TODO: check home dir?
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		// defaultPath not exists
	} else {
		GoWithPath(defaultPath)
	}
}

// GoWithPath loads environment variables from `filename`.
func GoWithPath(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.Index(line, commentStartsWith) == 0 {
			continue
		}
		matches := confLine.FindStringSubmatch(line)
		if len(matches) == 3 {
			key := matches[1]
			value := maybeParseQuotes(matches[2])
			err = os.Setenv(key, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func maybeParseQuotes(str string) string {
	if unquoted, err := strconv.Unquote(str); err == nil {
		return unquoted
	}

	return str
}
