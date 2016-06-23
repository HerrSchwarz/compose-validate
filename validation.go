package main

import "strings"

type RuleSet struct {
  Rules map[string]Rule
}

type Rule struct {
  Services     []string
  Networks     []string
  Labels       []string
  Network_mode string
}


func checkRuleSet(rules RuleSet) (int) {
  var errors int = 0
  for _, r := range rules.Rules {
    errors += checkRule(r)
  }

  return errors
}

func checkRule (r Rule) (int) {
  var errors int = 0
  if networkModeInvalid(r.Network_mode) {
    errors++
  }
  return errors
}

func networkModeInvalid(net_mode string) (bool) {
  net_mode = strings.TrimSpace(net_mode)
  var networkModes = [5]string{"host", "bridge", "none", "overlay", "" }
  for _, mode := range networkModes {
    if net_mode == mode {
      return false
    }
  }
  logError("Network_mode " +  net_mode + " is not a valid docker network mode")
  return true
}

