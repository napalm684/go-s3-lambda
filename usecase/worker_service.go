package usecase

import (
	"fmt"
	"log"

	"github.com/napalm684/mytest/domain"
	"github.com/napalm684/mytest/usecase/repository"
	"github.com/pkg/errors"
)

// WorkerService - provides consumer with the ability to work with cloud storage
type WorkerService struct {
	storageRepository repository.StorageRepository
}

// NewWorkerService - constructs worker service instance
func NewWorkerService(storageRepository repository.StorageRepository) *WorkerService {
	return &WorkerService{
		storageRepository: storageRepository,
	}
}

// Process - processes the event
func (w *WorkerService) Process(request domain.Event) error {
	data, err := w.storageRepository.GetObject(request)
	s := string(data)
	log.Printf("Object data: %v", s)

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to retrieve object from trigger: %v", request))
	}

	return nil
}
