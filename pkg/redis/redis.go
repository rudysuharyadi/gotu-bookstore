package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var Instance Redis

type Redis struct {
	pool *redis.Pool
}

func NewRedis(config RedisConfig) Redis {
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
			if err != nil {
				return nil, err
			}

			if config.Password != "" {
				_, err := c.Do("AUTH", config.Password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	Instance = Redis{redisPool}
	return Instance
}

// Set a key/value
func (n Redis) Set(key string, data interface{}) error {
	conn := n.pool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// SetWithExpiry a key/value
func (n Redis) SetWithExpiry(key string, data interface{}, time int) error {
	conn := n.pool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// SetnxWithExpiry a key/value if not exists
func (n Redis) SetnxWithExpiry(key string, data interface{}, time int) error {
	conn := n.pool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SETNX", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func (n Redis) Exists(keys ...string) bool {
	conn := n.pool.Get()
	defer conn.Close()

	var args = []interface{}{}
	for f, v := range keys {
		args = append(args, f, v)
	}
	exists, err := redis.Bool(conn.Do("EXISTS", args...))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func (n Redis) Get(key string, data interface{}) error {
	conn := n.pool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}

	err = json.Unmarshal(reply, &data)
	if err != nil {
		return err
	}

	return nil
}

// Delete delete a key
func (n Redis) Delete(key string) (bool, error) {
	conn := n.pool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// Incr increment a key value
func (n Redis) Incr(key string) (int64, error) {
	conn := n.pool.Get()
	defer conn.Close()

	reply, err := redis.Int64(conn.Do("INCR", key))
	if err != nil {
		return 0, err
	}

	return reply, nil

}
