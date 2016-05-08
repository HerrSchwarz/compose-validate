package param

import (
	"testing"
	"github.com/herrschwarz/compose-validate-local/param"
)

func TestInit(t *testing.T) {
	var params = param.Init()
	if *params.ConfigFile != "docker-compose.yml" {
		t.Fail()
	}

	if *params.RuleFile != "validate.yml" {
		t.Fail()
	}

	if *params.Verbose {
		t.Fail()
	}
}
