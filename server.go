package main

import (
	fiber "github.com/gofiber/fiber/v3"
	utils "github.com/parth469/go-restapi/utils"
)

func main() {
	utils.InitializedLogger()
	cnf, err := utils.LoadConfig()

	if err != nil {
		utils.LoggerMesssage("Failed to load environment variables", utils.Error)
		return
	}
	utils.LoggerMesssage("Environment variables loaded successfully", utils.Info)

	if error := utils.ConnectDB(cnf.Database); error != nil {
		utils.LoggerMesssage("fail to Connect DB"+error.Error(), utils.Error)
		return
	}
	utils.LoggerMesssage("Database connect Succfully", utils.Info)
	Server := fiber.New(fiber.Config{AppName: cnf.App})

	if err := Server.Listen(":" + cnf.Port); err != nil {
		utils.LoggerMesssage("Failed to start server: "+err.Error(), utils.Error)
	}
}
