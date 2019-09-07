package main

import (
	"fmt"

	"github.com/PC-SAFT/env"
)

func main() {
	fmt.Println("testing calculation app")
	env := env.GetAppEnv()
	fmt.Println("%v", env)
}
