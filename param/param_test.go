package param

import (
	"testing"
)

func TestInit(t *testing.T) {
	var p Parameter = Init()
	const configFileDefault = "docker-compose.yml"
	if *p.ConfigFile != configFileDefault {
		t.Error("default value for configFile is ", *p.ConfigFile,
		"Expected", configFileDefault)
	}

	const ruleFileDefault = "validation.yml"
	if *p.RuleFile != ruleFileDefault {
		t.Error("default value for ruleFile is ", *p.RuleFile,
			"Expected", ruleFileDefault)
	}

	if *p.Verbose {
		t.Error("default value for verbose is true, expected false")
	}
}
