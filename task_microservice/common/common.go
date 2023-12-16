package common

import (
	"encoding/json"
	"os"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

// Configuration stores setting values
type Configuration struct {
	Port                string `json:"port"`
	EnableGinConsoleLog bool   `json:"enableGinConsoleLog"`
	EnableGinFileLog    bool   `json:"enableGinFileLog"`

	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`

	DbAddrsWaizly       string `json:"db_addrs_waizly"`
	DbAddrsWaizlyPublic string `json:"db_addrs_waizly_public"`
	DbNameWaizly        string `json:"db_name_waizly"`
	DbUsernameWaizly    string `json:"db_username_waizly"`
	DbPasswordWaizly    string `json:"db_password_waizly"`
	DBPortWaizly        string `json:"db_port_waizly"`

	JwtSecretPassword string `json:"jwtSecretPassword"`
	Issuer            string `json:"issuer"`

	//ServiceHelper  string `json:"service_helper"`
	//ServiceImage   string `json:"service_image"`
	//ServicePublish string `json:"service_publish"`
	//ServiceMaster  string `json:"service_master"`
	//ServiceLog     string `json:"service_log"`
	//
	//UrlAuth string `json:"url_auth"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

const (
	StatusErrCreate      = "Gagal Menyimpan"
	StatusErrDelete      = "Gagal Menghapus"
	StatusErrUpdate      = "Gagal Mengedit"
	StatusSuccCreate     = "Berhasil Menyimpan"
	StatusSukses         = "Sukses"
	StatusGetNoData      = "Tidak ada data"
	StatusGetExistedData = "Sudah ada data"
	Error                = "Error"
)

// LoadConfig loads configuration from the config file
func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	// Setting Service Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   Config.LogFilename,
		MaxSize:    Config.LogMaxSize,    // megabytes after which new file is created
		MaxBackups: Config.LogMaxBackups, // number of backups
		MaxAge:     Config.LogMaxAge,     // days
	})
	log.SetLevel(log.DebugLevel)

	// log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
