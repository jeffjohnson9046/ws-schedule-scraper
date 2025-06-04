package config

type AppConfig struct {
	CredentialsFile string `env:"GOOGLE_CREDENTIALS_FILE,required"`
	CalendarId      string `env:"GOOGLE_CALENDAR_ID,required"`
	MaxResults      int64  `env:"GOOGLE_CALENDAR_MAX_RESULTS,required"`
	TimeZone        string `env:"GOOGLE_CALENDAR_TIMEZONE,required"`
	TargetUrl       string `env:"SCRAPER_TARGET_URL,required"`
	ScraperTimeout  int    `env:"SCRAPER_TIMEOUT,required"`
}
