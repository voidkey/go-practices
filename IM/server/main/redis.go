package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func initPool(addr string, poolsize int, idleTimeout time.Duration) {
	opt := &redis.Option{
		Addr:        addr,
		Password:    "",       // no password set
		DB:          0,        // use default DB
		PoolSize:    poolsize, // 连接池大小
		IdleTimeout: idleTimeout,
	}
	ConnPool = pool.NewConnPool(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// _, err = rdb.Ping(ctx).Result()
	// if err!=nil {
	// 	return err
	// }

	ConnPool.NewConn(ctx)

}

func initClient(addr string, poolSize int, idleTimeout time.Duration) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    "",       // no password set
		DB:          0,        // use default DB
		PoolSize:    poolSize, // 连接池大小
		IdleTimeout: idleTimeout,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}
