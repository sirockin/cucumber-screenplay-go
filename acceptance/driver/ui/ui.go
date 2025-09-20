package ui

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/domain/entities"
)

type AcceptanceTestDriver struct {
	browser     playwright.Browser
	context     playwright.BrowserContext
	page        playwright.Page
	frontendURL string
}

func New(frontendURL string) (*AcceptanceTestDriver, error) {
	err := playwright.Install()
	if err != nil {
		return nil, fmt.Errorf("failed to install playwright: %w", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to start playwright: %w", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true), // Run in headless mode for CI
	})
	if err != nil {
		return nil, fmt.Errorf("failed to launch browser: %w", err)
	}

	context, err := browser.NewContext()
	if err != nil {
		return nil, fmt.Errorf("failed to create browser context: %w", err)
	}

	page, err := context.NewPage()
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %w", err)
	}

	return &AcceptanceTestDriver{
		browser:     browser,
		context:     context,
		page:        page,
		frontendURL: frontendURL,
	}, nil
}

// Close cleans up browser resources
func (u *AcceptanceTestDriver) Close() error {
	if u.browser != nil {
		return u.browser.Close()
	}
	return nil
}

// verify that AcceptanceTestDriver implements AcceptanceTestDriver
var _ driver.AcceptanceTestDriver = (*AcceptanceTestDriver)(nil)

func (u *AcceptanceTestDriver) CreateAccount(name string) error {
	log.Printf("UI: Creating account for %s", name)

	// Navigate to the account creation page
	_, err := u.page.Goto(u.frontendURL + "/signup")
	if err != nil {
		return fmt.Errorf("failed to navigate to signup page: %w", err)
	}

	// Wait for page to load
	_, err = u.page.WaitForSelector("input[name='name']", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return fmt.Errorf("signup form not found: %w", err)
	}

	// Fill in the name field
	err = u.page.Fill("input[name='name']", name)
	if err != nil {
		return fmt.Errorf("failed to fill name field: %w", err)
	}

	// Click create account button
	err = u.page.Click("button[type='submit']")
	if err != nil {
		return fmt.Errorf("failed to click create account button: %w", err)
	}

	// Wait for success message or redirect
	_, err = u.page.WaitForSelector(".success, .message", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return fmt.Errorf("account creation failed or timed out: %w", err)
	}

	return nil
}

func (u *AcceptanceTestDriver) ClearAll() {
	log.Println("UI: Clearing all data")
	// For UI testing, we'll navigate to a reset page or use the API directly
	// For simplicity, we'll use a dedicated clear endpoint
	_, err := u.page.Goto(u.frontendURL + "/admin/clear")
	if err != nil {
		log.Printf("Warning: Failed to clear data via UI: %v", err)
	}
}

func (u *AcceptanceTestDriver) GetAccount(name string) (entities.Account, error) {
	log.Printf("UI: Getting account for %s", name)

	// Navigate to account details page
	_, err := u.page.Goto(u.frontendURL + "/account/" + name)
	if err != nil {
		return entities.Account{}, fmt.Errorf("failed to navigate to account page: %w", err)
	}

	// Wait for account data to load
	_, err = u.page.WaitForSelector(".account-info", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return entities.Account{}, fmt.Errorf("account not found: %s", name)
	}

	// Extract account information from the page
	activated, err := u.page.IsVisible(".status-activated")
	if err != nil {
		activated = false
	}

	authenticated, err := u.page.IsVisible(".status-authenticated")
	if err != nil {
		authenticated = false
	}

	// Create domain account
	domainAccount := entities.NewAccount(name)
	domainAccount.SetActivated(activated)
	domainAccount.SetAuthenticated(authenticated)

	return *domainAccount, nil
}

func (u *AcceptanceTestDriver) Authenticate(name string) error {
	log.Printf("UI: Authenticating %s", name)

	// Navigate to login page
	_, err := u.page.Goto(u.frontendURL + "/login")
	if err != nil {
		return fmt.Errorf("failed to navigate to login page: %w", err)
	}

	// Wait for login form
	_, err = u.page.WaitForSelector("input[name='name']", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return fmt.Errorf("login form not found: %w", err)
	}

	// Fill in the name field
	err = u.page.Fill("input[name='name']", name)
	if err != nil {
		return fmt.Errorf("failed to fill name field: %w", err)
	}

	// Click login button
	err = u.page.Click("button[type='submit']")
	if err != nil {
		return fmt.Errorf("failed to click login button: %w", err)
	}

	// Check for error message
	time.Sleep(1 * time.Second) // Give time for response
	errorVisible, _ := u.page.IsVisible(".error")
	if errorVisible {
		errorText, _ := u.page.TextContent(".error")
		return fmt.Errorf("%s", errorText)
	}

	return nil
}

func (u *AcceptanceTestDriver) IsAuthenticated(name string) bool {
	log.Printf("UI: Checking authentication status for %s", name)

	// Navigate to account page and check authentication status
	_, err := u.page.Goto(u.frontendURL + "/account/" + name)
	if err != nil {
		return false
	}

	// Check if authenticated indicator is visible
	authenticated, err := u.page.IsVisible(".status-authenticated")
	if err != nil {
		return false
	}

	return authenticated
}

func (u *AcceptanceTestDriver) Activate(name string) error {
	log.Printf("UI: Activating account for %s", name)

	// Navigate to activation page
	_, err := u.page.Goto(u.frontendURL + "/activate/" + name)
	if err != nil {
		return fmt.Errorf("failed to navigate to activation page: %w", err)
	}

	// Wait for activation button
	_, err = u.page.WaitForSelector("button.activate", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return fmt.Errorf("activation button not found: %w", err)
	}

	// Click activate button
	err = u.page.Click("button.activate")
	if err != nil {
		return fmt.Errorf("failed to click activate button: %w", err)
	}

	// Wait for success message
	_, err = u.page.WaitForSelector(".success", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return fmt.Errorf("activation failed or timed out: %w", err)
	}

	return nil
}

func (u *AcceptanceTestDriver) CreateProject(name string) error {
	log.Printf("UI: Creating project for %s", name)

	// Navigate to projects page
	_, err := u.page.Goto(u.frontendURL + "/account/" + name + "/projects")
	if err != nil {
		return fmt.Errorf("failed to navigate to projects page: %w", err)
	}

	// Wait for create project button
	_, err = u.page.WaitForSelector("button.create-project", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return fmt.Errorf("create project button not found: %w", err)
	}

	// Click create project button
	err = u.page.Click("button.create-project")
	if err != nil {
		return fmt.Errorf("failed to click create project button: %w", err)
	}

	// Wait for success message
	_, err = u.page.WaitForSelector(".project-created", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return fmt.Errorf("project creation failed or timed out: %w", err)
	}

	return nil
}

func (u *AcceptanceTestDriver) GetProjects(name string) ([]entities.Project, error) {
	log.Printf("UI: Getting projects for %s", name)

	// Navigate to projects page
	_, err := u.page.Goto(u.frontendURL + "/account/" + name + "/projects")
	if err != nil {
		return nil, fmt.Errorf("failed to navigate to projects page: %w", err)
	}

	// Wait for projects list
	_, err = u.page.WaitForSelector(".projects-list", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return nil, fmt.Errorf("projects list not found: %w", err)
	}

	// Count project items
	projectElements, err := u.page.QuerySelectorAll(".project-item")
	if err != nil {
		return nil, fmt.Errorf("failed to find project items: %w", err)
	}

	// Create project entities (simplified - just return count)
	projects := make([]entities.Project, len(projectElements))
	for i := range projectElements {
		projects[i] = entities.Project{} // Minimal implementation
	}

	return projects, nil
}
