package entities

type Project struct {
}

type Account struct {
	name          string
	activated     bool
	authenticated bool
}

func NewAccount(name string) *Account {
	return &Account{
		name:          name,
		activated:     false,
		authenticated: false,
	}
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) IsActivated() bool {
	return a.activated
}

func (a *Account) IsAuthenticated() bool {
	return a.authenticated
}

func (a *Account) SetActivated(activated bool) {
	a.activated = activated
}

func (a *Account) SetAuthenticated(authenticated bool) {
	a.authenticated = authenticated
}
