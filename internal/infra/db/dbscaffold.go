package db

import (
	"fmt"
	"os"
	"package/gqlnet/internal/infra" // update to your actual import path
)

type DbScaffold struct {
	Executor infra.Executor
}

func NewDbScaffold(exec infra.Executor) *DbScaffold {
	return &DbScaffold{Executor: exec}
}

func (s *DbScaffold) ScaffoldDatabase(connStr string) error {
	if err := os.MkdirAll("Domain", 0755); err != nil {
		return err
	}

	cmd := fmt.Sprintf(`~/.dotnet/tools/dotnet-ef dbcontext scaffold "%s" Microsoft.EntityFrameworkCore.SqlServer --context EStatementsContext --namespace Api.Domain.Models --context-namespace Api.Domain.Context --context-dir Domain/Context --output-dir Domain/Models`, connStr)

	return s.Executor.Exec(cmd)
}
