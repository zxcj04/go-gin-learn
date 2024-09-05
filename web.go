package main

import (
	"github.com/gin-gonic/gin"

	"github.com/zxcj04/go-gin-learn/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/deposit/:input", handlers.Deposit)
	router.GET("/withdraw/:input", handlers.Withdraw)
	router.GET("/balance/", handlers.GetBalance)

	// 多筆存款
	router.GET("/multiDeposit/:input", handlers.MultiDeposit)

	router.Run()
}
