package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Host              string
	Debug             bool
	Port              int
	Docker            bool
	DbDns             string
	TokenHourLifespan int
	TokenSecret       string
)

func InitConfig() {

	godotenv.Load()

	flag.BoolVar(&Debug, "debug", getEnvBool("DEBUG", false), "Debug mode")
	flag.StringVar(&Host, "host	", getEnvString("HOST", "localhost"), "Server hoat")
	flag.IntVar(&Port, "port", getEnvInt("PORT", 5555), "Listening port")
	flag.BoolVar(&Docker, "docker", getEnvBool("DOCKER", false), "Running in docker")
	flag.StringVar(&DbDns, "db-dns", getEnvString("DB_DNS", ""), "Database DNS string")
	flag.IntVar(&TokenHourLifespan, "token-hour-lifespan", getEnvInt("TOKEN_HOUR_LIFESPAN", 1), "Token lifespan (hours)")
	flag.StringVar(&TokenSecret, "token-secret", getEnvString("TOKEN_SECRET", ""), "Token secret")

	flag.Parse()
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
