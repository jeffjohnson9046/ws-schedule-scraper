package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	CredentialsFile    string `env:"GOOGLE_CREDENTIALS_FILE,required"`
	CalendarId         string `env:"GOOGLE_CALENDAR_ID,required"`
	MaxResults         int64  `env:"GOOGLE_CALENDAR_MAX_RESULTS,required"`
	TimeZone           string `env:"GOOGLE_CALENDAR_TIMEZONE,required"`
	WebsiteScheduleUrl string `env:"WEBSITE_SCHEDULE_URL,required"`
}

func NewAppConfig() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	appConfig := AppConfig{}

	err = env.Parse(&appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}
