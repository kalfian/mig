package queue

import (
	"accounting_migration/config"
	"accounting_migration/config/logger"

	"github.com/go-redis/redis"
)

func Enqueue(msg MessageQueue) error {
	logger.Info().Msg("Publishing event to stream " + config.RedisConfig.StreamName)

	err := config.Redis.XAdd(&redis.XAddArgs{
		Stream:       config.RedisConfig.StreamName,
		MaxLen:       int64(config.RedisConfig.StreanMaxLen),
		MaxLenApprox: int64(config.RedisConfig.StreanMaxLenApprox),
		ID:           "",
		Values: map[string]interface{}{
			"type":      msg.Type,
			"payload":   msg.Payload,
			"user_id":   msg.UserID,
			"cabang_id": msg.CabangID,
			"counter":   1,
		},
	}).Err()

	return err
}
