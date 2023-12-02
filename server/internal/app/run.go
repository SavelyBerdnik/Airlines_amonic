package app

import "github.com/SavelyBerdnik/Airlines_amonic/server/internal/server"

func (a *application) Run() error {
	s, err := server.New(a.logger, a.cfg)

	if err != nil {
		return err
	}

	return s.Run()
}
