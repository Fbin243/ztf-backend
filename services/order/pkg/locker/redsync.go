package locker

import (
	"context"
	"sync"
	"time"

	"github.com/go-redsync/redsync/v4"

	"github.com/go-redsync/redsync/v4/redis"
	redsync_goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredis "github.com/redis/go-redis/v9"
)

var (
	redsyncInstance *redsync.Redsync
	onceRedsyn      sync.Once
)

func InitRedsync() {
	onceRedsyn.Do(func() {
		redisAddresses := []string{
			"localhost:30001",
			"localhost:30002",
			"localhost:30003",
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		pools := []redis.Pool{}
		for _, addr := range redisAddresses {
			client := goredis.NewClient(&goredis.Options{
				Addr:     addr,
				Password: "", // no password set
				DB:       0,  // use default DB
			})
			if client.Ping(ctx).Err() != nil {
				panic("Failed to connect to Redis at " + addr)
			}
			pools = append(pools, redsync_goredis.NewPool(client))
		}

		redsyncInstance = redsync.New(pools...)
	})
}

func GetRedsync() *redsync.Redsync {
	if redsyncInstance == nil {
		InitRedsync()
	}
	return redsyncInstance
}
