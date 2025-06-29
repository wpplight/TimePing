package db

import "os"

//每一页大小为2KB

// 用户表项 16B 2^7项
type UsrItem struct {
	Used bool //1B 使用状态
	Page uint8 //1B 页码
	Set  uint16 //2B 页内偏移
	Name string // 12B 用户名，长度12
}

var(
	// 数据库文件路径
	path string
	// 数据库文件句柄
	auth *os.File
	shadow *os.File
	authindex *os.File
	datafile *os.File
	// auth索引用户数量
	asize uint16
	// 用户表项
	usrtable []UsrItem
	logintable map[string] uint16
)