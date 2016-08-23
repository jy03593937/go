package main
import (
    "fmt"
    
	"time"
	
)
func main(){
		
		t := time.Now()
		
		fmt.Println(t.Hour())
		fmt.Println(t.Minute())
		Hour1 :=fmt.Sprintf("%d",t.Hour())
		fmt.Println(Hour1)
		switch Hour1{
		case "1" :
			fmt.Println("yes")
			fallthrough
		case "2" :
			fmt.Println("no")
		
		}
		
}