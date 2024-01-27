package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	solutionName := "d"
	projectName := "d"

	if err := os.MkdirAll(solutionName, 0755); err != nil {
		return
	}

	if err := os.Chdir(solutionName); err != nil {
		return
	}

	commands := []string{
		fmt.Sprintf("dotnet new sln -n %s", solutionName),
		fmt.Sprintf("dotnet new web -n %s", projectName),
		fmt.Sprintf("dotnet sln %s.sln add %s/%s.csproj", solutionName, projectName, projectName),
	}

	for _, cmd := range commands {
		fmt.Println("›", cmd)

		cmd := exec.Command("bash", "-c", cmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("command failed: %s\n%v", cmd, err)
		}
	}

	if err := os.Chdir(projectName); err != nil {
		return
	}

	commands = []string{
		"dotnet add package Microsoft.EntityFrameworkCore.SqlServer > /dev/null 2>&1",
		"dotnet add package Microsoft.EntityFrameworkCore.Design > /dev/null 2>&1",
		"dotnet add package HotChocolate.AspNetCore > /dev/null 2>&1",
		"dotnet add package HotChocolate.Data > /dev/null 2>&1",
		"dotnet add package HotChocolate.Data.EntityFramework > /dev/null 2>&1",
		"dotnet add package HotChocolate.Types.Analyzers > /dev/null 2>&1",
	}

	for _, cmd := range commands {
		fmt.Println("›", cmd)

		cmd := exec.Command("bash", "-c", cmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("command failed: %s\n%v", cmd, err)
		}
	}
}
