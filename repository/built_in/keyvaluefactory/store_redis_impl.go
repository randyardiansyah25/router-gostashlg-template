package keyvaluefactory

import (
	"context"
	"router-gostashlg-template/entities/app"

	"time"

	"github.com/randyardiansyah25/libpkg/security/aes"
	"github.com/randyardiansyah25/libpkg/util/env"
	"github.com/redis/go-redis/v9"
)

func newRedisImpl() Store {
	return &redisImpl{}
}

type redisImpl struct {
	prefix string
	conn   *redis.Client
}

func (r *redisImpl) Open() (er error) {
	host := env.GetString(r.prefix + "redis.host")
	pass := env.GetString(r.prefix + "redis.passwword")
	if pass != "" {
		pvKey := []byte(app.PrivateKey)
		pass, er = aes.Decrypt(pvKey, pvKey, pass)
		if er != nil {
			return er
		}
	}
	dbIndex := env.GetInt(r.prefix+"redis.db_index", 0)
	poolSize := env.GetInt(r.prefix+"redis.max_poolsize", 100)
	maxIdle := env.GetInt(r.prefix+"redis.max_idle", 30)
	minIdle := env.GetInt(r.prefix+"redis.min_idle", 4)
	connMaxLifetime := env.GetInt(r.prefix+"max_lifetime", 12)
	r.conn = redis.NewClient(&redis.Options{
		Addr:            host,
		Password:        pass,
		DB:              dbIndex,
		PoolSize:        poolSize,
		MaxIdleConns:    maxIdle,
		MinIdleConns:    minIdle,
		ConnMaxLifetime: time.Duration(connMaxLifetime) * time.Minute,
	})

	return nil
}

func (r *redisImpl) Echo() error {
	ctx := context.Background()
	return r.conn.Ping(ctx).Err()
}

func (r *redisImpl) GetStore() interface{} {
	return r.conn
}

func (r *redisImpl) GetDriverName() string {
	return DRIVER_REDIS
}

func (r *redisImpl) Close() {
	r.conn.Close()
}
