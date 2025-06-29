package db

import (
	"errors"
	"os"
	"timeping/internal/tlog"
	"timeping/internal/bufferpool"
	"timeping/pkg/ostools"
)

//密码加密影子文件

//shadow 文件 RowSize = 16

func createShadow() error{
	p:=path+"shadow.tp"
	err:=ostools.CreateFile(p)
	if err!=nil{
		return err
	}
	return nil
}

func changePassStatus(offset int64,status []byte) error{ 
	if _,err:=shadow.Seek(offset,0);err!=nil{
		return err
	}
	if _,err:=shadow.Write(status);err!=nil{
		return err
	}
	return nil
}

//检查密码文件匹配程度
func checkPass() error{
	//安全性审核
	if usrtable==nil {
		tlog.Log_all("user table is nil","Error","auth")
		return errors.New("user table is nil")
	}
	if shadow==nil {
		tlog.Log_all("shadow file is nil","Error","auth")
		return errors.New("shadow file is nil")
	}

	//缓冲buffer池中获取缓冲块
	blk:=bufferpool.Buffer.Pool1k.Get().([]byte)
	defer bufferpool.Buffer.Pool1k.Put(blk)

	//使用缓冲块加载
	offset,len:=int64(0),1<<6
	if err:=loadBlock(shadow,offset,blk);err!=nil{
		tlog.Log_all("shadow file block load error"+err.Error(),"Error","auth")
		return err
	}

	for i,v:=range usrtable{
		//判断是否未使用
		if v.Used==false{
			continue
		}
		//判断是否命中缓存
		for i>=len{
			offset+=1<<10
			if err:=loadBlock(shadow,offset,blk);err!=nil{
				tlog.Log_all("shadow file block load error"+err.Error(),"Error","auth")
				return err
			}
			len+=1<<6
		}
		//处理状态位
		if blk[i%64]&0x80==0{
			blk[i%64]|=0x80
			if err:=changePassStatus(offset+int64(i%64),blk[i%64:i%64+1]);err!=nil{
				tlog.Common("用户修改状态错误")
				tlog.Err_in("用户修改状态错误"+err.Error(),"Error","auth")
				return err
			}
			tlog.Common("用户状态错误","Error","auth")
			tlog.Err_in("用户状态错误,"+string(i)+"与密码文件不同步,已fixed","Error","auth")
		}
		

	}
	return nil

}



//载入用户文件并校验密码文件完整性
func authLoadCheck() error {


	//用户文件尺寸读取
	Asize,err:=os.Stat(path+"auth.tp")
	if err!=nil{
		tlog.Common("auth file is not exist","Error","auth")
		tlog.Err_in("auth file is not exist"+err.Error(),"Error","auth")
		return errors.New("auth file is not exist")
	}
	//影子文件尺寸读取
	Ssize,err:=os.Stat(path+"shadow.tp")
	if err!=nil{
		tlog.Common("shadow file is not exist","Error","auth")
		tlog.Err_in("shadow file is not exist"+err.Error(),"Error","auth")
		return errors.New("shadow file is not exist")
	}

	asize=0
	//检测是否init过
	if Asize.Size()==0{
		if err=setAuthLen(0);err!=nil{
			tlog.Common("auth file init error","Error","auth")
			tlog.Err_in("auth file init error"+err.Error(),"Error","auth")
			return err
		}
		if Ssize.Size() == 0{
			//该软件还没有初始化
			return nil
		}
	}

	//如果读取的密码文件尺寸为0，一定是密码文件数据缺失
	if Ssize.Size()==0{
		tlog.Log_all("文件数据不匹配，密码文件为空，数据缺失","Error","authdb")
		return errors.New("no use")
	}

	//读取用户
	asize,err=getAuthLen()
	if err!=nil{
		tlog.Log_all("auth file get length error","Error","authdb")
		return err
	}
	//分配用户表
	usrtable=make([]UsrItem,asize,max(80,asize+40))
	//载入用户
	if err=loadUsr();err!=nil{
		tlog.Log_all("auth file load error","Error","authdb")
		return err
	}

	return nil
}