package main

import (
	"testing"
	flag "github.com/ogier/pflag"
	"os"
)

func TestInit(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
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

func TestPrintConfig(t *testing.T) {
	var config = "config"
	var rule = "rules"
	var verbose = true
	var p Parameter = Parameter{&config, &rule, &verbose}
	printConfig(p)
}
