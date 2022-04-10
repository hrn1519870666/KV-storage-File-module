/**
  @author: 黄睿楠
  @since: 2022/3/24
  @desc: 文件类
**/

package fileModule

import "fmt"

type File struct {
	dataShards int    // 数据分片数
	parShards  int    // 校验分片数
	fileName   string // 文件全路径名
	copyName   string // 备份文件全路径名
}

// 结构体工厂
func NewFile(dataShards int, parShards int, fileName string, copyName string) *File {
	if dataShards <= 0 || parShards <= 0 {
		// err,panic
		fmt.Errorf("数据分片数和校验分片数必须大于0")
		return nil
	}

	if parShards >= dataShards {
		fmt.Errorf("校验分片数必须小于数据分片数")
		return nil
	}

	return &File{dataShards, parShards, fileName, copyName}
}
