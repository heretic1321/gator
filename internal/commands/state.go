package commands

import (
	"github.com/heretic1321/gator/internal/config"
	"github.com/heretic1321/gator/internal/database"
)

type State struct {
	DB *database.Queries
	Cfg *config.Config 
}
