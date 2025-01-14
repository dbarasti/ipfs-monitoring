package main

import (
	"ipfs_monitoring/bwmonitor"
	"ipfs_monitoring/swarmmonitor"
	"log"
	"os"
	"sync"
)

func main() {
	logFile := setupLog()
	defer logFile.Close()
	log.Print("[MAIN] Starting execution")

	var wg sync.WaitGroup
	wg.Add(2)

	go bwmonitor.RunMonitor(&wg)
	go swarmmonitor.RunMonitor(&wg)

	wg.Wait()

}

func setupLog() *os.File {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return file
}
