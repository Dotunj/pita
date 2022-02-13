package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pita",
	Short: "Pita generates random quotes for you",
	Run: func(cmd *cobra.Command, args []string) {

		// use retryable http
		client := retryablehttp.NewClient()
		client.RetryMax = 10

		// get current time and set the date and month format
		currentTime := time.Now()
		month := currentTime.Format("1")
		date := currentTime.Format("2")

		req, err := retryablehttp.NewRequest(http.MethodGet, "http://numbersapi.com/"+month+"/"+date+"/date?json", nil)
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		var response map[string]interface{}

		fail := json.NewDecoder(res.Body).Decode(&response)
		if fail != nil {
			log.Fatalln(err)
		}

		resp, err := json.Marshal(response)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(resp))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
