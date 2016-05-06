package param

import ("fmt"
	"strings"
	flag "github.com/ogier/pflag")

type Params struct {
  ConfigFile *string
  RuleFile *string
  Verbose *bool
}

func Init() (Params) {
  var configFile = flag.StringP("config", "c", "docker-compose.yml", "docker-compose file to validate")
  var ruleFile = flag.StringP("rules", "r", "validation.yml", "file describing the validation rules")
  var verbose = flag.BoolP("verbose", "v", false, "more output")
  params := Params{configFile, ruleFile, verbose}

  flag.Usage = func() {
    fmt.Println("Usage: compose-validate --config <docker-compose config> --rules <rule file>")
    flag.PrintDefaults()
  }

  flag.Parse()

  if (*params.Verbose) {
    var l int = 9 + max(len(*params.ConfigFile), len(*params.RuleFile))
    fmt.Printf("\n%s\n", strings.Repeat("=", l))
    fmt.Printf("config : %s\n", *params.ConfigFile)
    fmt.Printf(" rules : %s\n", *params.RuleFile)
    fmt.Printf("%s\n", strings.Repeat("=", l))
  }

  return params
}

func max(a, b int) (int) {
  if (a > b) {
    return a
  }
  return b
}
