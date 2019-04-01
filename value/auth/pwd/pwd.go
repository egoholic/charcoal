package pwd

type EncryptedPassword string

// TODO: use scrypt or bcrypt.
func New(p string) *EncryptedPassword {
	ep := EncryptedPassword(p)
	return &ep
}

func (ep *EncryptedPassword) ToString() string {
	return string(*ep)
}
