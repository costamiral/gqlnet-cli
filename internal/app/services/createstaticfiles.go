package services

import (
	"fmt"
	"os"
	"package/gqlnet/internal/app/utils"
	"strings"
)

func GenerateProgramCS(connStr string) error {
	content := fmt.Sprintf(`using Api.Domain.Context;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddDbContextPool<EStatementsContext>(o => o
    .UseSqlServer("%s"));

builder.Services.AddGraphQLServer()
    .AddGraphQLTypes()
    .AddProjections()
    .AddFiltering()
    .AddSorting();

var app = builder.Build();

app.MapGraphQL("/graphql");
app.Run();
`, connStr)

	return os.WriteFile("Program.cs", []byte(content), 0644)
}

func GenerateQuerieResolvers() error {
	modelToDbSet, err := utils.GetModelToDbSetMap()
	if err != nil {
		return err
	}

	if err := os.MkdirAll("Resolvers", 0755); err != nil {
		return err
	}
	if err := os.Chdir("Resolvers"); err != nil {
		return err
	}
	defer os.Chdir("..")

	var builder strings.Builder
	builder.WriteString(`using Api.Domain.Context;
using Api.Domain.Models;
using Microsoft.EntityFrameworkCore;

namespace Api.Repository;

[QueryType]
public static class Queries
{
`)
	for model, dbSet := range modelToDbSet {
		method := fmt.Sprintf(`
    [UseProjection]
    [UseFiltering]
    [UseSorting]
    public static IQueryable<%s> %sQuery(EStatementsContext context)
    {
        return context.%s.AsNoTracking();
    }
`, model, model, dbSet)
		builder.WriteString(method)
	}

	builder.WriteString("}\n")
	return os.WriteFile("Queries.cs", []byte(builder.String()), 0644)
}
