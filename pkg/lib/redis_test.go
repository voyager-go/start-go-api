package lib

import (
	"context"
	"testing"
)

func TestNewRedis(t *testing.T) {
	args := RdsConfig{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DbNum:    0,
	}
	t.Run("测试", func(t *testing.T) {
		client := NewRedis(args)
		testKey := "test-redis-1"
		testValue := "redis test result value"
		err := client.Set(context.Background(), testKey, testValue, 0).Err()
		if err != nil {
			t.Errorf("Redis set key %s failure, err: %v", testKey, err)
		}
		value, err := client.Get(context.Background(), testKey).Result()
		if err != nil {
			t.Errorf("Redis get key %s failure, err: %v", testKey, err)
		}
		if value != testValue {
			t.Errorf("Redis get value error, value expect %s, but: %s", testValue, value)
		}
	})
}
