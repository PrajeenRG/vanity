package config

type Config struct {
	Domain  string              `json:"domain"`
	Modules []ModuleDescription `json:"modules"`
}

type ModuleDescription struct {
	Name    string `json:"name"`
	VCSType string `json:"vcs_type"`
	URL     string `json:"url"`
}
