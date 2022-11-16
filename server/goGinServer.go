package server

import (
	"log"
	"sync"

	"contibit_test/config"
	"contibit_test/controller"
	"contibit_test/service"
	"contibit_test/tools"

	"github.com/gin-gonic/gin"
)

func GoGinService(wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("")

	// 使用默认engine
	router := gin.Default()
	router.SetTrustedProxies(nil)
	// 配置路由

	// 會員
	go ChannelInfo()

	router.POST("/transfer", controller.Transfer) // 大量POST

	router.Run("localhost:8888")
}

func ChannelInfo() {
	var lastApiTimestamp int64 = 0
	// 取得 redis 內 最後一次api時間 檢查redis表內有無尚未完成帳變資料
	var transfer_num int = 0
	var now int64 = 0

	// 檢查redis 是否有殘存資料未處理 若有資料 優先處理舊單
	service.StartTransfer()

	for {
		select {
		case data := <-config.TRANSFER_CHAN:
			go service.TransferInfoToRedis(data.ApiTimestamp, data.TransferInfo)
			config.RWMUTEX.Lock()
			config.LAST_API_TIMESTAMP = data.ApiTimestamp
			config.RWMUTEX.Unlock()
			transfer_num++
			// 收取資料滿200筆 開始計算帳變
			if transfer_num >= 100 {
				service.StartTransfer()
				transfer_num = 0
			}
		default:
			now = tools.TimeNow13()
			config.RWMUTEX.RLock()
			lastApiTimestamp = config.LAST_API_TIMESTAMP
			config.RWMUTEX.RUnlock()
			if lastApiTimestamp > 0 && now-lastApiTimestamp > 5000 {
				config.RWMUTEX.Lock()
				config.LAST_API_TIMESTAMP = 0
				config.RWMUTEX.Unlock()
				service.StartTransfer()
			}
			continue
		}
	}
}
