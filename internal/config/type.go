package config

type cnf struct{
	Timeinterval uint16//时间轮的间隔
	TaskPoolSize uint16 //任务池的大小
	TimeWheelSize uint16//时间轮的大小
	Port  uint16//使用端口
	Timelevel int
};

var (
	Conf cnf
)