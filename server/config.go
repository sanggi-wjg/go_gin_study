package server

import (
	"os"
	"time"

	"github.com/go-ini/ini"
)

type Server struct {
	RunMode       string
	Debug         bool
	HttpPort      string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	FileAllowExts []string
	BasePath      string
}

type Database struct {
	Type         string
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

type App struct {
	AppTitle         string
	AppDeveloper     string
	AppDeveloperMail string
	FavIconURL       string
	FavIconPath      string
	MediaURL         string
	MediaRoot        string
	StaticURL        string
	StaticRoot       string
	TemplateRoot     string
}

var ServerConfig = &Server{}
var DatabaseConfig = &Database{}
var AppConfig = &App{}

func InitConfig() {
	// TODO : env 비교해서 ini 정해주자 env := os.Getenv("ENV")
	cfg := loadAppIni("app_debug.ini")

	cfg.Section("Server").MapTo(ServerConfig)
	cfg.Section("Database").MapTo(DatabaseConfig)
	cfg.Section("App").MapTo(AppConfig)

	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.BasePath = getBasePath()

	if ServerConfig.RunMode == "debug" {
		ServerConfig.Debug = true
	} else {
		ServerConfig.Debug = false
	}
}

func loadAppIni(path string) *ini.File {
	cfg, err := ini.Load(path)
	if err != nil {
		panic("fail to parse 'app.ini'")
		os.Exit(1)
	}
	return cfg
}

func getBasePath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	return dir
}
