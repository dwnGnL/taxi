package http

import (
	"fmt"
	"net/http"
	"taxi/internal/auth"
	"taxi/internal/auth/models"

	"github.com/gin-gonic/gin"
)

// Handler-used for data interaction over the http protocol
type UserHandler struct {
	// userImp - user repository implementation
	Impl auth.AuthUseCase
}

// NewHandler creates a new Handler structure
func NewUserHandler(Impl auth.AuthUseCase) *UserHandler {

	return &UserHandler{
		Impl: Impl,
	}
}

// signRequest - structure that is sent to sign-in/up url
type signRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (UserHandler *UserHandler) RegisterUser(groupname string, router *gin.Engine) {
	// user Implementation

	authEndpoints := router.Group(groupname)
	{
		authEndpoints.POST("/sign-in", UserHandler.signIn)
		authEndpoints.POST("/sign-up", UserHandler.signUp)
	}

	tokencheck := router.Group("safe")
	{
		tokencheck.Use(NewAuthTokenCheckMiddleware(UserHandler.Impl))
		tokencheck.GET("/test", func(c *gin.Context) {
			c.JSON(200, "hello")
		})
	}
}

// signUp creates an account for the user
func (userHandler *UserHandler) signUp(c *gin.Context) {
	input := new(signRequest)
	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := userHandler.Impl.SignUp(toModelsUser(input))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "the user was registered",
	})
	c.Status(http.StatusOK)
}

// signIn checks the users login and password and returns token if the user exists
func (userHandler *UserHandler) signIn(c *gin.Context) {
	input := new(signRequest)
	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println("user=>", input)

	token, err := userHandler.Impl.SignIn(toModelsUser(input))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": *token,
	})
}

type signInResponse struct {
	Token string `json:"token"`
}

func toSignInput(u *models.User) *signRequest {
	return &signRequest{
		Login:    u.Login,
		Password: u.Password,
	}
}

func toModelsUser(u *signRequest) *models.User {
	return &models.User{
		Login:    u.Login,
		Password: u.Password,
	}
}
