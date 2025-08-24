package main

import (
	"fmt"
	"log"
	"os"
	"github.com/heretic1321/gator/cli"
	"github.com/heretic1321/gator/internal/config"
)

func main(){
	conf, err := config.New()

	if err != nil {
		fmt.Println(err)
	}
	app := cli.New(&conf)

	args := os.Args[1:]	
	err = app.Run(args)

	if err != nil {
		log.Fatal(err.Error())
	}
}
