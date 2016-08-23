package main
import (
    "fmt"
    // 导入redigo扩展包
    "github.com/garyburd/redigo/redis"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
	
)
// 从配置文件获取redis配置并连接
func main() {
	t := time.Now()
	h :=fmt.Sprintf("%d",t.Hour())
	for{
//host := beego.AppConfig.String("redis_host")
		db, _ := beego.AppConfig.Int("redis_db")
// tcp连接redis
		rs, err := redis.Dial("tcp", "127.0.0.1:6379")
		time.Sleep(time.Second*60)

// 操作完后自动关闭
//		defer rs.Close()
// 若连接出错，则打印错误信息，返回
		if err != nil {
			fmt.Println(err)
			fmt.Println("redis connect error")
			return
		}
// 选择db
		rs.Do("SELECT", db)
		

// 取json数据  
// 先声明imap用来装数据  
		var imap map[string]string  
		key1 := "1"  
// json数据在go中是[]byte类型，所以此处用redis.Bytes转换  
		value, err := redis.Bytes(rs.Do("GET", key1))  
		if err != nil {  
			fmt.Println(err)  
		}  
// 将json解析成map类型  
		errShal := json.Unmarshal(value, &imap)  
		if errShal != nil {  
			fmt.Println(err)  
		}  
//		fmt.Println(imap["101"])  
//		fmt.Println(imap["102"])
		fmt.Println(imap)
		
		for{
			switch h {
			
			case "0" :
//存0点json数据		
				key2 :=key1+"t0"
				
				imap2 := map[string]string{"101": imap["101"], "102": imap["102"], "103": imap["103"],"104": imap["104"],"105": imap["105"],"106": imap["106"],"107": imap["107"],"108": imap["108"],"109": imap["109"],"110": imap["110"],"111": imap["111"],"112": imap["112"],"301": imap["301"],"302": imap["302"],"303": imap["303"],"304": imap["304"],"305": imap["305"],"306": imap["306"],"307": imap["307"],"308": imap["308"],"309": imap["309"],"310": imap["310"],"311": imap["311"],"312": imap["312"]}  
				
// 将map转换成json数据  
				value2, _ := json.Marshal(imap2)  
// 存入redis  
				_, err := rs.Do("SET", key2, value2)  
				if err != nil {  
					fmt.Println(err)  
				} 
				_, _ = rs.Do("EXPIRE", key2, 24*3600)	
				fmt.Println("success",key1)
//1路电流				
				dkey1 :=key1+"d1"
				dimap1 := map[string]string{"200":imap2["101"],"300":imap2["301"]}
				dvalue1, _ := json.Marshal(dimap1)
				_, err = rs.Do("SET", dkey1, dvalue1)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("1路电流")
				
				dkey2 :=key1+"d2"
				dimap2 := map[string]string{"200":imap2["102"],"300":imap2["302"]}
				dvalue2, _ := json.Marshal(dimap2)
				_, err = rs.Do("SET", dkey2, dvalue2)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("2路电流")
				
				dkey3 :=key1+"d3"
				dimap3 := map[string]string{"200":imap2["103"],"300":imap2["303"]}
				dvalue3, _ := json.Marshal(dimap3)
				_, err = rs.Do("SET", dkey3, dvalue3)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("3路电流")
				
				dkey4 :=key1+"d4"
				dimap4 := map[string]string{"200":imap2["104"],"300":imap2["304"]}
				dvalue4, _ := json.Marshal(dimap4)
				_, err = rs.Do("SET", dkey4, dvalue4)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4路电流")
				
				dkey5 :=key1+"d5"
				dimap5 := map[string]string{"200":imap2["105"],"300":imap2["305"]}
				dvalue5, _ := json.Marshal(dimap5)
				_, err = rs.Do("SET", dkey5, dvalue5)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("5路电流")
				
				dkey6 :=key1+"d6"
				dimap6 := map[string]string{"200":imap2["106"],"300":imap2["306"]}
				dvalue6, _ := json.Marshal(dimap6)
				_, err = rs.Do("SET", dkey6, dvalue6)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("6路电流")
				
				dkey7 :=key1+"d7"
				dimap7 := map[string]string{"200":imap2["107"],"300":imap2["307"]}
				dvalue7, _ := json.Marshal(dimap7)
				_, err = rs.Do("SET", dkey7, dvalue7)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("7路电流")
				
				dkey8 :=key1+"d8"
				dimap8 := map[string]string{"200":imap2["108"],"300":imap2["308"]}
				dvalue8, _ := json.Marshal(dimap8)
				_, err = rs.Do("SET", dkey8, dvalue8)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8路电流")
				
				dkey9 :=key1+"d9"
				dimap9 := map[string]string{"200":imap2["109"],"300":imap2["309"]}
				dvalue9, _ := json.Marshal(dimap9)
				_, err = rs.Do("SET", dkey9, dvalue9)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("9路温度")
				
				dkey10 :=key1+"d10"
				dimap10 := map[string]string{"200":imap2["110"],"300":imap2["310"]}
				dvalue10, _ := json.Marshal(dimap10)
				_, err = rs.Do("SET", dkey10, dvalue10)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("10路温度")
				
				dkey11 :=key1+"d11"
				dimap11 := map[string]string{"200":imap2["111"],"300":imap2["311"]}
				dvalue11, _ := json.Marshal(dimap11)
				_, err = rs.Do("SET", dkey11, dvalue11)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("11路温度")
				
				dkey12 :=key1+"d12"
				dimap12 := map[string]string{"200":imap2["112"],"300":imap2["312"]}
				dvalue12, _ := json.Marshal(dimap12)
				_, err = rs.Do("SET", dkey12, dvalue12)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12路温度")
				time.Sleep(time.Second*60)
				  
			  
			
			case "4" :
//存4点数值			
				key3 :=key1+"t4"
				imap3 := map[string]string{"101": imap["101"], "102": imap["102"], "103": imap["103"],"104": imap["104"],"105": imap["105"],"106": imap["106"],"107": imap["107"],"108": imap["108"],"109": imap["109"],"110": imap["110"],"111": imap["111"],"112": imap["112"],"301": imap["301"],"302": imap["302"],"303": imap["303"],"304": imap["304"],"305": imap["305"],"306": imap["306"],"307": imap["307"],"308": imap["308"],"309": imap["309"],"310": imap["310"],"311": imap["311"],"312": imap["312"]}
				value3, _ := json.Marshal(imap3)
				_, err := rs.Do("SETNX", key3, value3)
				if err != nil {  
					fmt.Println(err)  
				}  
				_, _ = rs.Do("EXPIRE", key3, 24*3600)  
				fmt.Println("success",key3)
				
//取4点前各回路数值
				var imap_0 map[string]string
				key_0 := key1+"t0"
				value_0,err := redis.Bytes(rs.Do("GET", key_0))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal := json.Unmarshal(value_0, &imap_0)  
				if errShal != nil {  
					fmt.Println(err)  
				} 


				
//存入4点数值各回路历史值				
				dkey1 :=key1+"d1"
				dimap1 := map[string]string{"200":imap_0["101"],"204":imap3["101"],"304":imap3["301"]}
				dvalue1, _ := json.Marshal(dimap1)
				_, err = rs.Do("SET", dkey1, dvalue1)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点1路电流")
				
				dkey2 :=key1+"d2"
				dimap2 := map[string]string{"200":imap_0["102"],"204":imap3["102"],"304":imap3["302"]}
				dvalue2, _ := json.Marshal(dimap2)
				_, err = rs.Do("SET", dkey2, dvalue2)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点2路电流")
				
				dkey3 :=key1+"d3"
				dimap3 := map[string]string{"200":imap_0["103"],"204":imap3["103"],"304":imap3["303"]}
				dvalue3, _ := json.Marshal(dimap3)
				_, err = rs.Do("SET", dkey3, dvalue3)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点3路电流")
				
				dkey4 :=key1+"d4"
				dimap4 := map[string]string{"200":imap_0["104"],"204":imap3["104"],"304":imap3["304"]}
				dvalue4, _ := json.Marshal(dimap4)
				_, err = rs.Do("SET", dkey4, dvalue4)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点4路电流")
				
				dkey5 :=key1+"d5"
				dimap5 := map[string]string{"200":imap_0["105"],"204":imap3["105"],"304":imap3["305"]}
				dvalue5, _ := json.Marshal(dimap5)
				_, err = rs.Do("SET", dkey5, dvalue5)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点5路电流")
				
				dkey6 :=key1+"d6"
				dimap6 := map[string]string{"200":imap_0["106"],"204":imap3["106"],"304":imap3["306"]}
				dvalue6, _ := json.Marshal(dimap6)
				_, err = rs.Do("SET", dkey6, dvalue6)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点6路电流")
				
				dkey7 :=key1+"d7"
				dimap7 := map[string]string{"200":imap_0["107"],"204":imap3["107"],"304":imap3["307"]}
				dvalue7, _ := json.Marshal(dimap7)
				_, err = rs.Do("SET", dkey7, dvalue7)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点7路电流")
				
				dkey8 :=key1+"d8"
				dimap8 := map[string]string{"200":imap_0["108"],"204":imap3["108"],"304":imap3["308"]}
				dvalue8, _ := json.Marshal(dimap8)
				_, err = rs.Do("SET", dkey8, dvalue8)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点8路电流")
				
				dkey9 :=key1+"d9"
				dimap9 := map[string]string{"200":imap_0["109"],"204":imap3["109"],"304":imap3["309"]}
				dvalue9, _ := json.Marshal(dimap9)
				_, err = rs.Do("SET", dkey9, dvalue9)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点9路温度")
				
				dkey10 :=key1+"d10"
				dimap10 := map[string]string{"200":imap_0["110"],"204":imap3["110"],"304":imap3["310"]}
				dvalue10, _ := json.Marshal(dimap10)
				_, err = rs.Do("SET", dkey10, dvalue10)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点10路温度")
				
				dkey11 :=key1+"d11"
				dimap11 := map[string]string{"200":imap_0["111"],"204":imap3["111"],"304":imap3["311"]}
				dvalue11, _ := json.Marshal(dimap11)
				_, err = rs.Do("SET", dkey11, dvalue11)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点11路温度")
				
				dkey12 :=key1+"d12"
				dimap12 := map[string]string{"200":imap_0["112"],"204":imap3["112"],"304":imap3["312"]}
				dvalue12, _ := json.Marshal(dimap12)
				_, err = rs.Do("SET", dkey12, dvalue12)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("4点12路温度")
				
				time.Sleep(time.Second*60)
				
				
				
				
				
				
				
				case "10" :
//存8点时即时数据
				key4 :=key1+"t8"
				imap4 := map[string]string{"101": imap["101"], "102": imap["102"], "103": imap["103"],"104": imap["104"],"105": imap["105"],"106": imap["106"],"107": imap["107"],"108": imap["108"],"109": imap["109"],"110": imap["110"],"111": imap["111"],"112": imap["112"],"301": imap["301"],"302": imap["302"],"303": imap["303"],"304": imap["304"],"305": imap["305"],"306": imap["306"],"307": imap["307"],"308": imap["308"],"309": imap["309"],"310": imap["310"],"311": imap["311"],"312": imap["312"]}
				value4, _ := json.Marshal(imap4)
				_, err := rs.Do("SETNX", key4, value4)
				if err != nil {  
					fmt.Println(err)  
				}  
				_, _ = rs.Do("EXPIRE", key4, 24*3600) 
				fmt.Println("success",key4)
//取8点前各回路数值 （0点和4点）
			var imap_0 map[string]string
				key_0 := key1+"t0"
				value_0,err := redis.Bytes(rs.Do("GET", key_0))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal := json.Unmarshal(value_0, &imap_0)  
				if errShal != nil {  
					fmt.Println(err)  
				} 	
			var imap_4 map[string]string
				key_4 := key1+"t4"
				value_4,err := redis.Bytes(rs.Do("GET", key_4))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal2 := json.Unmarshal(value_4, &imap_4)  
				if errShal2 != nil {  
					fmt.Println(err)  
				}
//存入8点数值各回路历史值
				dkey1 :=key1+"d1"
				dimap1 := map[string]string{"200":imap_0["101"],"204":imap_4["101"],"208":imap4["101"],"308":imap4["301"]}
				dvalue1, _ := json.Marshal(dimap1)
				_, err = rs.Do("SET", dkey1, dvalue1)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点1路电流")
				
				dkey2 :=key1+"d2"
				dimap2 := map[string]string{"200":imap_0["102"],"204":imap_4["102"],"208":imap4["102"],"308":imap4["302"]}
				dvalue2, _ := json.Marshal(dimap2)
				_, err = rs.Do("SET", dkey2, dvalue2)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点2路电流")
				
				dkey3 :=key1+"d3"
				dimap3 := map[string]string{"200":imap_0["103"],"204":imap_4["103"],"208":imap4["103"],"308":imap4["303"]}
				dvalue3, _ := json.Marshal(dimap3)
				_, err = rs.Do("SET", dkey3, dvalue3)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点3路电流")
				
				dkey4 :=key1+"d4"
				dimap4 := map[string]string{"200":imap_0["104"],"204":imap_4["104"],"208":imap4["104"],"308":imap4["304"]}
				dvalue4, _ := json.Marshal(dimap4)
				_, err = rs.Do("SET", dkey4, dvalue4)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点4路电流")
				
				dkey5 :=key1+"d5"
				dimap5 := map[string]string{"200":imap_0["105"],"204":imap_4["105"],"208":imap4["105"],"308":imap4["305"]}
				dvalue5, _ := json.Marshal(dimap5)
				_, err = rs.Do("SET", dkey5, dvalue5)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点5路电流")
				
				dkey6 :=key1+"d6"
				dimap6 := map[string]string{"200":imap_0["106"],"204":imap_4["106"],"208":imap4["106"],"308":imap4["306"]}
				dvalue6, _ := json.Marshal(dimap6)
				_, err = rs.Do("SET", dkey6, dvalue6)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点6路电流")
				
				dkey7 :=key1+"d7"
				dimap7 := map[string]string{"200":imap_0["107"],"204":imap_4["107"],"208":imap4["107"],"308":imap4["307"]}
				dvalue7, _ := json.Marshal(dimap7)
				_, err = rs.Do("SET", dkey7, dvalue7)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点7路电流")
				
				dkey8 :=key1+"d8"
				dimap8 := map[string]string{"200":imap_0["108"],"204":imap_4["108"],"208":imap4["108"],"308":imap4["308"]}
				dvalue8, _ := json.Marshal(dimap8)
				_, err = rs.Do("SET", dkey8, dvalue8)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点8路电流")
				
				dkey9 :=key1+"d9"
				dimap9 := map[string]string{"200":imap_0["109"],"204":imap_4["109"],"208":imap4["109"],"308":imap4["309"]}
				dvalue9, _ := json.Marshal(dimap9)
				_, err = rs.Do("SET", dkey9, dvalue9)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点9路温度")
				
				dkey10 :=key1+"d10"
				dimap10 := map[string]string{"200":imap_0["110"],"204":imap_4["110"],"208":imap4["110"],"308":imap4["310"]}
				dvalue10, _ := json.Marshal(dimap10)
				_, err = rs.Do("SET", dkey10, dvalue10)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点10路温度")
				
				dkey11 :=key1+"d11"
				dimap11 := map[string]string{"200":imap_0["111"],"204":imap_4["111"],"208":imap4["111"],"308":imap4["311"]}
				dvalue11, _ := json.Marshal(dimap11)
				_, err = rs.Do("SET", dkey11, dvalue11)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点11路温度")
				
				dkey12 :=key1+"d12"
				dimap12 := map[string]string{"200":imap_0["112"],"204":imap_4["112"],"208":imap4["112"],"308":imap4["312"]}
				dvalue12, _ := json.Marshal(dimap12)
				_, err = rs.Do("SET", dkey12, dvalue12)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("8点12路温度")

				time.Sleep(time.Second*60)







				
			case "12" :
//存入12点即时数据
				key5 :=key1+"t12"
				imap5 := map[string]string{"101": imap["101"], "102": imap["102"], "103": imap["103"],"104": imap["104"],"105": imap["105"],"106": imap["106"],"107": imap["107"],"108": imap["108"],"109": imap["109"],"110": imap["110"],"111": imap["111"],"112": imap["112"],"301": imap["301"],"302": imap["302"],"303": imap["303"],"304": imap["304"],"305": imap["305"],"306": imap["306"],"307": imap["307"],"308": imap["308"],"309": imap["309"],"310": imap["310"],"311": imap["311"],"312": imap["312"]}
				value5, _ := json.Marshal(imap5)
				_, err := rs.Do("SETNX", key5, value5)
				if err != nil {  
					fmt.Println(err)  
				}  
				_, _ = rs.Do("EXPIRE", key5, 24*3600)  
				fmt.Println("success",key5)

//取12点前各回路数值（0点、4点、8点）
			var imap_0 map[string]string
				key_0 := key1+"t0"
				value_0,err := redis.Bytes(rs.Do("GET", key_0))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal := json.Unmarshal(value_0, &imap_0)  
				if errShal != nil {  
					fmt.Println(err)  
				} 	
			var imap_4 map[string]string
				key_4 := key1+"t4"
				value_4,err := redis.Bytes(rs.Do("GET", key_4))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal2 := json.Unmarshal(value_4, &imap_4)  
				if errShal2 != nil {  
					fmt.Println(err)  
				}
			var imap_8 map[string]string
				key_8 := key1+"t8"
				value_8,err := redis.Bytes(rs.Do("GET", key_8))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal3 := json.Unmarshal(value_8, &imap_8)  
				if errShal3 != nil {  
					fmt.Println(err)  
				}
// 存入12点数值各回路历史值
				dkey1 :=key1+"d1"
				dimap1 := map[string]string{"200":imap_0["101"],"204":imap_4["101"],"208":imap_8["101"],"212":imap5["101"],"312":imap5["301"]}
				dvalue1, _ := json.Marshal(dimap1)
				_, err = rs.Do("SET", dkey1, dvalue1)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点1路电流")
				
				dkey2 :=key1+"d2"
				dimap2 := map[string]string{"200":imap_0["102"],"204":imap_4["102"],"208":imap_8["102"],"212":imap5["102"],"312":imap5["302"]}
				dvalue2, _ := json.Marshal(dimap2)
				_, err = rs.Do("SET", dkey2, dvalue2)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点2路电流")
				
				dkey3 :=key1+"d3"
				dimap3 := map[string]string{"200":imap_0["103"],"204":imap_4["103"],"208":imap_8["103"],"212":imap5["103"],"312":imap5["303"]}
				dvalue3, _ := json.Marshal(dimap3)
				_, err = rs.Do("SET", dkey3, dvalue3)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点3路电流")
				
				dkey4 :=key1+"d4"
				dimap4 := map[string]string{"200":imap_0["104"],"204":imap_4["104"],"208":imap_8["104"],"212":imap5["104"],"312":imap5["304"]}
				dvalue4, _ := json.Marshal(dimap4)
				_, err = rs.Do("SET", dkey4, dvalue4)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点4路电流")
				
				dkey5 :=key1+"d5"
				dimap5 := map[string]string{"200":imap_0["105"],"204":imap_4["105"],"208":imap_8["105"],"212":imap5["105"],"312":imap5["305"]}
				dvalue5, _ := json.Marshal(dimap5)
				_, err = rs.Do("SET", dkey5, dvalue5)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点5路电流")
				
				dkey6 :=key1+"d6"
				dimap6 := map[string]string{"200":imap_0["106"],"204":imap_4["106"],"208":imap_8["106"],"212":imap5["106"],"312":imap5["306"]}
				dvalue6, _ := json.Marshal(dimap6)
				_, err = rs.Do("SET", dkey6, dvalue6)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点6路电流")
				
				dkey7 :=key1+"d7"
				dimap7 := map[string]string{"200":imap_0["107"],"204":imap_4["107"],"208":imap_8["107"],"212":imap5["107"],"312":imap5["307"]}
				dvalue7, _ := json.Marshal(dimap7)
				_, err = rs.Do("SET", dkey7, dvalue7)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点7路电流")
				
				dkey8 :=key1+"d8"
				dimap8 := map[string]string{"200":imap_0["108"],"204":imap_4["108"],"208":imap_8["108"],"212":imap5["108"],"312":imap5["308"]}
				dvalue8, _ := json.Marshal(dimap8)
				_, err = rs.Do("SET", dkey8, dvalue8)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点8路电流")
				
				dkey9 :=key1+"d9"
				dimap9 := map[string]string{"200":imap_0["109"],"204":imap_4["109"],"208":imap_8["109"],"212":imap5["109"],"312":imap5["309"]}
				dvalue9, _ := json.Marshal(dimap9)
				_, err = rs.Do("SET", dkey9, dvalue9)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点9路温度")
				
				dkey10 :=key1+"d10"
				dimap10 := map[string]string{"200":imap_0["110"],"204":imap_4["110"],"208":imap_8["110"],"212":imap5["110"],"312":imap5["310"]}
				dvalue10, _ := json.Marshal(dimap10)
				_, err = rs.Do("SET", dkey10, dvalue10)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点10路温度")
				
				dkey11 :=key1+"d11"
				dimap11 := map[string]string{"200":imap_0["111"],"204":imap_4["111"],"208":imap_8["111"],"212":imap5["111"],"312":imap5["311"]}
				dvalue11, _ := json.Marshal(dimap11)
				_, err = rs.Do("SET", dkey11, dvalue11)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点11路温度")
				
				dkey12 :=key1+"d12"
				dimap12 := map[string]string{"200":imap_0["112"],"204":imap_4["112"],"208":imap_8["112"],"212":imap5["112"],"312":imap5["312"]}
				dvalue12, _ := json.Marshal(dimap12)
				_, err = rs.Do("SET", dkey12, dvalue12)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("12点12路温度")


				time.Sleep(time.Second*60)

				
				
				

			case "16" :
//存入16点即时数据
				key6 :=key1+"t16"
				imap6 := map[string]string{"101": imap["101"], "102": imap["102"], "103": imap["103"],"104": imap["104"],"105": imap["105"],"106": imap["106"],"107": imap["107"],"108": imap["108"],"109": imap["109"],"110": imap["110"],"111": imap["111"],"112": imap["112"],"301": imap["301"],"302": imap["302"],"303": imap["303"],"304": imap["304"],"305": imap["305"],"306": imap["306"],"307": imap["307"],"308": imap["308"],"309": imap["309"],"310": imap["310"],"311": imap["311"],"312": imap["312"]}
				value6, _ := json.Marshal(imap6)
				_, err := rs.Do("SETNX", key6, value6)
				if err != nil {  
					fmt.Println(err)  
				}  
				_, _ = rs.Do("EXPIRE", key6, 24*3600)  
				fmt.Println("success",key6)
//取16点前各回路数值（0点、4点、8点、12点）
			var imap_0 map[string]string
				key_0 := key1+"t0"
				value_0,err := redis.Bytes(rs.Do("GET", key_0))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal := json.Unmarshal(value_0, &imap_0)  
				if errShal != nil {  
					fmt.Println(err)  
				} 	
			var imap_4 map[string]string
				key_4 := key1+"t4"
				value_4,err := redis.Bytes(rs.Do("GET", key_4))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal2 := json.Unmarshal(value_4, &imap_4)  
				if errShal2 != nil {  
					fmt.Println(err)  
				}
			var imap_8 map[string]string
				key_8 := key1+"t8"
				value_8,err := redis.Bytes(rs.Do("GET", key_8))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal3 := json.Unmarshal(value_8, &imap_8)  
				if errShal3 != nil {  
					fmt.Println(err)  
				}
			var imap_12 map[string]string
				key_12 := key1+"t12"
				value_12,err := redis.Bytes(rs.Do("GET", key_12))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal4 := json.Unmarshal(value_12, &imap_12)  
				if errShal4 != nil {  
					fmt.Println(err)  
				}
//存入16点数值各回路历史值
				dkey1 :=key1+"d1"
				dimap1 := map[string]string{"200":imap_0["101"],"204":imap_4["101"],"208":imap_8["101"],"212":imap_12["101"],"216":imap6["101"],"316":imap6["301"]}
				dvalue1, _ := json.Marshal(dimap1)
				_, err = rs.Do("SET", dkey1, dvalue1)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点1路电流")
				
				dkey2 :=key1+"d2"
				dimap2 := map[string]string{"200":imap_0["102"],"204":imap_4["102"],"208":imap_8["102"],"212":imap_12["102"],"216":imap6["102"],"316":imap6["302"]}
				dvalue2, _ := json.Marshal(dimap2)
				_, err = rs.Do("SET", dkey2, dvalue2)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点2路电流")
				
				dkey3 :=key1+"d3"
				dimap3 := map[string]string{"200":imap_0["103"],"204":imap_4["103"],"208":imap_8["103"],"212":imap_12["103"],"216":imap6["103"],"316":imap6["303"]}
				dvalue3, _ := json.Marshal(dimap3)
				_, err = rs.Do("SET", dkey3, dvalue3)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点3路电流")
				
				dkey4 :=key1+"d4"
				dimap4 := map[string]string{"200":imap_0["104"],"204":imap_4["104"],"208":imap_8["104"],"212":imap_12["104"],"216":imap6["104"],"316":imap6["304"]}
				dvalue4, _ := json.Marshal(dimap4)
				_, err = rs.Do("SET", dkey4, dvalue4)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点4路电流")
				
				dkey5 :=key1+"d5"
				dimap5 := map[string]string{"200":imap_0["105"],"204":imap_4["105"],"208":imap_8["105"],"212":imap_12["105"],"216":imap6["105"],"316":imap6["305"]}
				dvalue5, _ := json.Marshal(dimap5)
				_, err = rs.Do("SET", dkey5, dvalue5)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点5路电流")
				
				dkey6 :=key1+"d6"
				dimap6 := map[string]string{"200":imap_0["106"],"204":imap_4["106"],"208":imap_8["106"],"212":imap_12["106"],"216":imap6["106"],"316":imap6["306"]}
				dvalue6, _ := json.Marshal(dimap6)
				_, err = rs.Do("SET", dkey6, dvalue6)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点6路电流")
				
				dkey7 :=key1+"d7"
				dimap7 := map[string]string{"200":imap_0["107"],"204":imap_4["107"],"208":imap_8["107"],"212":imap_12["107"],"216":imap6["107"],"316":imap6["307"]}
				dvalue7, _ := json.Marshal(dimap7)
				_, err = rs.Do("SET", dkey7, dvalue7)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点7路电流")
				
				dkey8 :=key1+"d8"
				dimap8 := map[string]string{"200":imap_0["108"],"204":imap_4["108"],"208":imap_8["108"],"212":imap_12["108"],"216":imap6["108"],"316":imap6["308"]}
				dvalue8, _ := json.Marshal(dimap8)
				_, err = rs.Do("SET", dkey8, dvalue8)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点8路电流")
				
				dkey9 :=key1+"d9"
				dimap9 := map[string]string{"200":imap_0["109"],"204":imap_4["109"],"208":imap_8["109"],"212":imap_12["109"],"216":imap6["109"],"316":imap6["309"]}
				dvalue9, _ := json.Marshal(dimap9)
				_, err = rs.Do("SET", dkey9, dvalue9)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点9路温度")
				
				dkey10 :=key1+"d10"
				dimap10 := map[string]string{"200":imap_0["110"],"204":imap_4["110"],"208":imap_8["110"],"212":imap_12["110"],"216":imap6["110"],"316":imap6["310"]}
				dvalue10, _ := json.Marshal(dimap10)
				_, err = rs.Do("SET", dkey10, dvalue10)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点10路温度")
				
				dkey11 :=key1+"d11"
				dimap11 := map[string]string{"200":imap_0["111"],"204":imap_4["111"],"208":imap_8["111"],"212":imap_12["111"],"216":imap6["111"],"316":imap6["311"]}
				dvalue11, _ := json.Marshal(dimap11)
				_, err = rs.Do("SET", dkey11, dvalue11)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点11路温度")
				
				dkey12 :=key1+"d12"
				dimap12 := map[string]string{"200":imap_0["112"],"204":imap_4["112"],"208":imap_8["112"],"212":imap_12["112"],"216":imap6["112"],"316":imap6["312"]}
				dvalue12, _ := json.Marshal(dimap12)
				_, err = rs.Do("SET", dkey12, dvalue12)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("16点12路温度")


				time.Sleep(time.Second*60)








			case "20" :
//存入20点即时数据
				key7 :=key1+"t20"
				imap7 := map[string]string{"101": imap["101"], "102": imap["102"], "103": imap["103"],"104": imap["104"],"105": imap["105"],"106": imap["106"],"107": imap["107"],"108": imap["108"],"109": imap["109"],"110": imap["110"],"111": imap["111"],"112": imap["112"],"301": imap["301"],"302": imap["302"],"303": imap["303"],"304": imap["304"],"305": imap["305"],"306": imap["306"],"307": imap["307"],"308": imap["308"],"309": imap["309"],"310": imap["310"],"311": imap["311"],"312": imap["312"]}
				value7, _ := json.Marshal(imap7)
				_, err := rs.Do("SETNX", key7, value7)
				if err != nil {  
					fmt.Println(err)  
				}  
				_, _ = rs.Do("EXPIRE", key7, 24*3600)  
				fmt.Println("success",key7)
//取20点前各回路数值（0点、4点、8点、12点、16点）				
			
			var imap_0 map[string]string
				key_0 := key1+"t0"
				value_0,err := redis.Bytes(rs.Do("GET", key_0))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal := json.Unmarshal(value_0, &imap_0)  
				if errShal != nil {  
					fmt.Println(err)  
				} 	
			var imap_4 map[string]string
				key_4 := key1+"t4"
				value_4,err := redis.Bytes(rs.Do("GET", key_4))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal2 := json.Unmarshal(value_4, &imap_4)  
				if errShal2 != nil {  
					fmt.Println(err)  
				}
			var imap_8 map[string]string
				key_8 := key1+"t8"
				value_8,err := redis.Bytes(rs.Do("GET", key_8))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal3 := json.Unmarshal(value_8, &imap_8)  
				if errShal3 != nil {  
					fmt.Println(err)  
				}
			var imap_12 map[string]string
				key_12 := key1+"t12"
				value_12,err := redis.Bytes(rs.Do("GET", key_12))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal4 := json.Unmarshal(value_12, &imap_12)  
				if errShal4 != nil {  
					fmt.Println(err)
				}
			
			var imap_16 map[string]string
				key_16 := key1+"t16"
				value_16,err := redis.Bytes(rs.Do("GET", key_16))
				if err != nil {  
					fmt.Println(err)  
				}
				errShal5 := json.Unmarshal(value_16, &imap_16)  
				if errShal5 != nil {  
					fmt.Println(err)
				}
			
//存入20点数值各回路历史值
				dkey1 :=key1+"d1"
				dimap1 := map[string]string{"200":imap_0["101"],"204":imap_4["101"],"208":imap_8["101"],"212":imap_12["101"],"216":imap_16["101"],"220":imap7["101"],"320":imap7["301"]}
				dvalue1, _ := json.Marshal(dimap1)
				_, err = rs.Do("SET", dkey1, dvalue1)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点1路电流")
				
				dkey2 :=key1+"d2"
				dimap2 := map[string]string{"200":imap_0["102"],"204":imap_4["102"],"208":imap_8["102"],"212":imap_12["102"],"216":imap_16["102"],"220":imap7["102"],"320":imap7["302"]}
				dvalue2, _ := json.Marshal(dimap2)
				_, err = rs.Do("SET", dkey2, dvalue2)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点2路电流")
				
				dkey3 :=key1+"d3"
				dimap3 := map[string]string{"200":imap_0["103"],"204":imap_4["103"],"208":imap_8["103"],"212":imap_12["103"],"216":imap_16["103"],"220":imap7["103"],"320":imap7["303"]}
				dvalue3, _ := json.Marshal(dimap3)
				_, err = rs.Do("SET", dkey3, dvalue3)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点3路电流")
				
				dkey4 :=key1+"d4"
				dimap4 := map[string]string{"200":imap_0["104"],"204":imap_4["104"],"208":imap_8["104"],"212":imap_12["104"],"216":imap_16["104"],"220":imap7["104"],"320":imap7["304"]}
				dvalue4, _ := json.Marshal(dimap4)
				_, err = rs.Do("SET", dkey4, dvalue4)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点4路电流")
				
				dkey5 :=key1+"d5"
				dimap5 := map[string]string{"200":imap_0["105"],"204":imap_4["105"],"208":imap_8["105"],"212":imap_12["105"],"216":imap_16["105"],"220":imap7["105"],"320":imap7["305"]}
				dvalue5, _ := json.Marshal(dimap5)
				_, err = rs.Do("SET", dkey5, dvalue5)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点5路电流")
				
				dkey6 :=key1+"d6"
				dimap6 := map[string]string{"200":imap_0["106"],"204":imap_4["106"],"208":imap_8["106"],"212":imap_12["106"],"216":imap_16["106"],"220":imap7["106"],"320":imap7["306"]}
				dvalue6, _ := json.Marshal(dimap6)
				_, err = rs.Do("SET", dkey6, dvalue6)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点6路电流")
				
				dkey7 :=key1+"d7"
				dimap7 := map[string]string{"200":imap_0["107"],"204":imap_4["107"],"208":imap_8["107"],"212":imap_12["107"],"216":imap_16["107"],"220":imap7["107"],"320":imap7["307"]}
				dvalue7, _ := json.Marshal(dimap7)
				_, err = rs.Do("SET", dkey7, dvalue7)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点7路电流")
				
				dkey8 :=key1+"d8"
				dimap8 := map[string]string{"200":imap_0["108"],"204":imap_4["108"],"208":imap_8["108"],"212":imap_12["108"],"216":imap_16["108"],"220":imap7["108"],"320":imap7["308"]}
				dvalue8, _ := json.Marshal(dimap8)
				_, err = rs.Do("SET", dkey8, dvalue8)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点8路电流")
				
				dkey9 :=key1+"d9"
				dimap9 := map[string]string{"200":imap_0["109"],"204":imap_4["109"],"208":imap_8["109"],"212":imap_12["109"],"216":imap_16["109"],"220":imap7["109"],"320":imap7["309"]}
				dvalue9, _ := json.Marshal(dimap9)
				_, err = rs.Do("SET", dkey9, dvalue9)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点9路温度")
				
				dkey10 :=key1+"d10"
				dimap10 := map[string]string{"200":imap_0["110"],"204":imap_4["110"],"208":imap_8["110"],"212":imap_12["110"],"216":imap_16["110"],"220":imap7["110"],"320":imap7["310"]}
				dvalue10, _ := json.Marshal(dimap10)
				_, err = rs.Do("SET", dkey10, dvalue10)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点10路温度")
				
				dkey11 :=key1+"d11"
				dimap11 := map[string]string{"200":imap_0["111"],"204":imap_4["111"],"208":imap_8["111"],"212":imap_12["111"],"216":imap_16["111"],"220":imap7["111"],"320":imap7["311"]}
				dvalue11, _ := json.Marshal(dimap11)
				_, err = rs.Do("SET", dkey11, dvalue11)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点11路温度")
				
				dkey12 :=key1+"d12"
				dimap12 := map[string]string{"200":imap_0["112"],"204":imap_4["112"],"208":imap_8["112"],"212":imap_12["112"],"216":imap_16["112"],"220":imap7["112"],"320":imap7["312"]}
				dvalue12, _ := json.Marshal(dimap12)
				_, err = rs.Do("SET", dkey12, dvalue12)  
				if err != nil {  
					fmt.Println(err)  
				} 
				fmt.Println("20点12路温度")
				time.Sleep(time.Second*60)
			}
		}
	}
}

