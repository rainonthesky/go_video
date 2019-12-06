package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)
/**
定义
 */
func RegisterHandles()*httprouter.Router{
	router:=httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/user/:user_name",Login)
	return router
}
func main() {
	r:=RegisterHandles()
	fmt.Println("开始")
	//注册的函数
	http.ListenAndServe(":8091",r)


	
}
