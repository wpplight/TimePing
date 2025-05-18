package ostools
import (
    "os"
)
func FileExists(path string) bool {
    if _, err := os.Stat(path); err == nil {
        return true  // 文件存在
    } else if os.IsNotExist(err) {
        return false // 文件不存在
    } else {
        // 其他错误（权限、IO 错误等），根据需要自行处理
        return false
    }
}

func CreateFile(path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    return nil
}


