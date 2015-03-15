package dotenv

import (
	"fmt"
	"os"
	"testing"
)

func printEnv() {
	fmt.Println(os.Environ())
}

func TestGoPlainWithTrimming(t *testing.T) {
	err := GoWithPath("fixtures/plain.env")
	if err != nil {
		t.Error(err)
	}

	if os.Getenv("PLAIN") != "true" {
		t.Error("Wrong value PLAIN, should be 'true'")
	}

	if os.Getenv("OPTION_A") != "1" {
		t.Error("Wrong value OPTION_A, should be '1'")
	}

	if os.Getenv("OPTION_B") != "2" {
		t.Error("Wrong value OPTION_B, should be '2'")
	}

	if os.Getenv("OPTION_C") != "3" {
		t.Error("Wrong value OPTION_C, should be '3'")
	}

	if os.Getenv("OPTION_D") != "4" {
		t.Error("Wrong value OPTION_D, should be '4'")
	}

	if os.Getenv("OPTION_E") != "5" {
		t.Error("Wrong value OPTION_E, should be '5'")
	}
}
