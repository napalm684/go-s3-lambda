package lambda

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/napalm684/mytest/usecase"
)

// S3TriggerHandler - S3 Trigger Handler structure
type S3TriggerHandler struct {
	interactor usecase.WorkerService
}

// NewS3TriggerHandler - constructs a new S3 trigger handler
func NewS3TriggerHandler(interactor *usecase.WorkerService) *S3TriggerHandler {
	return &S3TriggerHandler{
		interactor: *interactor,
	}
}

// StartHandler - Starts the S3Trigger Handler for the event
func (handler *S3TriggerHandler) StartHandler() {
	lambda.Start(func(request events.S3Event) error {
		return handler.interactor.Process(request)
	})
}
