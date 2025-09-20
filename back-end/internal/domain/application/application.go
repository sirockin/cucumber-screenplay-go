package application

import (
	"fmt"

	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/domain/entities"
)

// Service provides business operations for the application
type Service struct {
	accounts map[string]*entities.Account
	projects map[entities.Account][]entities.Project
}

// New creates a new service
func New() *Service {
	d := &Service{}
	d.ClearAll()
	return d
}

// ClearAll removes all data
func (d *Service) ClearAll() {
	d.accounts = make(map[string]*entities.Account)
	d.projects = make(map[entities.Account][]entities.Project)
}

// CreateAccount creates a new account
func (d *Service) CreateAccount(name string) error {
	d.accounts[name] = entities.NewAccount(name)
	return nil
}

// GetAccount retrieves an account by name
func (d *Service) GetAccount(name string) (entities.Account, error) {
	account, exists := d.accounts[name]
	if !exists {
		return entities.Account{}, fmt.Errorf("Account not found: %s", name)
	}
	return *account, nil
}

// Activate activates an account and also authenticates the user
func (d *Service) Activate(name string) error {
	account := d.accounts[name]
	if account == nil {
		return fmt.Errorf("account not found: %s", name)
	}
	account.SetActivated(true)
	account.SetAuthenticated(true) // Activation also authenticates the user
	return nil
}

// IsActivated checks if an account is activated
func (d *Service) IsActivated(name string) bool {
	account, err := d.GetAccount(name)
	if err != nil {
		return false
	}
	return account.IsActivated()
}

// Authenticate authenticates an account (requires activation first)
func (d *Service) Authenticate(name string) error {
	account := d.accounts[name]
	if account == nil {
		return fmt.Errorf("account not found: %s", name)
	}
	if !account.IsActivated() {
		return fmt.Errorf("%s, you need to activate your account", name)
	}
	account.SetAuthenticated(true)
	return nil
}

// IsAuthenticated checks if an account is authenticated
func (d *Service) IsAuthenticated(name string) bool {
	account, err := d.GetAccount(name)
	if err != nil {
		return false
	}
	return account.IsAuthenticated()
}

// GetProjects retrieves projects for an account
func (d *Service) GetProjects(name string) ([]entities.Project, error) {
	account, err := d.GetAccount(name)
	if err != nil {
		return nil, err
	}
	return d.projects[account], nil
}

// CreateProject creates a project for an account
func (d *Service) CreateProject(name string) error {
	account, err := d.GetAccount(name)
	if err != nil {
		return err
	}
	d.projects[account] = append(d.projects[account], entities.Project{})
	return nil
}
