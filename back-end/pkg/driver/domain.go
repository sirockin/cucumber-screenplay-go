package driver

import (
	"github.com/sirockin/cucumber-screenplay-go/back-end/internal/domain/application"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/domain/entities"
)

// New creates a new domain test driver that wraps the actual domain
func New() *DomainDriver {
	return &DomainDriver{
		appService: application.New(),
	}
}

// DomainDriver is a test driver that delegates to the actual domain
// It implements the AcceptanceTestDriver interface implicitly
type DomainDriver struct {
	appService *application.Service
}

func (t *DomainDriver) ClearAll() {
	t.appService.ClearAll()
}

func (t *DomainDriver) CreateAccount(name string) error {
	return t.appService.CreateAccount(name)
}

func (t *DomainDriver) GetAccount(name string) (entities.Account, error) {
	return t.appService.GetAccount(name)
}

func (t *DomainDriver) Activate(name string) error {
	return t.appService.Activate(name)
}

func (t *DomainDriver) IsActivated(name string) bool {
	return t.appService.IsActivated(name)
}

func (t *DomainDriver) Authenticate(name string) error {
	return t.appService.Authenticate(name)
}

func (t *DomainDriver) IsAuthenticated(name string) bool {
	return t.appService.IsAuthenticated(name)
}

func (t *DomainDriver) GetProjects(name string) ([]entities.Project, error) {
	return t.appService.GetProjects(name)
}

func (t *DomainDriver) CreateProject(name string) error {
	return t.appService.CreateProject(name)
}