package middleware

type tokenManager interface {
	Parse(token string) (string, error)
}

type Middleware struct {
	tokenManager tokenManager
}

func New(
	manager tokenManager,
) *Middleware {
	return &Middleware{
		tokenManager: manager,
	}
}
