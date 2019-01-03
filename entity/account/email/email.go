package email

type Email string

func New(email string) Email {
	return Email(email)
}
