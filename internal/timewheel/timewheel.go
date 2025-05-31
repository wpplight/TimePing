package timewheel

import (
	"container/list"
	"log"
	"time"
	"timeping/internal/config"
	"timeping/internal/task"
	"timeping/pkg/tlist"
)

var Tw []*tlist.Tlist
func InitialTimeWheel() {
	Tw=make([]*tlist.Tlist,config.Conf.TimeWheelSize)
	for i:=0;i<int(config.Conf.TimeWheelSize);i++{
		for j:=0;j<config.Conf.Timelevel;j++{		
			L:=list.New()
		}
	}
}


func TimeTicker() {
	ticker := time.NewTicker(time.Duration(config.Conf.Timeinterval) * time.Second)
	defer ticker.Stop()
	var index uint16=0;
	for {
		select {
		case <-ticker.C:
			//判断是否是当前轮的节点，如果是则运行,待写入
			index++;
			if(index==config.Conf.TimeWheelSize){
				index=0;
			}
		}
	}
}
