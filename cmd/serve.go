package cmd

import (
	"log"
	"soup/internal/infrastructure"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	port            int
	host            string
	environment     string
	gracefulTimeout int
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API server",
	Long: `Start the HTTP API server with the specified configuration.
	
The server will handle graceful shutdown on SIGINT/SIGTERM signals.`,
	Example: `  # Start server on default port (8080)
  soup serve
  # Start server on custom port
  soup serve --port 3000
  # Start server with specific host
  soup serve --host 0.0.0.0 --port 8080
  # Start in production mode
  soup serve --env production`,
	RunE: runServe,
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 0, "Port to run server on (default: from env or 8080)")
	serveCmd.Flags().StringVarP(&host, "host", "H", "", "Host to bind to (default: from env or 0.0.0.0)")
	serveCmd.Flags().StringVarP(&environment, "env", "e", "development", "Environment (development|production)")
	serveCmd.Flags().IntVar(&gracefulTimeout, "graceful-timeout", 15, "Graceful shutdown timeout in seconds")
}

func runServe(cmd *cobra.Command, args []string) error {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	server := infrastructure.NewServer()
	server.Run()
	return nil
}
