package features

import "strings"

func normalize(domain string) string {
	domain = strings.ToLower(domain)
	parts := strings.Split(domain, ".")
	return parts[0]
}
