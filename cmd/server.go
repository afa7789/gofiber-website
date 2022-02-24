package cmd

import (
	"afa7789/site/internal/domain"
	"afa7789/site/internal/server"
)

func ServerExecute(f domain.Flags) error {
	s := server.New()
	s.Start(*f.Port)
	return nil
}
