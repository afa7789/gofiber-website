package cmd

import (
	"afa7789/site/internal/database"
	"afa7789/site/internal/domain"
	"afa7789/site/internal/server"
)

func ServerExecute(f domain.Flags) error {
	// Setup Repositories
	r := database.NewRepositories()
	si := &domain.ServerInput{
		repositories: r,
	}
	// Setup and start server
	s := server.New(si)
	s.Start(*f.Port)
	return nil
}
