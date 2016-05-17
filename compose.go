package main

type Config struct {
  Version string
  Services map[string]Service
}

type Service struct {}
