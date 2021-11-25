package config

import (
	"time"

	"github.com/go-ini/ini"
)

type Server struct {
	RunMode      string
	Debug        bool
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
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
	AppTitle             string
	AppDeveloper         string
	AppDeveloperMail     string
	JwtSecret            string
	UploadMaxSize        int
	ExcelUploadAllowExts []string
	PageSize             int

	SessPath     string
	SessMaxAge   int
	LogSavePath  string
	LogFileExt   string
	FavIconURL   string
	FavIconPath  string
	MediaURL     string
	MediaRoot    string
	StaticURL    string
	StaticRoot   string
	TemplateRoot string
}

var ServerConfig = &Server{}
var DatabaseConfig = &Database{}
var AppConfig = &App{}

var cfg *ini.File

func Init() {
	// TODO : env 비교해서 ini 정해주자 env := os.Getenv("ENV")
	cfg = loadAppIni("conf/app_debug.ini")

	mapTo("server", ServerConfig)
	mapTo("database", DatabaseConfig)
	mapTo("app", AppConfig)

	configToGoReadable()
}

func configToGoReadable() {
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second

	if ServerConfig.RunMode == "debug" {
		ServerConfig.Debug = true
	} else {
		ServerConfig.Debug = false
	}
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
