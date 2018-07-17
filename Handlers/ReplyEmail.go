package Handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
	"EmailApi/Config"
	"net/http"
	"strconv"
	"fmt"
)
type Log struct {

	Json  []map[string]string `form:"json" json:"json" binding:"required"`
}
func ReplyEmail(context *gin.Context){
	c, err := redis.Dial("tcp", Config.REDIS_SERVER,redis.DialDatabase(Config.REDIS_DB))
	defer c.Close()
	if err != nil {
		panic("connect redis server faild -- " + err.Error())
	}
	// bind JSON数据
	var json Log
	if context.BindJSON(&json) == nil {
		hhy :=json.Json
		for _, value := range hhy{
			for k,v :=range value{
				//可能需要改为value['k']
				b,error := strconv.Atoi(v)
				if error != nil{
					fmt.Println("字符串转换成整数失败")
				}
				//k为收信人  b为状态
				//n为邮箱 v为状态
				//(测试为0 是否可以)
				//全部状态(个数)
				c.Do("INCRBY", "count",1 )
				c.Do("INCRBY", "num:" +v,1)
				c.Do("INCRBY", "today:" +v,1)
				c.Do("HSET", "email:"+k, "status", b)
			}
		}
		context.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"回执成功",
		})

	} else {
		context.JSON(http.StatusSeeOther,gin.H{
			"code":203,
			"message":"no",
		})
		return
		context.JSON(404, gin.H{"JSON=== status": "binding JSON error!"})
	}
}
