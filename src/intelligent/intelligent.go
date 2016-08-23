package main

import (
	"fmt"
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
		fmt.Println("error：", err)
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
			fmt.Println("error：", err2)
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
		fmt.Println("addr：", remoteAddr)
		go handleConnection(conn)
		for {
			_, err := conn.Write([]byte("\x41\x00\x01\x04\x00\x45"))		
			time.Sleep(time.Second*5)
			
			if err != nil {
			fmt.Println("error: ",err)
			return
			}
		
		}
	}
}

//下位机上传数据
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
//遍历		
		str := ([]byte(Data))
		var str2 byte
		lenstr := len(str)
		for i := 0; i < lenstr-1; i++ {
		
		str2 += str[i]
		}
		if str2 == str[lenstr-1] {

			fmt.Printf("receive data string OK:%X \n", string(Data))
			
//设备终端编号Tid
			Tid := fmt.Sprintf("%X", string(str[1:3]))
			

//设备终端编号 10进制

			Ttid,_ := strconv.ParseInt(string(Tid), 16, 64)
			Td := fmt.Sprintf("2015detector_%d",Ttid)

			
//漏电值 Leaked （16进制）
			Leaked1 := fmt.Sprintf("%X", string(str[6:8]))			
			Leaked2 := fmt.Sprintf("%X", string(str[13:15]))
			Leaked3 := fmt.Sprintf("%X", string(str[20:22]))
			Leaked4 := fmt.Sprintf("%X", string(str[27:29]))
			Leaked5 := fmt.Sprintf("%X", string(str[34:36]))
			Leaked6 := fmt.Sprintf("%X", string(str[41:43]))
			Leaked7 := fmt.Sprintf("%X", string(str[48:50]))
			Leaked8 := fmt.Sprintf("%X", string(str[55:57]))

//漏电值 10进制	

			Tleaked1,_ :=strconv.ParseInt(string(Leaked1), 16, 64)
			Td1 := fmt.Sprintf("%d",Tleaked1)
			Tleaked2,_ :=strconv.ParseInt(string(Leaked2), 16, 64)
			Td2 := fmt.Sprintf("%d",Tleaked2)
			Tleaked3,_ :=strconv.ParseInt(string(Leaked3), 16, 64)
			Td3 := fmt.Sprintf("%d",Tleaked3)
			Tleaked4,_ :=strconv.ParseInt(string(Leaked4), 16, 64)
			Td4 := fmt.Sprintf("%d",Tleaked4)
			Tleaked5,_ :=strconv.ParseInt(string(Leaked5), 16, 64)
			Td5 := fmt.Sprintf("%d",Tleaked5)
			Tleaked6,_ :=strconv.ParseInt(string(Leaked6), 16, 64)
			Td6 := fmt.Sprintf("%d",Tleaked6)
			Tleaked7,_ :=strconv.ParseInt(string(Leaked7), 16, 64)
			Td7 := fmt.Sprintf("%d",Tleaked7)
			Tleaked8,_ :=strconv.ParseInt(string(Leaked8), 16, 64)
			Td8 := fmt.Sprintf("%d",Tleaked8)

			
			
//温度值 Temperature	（16进制）
			Temperature1 := fmt.Sprintf("%X", string(str[83:85]))
			Temperature2 := fmt.Sprintf("%X", string(str[90:92]))
			Temperature3 := fmt.Sprintf("%X", string(str[97:99]))
			Temperature4 := fmt.Sprintf("%X", string(str[104:106]))

//温度值 10进制

			Ttemperature1,_ :=strconv.ParseInt(string(Temperature1), 16, 64)
			Te1 := fmt.Sprintf("%d",Ttemperature1)
			Ttemperature2,_ :=strconv.ParseInt(string(Temperature2), 16, 64)
			Te2 := fmt.Sprintf("%d",Ttemperature2)
			Ttemperature3,_ :=strconv.ParseInt(string(Temperature3), 16, 64)
			Te3 := fmt.Sprintf("%d",Ttemperature3)
			Ttemperature4,_ :=strconv.ParseInt(string(Temperature4), 16, 64)
			Te4 := fmt.Sprintf("%d",Ttemperature4)


			
//报警值  Output （16进制）
			Output1 := fmt.Sprintf("%X", string(str[8:10]))
			Output2 := fmt.Sprintf("%X", string(str[15:17]))
			Output3 := fmt.Sprintf("%X", string(str[22:24]))
			Output4 := fmt.Sprintf("%X", string(str[29:31]))
			Output5 := fmt.Sprintf("%X", string(str[36:38]))
			Output6 := fmt.Sprintf("%X", string(str[43:45]))
			Output7 := fmt.Sprintf("%X", string(str[50:52]))
			Output8 := fmt.Sprintf("%X", string(str[57:59]))
			Output9 := fmt.Sprintf("%X", string(str[85:87]))
			Output10 := fmt.Sprintf("%X", string(str[92:94]))
			Output11 := fmt.Sprintf("%X", string(str[99:101]))
			Output12 := fmt.Sprintf("%X", string(str[106:108]))
//报警值 10进制
			Toutput1,_:=strconv.ParseInt(string(Output1), 16, 64)
			Tt1 := fmt.Sprintf("%d",Toutput1)
			Toutput2,_:=strconv.ParseInt(string(Output2), 16, 64)
			Tt2 := fmt.Sprintf("%d",Toutput2)
			Toutput3,_:=strconv.ParseInt(string(Output3), 16, 64)
			Tt3 := fmt.Sprintf("%d",Toutput3)
			Toutput4,_:=strconv.ParseInt(string(Output4), 16, 64)
			Tt4 := fmt.Sprintf("%d",Toutput4)
			Toutput5,_:=strconv.ParseInt(string(Output5), 16, 64)
			Tt5 := fmt.Sprintf("%d",Toutput5)
			Toutput6,_:=strconv.ParseInt(string(Output6), 16, 64)
			Tt6 := fmt.Sprintf("%d",Toutput6)
			Toutput7,_:=strconv.ParseInt(string(Output7), 16, 64)
			Tt7 := fmt.Sprintf("%d",Toutput7)
			Toutput8,_:=strconv.ParseInt(string(Output8), 16, 64)
			Tt8 := fmt.Sprintf("%d",Toutput8)
			Toutput9,_:=strconv.ParseInt(string(Output9), 16, 64)
			Tt9 := fmt.Sprintf("%d",Toutput9)
			Toutput10,_:=strconv.ParseInt(string(Output10), 16, 64)
			Tt10 := fmt.Sprintf("%d",Toutput10)
			Toutput11,_:=strconv.ParseInt(string(Output11), 16, 64)
			Tt11 := fmt.Sprintf("%d",Toutput11)
			Toutput12,_:=strconv.ParseInt(string(Output12), 16, 64)
			Tt12 := fmt.Sprintf("%d",Toutput12)
			
			
			

//数据存入redis
			key := Td
			
			imap := map[string]string{"101":Td1 , "102":Td2 , "103":Td3 , "104":Td4 , "105":Td5 , "106":Td6 , "107":Td7 , "108":Td8 , "109":Te1 , "110":Te2 , "111":Te3 , "112":Te4 , "301":Tt1 , "302":Tt2 , "303":Tt3 , "304":Tt4 , "305":Tt5 , "306":Tt6 , "307":Tt7 , "308":Tt8 , "309":Tt9 , "310":Tt10 , "311":Tt11 , "312":Tt12}
//			imap := map[string]string{"101":Td1 , "102":Td2 , "103":Td3 , "104":Td4 , "105":Td5 , "106":Td6 , "107":Td7 , "108":Td8 , "109":Te1 , "110":Te2 , "111":Te3 , "112":Te4}

// 将map转换成json数据
			value, _ := json.Marshal(imap)
// 存入redis
			_, err := rs.Do("SET", key, value)
			if err != nil {
				fmt.Println(err)
			}
			_, _ = rs.Do("EXPIRE", key, 24*3600)
			
		
		
//判断自动报警
			if Leaked1 > Output1 || Leaked2 > Output2 || Leaked3 > Output3 || Leaked4 > Output4 || Leaked5 > Output5 || Leaked6 > Output6 || Leaked7 > Output7 || Leaked8 > Output8 || Temperature1 > Output9 || Temperature2 > Output10 || Temperature3 > Output11 || Temperature4 > Output12 {


// 转换ascii		
				num := fmt.Sprintf("%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X",string(str[1:3]),string(str[6:8]),string(str[13:15]),string(str[20:22]),string(str[27:29]),string(str[34:36]),string(str[41:43]),string(str[48:50]),string(str[55:57]),string(str[83:85]),string(str[90:92]),string(str[97:99]),string(str[104:106]),string(str[8:10]),string(str[15:17]),string(str[22:24]),string(str[29:31]),string(str[36:38]),string(str[43:45]),string(str[50:52]),string(str[57:59]),string(str[85:87]),string(str[92:94]),string(str[99:101]),string(str[106:108]))
// API	get传值	
				u, _ := url.Parse("http://xxfgjapi.36ht.com/api/Intelligent/AlarmGet.json")
				q := u.Query()
		


				q.Set("key",num)
			

		
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
			fmt.Printf("receive data string NO:%X \n", string(Data))
		
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

