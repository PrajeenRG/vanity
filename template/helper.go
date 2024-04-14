package template

import (
	"strings"

	"go.prajeen.com/vanity/config"
)

func pkgDocURL(pkg string) string {
	return "https://pkg.go.dev/" + pkg
}

func pkgBadgeURL(pkg string) string {
	return "https://pkg.go.dev/badge/" + pkg + ".svg"
}

func url(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "https://" + url
}

func urlWithoutScheme(url string) string {
	if !strings.ContainsAny(url, "://") {
		return url
	}
	return strings.SplitAfter(url, "://")[1]
}

func importContent(info config.PackageInfo) string {
	return info.ImportName + " " + info.VCSType + " " + info.URL
}

func sourceContent(info config.PackageInfo) string {
	return info.ImportName + " " + info.URL + " " + info.TreeURL + " " + info.BlobURL
}
