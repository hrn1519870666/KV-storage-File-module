/**
  @author: 黄睿楠
  @since: 2022/3/24
  @desc:文件模块测试
**/

package fileModule

import (
	"fmt"
	"testing"
)

// 解决go test文件路径问题：修改build的working directory为项目根目录
func TestCopyFile(t *testing.T) {

	/*
		../ 表示当前所处文件夹的上一级文件夹的绝对路径
		./ 表示当前所处文件夹的绝对路径
	*/

	file := NewFile(4, 2, "./data/data.txt", "./testData/testData.txt")
	written, err := file.CopyFile()
	fmt.Println(written)
	if err != nil {
		t.Errorf("CopyFile failed!")
	}
}

func TestFile_Encoder(t *testing.T) {
	file := NewFile(4, 2, "./data/data.txt", "./testData/testData.txt")
	file.Encoder()
}

func TestDecoderSuccess(t *testing.T) {

	file := NewFile(4, 2, "./data/data.txt", "./testData/testData.txt")
	file.Encoder()
	file.DeleteShard("./RS_Encode/0_data.txt")
	file.DeleteShard("./RS_Encode/4_data.txt")
	file.Decoder()
}

func TestDecoderFail(t *testing.T) {

	file := NewFile(4, 2, "./data/data.txt", "./testData/testData.txt")
	file.Encoder()
	// 删除了3个分片，所以本测试不通过
	file.DeleteShard("./RS_Encode/0_data.txt")
	file.DeleteShard("./RS_Encode/1_data.txt")
	file.DeleteShard("./RS_Encode/4_data.txt")
	file.Decoder()
}
