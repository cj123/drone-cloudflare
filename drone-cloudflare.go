package main

import (
	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

var (
	buildCommit string
)

func main() {
	fmt.Printf("Drone CloudFlare Plugin built from %s\n", buildCommit)

	var (
		apikey = os.Getenv("PLUGIN_APIKEY")
		email  = os.Getenv("PLUGIN_EMAIL")
		domain = os.Getenv("PLUGIN_DOMAIN")
	)

	if apikey == "" || email == "" || domain == "" {
		fmt.Println("Incorrect parameters specified")
		os.Exit(1)
	}

	cf, err := cloudflare.New(apikey, email)

	if err != nil {
		fmt.Printf("Unable to initialise client: %s\n", err.Error())
		os.Exit(1)
	}

	// find the zone ID for the domain
	zones, err := cf.ListZones()

	if err != nil {
		fmt.Printf("Unable to list zones: %s\n", err.Error())
		os.Exit(1)
	}

	var zoneID string

	for _, zone := range zones {
		if zone.Name == domain {
			zoneID = zone.ID
		}
	}

	if len(zoneID) == 0 {
		fmt.Printf("Unable to find zone for: %s\n", domain)
		os.Exit(1)
	}

	pcr, err := cf.PurgeEverything(zoneID)

	if err != nil || !pcr.Success {
		fmt.Printf("Unable to purge cache for zone ID: %s\n", zoneID)
		os.Exit(1)
	}

	fmt.Printf("Cache successfully purged for %s (zone ID: %s)\n", domain, zoneID)
}
