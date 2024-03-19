package main

import (
	"github.com/junioralcant/api-stores-go/infra/config"
	"github.com/junioralcant/api-stores-go/main/router"
)

func main() {
	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
