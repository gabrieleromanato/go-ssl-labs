package main

import (
	"flag"
	"gabrieleromanato/ssl-labs/cmd"
	"os"
)

func main() {

	domain := flag.String("domain", "google.com", "Domain to check")
	flag.Parse()
	domainToCheck := *domain
	cmd.Run(domainToCheck)
	os.Exit(0)
}
