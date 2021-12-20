package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/l1f/blockornot/internal/application"
)

func Start(appCtx *application.Context) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", appCtx.Config.Application.Port),
		Handler:      router(appCtx),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		appCtx.Logger.Println("caught signal", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		appCtx.Logger.Println("Completing background tasks", map[string]string{
			"addr": srv.Addr,
		})

		appCtx.Wg.Wait()
		shutdownError <- nil
	}()

	appCtx.Logger.Println("Stating server", map[string]string{
		"add": srv.Addr,
		"env": string(appCtx.Config.Application.Env),
	})

	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	appCtx.Logger.Println("Server stopped", map[string]string{
		"addr": srv.Addr,
	})

	return nil
}
