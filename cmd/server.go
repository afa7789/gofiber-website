package cmd

import (
	"afa7789/site/internal/database"
	"afa7789/site/internal/domain"
	"afa7789/site/internal/server"
)

func ServerExecute(f domain.Flags) error {
	// Setup Repositories
	r := database.NewRepositories()
	if r == nil {
		print("db as nil")
	}

	si := &domain.ServerInput{
		Reps: r,
	}

	// posts, _ := si.Reps.PostRep.LastThreePosts()
	// print("\n", len(posts), "\n")
	// for _, p := range posts {
	// 	print("title:", p.Title, "\n")
	// 	print("date:", p.CreatedAt.String(), "\n")
	// }

	// Setup and start server
	s := server.New(si)
	s.Start(*f.Port)
	return nil
}
