package main

type RuleSet struct {
  Rules map[string]Rule
}

type Rule struct {
  Services []string
  Networks []string
  Labels   []string
}
