package config

import (
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log/slog"
	"time"
)

type Database struct {
	User     string `mapstructure:"user" yaml:"user" json:"user"`
	Password string `mapstructure:"password" yaml:"password" json:"password"`
	Host     string `mapstructure:"host" yaml:"host" json:"host"`
	Name     string `mapstructure:"name" yaml:"name" json:"name"`
	Port     int    `mapstructure:"port" yaml:"port" json:"port"`
	SSLMode  bool   `mapstructure:"ssl_mode" yaml:"ssl_mode" json:"ssl_mode"`
}

func (db Database) Connect() (*sql.DB, error) {
	SSLMode := "require"
	if !db.SSLMode {
		SSLMode = "disable"
	}

	if db.Port == 0 {
		db.Port = 5432
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=%s", db.User, db.Password, db.Host, db.Name, db.Port, SSLMode)
	return sql.Open("postgres", connStr)
}

func (db Database) ConnectAndPing(log *slog.Logger) (*sql.DB, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	for {
		log.Info("Trying to ping database...")
		if err := conn.Ping(); err == nil {
			log.Info("OK, connected to database")
			break
		} else {
			log.Error("Failed to connect to database", "error", err)
		}
		time.Sleep(5 * time.Second)
	}
	return conn, nil
}

type Server struct {
	Mode         string
	Hostname     string
	Url          string
	Port         string
	ProblemsDir  string
	TemplatesDir string

	CookieSecret string

	GoogleAuth struct {
		Enabled   bool   `json:"enabled" mapstructure:"enabled"`
		ClientKey string `json:"client_key" mapstructure:"client_key"`
		Secret    string `json:"secret" mapstructure:"secret"`
		Callback  string `json:"callback" mapstructure:"callback"`
	} `json:"googleAuth" mapstructure:"googleAuth"`

	Sendgrid struct {
		Enabled       bool   `json:"enabled" mapstructure:"enabled"`
		ApiKey        string `json:"api_key" mapstructure:"api_key"`
		SenderName    string `json:"sender_name" mapstructure:"sender_name"`
		SenderAddress string `json:"sender_address" mapstructure:"sender_address"`
	} `json:"sendgrid" mapstructure:"sendgrid"`

	SMTP struct {
		Enabled             bool   `json:"enabled" mapstructure:"enabled"`
		MailAccount         string `json:"mail_account" mapstructure:"mail_account"`
		MailServerHost      string `json:"mail_server" mapstructure:"mail_server"`
		MailServerPort      string `json:"mail_port" mapstructure:"mail_port"`
		MailAccountPassword string `json:"mail_password" mapstructure:"mail_password"`
	} `json:"smtp" mapstructure:"smtp"`

	Database `json:"database" mapstructure:"database"`

	Keys Keys
}

type Keys struct {
	PrivateKeyLocation string `json:"private_key" mapstructure:"private_key"`
	PublicKeyLocation  string `json:"public_key" mapstructure:"public_key"`
	PrivateKey         *rsa.PrivateKey
	PublicKey          *rsa.PublicKey
}

func (k *Keys) Parse() error {
	if k.PrivateKeyLocation != "" {
		if k.PublicKeyLocation == "" {
			return errors.New("private key filled, public not")
		}

		privateKeyContents, err := ioutil.ReadFile(k.PrivateKeyLocation)
		if err != nil {
			return err
		}

		block, _ := pem.Decode(privateKeyContents)
		if block == nil {
			return fmt.Errorf("can't parse pem private key file: %s", k.PrivateKeyLocation)
		}

		var pKey any
		if pKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return err
		}
		k.PrivateKey = pKey.(*rsa.PrivateKey)

		publicKeyContents, err := ioutil.ReadFile(k.PublicKeyLocation)
		if err != nil {
			return err
		}

		block, _ = pem.Decode(publicKeyContents)
		if block == nil {
			return fmt.Errorf("can't parse pem public key file: %s", k.PrivateKeyLocation)
		}

		if k.PublicKey, err = x509.ParsePKCS1PublicKey(block.Bytes); err != nil {
			return err
		}
	}

	return nil
}

func (k *Keys) MustParse() {
	if err := k.Parse(); err != nil {
		panic(err)
	}
}
