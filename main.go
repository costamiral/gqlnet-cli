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
}
