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

var cfg *ini.File

func InitConfig() {
	// TODO : env 비교해서 ini 정해주자 env := os.Getenv("ENV")
	cfg = loadAppIni("conf/app_debug.ini")

	mapTo("Server", ServerConfig)
	mapTo("Database", DatabaseConfig)
	mapTo("App", AppConfig)

	configToGoReadable()
}

func loadAppIni(path string) *ini.File {
	cfg, err := ini.Load(path)
	if err != nil {
		panic("fail to parse: " + path)
	}
	return cfg
}

func mapTo(section string, config interface{}) {
	err := cfg.Section(section).MapTo(config)
	if err != nil {
		panic("fail to load section: " + section)
	}
}

func configToGoReadable() {
	ServerConfig.BasePath = getBasePath()
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second

	if ServerConfig.RunMode == "debug" {
		ServerConfig.Debug = true
	} else {
		ServerConfig.Debug = false
	}
}

func getBasePath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
