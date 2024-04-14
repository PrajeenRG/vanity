package template

import "strings"

func PkgDocURL(pkg string) string {
	return "https://pkg.go.dev/" + pkg
}

func PkgBadgeURL(pkg string) string {
	return "https://pkg.go.dev/badge/" + pkg + ".svg"
}

func URL(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "https://" + url
}

func URLWithoutScheme(url string) string {
	if !strings.ContainsAny(url, "://") {
		return url
	}
	return strings.SplitAfter(url, "://")[1]
}
