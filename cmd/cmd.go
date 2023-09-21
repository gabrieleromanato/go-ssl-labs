package cmd

import (
	"fmt"
	"gabrieleromanato/ssl-labs/api"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

const (
	interval = 2
)

func Run(domain string) {
	tm.Clear()
	for range time.Tick(interval * time.Second) {
		ssllabsResponse, err := api.GetSSLLabsResponse(domain)
		if err != nil {
			tm.Println(err)
			break
		}
		response := api.Response{}
		if len(ssllabsResponse.Endpoints) > 0 {
			response.Status = ssllabsResponse.Endpoints[0].StatusMessage
			response.Progress = ssllabsResponse.Endpoints[0].Progress
			response.Grade = ssllabsResponse.Endpoints[0].Grade
		} else {
			response.Status = ssllabsResponse.Status
			response.Progress = 0
			response.Grade = ""
		}

		percent := fmt.Sprintf("%d%%", response.Progress)
		outputTable := tm.NewTable(0, 10, 5, ' ', 0)
		fmt.Fprintf(outputTable, "Domain\tStatus\tProgress\tGrade\n")
		fmt.Fprintf(outputTable, "%s\t%s\t%s\t%s\n", domain, response.Status, percent, response.Grade)
		tm.Println(outputTable)
		tm.Flush()
		if strings.ToUpper(response.Status) == "ERROR" || strings.ToUpper(response.Status) == "READY" || response.Progress == 100 {
			break
		}

	}

}
