package main

import (
	"github.com/KovalMax/urlshortener/redis"
	"github.com/KovalMax/urlshortener/routes"
)

func main() {
	redis.CheckConnection()
	routes.Handle()
}
