package dotenv

import (
	"os"
	"testing"
)

func assertEnvValue(t *testing.T, name, want string) {
	got := os.Getenv(name)
	if got != want {
		t.Errorf("Want <%v> but got <%v>\n", want, got)
	}
	os.Unsetenv(name)
}

func TestGoQuoted(t *testing.T) {
	err := GoWithPath("fixtures/quoted.env")
	if err != nil {
		t.Error(err)
	}

	assertEnvValue(t, "QUOTED", "true")
	assertEnvValue(t, "OPTION_A", "1")
	assertEnvValue(t, "OPTION_B", "2")
	assertEnvValue(t, "OPTION_C", "")
	assertEnvValue(t, "OPTION_D", "\n")
	assertEnvValue(t, "OPTION_E", "1")
	assertEnvValue(t, "OPTION_F", "2")
	assertEnvValue(t, "OPTION_G", "")
	assertEnvValue(t, "OPTION_H", "\n")
}

func TestGoComments(t *testing.T) {
	err := GoWithPath("fixtures/comments.env")
	if err != nil {
		t.Error(err)
	}

	assertEnvValue(t, "COMMENT", "")
	assertEnvValue(t, "#COMMENT", "")
	assertEnvValue(t, "PLAIN", "true")
	assertEnvValue(t, "ARG", "1")
}
