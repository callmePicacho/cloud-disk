package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       1,  // use default DB
})

func TestRedisSet(t *testing.T) {
	err := rdb.Set(context.Background(), "k", "v", time.Second*10).Err()
	if err != nil {
		t.Error(err)
	}
}

func TestRedisGet(t *testing.T) {
	val, err := rdb.Get(context.Background(), "k").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}
