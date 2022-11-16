package controller

import (
	"contibit_test/config"
	"contibit_test/pkg"
	"contibit_test/tools"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Transfer(c *gin.Context) {
	now := tools.TimeNow13()

	json := make(map[string]string)
	c.ShouldBind(&json)

	info := fmt.Sprintf(`{"user_id": %s, "target_id": %s, "type": "%s", "coin": "%s", "amount": %s, "api_timestamp": %s, "api_datetime": "%s"}`, json["user_id"], json["target_id"], json["type"], json["coin"], json["amount"], tools.Int64ToString(now), tools.Timestamp13ToDatetime(now))
	log.Println("+++000+++", json, info)

	// ToDo 檢查收到的資料格式是否正確
	check := true

	// 如果資料正確 直接回傳成功 資料格式錯誤 回傳失敗
	if check {
		config.TRANSFER_CHAN <- pkg.TransferChannel{
			ApiTimestamp: now,
			TransferInfo: info,
		}
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "timestamp": now})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "false data", "timestamp": now})
	}

}
