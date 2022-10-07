package main

import (
	"context"

	"github.com/cucumber/godog"
)

type Driver interface{
	CreateAccount(name string)error
	ClearAccounts()
}

type Abilities struct {
	name string
	app Driver
}

type Actor struct {
	abilities Abilities
}

type Action func(Abilities)error

func(a *Actor) AttemptsTo(actions ...Action)error{
	for i:=0; i<len(actions); i++{
		err := actions[i](a.abilities)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewActor(name string, app Driver)*Actor{
	return &Actor{
		abilities:Abilities{
			name: name, 
			app: app,	
		},
	}
}

var CreateAccount = struct {
	forThemselves Action
}{
	forThemselves: func(abilities Abilities)error{
		return abilities.app.CreateAccount(abilities.name)
	},
}


type accountFeature struct {
	actors map[string]*Actor
	app Driver
}

func(a *accountFeature) reset(){
	a.actors = make(map[string]*Actor)
	a.app.ClearAccounts()
}

func(a *accountFeature) Actor(name string)*Actor{
	if a.actors[name]==nil {
		a.actors[name]=NewActor(name, a.app)
	}
	return a.actors[name]	
}

func (a *accountFeature) personHasCreatedAnAccount(name string) error {
	return a.Actor(name).AttemptsTo(CreateAccount.forThemselves)
}

func (a *accountFeature) personHasSignedUp(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personShouldNotBeAuthenticated(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personShouldNotSeeAnyProjects(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personTriesToSignIn(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personCreatesAProject(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personShouldSeeTheirProject(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personActivatesTheirAccount(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personShouldBeAuthenticated(name string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	af := &accountFeature{
		app: NewDomainDriver(),
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {		
		af.reset()
		return ctx, nil
	})

	ctx.Step(`^(Bob|Tanya|Sue) has created an account$`, af.personHasCreatedAnAccount)
	ctx.Step(`^(Bob|Tanya|Sue) has signed up$`, af.personHasSignedUp)
	ctx.Step(`^(Bob|Tanya|Sue) should not be authenticated$`, af.personShouldNotBeAuthenticated)
	ctx.Step(`^(Bob|Tanya|Sue) should not see any projects$`, af.personShouldNotSeeAnyProjects)
	ctx.Step(`^(Bob|Tanya|Sue) should see an error telling (him|her|them) to activate the account$`, af.personShouldSeeAnErrorTellingThemToActivateTheAccount)
	ctx.Step(`^(Bob|Tanya|Sue) tries to sign in$`, af.personTriesToSignIn)
	ctx.Step(`^(Bob|Tanya|Sue) creates a project$`, af.personCreatesAProject)
	ctx.Step(`^(Bob|Tanya|Sue) should see (his|her|the) project$`, af.personShouldSeeTheirProject)
	ctx.Step(`^(Bob|Tanya|Sue) activates (his|her) account$`, af.personActivatesTheirAccount)
	ctx.Step(`^(Bob|Tanya|Sue) should be authenticated$`, af.personShouldBeAuthenticated)
}
