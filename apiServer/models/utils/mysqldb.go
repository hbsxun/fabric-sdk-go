package utils

import (
	"sync"

	"github.com/op/go-logging"
	"github.com/widuu/gomysql"
)

var dbUtils *gomysql.Model
var once sync.Once

var dbLogger = logging.MustGetLogger("apiServer_mysql")

func GetDBInstance() *gomysql.Model {
	once.Do(func() {
		if dbUtils == nil {
			var err error
			dbUtils, err = gomysql.SetConfig("./config/db.ini")
			if err != nil {
				dbLogger.Panicf("db initialize, err [%s]/n", err)
			}
		}
	})
	return dbUtils
}
