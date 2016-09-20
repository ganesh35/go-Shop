package lib

import (
	"os"
	"encoding/json"
	"log"
)

type GConfig struct{
	DbSettings TypeDbSettings `json:"DbSettings"`
	HttpSettings TypeHttpSettings `json:"HttpSettings"`
	LogSettings TypeLogSettings `json:"LogSettings"`
	LanguageSettings TypeLanguageSettings `json:"LanguageSettings"`
	SystemSettings TypeSystemSettings `json:"SystemSettings"`
	MailSettings TypeMailSettings `json:"MailSettings"`
	SmtpSettings TypeSmtpSettings `json:"SmtpSettings"`	
}

// strudct for SystemSettings
type TypeSystemSettings struct{
	DefaultRole string	`json:"DefaultRole"`
	DefaultManager string		`json:"DefaultManager"`
}

// struct for HttpSettings
type TypeHttpSettings struct{
	Domain string	`json:"Domain"`
	Port string		`json:"Port"`
}

type TypeLogSettings struct{
	LogFile string `json:LogFile`
	LogFolder string `json:LogFolder`
	LogFormat string `json:LogFormat`	// Supported formats: json. txt
}


type TypeLanguageSettings struct {
	Enabled bool `json:Enabled`
	Lang string `json:Lang`
	Debug bool `json:Debug`
	Languages map[string]string `json:Languages`
	Folders []string `json:Folders`
}

// struct for DbSettings
type TypeDbSettings struct{
	Domain string	`json:"Domain"`
	Port string		`json:"Port"`
	Username string		`json:"Username"`
	Password string		`json:"Password"`
	Database string		`json:"Database"`
}

// struct for smtpSettings
type TypeSmtpSettings struct{
	Smtp_enabled bool	`json:"Smtp_enabled"`
    Host string			`json:"Host"`
    Username string		`json:"Username"`
    Password string		`json:"Password"`
    Port int			`json:"Port"`
    Secure string		`json:"Secure"`
}

type TypeMailSettings struct {
	Alert_email string 	`json:"Alert_email"`
	Sender_name string `json:"Sender_name"`
	Sender_email string	`json:"Sender_email"`
}

func (p *GConfig) LoadFile(jsonfile string) (error){
	file, _ := os.Open(jsonfile)			// "config/config.json"
	decoder := json.NewDecoder(file)
	err:=decoder.Decode(&p)
	if err != nil {
		log.Println("gconfig.go::LoadFile : Could not load configuration file  : ", err)
	}
	return err;
}


// Usage
/*

Requirements:
config/config.json


var GConfig lib.GConfig
GConfig.LoadFile("config/config.json")
fmt.Println(GsConfig.HttpSettings.Domain)

*/