package api

import (
	"encoding/json"
	"errors"
	"gabrieleromanato/ssl-labs/validation"
	"io"
	"net/http"
)

type Endpoint struct {
	IPAddress            string `json:"ipAddress"`
	ServerName           string `json:"serverName"`
	StatusMessage        string `json:"statusMessage"`
	StatusDetails        string `json:"statusDetails"`
	StatusDetailsMessage string `json:"statusDetailsMessage"`
	Grade                string `json:"grade"`
	GradeTrustIgnored    string `json:"gradeTrustIgnored"`
	HasWarnings          bool   `json:"hasWarnings"`
	IsExceptional        bool   `json:"isExceptional"`
	Progress             int    `json:"progress"`
	Duration             int    `json:"duration"`
	ETA                  int    `json:"eta"`
	Delegation           int    `json:"delegation"`
}

type SSLLabsQuery struct {
	Domain string
}

type SSLLabsResponse struct {
	Host            string     `json:"host"`
	Port            int        `json:"port"`
	Protocol        string     `json:"protocol"`
	IsPublic        bool       `json:"isPublic"`
	Status          string     `json:"status"`
	StartTime       int        `json:"startTime"`
	TestTime        int        `json:"testTime"`
	EngineVersion   string     `json:"engineVersion"`
	CriteriaVersion string     `json:"criteriaVersion"`
	Endpoints       []Endpoint `json:"endpoints"`
}

type Response struct {
	Status   string
	Progress int
	Grade    string
}

const (
	SSLLabsUrl = "https://api.ssllabs.com/api/v3/analyze?host="
)

func GetSSLLabsResponse(domain string) (SSLLabsResponse, error) {
	var ssllabsResponse SSLLabsResponse
	if !validation.IsFQDN(domain) {
		return ssllabsResponse, errors.New("Invalid domain")
	}
	request, err := http.NewRequest("GET", SSLLabsUrl+domain, nil)
	if err != nil {
		return ssllabsResponse, errors.New("Invalid request")
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ssllabsResponse, errors.New("Invalid request")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ssllabsResponse, errors.New("Invalid response")
	}

	err = json.Unmarshal(body, &ssllabsResponse)
	if err != nil {
		return ssllabsResponse, errors.New("Invalid JSON")
	}
	return ssllabsResponse, nil
}
