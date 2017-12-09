package dao

import (
	"net"
	"fmt"
	"time"
	"config"
	"log"
)


/**
	处理读出消息
 */
func HandleRead(conn net.Conn)  {
	defer conn.Close()
	for{
		recvBytes := make([]byte,1024)
		
	}
}

/**
	处理写入消息
 */
func HandleWrite(conn net.Conn)  {
	defer conn.Close()
	for{
		meterOrders := config.ReadConfig().MeterIds
		for i:=0;i<len(meterOrders);i++{
			for j:=0;j<len(meterOrders[i]);j++ {
				log.Printf("write:% X\n",meterOrders[i][j])
				_, e := conn.Write(meterOrders[i][j])
				if e != nil {
					log.Println("write err:", e)
					return;
				}
				time.Sleep(60*time.Second)
			}
		}
		time.Sleep(5*time.Second)
	}
}