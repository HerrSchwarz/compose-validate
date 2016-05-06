package main

import ("fmt"
	"io/ioutil"
	"os"

	"github.com/herrschwarz/compose-validate/validation"
	"github.com/herrschwarz/compose-validate/compose"
	"github.com/herrschwarz/compose-validate/param"

        "gopkg.in/yaml.v2"
       )

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  var params = param.Init()

  configData, errC := ioutil.ReadFile(*params.ConfigFile)
  validate, errV := ioutil.ReadFile(*params.RuleFile)

  check(errC)
  check(errV)

  var rules validation.Rule
  var config compose.Config
  yaml.Unmarshal(validate, &rules)
  yaml.Unmarshal(configData, &config)

  var errors int
  for _, s := range rules.Services {
    errors += validateServices(config.Services, s, *params.Verbose)
    for _, l := range rules.Labels {
      errors += validateLabel(config.Services[s], l, *params.Verbose)
    }
  }

  if errors > 0 {
    fmt.Printf("validation failed\n\n")
    os.Exit(1)
  } else {
    fmt.Printf("validation successful\n\n")
  }
}

func validateServices(services map[string]compose.Service, s string, verbose bool) (int) {
  var errors int
  if _, present := services[s]; present {
    if verbose {
      fmt.Printf("\nservice %s found\n", s)
    }
  } else {
    fmt.Printf("service %s not found\n", s)
    errors++
  }
  return errors
}

func validateLabel(s compose.Service, l string, verbose bool) (int) {
  var errors int
  if _, present := s.Labels[l]; present {
    if verbose {
      fmt.Printf("service %s has label %s\n", s, l)
    }
  } else {
    fmt.Printf("service %s should have label %s, but label is not present!\n", s, l)
    errors++
  }
  return errors
}

