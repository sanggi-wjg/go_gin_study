package server

import "runtime"

type Settings struct {
	Debug         bool
	Port          string
	SiteTitle     string
	SiteDeveloper string
	FavIconURL    string
	FavIconPath   string
	StaticURL     string
	StaticRoot    string
	MediaURL      string
	MediaRoot     string
	TemplateRoot  string
}

func NewSettings() *Settings {
	settings := Settings{
		Debug:         true,
		Port:          ":8080",
		SiteTitle:     "AMK_DEMO",
		SiteDeveloper: "쫑꾸",
		FavIconURL:    "/favicon.ico",
		FavIconPath:   "./static/favicon.ico",
		StaticURL:     "/static",
		StaticRoot:    "./static",
		MediaURL:      "/media",
		MediaRoot:     "./media",
		TemplateRoot:  "templates/**/*",
	}

	if runtime.GOOS == "windows" {

	}
	return &settings
}
