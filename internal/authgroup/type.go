package authgroup

import (
	"timeping/internal/authgroup/jumptable"
)




var(
	//用户系统核心对象
	ag struct{
		maxid uint16
		AuthNum int 
		Uidtable *jumptable.JumpTable
	}
	//总的用户管理文件
	

)