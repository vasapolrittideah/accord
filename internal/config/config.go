package config

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	validate "github.com/vasapolrittideah/accord/internal/validator"
	"github.com/vasapolrittideah/accord/internal/validator/translations"
	"github.com/vasapolrittideah/accord/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Constants struct {
	Environment            string        `mapstructure:"ENVIRONMENT"`
	DBHost                 string        `mapstructure:"POSTGRES_HOST"`
	DBPort                 string        `mapstructure:"POSTGRES_PORT"`
	DBUserName             string        `mapstructure:"POSTGRES_USER"`
	DBUserPassword         string        `mapstructure:"POSTGRES_PASSWORD"`
	DBName                 string        `mapstructure:"POSTGRES_DB"`
	ServerPort             string        `mapstructure:"SERVER_PORT"`
	ServerHost             string        `mapstructure:"SERVER_HOST"`
	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
	ValidationTranslator   ut.Translator
	Validate               *validator.Validate
}

type Config struct {
	Constants
	DB *gorm.DB
}

func New() (config *Config, err error) {
	constants, err := loadEnvironmentVariables()
	if err != nil {
		return nil, err
	}

	db, err := connectDatabase(constants)
	if err != nil {
		return nil, err
	}

	log.Println("🎉 Connected successfully to the database")

	config = &Config{
		Constants: *constants,
		DB:        db,
	}
	if err := migrateDatabase(config); err != nil {
		return nil, err
	}

	return
}

func loadEnvironmentVariables() (constants *Constants, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetDefault("ServerPort", "8080")
	viper.Set("ValidationTranslator", translations.RegisterTranslations(validate.Validate))
	viper.Set("Validate", validate.Validate)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&constants); err != nil {
		return nil, err
	}

	return
}

func connectDatabase(constants *Constants) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		constants.DBHost,
		constants.DBUserName,
		constants.DBUserPassword,
		constants.DBName,
		constants.DBPort,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func migrateDatabase(config *Config) error {
	config.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	return config.DB.AutoMigrate(&models.User{})
}
