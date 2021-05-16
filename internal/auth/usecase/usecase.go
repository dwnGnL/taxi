package usecase

import (
	"crypto/sha1"
	"fmt"
	"taxi/internal/auth"
	"taxi/internal/auth/models"

	token "taxi/pkg/auth"
)

// AuthUseCase ...
type AuthUseCase struct {
	userRepo auth.UserRepository
	jwtImpl  token.TokenUseCase
}

// NewAuthUseCase ...
func NewAuthUseCase(userRepo auth.UserRepository, jwt token.TokenUseCase) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
		jwtImpl:  jwt,
	}
}

// ParseToken ...
func (a *AuthUseCase) ParseToken(token *string) (int, error) {
	return a.jwtImpl.ParseToken(token)
}

// SignIn ...
func (a *AuthUseCase) SignIn(user *models.User) (*string, error) {
	fmt.Println("(a *AuthUseCase) SignIn start")
	user.Password = *hash(&user.Password)

	id, err := a.userRepo.SingIn(user)
	if err != nil {
		return nil, err
	}
	fmt.Println("(a *AuthUseCase) SignIn end")
	return a.jwtImpl.NewJWT(id)
}

// SignUp ...
func (a *AuthUseCase) SignUp(user *models.User) error {

	user.Password = *hash(&user.Password)

	err := a.userRepo.SignUp(user)
	if err != nil {
		return err
	}
	return nil
}

func hash(pass *string) *string {

	h := sha1.New()

	h.Write([]byte(*pass))

	hash := fmt.Sprintf("%x\n", h.Sum(nil))

	return &hash
}
