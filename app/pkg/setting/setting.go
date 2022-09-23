package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App ****************************************************************
type App struct {
	JwtSecret string
	PageSize int
	RuntimeRootPath string
	PrefixUrl string

	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	ExportSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}

var AppSetting = &App{}

// Server ****************************************************************
type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

// Database ****************************************************************
type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}

// Kafka ****************************************************************
type Kafka struct {
	Topic string
	GroupId string
	BootStrapServers []string
	SecurityProtocol string
	SaslMechanism string
	SaslUsername string
	SaslPassword string
}

var KafkaSetting = &Kafka{}

// 配置文件的路径
const source = "app/conf/app.ini"

func Setup() {
	// 找到 配置文件
	Cfg, err := ini.Load(source)
	if err != nil {
		log.Fatalf("Fail to parse 'app/conf/app.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}

	err = Cfg.Section("kafka").MapTo(KafkaSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo KafkaSetting err: %v", err)
	}
}