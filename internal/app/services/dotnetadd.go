package services

import (
	"fmt"
	"os"
	"package/gqlnet/internal/infra"
)

type ProjectScaffolder struct {
	Executor infra.Executor
}

func NewProjectScaffolder(exec infra.Executor) *ProjectScaffolder {
	return &ProjectScaffolder{Executor: exec}
}

func (s *ProjectScaffolder) Runner(solutionName, projectName string) error {
	if err := os.MkdirAll(solutionName, 0755); err != nil {
		return err
	}

	if err := os.Chdir(solutionName); err != nil {
		return err
	}

	commands := []string{
		fmt.Sprintf("dotnet new sln -n %s", solutionName),
		fmt.Sprintf("dotnet new web -n %s", projectName),
		fmt.Sprintf("dotnet sln %s.sln add %s/%s.csproj", solutionName, projectName, projectName),
	}

	for _, cmd := range commands {
		if err := s.Executor.Exec(cmd); err != nil {
			return err
		}
	}

	if err := os.Chdir(projectName); err != nil {
		return err
	}

	if _, err := os.Stat("Properties"); os.IsNotExist(err) {
		if err := os.Mkdir("Properties", 0755); err != nil {
			fmt.Printf("Failed to create Properties directory: %v\n", err)
		}
	}

	commands = []string{
		"dotnet add package Microsoft.EntityFrameworkCore.SqlServer > /dev/null 2>&1",
		"dotnet add package Microsoft.EntityFrameworkCore.Design > /dev/null 2>&1",
		"dotnet add package HotChocolate.AspNetCore > /dev/null 2>&1",
		"dotnet add package HotChocolate.Data > /dev/null 2>&1",
		"dotnet add package HotChocolate.Data.EntityFramework > /dev/null 2>&1",
		"dotnet add package HotChocolate.Types.Analyzers > /dev/null 2>&1",
		"echo '[assembly: Module(\"GraphQLTypes\")]' > Properties/ModuleInfo.cs",
		"rm -f *.http appsettings*.json",
	}

	for _, cmd := range commands {
		if err := s.Executor.Exec(cmd); err != nil {
			return err
		}
	}

	if err := s.Executor.Exec("dotnet ef --version"); err != nil {
		s.Executor.Exec("dotnet tool install --global dotnet-ef > /dev/null 2>&1")
		s.Executor.Exec("export PATH=\"$PATH:$HOME/.dotnet/tools\"")
	}

	return nil
}
