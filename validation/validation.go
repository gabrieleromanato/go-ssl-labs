package validation

import "net/url"

func IsFQDN(domain string) bool {
	fullURL := "https://" + domain
	_, err := url.Parse(fullURL)
	if err != nil {
		return false
	}
	return true
}
