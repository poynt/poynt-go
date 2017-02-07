package main

import (
  "fmt"
  "os"
  "github.com/poynt/poynt-go/poynt"
)

const applicationId = "urn:aid:b49ac4b8-c32e-4ad8-a4e2-5f791ee33dc5"


func main() {
  fmt.Println("Hello World")

  pwd, err := os.Getwd()
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }

  poyntApi := new(poynt.PoyntApi)
  poyntApi.Init(applicationId, pwd + "/src/github.com/poynt/poynt-go/poynt/keys/private-key.pem")
  err = poyntApi.GetAccessToken()
  fmt.Println(err)
  fmt.Println(poyntApi)
}
