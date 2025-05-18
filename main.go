package main

import (
	"TimePing/config"
	"fmt"
	"time"
	"TimePing/config"

	"github.com/prometheus/alertmanager/config"
	"golang.org/x/tools/playground/socket"
)


	
func main() {
	TaskPool := make([]TaskNode, config.TaskPoolSize)
	currentTask:=&TaskPool[0]
	TimePing := make([]TaskNode, config.TimeWheelSize)
	go func(){
		currentIndex :=0
		ticker:=time.NewTicker(time.Second*time.Duration(config.Timeinterval))
		defer ticker.Stop()
		for currentIndex<config.TimeWheelSize{
			select{
			case <-ticker.C:
				if &TimePing[currentIndex].time<=time.Now()+time.Duration(config.TimeInterval)/2 && &TimePing[currentIndex].time>=time.Now()-time.Duration(config.TimeInterval)/2{
					fmt.Println("执行任务")
					//待补充
				}
				currentIndex++
				currentIndex=currentIndex%config.TimeWheelSize
				
			}
		}
	}

}