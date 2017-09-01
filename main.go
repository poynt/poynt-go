package main

import (
	"github.com/poynt/poynt-go/poynt"
)

func main() {
	var applicationId = ""
	var privateKeyPath = ""

	poyntApi := new(poynt.PoyntApi)
	poyntApi.Init(applicationId, privateKeyPath)
}
