package config

import (
	"accounting_migration/config/logger"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	Redis          *redis.Client
	RedisConfig    RedisOption
	redisVersionRE = regexp.MustCompile(`redis_version:(.+)`)
)

type RedisOption struct {
	Host string
	Port string
	Auth string
	DB   int

	StreamName         string
	StreanMaxLen       int
	StreanMaxLenApprox int
}

func LoadRedisConfig() RedisOption {
	return RedisOption{
		Host: viper.GetString("REDIS_HOST"),
		Port: viper.GetString("REDIS_PORT"),
		Auth: viper.GetString("REDIS_AUTH"),
		DB:   viper.GetInt("REDIS_DB"),

		StreamName:         viper.GetString("STREAM_NAME"),
		StreanMaxLen:       viper.GetInt("STREAM_MAX_LEN"),
		StreanMaxLenApprox: viper.GetInt("STREAM_MAX_LEN_APPROX"),
	}
}

func ConnectRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", RedisConfig.Host, RedisConfig.Port),
		Password: RedisConfig.Auth,
		DB:       RedisConfig.DB,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		logger.Fatal().Msg(err.Error())
		os.Exit(1)
	}

	if err := redisPreflightChecks(redisClient); err != nil {
		logger.Fatal().Msg(err.Error())
		os.Exit(1)
	}

	logger.Info().Msg("Redis Started!")
	return redisClient
}

func RedisPing() error {
	_, err := Redis.Ping().Result()
	if err != nil {
		Redis.Close()
		ConnectRedis()
	}
	return nil
}

func OpenRedis() {
	RedisConfig = LoadRedisConfig()
	Redis = ConnectRedis()
}

// offers the functionality we need. Specifically, it also that it can connect
// to the actual instance and that the instance supports Redis streams (i.e.
// it's at least v5).
func redisPreflightChecks(client redis.UniversalClient) error {
	info, err := client.Info("server").Result()
	if err != nil {
		return err
	}

	match := redisVersionRE.FindAllStringSubmatch(info, -1)
	if len(match) < 1 {
		return fmt.Errorf("could not extract redis version")
	}
	version := strings.TrimSpace(match[0][1])
	parts := strings.Split(version, ".")
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if major < 5 {
		return fmt.Errorf("redis streams are not supported in version %q", version)
	}

	return nil
}
