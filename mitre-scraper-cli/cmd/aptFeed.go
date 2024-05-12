/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

var (
	cliFlag string
)

// aptFeedCmd represents the aptFeed command
var AptFeedCmd = &cobra.Command{
	Use:   "aptFeed",
	Short: "--mitre\n-m techniques\n-m tactics\n-m mitigations\n-m relationships\n-m endpoints",
	Long: `For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cliFlag == "" {
			fmt.Println("Please provide a flag.")
			return
		} else if cliFlag != "" {
			listAllData(cliFlag)
		}
	},
}

func init() {
	AptFeedCmd.Flags().StringVarP(&cliFlag, "mitre", "m", "", "Type requested data.")
}

func listAllData(cliFlag string) {
	apiUrl, apiQuery := activitiesAPIURLQueryCreate()

	apiQuery.Set("aptFeed", cliFlag)
	apiUrl.RawQuery = apiQuery.Encode()
	URL = apiUrl.String()

	go requester()

	// Serve the websocket endpoint and wait for messages
	http.HandleFunc("/ws", handler)
	http.ListenAndServe("localhost:7778", nil)
}

func activitiesAPIURLQueryCreate() (url.URL, url.Values) {
	fmt.Println("[+] APT Feed module started...")
	// Creating an url for API request.
	ApiUrl := url.URL{
		Scheme: "http",
		Host:   "localhost:7777",
		Path:   "/api/v1/apt_feed",
	}
	// Adds query verbosity string into URL
	ApiQuery := ApiUrl.Query()

	return ApiUrl, ApiQuery
}
