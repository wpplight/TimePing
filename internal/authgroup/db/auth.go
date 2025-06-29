package db

import (
	"errors"
	"timeping/internal/tlog"
	"timeping/internal/bufferpool"
)

//设置用户长度
func setAuthLen(n uint16) error {
	if _,err:=auth.Seek(0,0);err!=nil{
		tlog.Common("查找auth文件错误","Error","auth")
		tlog.Err_in("auth file seek fail"+err.Error(),"Error","auth")
		return errors.New("auth file seek fail")
	}
	_,err:=auth.Write(uint16toByte(n))
	if err!=nil{
		tlog.Common("写入auth文件错误","Error","auth")
		tlog.Err_in("auth file write fail"+err.Error(),"Error","auth")
		return errors.New("auth file write fail")
	}
	if err=auth.Sync() ;err!=nil{
		tlog.Common("同步auth文件错误","Error","auth")
		tlog.Err_in("auth file sync fail"+err.Error(),"Error","auth")
		return errors.New("auth file sync fail")
	}
	return nil

}

//获取用户总id长度，包含懒删除的用户
func getAuthLen() (uint16,error) { 
	if auth==nil{
		tlog.Log_all("auth file is nil","Error","auth")
		return 0,errors.New("auth file is nil")
	}
	_,err:=auth.Seek(0,0)
	if err!=nil{
		tlog.Common("auth file seek fail","Error","auth")
		tlog.Err_in("auth file seek fail"+err.Error(),"Error","auth")
		return 0,errors.New("auth file seek fail")
	}
	b:=make([]byte,2)
	_,err=auth.Read(b)
	if err!=nil{
		tlog.Common("auth file read fail","Error","auth")
		tlog.Err_in("auth file read fail"+err.Error(),"Error","auth")
		return 0,errors.New("auth file read fail")
	}
	return byteToUint16(b),nil
}



//将用户表单载入到内存中
func loadUsr() error{ 
	//初始检测
	if usrtable==nil{
		tlog.Log_all("usr table is nil","Error","auth")
		return errors.New("usr table is nil")
	}
	if asize==0{
		tlog.Log_all("auth size is zero","Error","auth")
		return errors.New("auth size is zero")
	}

	//索引与文件偏移
	index,offset:=uint16(0),16	
	//缓冲池块分配
	blk:=bufferpool.Buffer.Pool1k.Get().([]byte)
	defer bufferpool.Buffer.Pool1k.Put(blk)

	//auth文件解析
	for index<asize{ 

		//块载入
		if err:=loadBlock(auth,int64(offset),blk);err!=nil{
			tlog.Common("auth file block analasy fail","Error","authdb")
			tlog.Err_in("auth file block analasy fail"+err.Error(),"Error","authdb")
			return errors.New("auth file block analasy fail")
		}
		//下一次块读取位置
		offset+=1<<10

		//块解析
		num:=1<<10
		for i:=0;i<num;i+=16{ 

			//判断是否懒删除跳过解析
			if blk[i+2]&0x80 == 0x80{
				//进行表项解析
				usrtable[index]=byteToUsrItem(blk[i:i+16])
			}else{
				usrtable[index].Used=false
			}
			index++
			if index>=asize{
				break
			}
		}
	}
	return nil
}


//创建新用户
func CreateAuth(usrname string ,pass string) error{

	//账户密码检测
	if len(pass)>15{
		return errors.New("pass to long")
	}
	if len(usrname)>12{
		return errors.New("usrname to long")
	}

	//用户表更新
	asize++
	usrtable=append(usrtable,UsrItem{Name:usrname,Page:0,Set:0,Used:true})

	//获取一个小块
	blk:=bufferpool.Buffer.Pool16.Get().([]byte)
	defer bufferpool.Buffer.Pool16.Put(blk)

	

}