package utils

import (
	"os/exec"
	"rasberryPi_api/models"
)

func Run(com *models.Command) (err error) {
	logger.LogObj("Running a command", com)
	var cmd *exec.Cmd

	if len(com.Args) > 0 {
		cmd = exec.Command(com.Command, com.Args...)
	} else {
		cmd = exec.Command(com.Command)
	}
	if err != nil {
		logger.LogErrorObj("Failed to get output of command", cmd)
		return err
	}

	out, err := cmd.Output()
	if err != nil {
		logger.LogErrorObj("Failed to get output", err)
		return
	}

	com.Output = string(out)
	logger.Log(com.Output)
	return nil
}
