package dao

import (
	"net"
	"time"
	"config"
	"log"
	"utils"
)


/**
	处理读出消息
 */
func HandleRead(conn net.Conn)  {
	defer conn.Close()
	for{
		recvBytes := make([]byte,1024)
		n,err :=conn.Read(recvBytes)
		if err!=nil{
			log.Println("read err:",err)
			break
		}
		if n<=0{
			continue
		}
		go utils.ParseData(recvBytes[:n])
	}
}

/**
	处理写入消息
 */
func HandleWrite(conn net.Conn)  {
	defer conn.Close()
	for{
		cmds := config.ReadConfig().Cmds
		for i:=0;i<len(cmds);i++{
			for j:=0;j<len(cmds[i]);j++ {
				log.Printf("write:% X\n",cmds[i][j])
				_, e := conn.Write(cmds[i][j])
				if e != nil {
					log.Println("write err:", e)
					return;
				}
				time.Sleep(120*time.Second)
			}
		}
		time.Sleep(5*time.Second)
	}
}