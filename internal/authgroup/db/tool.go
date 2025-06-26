package db

import (
	"errors"
	"timeping/internal/tlog"
	"timeping/pkg/ostools"
)

// auth.tp用户身份文件
// shadow.tp密码文件
// index.tp 索引文件
// data.tp 数据文件

func checkFileExit() error {
	if !ostools.FileExists(path+"auth.tp"){
		tlog.Common("没有用户总文件","Warning","auth")
		tlog.Err_in("没有用户总文件","Warning","auth")
		err:=create_allauth()
		if err!=nil{
			return err
		}
		return errors.New("没有用户总文件")
	} 

	if !ostools.FileExists(path+"shadow.tp"){
		tlog.Common("没有密码文件","Warning","auth")
		tlog.Err_in("没有密码文件","Warning","auth")
		err:=createShadow()
		if err!=nil{
			return err
		}
	}

	if !ostools.FileExists(path+"data.tp"){
		tlog.Common("没有用户数据文件","Warning","auth")
		tlog.Err_in("没有用户数据文件","Warning","auth")
		err:=createData()
		if err!=nil{
			return err
		}
	}

	return nil
}


func uint16toByte(num uint16) []byte {
	b := make([]byte, 2)
	b[0] = byte(num >> 8)
	b[1] = byte(num)
	return b
}

func byteToUint16(b []byte) uint16 {
	return uint16(b[0])<<8 | uint16(b[1])
}

func byteToUsrItem(b []byte) *UsrItem{ 
	l:=uint8(b[2] & 0x7F)
	return &UsrItem{
		Id:byteToUint16(b[0:2]),
		Used: b[2]>>7,
		Page: byteToUint16(b[3:5]),
		Set: uint8(b[5]),
		Name: string(b[6:6+l]),
	}
}


func usrItemToByte(usr *UsrItem) ([]byte,error){
	ans:=make([]byte,16)
	if usr==nil{
		tlog.Common("usr is nil","Error","auth")
		tlog.Err_in("usr is nil","Error","auth")
		return nil,errors.New("usr is nil")
	}
	
	// id
	ans[0]=byte(usr.Id>>8)
	ans[1]=byte(usr.Id)
	// 使用状况
	u:=usr.Used
	ans[2]=byte(u)
	//页号
	ans[3]=byte(usr.Page>>8)
	ans[4]=byte(usr.Set)
	//页内偏移
	ans[5]=byte(usr.Set)
	//用户名
	 l:=len(usr.Name)
	 for i:=0;i<l;i++{
		 ans[i+6]=usr.Name[i]
	 }
	 return ans,nil	
}