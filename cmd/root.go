package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
	envFile string
)
var rootCmd = &cobra.Command{
	Use:   "soup",
	Short: "Soup API Server",
	Long: `Soup is a REST API server built with Go and Gin.
	
It provides authentication, user management, and more.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if envFile != "" {
			if err := godotenv.Load(envFile); err != nil {
				fmt.Printf("Warning: Could not load env file '%s': %v\n", envFile, err)
			}
		} else {
			if err := godotenv.Load(); err != nil {
				fmt.Println("Warning: .env file not found, using system environment variables")
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.PersistentFlags().StringVar(&envFile, "env-file", "", "Path to .env file (default: .env)")
}

