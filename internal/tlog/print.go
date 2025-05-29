package tlog
import "fmt"

func Common(message string, tag ...string){
	fmt.Print("[TimePing]")
	for _,v:=range tag{
		fmt.Print("[",v,"]  ")
	}
	fmt.Println(message)
}