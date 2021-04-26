package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"accounting_migration/config"

	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "boot",
	Short: "Boot Accounting Migration http service.",
	Run: func(cmd *cobra.Command, args []string) {
		config.Start(".")

		e := echo.New()

		s := &http.Server{
			Addr:         ":" + viper.GetString("PORT"),
			ReadTimeout:  time.Duration(viper.GetInt64("HTTP_READ_TIMEOUT")) * time.Second,
			WriteTimeout: time.Duration(viper.GetInt64("HTTP_WRITE_TIMEOUT")) * time.Second,
		}

		go func() {
			if err := e.StartServer(s); err != nil {
				e.Logger.Info("Shutting down the server.")
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)

		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
