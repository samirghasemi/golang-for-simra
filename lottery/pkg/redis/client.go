package redis

import (
	"context"
	// "fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	rdb *redis.Client
}

func NewClient() *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       1,
	})

	return &Client{rdb: rdb}
}

func (c *Client) IncrementUserCount(ctx context.Context, userID string) error {
	key := "lottery:" + userID
	err := c.rdb.Incr(ctx, key).Err()
	if err != nil {
		return err
	}

	c.rdb.Expire(ctx, key, 24*time.Hour) // reset the count after 24 hours
	return nil
}

func (c *Client) GetUserCount(ctx context.Context, userID string) (int, error) {
	key := "lottery:" + userID
	res, err := c.rdb.Get(ctx, key).Int()
	// fmt.Println("err: ", err)
	// fmt.Println("res: ", res)
	if err != nil && err.Error() == "redis: nil" {
		err = c.rdb.Set(ctx, key, 0, 0).Err()
		return res, nil
	}
	return res, err
}
