package main

import (
	"fmt"

	 "github.com/garyburd/redigo/redis"
)

func main() {
	// 从配置文件获取redis配置并连接  
host := beego.AppConfig.String("redis_host")  
db, _ := beego.AppConfig.Int("redis_db")  
// tcp连接redis  
rs, err := redis.Dial("tcp", host)  
// 操作完后自动关闭  
defer rs.Close()  
// 若连接出错，则打印错误信息，返回  
if err != nil {  
    fmt.Println(err)  
    fmt.Println("redis connect error")  
    return  
}  
// 选择db  
rs.Do("SELECT", db)  

key := "aaa"  
value := {"value":"0","standard":"0","type":"1","status":4,"value":"0","standard":"0","type":"2","status":4,"value":"0","standard":"0","type":"2","status":4,"value":"0","standard":"0","type":"2","status":4}
// 操作redis时调用Do方法，第一个参数传入操作名称（字符串），然后根据不同操作传入key、value、数字等  
// 返回2个参数，第一个为操作标识，成功则为1，失败则为0；第二个为错误信息  
n, err := rs.Do("SETNX", key, value)  
// 若操作失败则返回  
if err != nil {  
    fmt.Println(err)  
    return  
}  


}
