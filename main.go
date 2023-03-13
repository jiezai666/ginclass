package main

import (
	"ginclass/common"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	//https: //gorm.io/zh_CN/docs/generic_interface.html#%E8%BF%9E%E6%8E%A5%E6%B1%A0
	defer func() {
		sqlDB, _ := db.DB() // 获取通用数据库对象 sql.DB，然后使用其提供的功能
		err := sqlDB.Close()
		if err != nil {
			return
		}
	}()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run("127.0.0.1:8888")) //为空就默认为8080端口

}
