package main

import (
	handler "github.com/napalm684/mytest/delivery/lambda"
	"github.com/napalm684/mytest/infrastructure/repository/storage"
	"github.com/napalm684/mytest/usecase"
	"github.com/pkg/errors"
)

// SetupHandler - constructs new handler for the serverless function
func SetupHandler() *handler.S3TriggerHandler {
	storageRepo, err := storage.NewS3Repository()
	if err != nil {
		panic(errors.Wrap(err, "Unable to construct storage repository, fatal error"))
	}

	interactor := usecase.NewWorkerService(storageRepo)
	handler := handler.NewS3TriggerHandler(interactor)
	return handler
}

func main() {
	SetupHandler().StartHandler()
}
