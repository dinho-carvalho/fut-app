package main

import (
	"log/slog"

	"fut-app/internal/database/gateway"
	"fut-app/internal/database/repositories"
	"fut-app/internal/usecase"

	"fut-app/internal/database"
)

type Dependencies struct {
	usecase.RegisterPlayerUseCase
}

func InjectDependencies(db *database.Database, logger *slog.Logger) Dependencies {
	repo := repositories.NewPlayer(db.DB, logger)
	rg := gateway.NewRegisterPlayerGateway(repo)
	p := usecase.NewPlayerUseCase(rg)

	return Dependencies{
		RegisterPlayerUseCase: p,
	}
}
