package config

import (
	"io/ioutil"
	"encoding/json"
	"runtime"
	"strings"
	"log"
)

var ConfigPath string

type Config struct {
	Port string
	Cmds [][][]byte
	MeterIds map[string] string
	DataSource string
	DataSource2 string
}

/**
	读取配置文件
 */
func ReadConfig()  Config{
	/**
	判断操作系统类型
	 */
	os := runtime.GOOS
	path :=""
	if strings.Compare(os,"windows")==0{
		path = "C:\\work\\go\\xiuzhouserver\\src\\config\\"+ConfigPath
	}else {
		path = "/etc/xiuzhou/config/"+ConfigPath
	}

	file,err:=ioutil.ReadFile(path)
	if err!=nil{
		log.Println("config file err:",err)
	}
	config := Config{}
	err2 := json.Unmarshal(file,&config)
	if err2!=nil{
		log.Println("json err:",err2)
	}
	return config
}


