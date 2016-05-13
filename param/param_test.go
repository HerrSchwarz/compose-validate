package param

import (
	"testing"
	"os"

	"github.com/herrschwarz/compose-validate/param"
)

var p param.Params

func TestMain(m *testing.M) {
	p = param.Init()
	os.Exit(m.Run())
}

func TestConfig(t *testing.T) {
	const configFileDefault = "docker-compose.yml"
	if *p.ConfigFile != configFileDefault {
		t.Error("default value for configFile is ", *p.ConfigFile,
		"Expected", configFileDefault)
	}
}

func TestRule(t *testing.T) {
	const ruleFileDefault = "validation.yml"
	if *p.RuleFile != ruleFileDefault {
		t.Error("default value for ruleFile is ", *p.RuleFile,
			"Expected", ruleFileDefault)
	}
}

func TestVerbode(t *testing.T) {
	if *p.Verbose {
		t.Error("default value for verbose is true, expected false")
	}
}
