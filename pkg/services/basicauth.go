package services

func BasicAuthService() *BasicAuth {
	return &BasicAuth{
		users: map[string]string{
			"foo": "bar",
		},
	}
}

type BasicAuth struct {
	users map[string]string
}

func (b *BasicAuth) Users() map[string]string {
	return b.users
}
