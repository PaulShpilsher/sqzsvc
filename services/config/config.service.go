package config

import (
	"flag"
	"log"
	"os"
	"strconv"
)

var (
	Debug             bool   = false
	Port              int    = 5555
	DbDns             string = ""
	TokenHourLifespan int    = 1
	TokenSecret       string = ""
)

func InitConfig() {

	flag.BoolVar(&Debug, "debug", getEnvBool("DEBUG", Debug), "Debug mode")
	flag.IntVar(&Port, "port", getEnvInt("PORT", Port), "Listening port")
	flag.StringVar(&DbDns, "db-dns", getEnvString("DB_DNS", DbDns), "Database DNS string")
	flag.IntVar(&TokenHourLifespan, "token-hour-lifespan", getEnvInt("TOKEN_HOUR_LIFESPAN", TokenHourLifespan), "Token lifespan (hours)")
	flag.StringVar(&TokenSecret, "token-secret", getEnvString("TOKEN_SECRET", TokenSecret), "Token secret")

	flag.Parse()

	log.Println("config.Debug:", Debug)
	log.Println("config.Port:", Port)
	log.Println("config.DbDns:", DbDns)
	log.Println("config.TokenHourLifespan:", TokenHourLifespan)
	log.Println("config.TokenSecret:", TokenSecret)
}

func getEnvString(key string, defaultValue string) string {
	if envValue, ok := os.LookupEnv(key); ok {
		return envValue
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if envValue, ok := os.LookupEnv(key); ok {
		if value, err := strconv.Atoi(envValue); err != nil {
			return value
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if envValue, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseBool(envValue); err != nil {
			return value
		}
	}
	return defaultValue
}
