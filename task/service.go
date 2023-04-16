package task

type Service struct {
	token string
}

func NewService(token string) Service {
	return Service{
		token: token,
	}
}
