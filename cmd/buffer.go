package main

import (
	"fmt"
	"github.com/coff/go-hellfire/pkg/config"
	"github.com/coff/go-hellfire/pkg/system"
)

func main() {
	cfg := config.Config{Filepath: "hellfire.yaml"}
	err := cfg.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bfr := system.Buffer{}
	err = bfr.Bootstrap(&cfg)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
