package config

import (
  "github.com/tkanos/gonfig"

  "fmt"
  "os"
)

type Configuration struct {
    Port int
    Url string
}

func GetConfig() Configuration {
  config := Configuration{}
  
  env := os.Getenv("PROJ_ENV")
  fmt.Println("get config", env)
  
  switch env {
    case "staging":
      if err := gonfig.GetConf("env/staging.json", &config); err != nil {
        fmt.Println("> err", err)
      }
    case "production":
      if err := gonfig.GetConf("env/prod.json", &config); err != nil {
        fmt.Println("> err", err)
      }
    default:
      if err := gonfig.GetConf("env/local.json", &config); err != nil {
        fmt.Println("> err", err)
      }
  }

  return config
}