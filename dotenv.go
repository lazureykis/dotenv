package dotenv

import (
	"github.com/lazureykis/dotenv/parser"
	"io/ioutil"
	"os"
)

const defaultPath = ".env"

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
	vars, err := parseFile(filename)
	if err != nil {
		return err
	}

	return setEnvVars(vars)
}

func parseFile(filename string) (map[string]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return parser.ParseEnvVars(string(data)), nil
}

func setEnvVars(vars map[string]string) error {
	for key, value := range vars {
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}
