package authgroup

import (
	"fmt"
	"os"
	"timeping/internal/tlog"
	"timeping/pkg/ostools"
)

//顺序读取校验总auth文件,使用os read
func init_authdb() error {
	//校验文件是否存在
	if !ostools.FileExists("./auth/auth.tp") {
		tlog.Common("no authdb file","Warning","auth")
		tlog.Err_in("no authdb file","Warning","auth")
		if err:=create_allauth();err!=nil{
			return err
		}
		return nil
	}
	var err error
	auall,err=os.Open("./auth/auth.tp")
	if err!=nil{
		tlog.Common("open authdb file fail","Error","auth")
		tlog.Err_in("open authdb file fail","Error","auth")
		return fmt.Errorf("open authdb file fail")
	}
	
    return nil
}

//释放资源
func kill_authdb() error {
	if auall==nil{
		return nil
	}
	if err:=auall.Close();err!=nil{
		tlog.Common("close authdb file fail","Error","auth")
		tlog.Err_in("close authdb file fail","Error","auth")
		return fmt.Errorf("close authdb file fail")
	}
	return nil
}