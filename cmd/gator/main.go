package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/heretic1321/gator/internal/cli"
	"github.com/heretic1321/gator/internal/commands"
	"github.com/heretic1321/gator/internal/config"
	"github.com/heretic1321/gator/internal/database"
	"github.com/heretic1321/gator/pkg/rss"
	_ "github.com/lib/pq"
)


func main(){
	conf, err := config.New()
	if err != nil {
		fmt.Println(err)
	}
	
	db, err := sql.Open("postgres", conf.DBURL)
	dbQueries := database.New(db)
	
	state := commands.State{
		DB: dbQueries,
		Cfg: &conf,
		Client: rss.New(&http.Client{}),
	}


	app := cli.New(state)


	args := os.Args[1:]	
	err = app.Run(args)

	if err != nil {
		log.Fatal(err.Error())
	}
}
