package parser

import (
	"io/ioutil"
	"testing"
)

func parseFixture(t *testing.T, filename string) map[string]string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	return ParseEnvVars(string(data))
}

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("Want <%v> but got <%v>\n", want, got)
	}
}

func TestParsePlainWithTrimming(t *testing.T) {
	vars := parseFixture(t, "../fixtures/plain.env")

	assertEqual(t, vars["PLAIN"], "true")
	assertEqual(t, vars["OPTION_A"], "1")
	assertEqual(t, vars["OPTION_B"], "2")
	assertEqual(t, vars["OPTION_C"], "3")
	assertEqual(t, vars["OPTION_D"], "4")
	assertEqual(t, vars["OPTION_E"], "5")
}

func TestGoQuoted(t *testing.T) {
	vars := parseFixture(t, "../fixtures/quoted.env")

	assertEqual(t, vars["QUOTED"], "true")
	assertEqual(t, vars["OPTION_A"], "1")
	assertEqual(t, vars["OPTION_B"], "2")
	assertEqual(t, vars["OPTION_C"], "")
	assertEqual(t, vars["OPTION_D"], "\n")
	assertEqual(t, vars["OPTION_E"], "1")
	assertEqual(t, vars["OPTION_F"], "2")
	assertEqual(t, vars["OPTION_G"], "")
	assertEqual(t, vars["OPTION_H"], "\n")
}

func TestGoComments(t *testing.T) {
	vars := parseFixture(t, "../fixtures/comments.env")

	assertEqual(t, vars["COMMENT"], "")
	assertEqual(t, vars["#COMMENT"], "")
	assertEqual(t, vars["PLAIN"], "true")
	assertEqual(t, vars["ARG"], "1")
}
