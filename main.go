package main

import (
	"contibit_test/config"
	"contibit_test/server"
	"log"
	"sync"
)

func main() {
	log.Println(config.NAME, config.THREADS, config.VERSION)

	wg := &sync.WaitGroup{}

	// 啟動http service
	wg.Add(1)
	go server.GoGinService(wg)

	wg.Wait()
}
