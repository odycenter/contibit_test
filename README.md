# contibit_test

contibit golang test

目前為了保證大量交易時的資料正確性問題，採用這個架構，由 api 伺服器接收使用者出金/轉帳請求後 ，將該任務轉換成 task 形式，並且存進 queue，再由訂單處理伺服器"一次性"處理兩百筆訂單，並將處理結果寫入 sql server。
麻煩請用 golang 搭配任意 queue 概念(db/redis...等)以及任意 sql server，完成該架構的最小 poc 並有以下幾點要求：

# 1.訂單處理伺服器需"一次性"於記憶體內處理兩百筆訂單，並且考慮 crash 後不得造成計算結果有誤。

    1. 使用gin套件 建立api環境
    2. 將所有收到訂單 使用channel 推送至redis 若需要更保險 會需要一個非同步寫入mysql內 當作log 儲存
    3. 紀錄最後api時間戳(LAST_API_TIMESTAMP) 若超過10s 無新資料輸入 則處理目前資料
    4. 若golang crash 重新啟動 會先檢測 redis 內是否有資料未處理 若是有 則進入處理訂單流程
    5. 處理訂單流程 鎖定 row lock 檢查會員餘額 -> 檢查目標會員資格(未凍結) -> 修改餘額 -> 寫入帳本

# 2.需有可以給 1000 人以下使用者提供互相轉帳功能。
