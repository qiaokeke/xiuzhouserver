package utils

import "log"

func checkLength(sbytes []byte)  bool{
	if  sbytes[0]!=0x68{
		return false
	}
	if sbytes[5]!=0x68{
		return false
	}
	if sbytes[len(sbytes)-1]!=0x16{
		return false
	}
	datalen:=(int(sbytes[1]) + int(sbytes[2])*16)/4
	log.Println("data len:",datalen)

	if(len(sbytes)<=datalen){
		return false
	}
	return true
}

func CheckData(sbytes []byte)  bool {
	if (!checkLength(sbytes)){
		return false
	}
	crcint := 0

	for i:=6;i<len(sbytes)-2;i++  {
		crcint += int(sbytes[i])
	}
	crcint %= 256
	if crcint!=int(sbytes[len(sbytes)-2]) {
		return false
	}
	return true
}
