package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Conn interface {
	connection()
}

type local struct {
	Redis *redis.Client
}

func (l local) connection() {
	l.Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

type Prod struct {
	Redis *redis.ClusterClient
}

func (p Prod) connection() {
	p.Redis = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"host1"},
		Password: "", // no password set
	})
}

func conn(arg Conn) {
	arg.connection()
}

var ctx = context.Background()

func test2() {

	conn(Prod{})

	// set := local{}
	// set.local = redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// set.local.sett.HSet(ctx, "1", 1)
}
