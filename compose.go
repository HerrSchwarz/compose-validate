package main

type Config struct {
  Version string
  Services map[string]Service
}

type Service struct {
  Labels map[string]string
  Networks map[string]string
  Network_mode string
}