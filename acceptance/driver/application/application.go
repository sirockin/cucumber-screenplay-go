package application

import (
	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/domain/application"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/domain/entities"
)

// New creates a new domain test driver that wraps the actual domain
func New() *AcceptanceTestDriver {
	return &AcceptanceTestDriver{
		appService: application.New(),
	}
}

// AcceptanceTestDriver is a test driver that delegates to the actual domain
type AcceptanceTestDriver struct {
	appService *application.Service
}

// verify that TestDriver implements AcceptanceTestDriver
var _ driver.AcceptanceTestDriver = (*AcceptanceTestDriver)(nil)

func (t *AcceptanceTestDriver) ClearAll() {
	t.appService.ClearAll()
}

func (t *AcceptanceTestDriver) CreateAccount(name string) error {
	return t.appService.CreateAccount(name)
}

func (t *AcceptanceTestDriver) GetAccount(name string) (entities.Account, error) {
	return t.appService.GetAccount(name)
}

func (t *AcceptanceTestDriver) Activate(name string) error {
	return t.appService.Activate(name)
}

func (t *AcceptanceTestDriver) IsActivated(name string) bool {
	return t.appService.IsActivated(name)
}

func (t *AcceptanceTestDriver) Authenticate(name string) error {
	return t.appService.Authenticate(name)
}

func (t *AcceptanceTestDriver) IsAuthenticated(name string) bool {
	return t.appService.IsAuthenticated(name)
}

func (t *AcceptanceTestDriver) GetProjects(name string) ([]entities.Project, error) {
	return t.appService.GetProjects(name)
}

func (t *AcceptanceTestDriver) CreateProject(name string) error {
	return t.appService.CreateProject(name)
}
