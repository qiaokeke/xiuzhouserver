package main

import (

	"dao"
	"config"
	"log"
	"net"
	"os"
	"entity"
)

func main()  {
	/**
	全局map,保存读出的数据
	 */
	entity.PowerInfoMap = make(map[string] *entity.PowerInfo)

	config.ConfigPath = "xiuzhou6030.json"
	port := config.ReadConfig().Port

	listener,err := net.Listen("tcp",""+":"+port)
	defer listener.Close()
	if err!=nil{
		log.Println("listen err:",err)
		os.Exit(1)
	}
	log.Println("listening on:",port)
	for{
		conn,err := listener.Accept()
		if err !=nil{
			log.Println("accept err:",err)
			break
		}
		log.Println("connect from :",conn.RemoteAddr(),conn.LocalAddr())
		go dao.HandleRead(conn)
		go dao.HandleWrite(conn)
	}


}