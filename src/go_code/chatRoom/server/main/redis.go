package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// 定义一个全局的pool
var pool *redis.Pool

func initPool(maxIdle, maxActive int, address string, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大空闲连接数
		MaxActive:   maxActive,   // 表示和数据库的最大连接数
		IdleTimeout: idleTimeout, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化连接的代码,连接哪个IP的redis
			return redis.Dial("tcp", address)
		},
	}
}
