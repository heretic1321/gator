package commands

import (
	"github.com/heretic1321/gator/internal/config"
	"github.com/heretic1321/gator/internal/database"
	"github.com/heretic1321/gator/pkg/rss"
)

type State struct {
	DB *database.Queries
	Cfg *config.Config 
	Client *rss.Client
}
