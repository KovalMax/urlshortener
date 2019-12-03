package redis

import (
    "fmt"
    "log"
)

func CheckConnection() {
    if err := GetRepository(GetConfig()).ping(); err != nil {
        log.Fatal("Redis unavailable: ", err)
    }

    fmt.Println("Connected to redis")
}
