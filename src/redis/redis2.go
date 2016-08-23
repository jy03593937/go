package main

import (
	"fmt"
//	"io/ioutil"
	"net"
//	"bytes"
	"time"
//	"net/http"
//	"net/url"
//	"os"
//  "strconv"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"github.com/astaxie/beego"
)
const RECV_BUF_LEN = 1024

func main() {
	localAddress, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:8383") 
	var tcpListener, err = net.ListenTCP("tcp", localAddress)      
	if err != nil {
		fmt.Println("error：", err)
		return
	}
	defer func() { //
		tcpListener.Close()
	}()
	fmt.Println("TCPAddr wait...")
		
	
	
	for {
		var conn, err2 = tcpListener.AcceptTCP() 
		if err2 != nil {
			fmt.Println("error：", err2)
			return
		}
		var remoteAddr = conn.RemoteAddr() 
		fmt.Println("addr：", remoteAddr)
		go handleConnection(conn)
		
		
        for{
			_, err2 = conn.Write([]byte("\x41\x00\x01\x04\x00\x45"))		
			time.Sleep(time.Second*10)
			if err2 != nil {
			fmt.Println("error: ",err2)
			return
			}
		}  	

		
	}
}		
	

//下位机上传数据 net.Conn
func handleConnection(conn net.Conn ) { 
	buffer := make([]byte, 1024)  
	
//redis	
	db, _ := beego.AppConfig.Int("redis_db")
// tcp连接redis
	rs, err := redis.Dial("tcp", "127.0.0.1:6379")
// 操作完后自动关闭
//	defer rs.Close()
// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
		return
	}
// 选择db
	rs.Do("SELECT", db)
	
	fmt.Println("redis wait...")
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
//       print( conn.RemoteAddr().String())
//遍历		
		str := ([]byte(Data))
		var str2 byte
		lenstr := len(str)
		for i := 0; i < lenstr-1; i++ {
		
		str2 += str[i]
		}
		if str2 == str[lenstr-1] {

 //		fmt.Printf("receive data string OK:%X \n", string(Data))
 //           num := fmt.Sprintf("%X", string(Data))
//		t := time.Now().Unix()
		
		// 取json数据
// 先声明imap用来装数据
		var imap map[string]string
		key := "0001"
// json数据在go中是[]byte类型，所以此处用redis.Bytes转换
		value, err := redis.Bytes(rs.Do("GET", key))
		if err != nil {
			fmt.Println(err)
		}
 //将json解析成map类型
		errShal := json.Unmarshal(value, &imap)
		if errShal != nil {
			fmt.Println(err)
		}
		fmt.Println(imap["12"])
//		fmt.Println(imap["key2"])

		

		

/*
// 转换ascii		
		num := fmt.Sprintf("%X", string(Data))
// API	get传值	
				u, _ := url.Parse("http://xfzxapi.36ht.com/api/Intelligent/intelligent.json")
				q := u.Query()
		


				q.Set("key",num)
			

		
				u.RawQuery = q.Encode()
				res, err := http.Get(u.String());
				if err != nil { 
      
				}
				result, err := ioutil.ReadAll(res.Body) 
				res.Body.Close() 
				if err != nil { 
     
				} 
				fmt.Printf("%s\n", result)
				
*/	
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