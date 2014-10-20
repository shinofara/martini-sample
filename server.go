package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/garyburd/redigo/redis"
	"github.com/codegangsta/martini-contrib/render"
)

var (
	redisAddress   = flag.String("redis-address", ":6379", "Address to the Redis server")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
)

func main() {
	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", *redisAddress)
		
		if err != nil {
			return nil, err
		}
		
		return c, err
	}, *maxConnections)
	
	defer redisPool.Close()
	
	m := martini.Classic()
	
	m.Map(redisPool)
	
	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout: "layout",
		Extensions: []string{".html", ".tmpl"},
		Charset: "utf-8",
	}))
	
	//redisから値をとってみるサンプル
	m.Get("/get/:key", getKey)
	
	m.Get("/", IndexRender)
	m.Get("/about", AboutRender)
	m.Run()
}

func getKey(r render.Render, pool *redis.Pool, params martini.Params) string {
	key := params["key"]
	
	c := pool.Get()
	defer c.Close()
	
	//アクセスカウントを+1
	//nilの場合はredis側で1で作成
	c.Do("INCR", key)
	
	//nilにはなることは無いが
	value, error := redis.String(c.Do("GET", key))
	
	if error != nil {
		return "0"
	}
	
	return "アクセス回数は、" + value + "回です。"
}
