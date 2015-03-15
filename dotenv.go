package dotenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	ConfLine = regexp.MustCompile("^\\s*?(\\w[\\w\\d_]+?)\\s*?=\\s*?([^\\s]+?)\\s*?$")
)

const DefaultPath = ".env"

// Load environment variables from '.env' file if that file exists.
func Go() {
	// TODO: check home dir?
	if fi, err := os.Stat(DefaultPath); os.IsNotExist(err) {
		fmt.Println(fi)
		// DefaultPath not exists
	} else {
		fmt.Println(err)
		GoWithPath(DefaultPath)
	}
}

// Load environment variables from `filename`.
// TODO: parse quoted values.
func GoWithPath(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		matches := ConfLine.FindStringSubmatch(line)
		if len(matches) == 3 {
			key := matches[1]
			value := matches[2]
			err = os.Setenv(key, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
