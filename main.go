package main

import (
	"fmt"

	"github.com/poynt/poynt-go/poynt"
)

func main() {
	fmt.Println("Hello World")
	var applicationId = ""
	var privateKeyPath = ""

	poyntApi := new(poynt.PoyntApi)
	poyntApi.Init(applicationId, privateKeyPath)
}
