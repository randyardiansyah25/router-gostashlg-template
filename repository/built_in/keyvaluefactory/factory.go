package keyvaluefactory

import (
	"github.com/randyardiansyah25/libpkg/util/env"
)

func GetStore() (Store, error) {
	driverName := env.GetString("app.keyvalue_driver")
	if driverName == DRIVER_REDIS {
		return newRedisImpl(), nil
	} else {
		return newStoreMockImpl(), nil
	}
}
