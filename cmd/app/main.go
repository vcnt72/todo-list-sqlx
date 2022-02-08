package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vcnt72/golang-boilerplate/config"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()

	defer logger.Sync()
	config.Init()
	g := gin.Default()

	port := viper.GetString("server.port")

	srv := http.Server{
		Addr:    ":" + port,
		Handler: g,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {

		logger.Sugar().Infof("Listen: %s", port)

		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.Sugar().Infof("Listen: %s", err)
		}

	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Sugar().Fatal("Server forced to shutdown:", err)
	}

	logger.Info("Server exiting")

}
