package shell

import (
	"fmt"
	"os"
	"os/exec"
)

type ShellExecutor struct{}

func (s ShellExecutor) Exec(command string) error {
	fmt.Println("›", command)

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command failed: %s\n%w", command, err)
	}
	return nil
}
