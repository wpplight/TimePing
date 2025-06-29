package db

import (
	"errors"
	"os"
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

func byteToUsrItem(b []byte) UsrItem{ 
	l:=uint8(b[0] & 0x7F)
	return UsrItem{
		Used: b[0]>>7 == 1,
		Page: uint8(b[1]),
		Set: uint16(byteToUint16(b[3:5])),
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
	
	// 使用状况
	if usr.Used{
		ans[0]=1<<7|byte(len(usr.Name))
	}else{
		ans[0]=0|byte(len(usr.Name))
	}
	
	//页号
	ans[1]=byte(usr.Page)
	//页内偏移
	ans[2]=byte(usr.Set>>8)
	ans[3]=byte(usr.Set)
	//用户名
	 l:=len(usr.Name)
	 for i:=0;i<l;i++{
		 ans[i+4]=usr.Name[i]
	 }
	 return ans,nil	
}

// 将文件中指定大小的数据读入块中
func loadBlock(file *os.File,offset int64,blk []byte) error {

	if _,err:=file.Seek(int64(offset),0);err!=nil{
		return err
	}
	if _,err:=file.Read(blk);err!=nil{
		return err
	}
	return nil
}

// 将块写入文件
func writeBlock(file *os.File,offset int64,blk []byte) error { 
	if _,err:=file.Seek(int64(offset),0);err!=nil{
		return err
	}
	if _,err:=file.Write(blk);err!=nil{
		return err
	}
	return nil
}