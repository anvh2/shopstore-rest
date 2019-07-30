package redis

import (
	"github.com/go-redis/redis"
)

func newPool() *redis.PoolStats {
	return nil
}

func initConn() {

}

//NewClient -
func NewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return client, nil
}
