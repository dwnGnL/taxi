package auth

// TokenUseCase ...
type TokenUseCase interface {
	NewJWT(userId int) (*string, error)
	ParseToken(accessToken *string) (int, error)
	//NewRefreshToken() (string, error)
}
