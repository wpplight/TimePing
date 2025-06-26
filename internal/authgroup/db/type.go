package db

import "os"

//每一页大小为2KB

// 用户表项 16B 2^7项
type UsrItem struct {
	Id   uint16 //2B 用户编号
	Used uint8 //1B 使用状态
	Page uint16 //2B 页码
	Set  uint8 //1B 页内偏移
	Name string // 10B 用户名，长度11
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
	usrmmap map[uint16] string
)