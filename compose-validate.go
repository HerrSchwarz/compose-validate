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
  var params = param.Init()

  configData, errC := ioutil.ReadFile(*params.ConfigFile)
  validate, errV := ioutil.ReadFile(*params.RuleFile)

  check(errC)
  check(errV)

  var rules validation.Rule
  var config compose.Config
  var errors int
  yaml.Unmarshal(validate, &rules)
  yaml.Unmarshal(configData, &config)

  for _, s := range rules.Services {
    if _, present := config.Services[s]; present {
      fmt.Printf("service %s found\n", s)
    } else {
      fmt.Printf("service %s not found\n", s)
      errors++
    }
  }

  if errors > 0 {
    os.Exit(1)
  }
}

