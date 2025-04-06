package main

import (
	"log"

	"package/gqlnet/internal/app/services"
	"package/gqlnet/internal/app/utils"
	"package/gqlnet/internal/domain/models"
	"package/gqlnet/internal/infra/db"
	"package/gqlnet/internal/infra/shell"
)

func main() {
	cfg, err := utils.Load("config.json")
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	err = handler(cfg)
	if err != nil {
		log.Fatalf("generation failed: %v", err)
	}
}

func handler(cfg *models.Config) error {
	connStr := utils.BuildConnectionString(cfg)

	executor := shell.ShellExecutor{}

	ronin := services.NewProjectScaffolder(executor)
	if err := ronin.Runner(cfg.SolutionName, cfg.ProjectName); err != nil {
		log.Fatal(err)
		return err
	}

	scaffolder := db.NewDbScaffold(executor)
	if err := scaffolder.ScaffoldDatabase(connStr); err != nil {
		return err
	}

	if err := services.GenerateProgramCS(connStr); err != nil {
		return err
	}

	if err := services.GenerateQuerieResolvers(); err != nil {
		return err
	}

	return nil
}
