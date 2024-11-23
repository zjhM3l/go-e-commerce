package dal

import (
	"github.com/zjhM3l/go-e-commerce/backend/app/frontend/biz/dal/mysql"
	"github.com/zjhM3l/go-e-commerce/backend/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
