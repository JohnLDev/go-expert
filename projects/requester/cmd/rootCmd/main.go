package rootCmd

import (
	"fmt"
	"os"

	"github.com/johnldev/requester/internal/handlers"
	"github.com/spf13/cobra"
)

var Url string
var Requests int
var Concurrency int

var rootCmd = &cobra.Command{
	Use:   "Stress Tester",
	Short: "This is a stress tester for websites. It will make a number of requests to a given url.",
	Long:  `This is a stress tester for websites. It will make a number of requests to a given url.`,

	Run: func(cmd *cobra.Command, args []string) {
		handlers.StressTestHandler(Url, Requests, Concurrency)
	},
}

func init() {
	// ? defining flags
	rootCmd.PersistentFlags().StringVar(&Url, "url", "", "url to request")
	rootCmd.MarkPersistentFlagRequired("url")

	rootCmd.PersistentFlags().IntVar(&Requests, "requests", 0, "amount of requests to make")
	rootCmd.MarkPersistentFlagRequired("requests")

	rootCmd.PersistentFlags().IntVar(&Concurrency, "concurrency", 1, "amount of concurrent requests to make")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
