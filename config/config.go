package config

type Config struct {
	Domain  string              `json:"domain"`
	Modules []ModuleDescription `json:"modules"`
}

type ModuleDescription struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	VCSType     string `json:"vcs_type"`
	URL         string `json:"url"`
}

type PackageInfo struct {
	ImportName  string
	Name        string
	Description string
	VCSType     string
	URL         string
	TreeURL     string
	BlobURL     string
}

func ProcessConfig(config Config) []PackageInfo {
	var packages []PackageInfo
	for _, v := range config.Modules {
		packages = append(packages, PackageInfo{
			ImportName:  config.Domain + "/" + v.Name,
			Name:        v.Name,
			Description: v.Description,
			VCSType:     v.VCSType,
			URL:         v.URL,
			TreeURL:     v.URL + "/tree/master{/dir}",
			BlobURL:     v.URL + "/blob/master{/dir}/{file}#L{line}",
		})
	}
	return packages
}
