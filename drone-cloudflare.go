package main

import (
	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/drone/drone-go/plugin"
)

var (
	buildCommit string
)

type CloudFlare struct {
	APIKey string `json:"apikey"`
	Email  string `json:"email"`
	Domain string `json:"domain"`
}

func main() {
	fmt.Printf("Drone CloudFlare Plugin built from %s\n", buildCommit)

	args := CloudFlare{}

	plugin.Param("vargs", &args)

	err := plugin.Parse()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(args.APIKey) == 0 || len(args.Email) == 0 || len(args.Domain) == 0 {
		fmt.Println("Incorrect parameters specified")
		os.Exit(1)
	}

	cf, err := cloudflare.New(args.APIKey, args.Email)

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
		if zone.Name == args.Domain {
			zoneID = zone.ID
		}
	}

	if len(zoneID) == 0 {
		fmt.Printf("Unable to find zone for: %s\n", args.Domain)
		os.Exit(1)
	}

	pcr, err := cf.PurgeEverything(zoneID)

	if err != nil || !pcr.Success {
		fmt.Printf("Unable to purge cache for zone ID: %s\n", zoneID)
		os.Exit(1)
	}

	fmt.Printf("Cache successfully purged for %s (zone ID: %s)\n", args.Domain, zoneID)
}