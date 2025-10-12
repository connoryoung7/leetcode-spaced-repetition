package services

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a AuthService) Login(email string, password string) error {
	return nil
}

func (a AuthService) Logout() {

}
