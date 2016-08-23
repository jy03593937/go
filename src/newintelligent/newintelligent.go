package main

import (
	"fmt"
	"crc16"
	"io/ioutil"
	"net"
//	"bytes"
	"time"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"github.com/astaxie/beego"
)
const RECV_BUF_LEN = 1024

func main() {
	
	localAddress,_ := net.ResolveTCPAddr("tcp4", "0.0.0.0:8383")
	
	tcpListener, err := net.ListenTCP("tcp", localAddress)      
	CheckError(err)
/*
	if err != nil {
		fmt.Println("error:", err)
		return
	}
*/
	defer func() { //
		tcpListener.Close()
	}()
	
	fmt.Println("tcpaddr wait...")
	for {
		var conn, err2 = tcpListener.AcceptTCP() 
		if err2 != nil {
			fmt.Println("error:", err2)
			continue
		}
		go handleClient(conn)
		  	

		
	
	}
}
func CheckError(err error) {  
    if err != nil {  
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())  
        os.Exit(1)  
    }
}	

func handleClient(conn net.Conn) {
    defer conn.Close()
    for{
		var remoteAddr = conn.RemoteAddr() 
		fmt.Println("addr:", remoteAddr)
		go handleConnection(conn)
		
		for {
		
			_, err := conn.Write([]byte("\x01\x04\x00\x00\x00\x0A\x70\x0D"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x02\x04\x00\x00\x00\x0A\x70\x3E"))
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x03\x04\x00\x00\x00\x0A\x71\xEF"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x04\x04\x00\x00\x00\x0A\x70\x58"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x05\x04\x00\x00\x00\x0A\x71\x89"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x06\x04\x00\x00\x00\x0A\x71\xBA"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x07\x04\x00\x00\x00\x0A\x70\x6B"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x08\x04\x00\x00\x00\x0A\x70\x94"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x09\x04\x00\x00\x00\x0A\x71\x45"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x0A\x04\x00\x00\x00\x0A\x71\x76"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x0B\x04\x00\x00\x00\x0A\x70\xA7"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x0C\x04\x00\x00\x00\x0A\x71\x10"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x0D\x04\x00\x00\x00\x0A\x70\xC1"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x0E\x04\x00\x00\x00\x0A\x70\xF2"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x0F\x04\x00\x00\x00\x0A\x71\x23"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x10\x04\x00\x00\x00\x0A\x73\x4C"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x11\x04\x00\x00\x00\x0A\x72\x9D"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x12\x04\x00\x00\x00\x0A\x72\xAE"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x13\x04\x00\x00\x00\x0A\x73\x7F"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x14\x04\x00\x00\x00\x0A\x72\xC8"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x15\x04\x00\x00\x00\x0A\x73\x19"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x16\x04\x00\x00\x00\x0A\x73\x2A"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x17\x04\x00\x00\x00\x0A\x72\xFB"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x18\x04\x00\x00\x00\x0A\x72\x04"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x19\x04\x00\x00\x00\x0A\x73\xD5"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x1A\x04\x00\x00\x00\x0A\x73\xE6"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x1B\x04\x00\x00\x00\x0A\x72\x37"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x1C\x04\x00\x00\x00\x0A\x73\x80"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x1D\x04\x00\x00\x00\x0A\x72\x51"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x1E\x04\x00\x00\x00\x0A\x72\x62"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x1F\x04\x00\x00\x00\x0A\x73\xB3"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x20\x04\x00\x00\x00\x0A\x76\xBC"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x21\x04\x00\x00\x00\x0A\x77\x6D"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x22\x04\x00\x00\x00\x0A\x77\x5E"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x23\x04\x00\x00\x00\x0A\x76\x8F"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x24\x04\x00\x00\x00\x0A\x77\x38"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x25\x04\x00\x00\x00\x0A\x76\xE9"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x26\x04\x00\x00\x00\x0A\x76\xDA"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x27\x04\x00\x00\x00\x0A\x77\x0B"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x28\x04\x00\x00\x00\x0A\x77\xF4"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x29\x04\x00\x00\x00\x0A\x76\x25"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x2A\x04\x00\x00\x00\x0A\x76\x16"))				
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x2B\x04\x00\x00\x00\x0A\x77\xC7"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x2C\x04\x00\x00\x00\x0A\x76\x70"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x2D\x04\x00\x00\x00\x0A\x77\xA1"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x2E\x04\x00\x00\x00\x0A\x77\x92"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x2F\x04\x00\x00\x00\x0A\x76\x43"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x30\x04\x00\x00\x00\x0A\x74\x2C"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x31\x04\x00\x00\x00\x0A\x75\xFD"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x32\x04\x00\x00\x00\x0A\x75\xCE"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x33\x04\x00\x00\x00\x0A\x74\x1F"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x34\x04\x00\x00\x00\x0A\x75\xA8"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x35\x04\x00\x00\x00\x0A\x74\x79"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x36\x04\x00\x00\x00\x0A\x74\x4A"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x37\x04\x00\x00\x00\x0A\x75\x9B"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x38\x04\x00\x00\x00\x0A\x75\x64"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x39\x04\x00\x00\x00\x0A\x74\xB5"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x3A\x04\x00\x00\x00\x0A\x74\x86"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x3B\x04\x00\x00\x00\x0A\x75\x57"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x3C\x04\x00\x00\x00\x0A\x74\xE0"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x3D\x04\x00\x00\x00\x0A\x75\x31"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x3E\x04\x00\x00\x00\x0A\x75\x02"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x3F\x04\x00\x00\x00\x0A\x74\xD3"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x40\x04\x00\x00\x00\x0A\x7F\x1C"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x41\x04\x00\x00\x00\x0A\x7E\xCD"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x42\x04\x00\x00\x00\x0A\x7E\xFE"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x43\x04\x00\x00\x00\x0A\x7F\x2F"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x44\x04\x00\x00\x00\x0A\x7E\x98"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x45\x04\x00\x00\x00\x0A\x7F\x49"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x46\x04\x00\x00\x00\x0A\x7F\x7A"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x47\x04\x00\x00\x00\x0A\x7E\xAB"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x48\x04\x00\x00\x00\x0A\x7E\x54"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x49\x04\x00\x00\x00\x0A\x7F\x85"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x4A\x04\x00\x00\x00\x0A\x7F\xB6"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x4B\x04\x00\x00\x00\x0A\x7E\x67"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x4C\x04\x00\x00\x00\x0A\x7F\xD0"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x4D\x04\x00\x00\x00\x0A\x7E\x01"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x4E\x04\x00\x00\x00\x0A\x7E\x32"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x4F\x04\x00\x00\x00\x0A\x7F\xE3"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x50\x04\x00\x00\x00\x0A\x7D\x8C"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x51\x04\x00\x00\x00\x0A\x7C\x5D"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x52\x04\x00\x00\x00\x0A\x7C\x6E"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x53\x04\x00\x00\x00\x0A\x7D\xBF"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x54\x04\x00\x00\x00\x0A\x7C\x08"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x55\x04\x00\x00\x00\x0A\x7D\xD9"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x56\x04\x00\x00\x00\x0A\x7D\xEA"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x57\x04\x00\x00\x00\x0A\x7C\x3B"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x58\x04\x00\x00\x00\x0A\x7C\xC4"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x59\x04\x00\x00\x00\x0A\x7D\x15"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x5A\x04\x00\x00\x00\x0A\x7D\x26"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x5B\x04\x00\x00\x00\x0A\x7C\xF7"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x5C\x04\x00\x00\x00\x0A\x7D\x40"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x5D\x04\x00\x00\x00\x0A\x7C\x91"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x5E\x04\x00\x00\x00\x0A\x7C\xA2"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x5F\x04\x00\x00\x00\x0A\x7D\x73"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x60\x04\x00\x00\x00\x0A\x78\x7C"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x61\x04\x00\x00\x00\x0A\x79\xAD"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x62\x04\x00\x00\x00\x0A\x79\x9E"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x63\x04\x00\x00\x00\x0A\x78\x4F"))		
			time.Sleep(time.Second*1)
			_, err = conn.Write([]byte("\x64\x04\x00\x00\x00\x0A\x79\xF8"))		
			time.Sleep(time.Second*1)
			
			
			
			
			
			
			
			if err != nil {
			fmt.Println("error: ",err)
			return
			}
		
		}

	}
}

func handleConnection(conn net.Conn) { 
	buffer := make([]byte, 1024)
//连接redis
	db, _ := beego.AppConfig.Int("redis_db")
	rs, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
		return
	}
	rs.Do("SELECT", db)
	fmt.Println("redis wait...")
	
//检测下位机数据
	for {  
        n, err2 := conn.Read(buffer)  
   
        if err2 != nil {
			fmt.Println("error: ",err2)
			return
		} 
        Data :=(buffer[:n])  
        messager := make(chan byte)  
//        postda :=make(chan byte)  
        //心跳计时  
//        go HeartBeating(conn,messager,timeout)  
        //检测每次Client是否有数据传来  
        go GravelChannel(Data,messager)  
        print( conn.RemoteAddr().String())
		
		str := ([]byte(Data))
		
		lenstr := len(str)
		l := fmt.Sprintf("%d", lenstr)
		
		if l == "25" {
		str2:= crc16.Crc16(str)
		crc1 := fmt.Sprintf("%04X",str2)
		
		crc2 := fmt.Sprintf("%X", str[lenstr-2:lenstr])
		

		if crc1 == crc2 {
			
			
						
			fmt.Printf("receive data string OK :%X \n", string(Data))
//设备终端编号Tid
			Tid := fmt.Sprintf("%X", string(str[0:1]))
			

//设备终端编号 10进制

			Ttid,_ := strconv.ParseInt(string(Tid), 16, 64)
			Td := fmt.Sprintf("2016detector_%d",Ttid)			
			
//漏电值 Leaked （16进制）
			Leaked1 := fmt.Sprintf("%X", string(str[7:9]))			
			

//漏电值 10进制	

			Tleaked1,_ :=strconv.ParseInt(string(Leaked1), 16, 64)
			Td1 := fmt.Sprintf("%d",Tleaked1)
			
//温度值 Temperature	（16进制）
			Temperature1 := fmt.Sprintf("%X", string(str[11:13]))
			Temperature2 := fmt.Sprintf("%X", string(str[15:17]))
			Temperature3 := fmt.Sprintf("%X", string(str[19:21]))
			

//温度值 10进制

			Ttemperature1,_ :=strconv.ParseInt(string(Temperature1), 16, 64)
			Te1 := fmt.Sprintf("%d",Ttemperature1)
			Ttemperature2,_ :=strconv.ParseInt(string(Temperature2), 16, 64)
			Te2 := fmt.Sprintf("%d",Ttemperature2)
			Ttemperature3,_ :=strconv.ParseInt(string(Temperature3), 16, 64)
			Te3 := fmt.Sprintf("%d",Ttemperature3)

//报警值  Output （16进制）
			Output1 := fmt.Sprintf("%X", string(str[9:11]))
			Output2 := fmt.Sprintf("%X", string(str[13:15]))
			Output3 := fmt.Sprintf("%X", string(str[17:19]))
			Output4 := fmt.Sprintf("%X", string(str[21:23]))
	
//报警值 10进制
			Toutput1,_:=strconv.ParseInt(string(Output1), 16, 64)
			Tt1 := fmt.Sprintf("%d",Toutput1)
			Toutput2,_:=strconv.ParseInt(string(Output2), 16, 64)
			Tt2 := fmt.Sprintf("%d",Toutput2)
			Toutput3,_:=strconv.ParseInt(string(Output3), 16, 64)
			Tt3 := fmt.Sprintf("%d",Toutput3)
			Toutput4,_:=strconv.ParseInt(string(Output4), 16, 64)
			Tt4 := fmt.Sprintf("%d",Toutput4)

			
//状态 state 
			
			state1 := fmt.Sprintf("%X", string(str[5:7]))
			state2, _ := strconv.ParseInt(string(state1), 16, 64)
			state3 := uint64(state2)
			state := fmt.Sprintf("%012s", strconv.FormatUint(state3, 2))
// 各路状态
			St4 := fmt.Sprintf("%s", state[0:3])
			St3 := fmt.Sprintf("%s", state[3:6])
			St2 := fmt.Sprintf("%s", state[6:9])
			St1 := fmt.Sprintf("%s", state[9:12])
			
			
			

//数据存入redis
			key := Td
			
//			imap := map[string]string{"101":Td1 , "102":Te1 , "103":Te2 , "104":Te3  , "201":Tt1 , "202":Tt2 , "203":Tt3, "204":Tt4, "301":St4, "302":St3, "303":St2, "304":St1}
			imap := map[string]string{"101value":Td1,"102standard":Tt1,"103type":"1","104status":St1,"121value":Te1,"122standard":Tt2,"123type":"2","124status":St2,"131value":Te2,"132standard":Tt3,"133type":"2","134status":St3,"141value":Te3,"142standard":Tt4,"143type":"2","144status":St4}
			

// 将map转换成json数据
			value, _ := json.Marshal(imap)
// 存入redis
			_, err := rs.Do("SET", key, value)
			if err != nil {
				fmt.Println(err)
			}
			_, _ = rs.Do("EXPIRE", key, 24*3600)
			
			
			
/*			
			} else if now == "03" {

			
			fmt.Printf("receive data string OK :%X \n", string(Data))
			
//设备终端编号Tid
			Tid := fmt.Sprintf("%X", string(str[0:1]))
			

//设备终端编号 10进制

			Ttid,_ := strconv.ParseInt(string(Tid), 16, 64)
			Td := fmt.Sprintf("2016detector_t_%d",Ttid)
//报警门槛值 Threshold	（16进制）
			Threshold1 := fmt.Sprintf("%X", string(str[3:5]))
			Threshold2 := fmt.Sprintf("%X", string(str[5:7]))
			Threshold3 := fmt.Sprintf("%X", string(str[7:9]))
			Threshold4 := fmt.Sprintf("%X", string(str[9:11]))
			

//报警门槛值 10进制

			Tthreshold1,_ :=strconv.ParseInt(string(Threshold1), 16, 64)
			Th1 := fmt.Sprintf("%d",Tthreshold1)
			Tthreshold2,_ :=strconv.ParseInt(string(Threshold2), 16, 64)
			Th2 := fmt.Sprintf("%d",Tthreshold2)
			Tthreshold3,_ :=strconv.ParseInt(string(Threshold3), 16, 64)
			Th3 := fmt.Sprintf("%d",Tthreshold3)
			Tthreshold4,_ :=strconv.ParseInt(string(Threshold4), 16, 64)
			Th4 := fmt.Sprintf("%d",Tthreshold4)

//数据存入redis
			key := Td
			
//			imap := map[string]string{"101":Td1 , "102":Td2 , "103":Td3 , "104":Td4 , "105":Td5 , "106":Td6 , "107":Td7 , "108":Td8 , "109":Te1 , "110":Te2 , "111":Te3 , "112":Te4 , "301":Tt1 , "302":Tt2 , "303":Tt3 , "304":Tt4 , "305":Tt5 , "306":Tt6 , "307":Tt7 , "308":Tt8 , "309":Tt9 , "310":Tt10 , "311":Tt11 , "312":Tt12}
			imap := map[string]string{"101":Th1 , "102":Th2 , "103":Th3 , "104":Th4 }

// 将map转换成json数据
			value, _ := json.Marshal(imap)
// 存入redis
			_, err := rs.Do("SET", key, value)
			if err != nil {
				fmt.Println(err)
			}
			_, _ = rs.Do("EXPIRE", key, 24*3600)			
				
			}
*/
			if St1 == "001" || St2 == "001" || St2 == "001" || St3 == "001" || St4 == "001" {
				num := fmt.Sprintf("%x%x%x%x%x%x%x%x%s%s%s%s",string(str[7:9]),string(str[11:13]),string(str[15:17]),string(str[19:21]),string(str[9:11]),string(str[13:15]),string(str[17:19]),string(str[21:23]),state[9:12],state[6:9],state[3:6],state[0:3])
			
// API	get传值	
				u, _ := url.Parse("http://xxfgjapi.36ht.com/api/Intelligent/AlarmGet.json")
				q := u.Query()
		

				q.Set("key",Td)
				q.Set("value",num)
			

		
				u.RawQuery = q.Encode()
				res, err := http.Get(u.String());
				if err != nil { 
					continue
				}
				result, err := ioutil.ReadAll(res.Body) 
				res.Body.Close() 
				if err != nil { 
					continue
				} 
				fmt.Printf("%s\n", result)
				
				}
			
						
			
		
		
		} else {
			fmt.Printf("receive data string 1NO :%X \n", string(Data))
		
		}
		}else {
			fmt.Printf("receive data string 2NO :%X \n", string(Data))
		}
		
		
		
	}	
	
}
func GravelChannel(n []byte,mess chan byte) {  
    for _ , v := range n{  
         mess <- v  
    }  
    close(mess)  
}  

func Log(v ...interface{}) {  
    fmt.Println(v...)  
} 		