package authgroup

import "os"




var(
	//用户系统核心对象
	ag struct{
		maxid uint16
		AuthNum int 
	}
	//总的用户管理文件
	auall *os.File
)