package driver

import (
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/domain/entities"
)

// AcceptanceTestDriver is our interface to the system under test
type AcceptanceTestDriver interface {
	CreateAccount(name string) error
	ClearAll()
	GetAccount(name string) (entities.Account, error)
	Authenticate(name string) error
	IsAuthenticated(name string) bool
	Activate(name string) error
	CreateProject(name string) error
	GetProjects(name string) ([]entities.Project, error)
}
