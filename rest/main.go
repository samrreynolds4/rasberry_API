package main

import (
	"net/http"
	"rasberryPi_api/models"
	"rasberryPi_api/rest/router"
	"rasberryPi_api/utils"
)

var logger utils.Logger

func CommandHandler(writer http.ResponseWriter, req *http.Request) {
	logger.Log("Running a command")
	var command models.Command
	err := utils.ReadRequest(req, &command)
	if err != nil {
		logger.LogError("Failed to read request")
	}

	err = utils.Run(&command)
	if err != nil {
		logger.LogErrorObj("Failed to run command", err)
	}
	err = utils.WriteResponse(writer, command)
	if err != nil {
		logger.LogErrorObj("failed to write response", err)
	}

	return
}

func main() {

	restRouter := router.NewManager(":8000")
	restRouter.AddFunc(router.PiCommand, CommandHandler, "POST")
	err := restRouter.Start()
	if err != nil {
		logger.LogErrorObj("Failed to start router", err)
	}

}
