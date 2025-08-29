package main

import (
	"github.com/JoshuaSE-git/gator/internal/config"
	"github.com/JoshuaSE-git/gator/internal/database"
)

type State struct {
	cfg *config.Config
	db  *database.Queries
}
