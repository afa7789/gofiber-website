package cmd

import (
	"afa7789/site/internal/database"
	"afa7789/site/internal/domain"
	"afa7789/site/internal/server"
)

// ServerExecute is the command function that setup initial structs and value.
// After it, it will start the server.
func ServerExecute(f domain.Flags) error {
	// Setup Repositories
	r := database.NewRepositories()
	if r == nil {
		print("db as nil")
	}

	si := &domain.ServerInput{
		Reps: r,
	}

	// Setup and start server
	s := server.New(si)
	if *f.TLS {
		s.StartTLS(*f.Port)
	} else {
		s.Start(*f.Port)
	}
	return nil
}
