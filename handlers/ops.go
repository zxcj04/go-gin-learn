package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/zxcj04/go-gin-learn/libs"
)

type Result struct {
	Amount  int    `json:"amount"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

var result = Result{}

// GetBalance 取得帳戶內餘額
func GetBalance(context *gin.Context) {
	result.Amount = libs.GetBalance() // 返回的Amount為當前餘額
	result.Status = "ok"              // 查詢時，可將預設狀態設為成功
	result.Message = ""               // 成功時Message提示為空
	context.JSON(http.StatusOK, result)
}

// Deposit 儲值、存款
func Deposit(context *gin.Context) {
	input := context.Param("input")
	amount, err := strconv.Atoi(input)

	result.Status = "failed" // 存款操作時，可將預設狀態設為失敗
	result.Message = ""

	if err == nil {
		if amount <= 0 {
			result.Amount = 0 // 操作未成功，返回金額為0
			result.Message = "操作失敗，存款金額需大於0元！"
		} else {
			balance := libs.GetBalance()
			balance = libs.SetBalance(balance + amount)
			result.Amount = balance // 操作成功，返回的Amount為儲值後的餘額
			result.Status = "ok"    // 操作成功
		}
	} else {
		result.Amount = 0 // 操作未成功，返回金額為0
		result.Message = "操作失敗，輸入有誤！"
	}
	context.JSON(http.StatusOK, result)
}

func addToBalance(amount int, c chan int) {
	balance := libs.GetBalance()
	c <- libs.SetBalance(balance + amount)
}

func MultiDeposit(context *gin.Context) {
	input := context.Param("input")
	// input = "100,200,300"
	inputs := strings.Split(input, ",")
	amounts := make([]int, len(inputs))
	for i, v := range inputs {
		amounts[i], _ = strconv.Atoi(v)
	}

	c := make(chan int)
	for _, amount := range amounts {
		go addToBalance(amount, c)
	}

	balance := 0

	for range amounts {
		balance = <-c
	}

	result.Amount = balance
	result.Status = "ok"
	result.Message = ""
	context.JSON(http.StatusOK, result)
}

// Withdraw 提款
func Withdraw(context *gin.Context) {
	result.Status = "failed" // 提款操作時，可將預設狀態設為失敗
	result.Message = ""

	input := context.Param("input")
	amount, err := strconv.Atoi(input)

	if err == nil {
		if amount <= 0 {
			result.Amount = 0 // 操作未成功，返回金額為0
			result.Message = "操作失敗，提款金額需大於0元！"
		} else {
			balance := libs.GetBalance()
			if balance-amount < 0 {
				result.Amount = 0 // 操作未成功，返回金額為0
				result.Message = "操作失敗，餘額不足！"
			} else {
				balance = libs.SetBalance(balance - amount)
				result.Amount = balance // 操作成功，返回的Amount為提款後的餘額
				result.Status = "ok"
			}
		}
	} else {
		result.Amount = 0 // 操作未成功，返回金額為0
		result.Message = "操作失敗，輸入有誤！"
	}
	context.JSON(http.StatusOK, result)
}
