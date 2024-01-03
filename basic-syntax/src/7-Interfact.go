package basicSyntax

type Account struct {
	Name string
	Age  int
}

type AccountService interface {
	SetProfile(pf profile)
	GetProfile() Account
}

type profile struct {
	Account *Account
}

func NewProfile(acc *Account) AccountService {
	return &profile{Account: acc}
}

func (p profile) GetProfile() Account {
	return *p.Account
}

func (p *profile) SetProfile(pf profile) {
	p.Account = pf.Account
}
