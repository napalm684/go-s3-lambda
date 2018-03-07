package repository

import (
	domain "github.com/napalm684/mytest/domain"
)

// StorageRepository - Interacts with storage solution to retrieve objects as bytes
type StorageRepository interface {
	GetObject(request domain.Event) ([]byte, error)
}
