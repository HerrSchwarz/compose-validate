package main

import ("fmt"
	"strings"
	flag "github.com/ogier/pflag")

type Parameter struct {
  ConfigFile *string
  RuleFile *string
  Verbose *bool
}

func Init() (Parameter) {
  params := createParams()
  setUsage()
  flag.Parse()
  printConfig(params)
  return params
}

func createParams() (Parameter) {
  var configFile = flag.StringP("config", "c", "docker-compose.yml", "docker-compose file to validate")
  var ruleFile = flag.StringP("rules", "r", "validation.yml", "file describing the validation rules")
  var verbose = flag.BoolP("verbose", "v", false, "more output")
  params := Parameter{configFile, ruleFile, verbose}
  return params
}

func setUsage() {
  flag.Usage = func() {
    fmt.Println("Usage: compose-validate --config <docker-compose config> --rules <rule file>")
    flag.PrintDefaults()
  }
}

func printConfig(params Parameter) {
  if (*params.Verbose) {
    var l int = 9 + max(len(*params.ConfigFile), len(*params.RuleFile))
    fmt.Printf("\n%s\n", strings.Repeat("=", l))
    fmt.Printf("config : %s\n", *params.ConfigFile)
    fmt.Printf(" rules : %s\n", *params.RuleFile)
    fmt.Printf("%s\n", strings.Repeat("=", l))
  }
}

func max(a, b int) (int) {
  if (a > b) {
    return a
  }
  return b
}
