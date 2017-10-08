package main

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.DialURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	keys, err := redis.Values(c.Do("KEYS", "*"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("KEYS: %d\n", len(keys))

	var skey string
	var slen = 1000000000000000
	for _, key := range keys {
		data, err := redis.String(c.Do("GET", fmt.Sprintf("%s", key)))
		if err != nil {
			panic(err)
		}

		if len(data) < slen {
			slen = len(data)
			skey = fmt.Sprintf("%s", key)
		}
		fmt.Print(".")
	}
	fmt.Printf("\n\nSmallest Key: %s Size: %d\n", skey, slen)
}
