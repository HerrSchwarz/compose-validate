package main

import ("fmt"
	"io/ioutil"
	"os"
        gc "github.com/daviddengcn/go-colortext"

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

  var configData = readData(*params.ConfigFile)
  var validate   = readData(*params.RuleFile)

  var ruleSet validation.RuleSet
  var config compose.Config
  yaml.Unmarshal(validate, &ruleSet)
  yaml.Unmarshal(configData, &config)
  var services = config.Services

  fmt.Println(config)

  if (*params.Verbose) {
    fmt.Printf("\ndocker-compose file version: %s\n", config.Version)
    fmt.Printf("Found %d rules\n", len(ruleSet.Rules))
  }

  var errors = validateRules(ruleSet.Rules, services, *params.Verbose)
  printResult(errors)

}

func readData(fileName string) ([]byte) {
  data, err := ioutil.ReadFile(fileName)
  check(err)
  return data
}


func validateServices(services map[string]compose.Service, s string, verbose bool) (int) {
  var errors int
  if _, present := services[s]; present {
    if verbose {
      fmt.Printf("\nservice %s found\n", s)
    }
  } else {
    fmt.Printf("\nservice %s not found\n", s)
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

func validateNetwork(s compose.Service, n string, verbose bool) (int) {
  var errors int
  if _, present := s.Networks[n]; present {
    if verbose {
      fmt.Printf("service %s has network %s\n", s, n)
    }
  } else {
    fmt.Printf("service %s should have network %s, but network is not present!\n", s, n)
    errors++
  }
  return errors
}

func validateRules(rules map[string]validation.Rule, services map[string]compose.Service, verbose bool) (int) {
  var errors int
  for name, rule := range rules {
    fmt.Printf("\nValidating %s:\n", name)
    for _, s := range rule.Services {
      errors += validateServices(services, s, verbose)
      for _, l := range rule.Labels {
        errors += validateLabel(services[s], l, verbose)
      }
      for _, n := range rule.Networks {
        errors += validateNetwork(services[s], n, verbose)
      }
    }
  }
  return errors
}

func printResult(errors int) {
  if errors > 0 {
    gc.Foreground(gc.Red, false)
    fmt.Printf("\nvalidation failed\n\n")
    gc.ResetColor()
    os.Exit(1)
  } else {
    gc.Foreground(gc.Green, false)
    fmt.Printf("\nvalidation successful\n\n")
    gc.ResetColor()
  }
}
