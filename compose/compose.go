package compose

type Config struct {
  Version string
  Services map[string]Service
}

type Service struct {
  Labels map[string]string
}
