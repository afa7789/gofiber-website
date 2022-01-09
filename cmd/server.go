package cmd

import "afa7789/site/internal/server"

func ServerExecute() error {
	s := server.New()
	s.Start()
	return nil
}
