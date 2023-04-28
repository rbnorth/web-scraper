/*
Copyright © 2023 NAME HERE rbnorthcutt@gmail.com
*/

package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var url string

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape a website",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Your scraping code goes here, for example:
		doc.Find("h1").Each(func(i int, s *goquery.Selection) {
			fmt.Printf("Heading %d: %s\n", i+1, s.Text())
		})
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)

	scrapeCmd.Flags().StringVar(&url, "url", "", "URL to scrape")
	scrapeCmd.MarkFlagRequired("url")
}

