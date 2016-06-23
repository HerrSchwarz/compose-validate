package main

import ("fmt"
	"io/ioutil"
	"os"
        "gopkg.in/yaml.v2"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  var params = Init()

  var configData = readData(*params.ConfigFile)
  var validate   = readData(*params.RuleFile)

  var ruleSet RuleSet
  var config Config
  yaml.Unmarshal(validate, &ruleSet)
  yaml.Unmarshal(configData, &config)
  var ruleSetErrors = checkRuleSet(ruleSet)
  if ruleSetErrors > 0 {
    logError("found errors in rule set. Please fix the errors and restart validation.")
    os.Exit(2)
  }
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


func validateServices(services map[string]Service, s string, verbose bool) (int) {
  var errors int
  if _, present := services[s]; present {
    logVerbose("service " + s + " found", verbose)
  } else {
    logError("service " + s + " not found\n")
    errors++
  }
  return errors
}

func validateLabel(name string, s Service, l string, verbose bool) (int) {
  var errors int
  if _, present := s.Labels[l]; present {
    logVerbose("service " + name + " has label " + l, verbose)
  } else {
    logError("service " + name + " should have label " + l + ", but label is not present!")
    errors++
  }
  return errors
}

func validateNetwork(name string, s Service, n string, verbose bool) (int) {
  var errors int
  if _, present := s.Networks[n]; present {
    logVerbose("servicehas network " + n, verbose)
  } else {
    logError("service " + name + " should have network " + n + ", but network is not present!")
    errors++
  }
  return errors
}

func validateNetworkMode(name string, s Service, net_mode string, verbose bool) (int) {
  var errors int
  if s.Network_mode == net_mode {
      logVerbose("service " + name + " has network mode " + net_mode, verbose)
  } else {
    logError("service " + name + " should have network mode " + net_mode + ", but is " + s.Network_mode)
    errors++
  }
  return errors
}

func validateRules(rules map[string]Rule, services map[string]Service, verbose bool) (int) {
  var errors int
  for name, rule := range rules {
    fmt.Printf("\nValidating %s:\n", name)
    for _, s := range rule.Services {
      errors += validateServices(services, s, verbose)
      for _, l := range rule.Labels {
        errors += validateLabel(s, services[s], l, verbose)
      }
      for _, n := range rule.Networks {
        errors += validateNetwork(s, services[s], n, verbose)
      }
      errors += validateNetworkMode(s, services[s], rule.Network_mode, verbose)
    }
  }
  return errors
}

func printResult(errors int) {
  if errors > 0 {
    logError("\nvalidation failed\n\n")
    os.Exit(1)
  } else {
    logSuccess("\nvalidation successful\n\n")
  }
}
