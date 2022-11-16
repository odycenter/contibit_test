package service

import (
	"contibit_test/config"
	"contibit_test/pkg"
	"database/sql"
	"encoding/json"
	"log"

	"contibit_test/tools"

	_ "github.com/go-sql-driver/mysql"
)

// api 收到的資料寫入 redis 保存
func TransferInfoToRedis(apiTimestamp int64, params string) {
	redisClient := config.RedisClient(config.REDIS_HOST, config.REDIS_PORT, config.REDIS_PWD, config.REDIS_DB)
	config.RWMUTEX.Lock()
	counter := config.COUNTER
	config.COUNTER++
	if config.COUNTER > 10000 {
		config.COUNTER = 0
	}
	config.RWMUTEX.Unlock()

	err := redisClient.HSet("TRANSFER", tools.Int64ToString(apiTimestamp)+"-"+tools.Int64ToString(counter), params).Err()
	if err != nil {
		log.Println(err)
	}
}

// api 收取資料滿 200 筆 or 100ms 無新資料寫入 開始計算帳變資料並寫入 redis用時間鎖 將資料鎖定
func StartTransfer() {
	redisClient := config.RedisClient(config.REDIS_HOST, config.REDIS_PORT, config.REDIS_PWD, config.REDIS_DB)
	// now := tools.Int64ToString(tools.TimeNow13())
	log.Println("開始計算帳變")
	var finish bool

	var infos map[string]string
	var infos_num int64 = 0
	infos, _ = redisClient.HGetAll("TRANSFER").Result()
	for redisKey, info := range infos {
		// log.Println(info)
		finish = TransferToMysql(redisKey, info)
		if finish {
			log.Println("完成轉帳", redisKey)
			redisClient.Del("TRANSFER", redisKey)
		}
		infos_num++
	}
	if infos_num == 0 {
		log.Println("redis無資料")
	}
}

// 處理轉帳訂單 寫入mysql
func TransferToMysql(redisKey string, info string) bool {
	var finish bool = false
	var data pkg.Transfer
	var checkMysqlTransferInfo bool = false
	err := json.Unmarshal([]byte(info), &data)
	if err != nil {
		log.Println(err.Error())
	}
	thisTime := tools.Timestamp10ToDatetime(tools.TimeNow10())
	db, dbErr := sql.Open("mysql", config.DB_CONNECT)

	var id int
	transferInfo := db.QueryRow("SELECT `id` FROM `cb_account_book` WHERE `api_kei`=?", data.ApiKey)
	switch err := transferInfo.Scan(&id); err {
	case sql.ErrNoRows:
		checkMysqlTransferInfo = true
	default:
		panic(err)
	}

	tools.CheckErr(dbErr)
	defer db.Close()

	if checkMysqlTransferInfo {
		tx, err := db.Begin()

		tools.CheckErr(err)
		// 撈取會員資料 鎖定表內row lock
		var user pkg.MysqlUser
		var user_num int64 = 0
		userInfo, err := db.Query("SELECT `id`, `balance_usdt` FROM `cb_user` WHERE (`id`=? or `id`=?) AND `user_status` = 0 FOR UPDATE", data.UserId, data.TargetId)
		var fromUser, toUser int64 = data.UserId, data.TargetId
		var fromUserBalance, toUserBalance float64 = 0.0, 0.0
		for userInfo.Next() {
			if err := userInfo.Scan(&user.Id, &user.BalanceUSDT); err != nil {
				log.Println(err)
			}
			// log.Println(user)
			if user.Id == fromUser {
				fromUserBalance = user.BalanceUSDT
			}
			if user.Id == toUser {
				toUserBalance = user.BalanceUSDT
			}
			user_num++
		}

		// 使用者帳戶未被凍結 且餘額足夠 開始轉帳
		if user_num == 2 && fromUserBalance > data.Amount {
			stmt0, err := tx.Prepare("UPDATE `cb_user` SET `balance_usdt` = ? WHERE `id` = ?")
			if err != nil {
				log.Println(err.Error())
			}
			stmt0.Exec(fromUserBalance-data.Amount, fromUser)
			stmt1, err := tx.Prepare("UPDATE `cb_user` SET `balance_usdt` = ? WHERE `id` = ?")
			if err != nil {
				log.Println(err.Error())
			}
			stmt1.Exec(toUserBalance+data.Amount, toUser)

			stmt2, err := tx.Prepare("INSERT INTO `cb_account_book` (`user_id`, `target_id`, `transfer_type`, `coin`, `amount`, `status`, `api_timestamp`, `api_datetime`, `api_kei`, `create_datetime`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			if err != nil {
				log.Println(err.Error())
			}
			var transferType int = 1
			if data.TransferType == "in" {
				transferType = 2
			}
			stmt2.Exec(data.UserId, data.TargetId, transferType, data.Coin, data.Amount, 1, data.ApiTimestamp, data.ApiDatetime, redisKey, thisTime)

			log.Println(data.UserId, data.TargetId, transferType, data.Coin, data.Amount, 1, tools.Int64ToString(data.ApiTimestamp), data.ApiDatetime, thisTime)

			finish = true
		} else {
			log.Println("使用者帳戶凍結")
		}

		defer func() {
			if err != nil {
				log.Println("錯誤回滾", err.Error())
				tx.Rollback()
				return
			}
		}()

		dbCommitErr := tx.Commit()
		tools.CheckErr(dbCommitErr)
	} else {
		// 資料重複 刪除redis數據
		finish = true
	}

	return finish
}
