package redis

import (
    "crypto/tls"
    "fmt"
    "log"
    "os"
    "strconv"
)

type Config struct {
    Host, Password string
    DbIndex        int
    TlsConfig      *tls.Config
}

func GetConfig() Config {
    host := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
    password := os.Getenv("REDIS_PASSWORD")
    dbIndex, err := strconv.Atoi(os.Getenv("REDIS_DB_INDEX"))
    if err != nil {
        log.Fatal("Can not get db index env", err)
    }

    return Config{host, password, dbIndex, nil}
}
