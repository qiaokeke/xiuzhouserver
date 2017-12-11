package utils

import (
	"log"
	"bytes"
	"strconv"
	"fmt"
	"config"
	"entity"
	"daosql"
)



func ParseData(sbytes []byte)  {
	defer func() {
		if x := recover();x!=nil{
			log.Println("parseData err")
		}
	}()
	log.Printf("parseData: % X\n",sbytes)
	ParseZXYGDN(sbytes)

}


/**
	解析正向有功总电能bytes
 */
func ParseZXYGDN(sbytes []byte)  {
	if !bytes.Contains(sbytes,[]byte{0x1A,0x0C}){
		return
	}
	log.Println("ParseZXYGDN")

	index:=bytes.Index(sbytes,[]byte{1,1,16})
	log.Println(index)
	if index<0{
		log.Println("cant find 01 01 10")
		return
	}
	for{
		if index > len(sbytes)-35{
			break
		}

		numByte := sbytes[index-1]
		meterId := config.ReadConfig().MeterIds[strconv.Itoa(int(numByte))]
		log.Printf("num:% X\tmeterId:%s\n",numByte,meterId)

		entity.PowerInfoMap[meterId] = &entity.PowerInfo{}
		entity.PowerInfoMap[meterId].PowerMeterId = meterId


		timeBytes := sbytes[index+3:index+8]
		log.Printf("time:% X\n",timeBytes)
		zxygdnBytes := sbytes[index+9:index+14]
		zxygdn:=parse5bytes2float(zxygdnBytes)
		log.Printf("zxygdn:% X\t%f\n",zxygdnBytes,zxygdn)
		entity.PowerInfoMap[meterId].Zxygdn=zxygdn

		zxygdn1Bytes := sbytes[index+14:index+19]
		zxygdn1 := parse5bytes2float(zxygdn1Bytes)
		log.Printf("zxydn1:% X\t%f\n",zxygdn1Bytes,zxygdn1)
		entity.PowerInfoMap[meterId].Zxygdn1=zxygdn1

		zxygdn2Bytes := sbytes[index+19:index+24]
		zxygdn2 := parse5bytes2float(zxygdn2Bytes)
		log.Printf("zxydn2:% X\t%f\n",zxygdn2Bytes,zxygdn2)
		entity.PowerInfoMap[meterId].Zxygdn2=zxygdn2

		zxygdn3Bytes := sbytes[index+24:index+29]
		zxygdn3 := parse5bytes2float(zxygdn3Bytes)
		log.Printf("zxydn3:% X\t%f\n",zxygdn3Bytes,zxygdn3)
		entity.PowerInfoMap[meterId].Zxygdn3=zxygdn3

		zxygdn4Bytes := sbytes[index+29:index+34]
		zxygdn4 := parse5bytes2float(zxygdn4Bytes)
		log.Printf("zxydn4:% X\t%f\n",zxygdn4Bytes,zxygdn4)
		entity.PowerInfoMap[meterId].Zxygdn4=zxygdn4

		index+= 35
	}

	daosql.InsertAllInfos()

}
/**
	将5个字节的bcd bytes转换为float
 */
func parse5bytes2float(sbytes []byte)  float64{
	r:=0.0
	for j:=0;j<=4;j++ {
		hexs := strconv.FormatInt(int64((sbytes[j])&0xff),16)
		//fmt.Println(hexs)
		i, err := strconv.Atoi(hexs)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		k := 0.0
		if j==0{
			k=0.0001
		}
		if j==1{
			k=0.01
		}
		if j==2{
			k=1
		}
		if j==3{
			k=100
		}
		if j==4{
			k=10000
		}
		r += k * float64(i)
	}
	return r
}
