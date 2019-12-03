package redis

import (
    "log"

    "github.com/go-redis/redis"
)

type Repository struct {
    c *redis.Client
}

func GetRepository(conf Config) Repository {
    return Repository{redis.NewClient(&redis.Options{
        Addr:      conf.Host,
        Password:  conf.Password,
        TLSConfig: conf.TlsConfig,
        DB:        conf.DbIndex,
    })}
}

func (r Repository) IsKeyExists(key string) (bool, error) {
    result := r.c.Exists(key)

    return result.Val() == 1, result.Err()
}

func (r Repository) SetValue(key, value string) error {
    return r.c.Set(key, value, 0).Err()
}

func (r Repository) GetValue(key string) (string, error) {
    return r.c.Get(key).Result()
}

func (r Repository) Close() error {
    return r.c.Close()
}

func (r Repository) ping() error {
    if _, err := r.c.Ping().Result(); err != nil {
        log.Fatal("Error with ping to redis: ", err)
        return err
    }

    return nil
}
