package pwd

type EncryptedPassword string

// TODO: use scrypt or bcrypt.
func New(p string) *EncryptedPassword {
	ep := EncryptedPassword(p)
	return &ep
}

func (ep *EncryptedPassword) IsEqualString(p string) bool {
	return string(*ep) == string(*New(p))
}
