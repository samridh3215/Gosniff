package main

import (
	// "fmt"

	"log"
	"net/http"
	"os"
)

var infoLog *log.Logger
var panicLog *log.Logger

func initLogs() {
	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	infoLog = log.New(logFile, "INFO: ", log.LstdFlags)
	panicLog = log.New(logFile, "ERROR: ", log.LstdFlags)
}

func main() {
	initLogs()

	go func() {

		http.HandleFunc("/ws", HandleConnections)
		infoLog.Println("http server started on :8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panicLog.Fatalln("ListenAndServe: ", err)
		}
	}()

	go Sniff()

	go func() {
		cnt := 0
		for packet := range packet_channel {
			// fmt.Println("Packet received:", packet)
			packet_json := ParsePacket(packet)

			BroadcastToConnections(packet_json)
			cnt++
		}
	}()

	select {}
}
