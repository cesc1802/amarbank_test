package cmd

import (
	"amarbank/cmd/handlers"
	"amarbank/pkg/httpserver"
	"amarbank/pkg/logger/consolelogger"
	"github.com/spf13/cobra"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "This command use to start http server",
	Long:  "This command use to start http server",
	RunE:  startServer,
}

func startServer(cmd *cobra.Command, args []string) error {

	appConsoleLogger := consolelogger.New()
	mode, _ := cmd.Flags().GetString("mode")
	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetString("port")
	log.Println("============================", mode)
	log.Println("============================", host)
	log.Println("============================", port)

	engine := httpserver.NewGinService(appConsoleLogger, "debug", host, port)

	engine.AddHandler(handlers.PublicAPI())
	engine.AddHandler(handlers.PrivateAPI())
	engine.AddHandler(handlers.AdminAPI())

	engine.Start()
	return nil
}
