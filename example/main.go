package main

import (
	"github.com/gin-gonic/gin"
	xxl_job_executor_gin "github.com/gin-middleware/xxl-job-executor"
	"github.com/konglong87/xxl-job-executor-go"
	"github.com/konglong87/xxl-job-executor-go/example/task"
	"log"
)

const Port = "9999"
const 	test_xxx_host = "https://job.testing4.hetao101.com/xxl-job-admin"


func main() {
	//初始化执行器
	exec := xxl.NewExecutor(
		xxl.ServerAddr("http://101.200.72.228:8080/xxl-job-admin"),
		xxl.AccessToken(""),            //请求令牌(默认为空)
		//xxl.ExecutorIp("127.0.0.1"),    //可自动获取
		//xxl.ExecutorPort("8888"),         //默认9999（此处要与gin服务启动port必需一至）
		xxl.RegistryKey("oaakl"), //执行器名称
	)
	exec.Init()



	//添加到gin路由
	r := gin.New()
	xxl_job_executor_gin.XxlJobMux(r, exec)


	//注册gin的handler
	//r.GET("ping", func(cxt *gin.Context) {
	//	cxt.JSON(200, "pong11112222222")
	//})
	//r.GET("p2", func(cxt *gin.Context) {
	//	cxt.JSON(200, "test-panic-2")
	//})
	//
	//r.GET("p3", func(cxt *gin.Context) {
	//	panic("see what p3")
	//	cxt.JSON(200, "test-panic-3")
	//})
	//
	//r.GET("p4", gin.Recovery(),func(cxt *gin.Context) {
	//	panic("see what p4")
	//	cxt.JSON(200, "test-panic-4")
	//})
	//
	//r("p5", func(cxt *gin.Context) {
	//	panic("see what p5, later recoverry")
	//	cxt.JSON(200, "test-panic-5")
	//},gin.Recovery())

	//注册任务handler
	exec.RegTask("task.test", task.Test)
	exec.RegTask("task.test2", task.Test2)
	exec.RegTask("task.panic", task.Panic)


	//go exec.Run()

	log.Fatal(r.Run(":" + Port))
}
