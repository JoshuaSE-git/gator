package main

import (
	"fmt"

	"github.com/JoshuaSE-git/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	cfg.SetUser("Joshua")

	cfg, err = config.Read()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", *cfg)
}
