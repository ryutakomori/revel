package redis

import (
	log "api/app/util/log"
	"strconv"
	"time"

	define "api/app/util/define"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/revel/revel"
	"github.com/vmihailenco/msgpack"
)

type Object struct {
	FunctionName string
	Count        int
}

func connection() *cache.Codec {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": "redis:6379",
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	return codec
}

func Limit(user_id int, fn string) {
	cnt := Get(user_id, fn)

	if cnt == -1 {
		cnt = 1
	} else {
		cnt = cnt + 1
	}

	codec := connection()

	obj := &Object{
		FunctionName: fn,
		Count:        cnt,
	}

	codec.Set(&cache.Item{
		Key:        strconv.Itoa(user_id) + fn,
		Object:     obj,
		Expiration: time.Second * define.LIMIT_TIME_SECOND,
	})
}

func LimitIp(req *revel.Request, fn string) {
	cnt := GetIp(req, fn)

	if cnt == -1 {
		cnt = 1
	} else {
		cnt = cnt + 1
	}

	codec := connection()

	obj := &Object{
		FunctionName: fn,
		Count:        cnt,
	}

	codec.Set(&cache.Item{
		Key:        req.RemoteAddr + fn,
		Object:     obj,
		Expiration: time.Second * define.LIMIT_TIME_SECOND,
	})
}

func Get(user_id int, fn string) int {
	codec := connection()

	var wanted Object
	if err := codec.Get(strconv.Itoa(user_id)+fn, &wanted); err != nil {
		log.Println("Redis Get Error")
		return -1
	}

	return wanted.Count
}

func GetIp(req *revel.Request, fn string) int {
	codec := connection()

	var wanted Object
	if err := codec.Get(req.RemoteAddr+fn, &wanted); err != nil {
		log.Println("Redis Get Error")
		return -1
	}

	return wanted.Count
}

func Check(user_id int, fn string) bool {

	codec := connection()

	var wanted Object
	if err := codec.Get(strconv.Itoa(user_id)+fn, &wanted); err != nil {
		log.Println("Redis Check Error")
		return false
	}
	log.Println("wanted.Count ", wanted.Count)

	if wanted.Count > define.LIMIT_RATE {
		return false
	}

	Limit(user_id, fn)

	return true
}
