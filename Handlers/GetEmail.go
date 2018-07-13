package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"EmailApi/Config"
)

func GetEmail(context *gin.Context){
	num,ok :=context.GetQuery("num")
	if !ok {
		context.JSON(http.StatusSeeOther,gin.H{
			"code":203,
			"message":"num is required",
		})
		return
	}
	number,_ := strconv.Atoi(num)
	fmt.Print(num)
	//链接redis
	c, err := redis.Dial("tcp", Config.REDIS_SERVER,redis.DialDatabase(Config.REDIS_DB))
	defer c.Close()
	if err != nil {
		log.Panic("connect redis server faild --- " + err.Error())
	}
	len, _ := redis.Int(c.Do("llen", "list:email:0"))
	if len<1{
		context.JSON(http.StatusSeeOther,gin.H{
			"code":203,
			"message":"合适的账号不存在",
		})
		return
	}
	if len < number{
		number =len
	}
	var email = make([]string,0,number)
	for a := 0; a < number; a++ {
		hh, _ := redis.String(c.Do("lpop", "list:email:0"))
		email = append(email,hh)
	}
	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"email": email,
	})

}
