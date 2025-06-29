package tlog
import "log"
//输入报错内容和多个标签
func Common(message string, tag ...string){
	s:=""
	for _,v:=range tag{
		s+="["+v+"]"
	}
	log.Println(s+message)
}