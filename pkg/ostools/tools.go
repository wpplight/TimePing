package ostools
import (
    "os"
)

// 判断文件是否可以打开,可以则返回true
func FileExists(path string) bool {
    if _, err := os.Stat(path); err == nil {
        return true  // 文件存在
    } 
    return false
}

func OpenFile(path string) (*os.File, error) {
    file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    return file, nil
}

func CreateFile(path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    return nil
}


