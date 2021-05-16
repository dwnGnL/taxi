package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	authhttp "taxi/internal/auth/delivery/http"
	authpostgres "taxi/internal/auth/repository/postgres"
	authusecase "taxi/internal/auth/usecase"
	"taxi/internal/config"
	"taxi/pkg/auth/jwtImpl"
	"time"

	"github.com/gin-gonic/gin"
)

// Run ...
func Run(confpath string) error {
	conf, err := config.Init(confpath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rout := gin.Default()

	jwtToken, err := jwtImpl.NewJWT(conf.JWT.SigningKey, int64(conf.JWT.ExpiredDuration))

	if err != nil {
		fmt.Println(err)
		return err
	}
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dburl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", conf.Database.DBMS, conf.Database.Username, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.Dbname)
	fmt.Println("db url = ", dburl)

	dbRepo := authpostgres.NewDatabase(dburl)

	userRepo := authpostgres.NewUserRepository(dbRepo)

	authUseCase := authusecase.NewAuthUseCase(userRepo, jwtToken)

	userhandler := authhttp.NewUserHandler(authUseCase)

	userhandler.RegisterUser("auth", rout)

	httpServer := &http.Server{
		Addr:           conf.Http.Port,
		Handler:        rout,
		ReadTimeout:    conf.Http.ReadTimeout,
		WriteTimeout:   conf.Http.WriteTimeout,
		MaxHeaderBytes: conf.Http.MaxHeaderMegabytes << 20,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return httpServer.Shutdown(ctx)
}
