package password

type Password struct {
	origin    *string // could'd be nil
	encrypted string
}

func New(pwd string) Password {
	return Password{&pwd, pwd}
}

func (p *Password) ToString() string {
	return p.encrypted
}

func (p *Password) IsEqual(pwd string) bool {
	newPwd := New(pwd)
	return p.ToString() == newPwd.ToString()
}
