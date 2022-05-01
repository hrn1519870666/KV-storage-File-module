/**
  @author: 黄睿楠
  @since: 2022/3/16
  @desc: 文件备份模块
**/

package fileModule

import (
	"fmt"
	"github.com/klauspost/reedsolomon"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (file *File) Encoder() {
	// 生成编码矩阵
	enc, err := reedsolomon.New(file.dataShards, file.parShards)
	checkErr(err)

	fmt.Println("Opening", file.fileName)
	inputFile, err := ioutil.ReadFile(file.fileName)
	checkErr(err)

	// 将inputFile分割成dataShards个分片，并创建parShards个空的校验分片
	shards, err := enc.Split(inputFile)
	checkErr(err)
	fmt.Printf("File split into %d data+parity shards with %d bytes per shard.\n", len(shards), len(shards[0]))

	// 生成校验分片，enc是编码矩阵
	err = enc.Encode(shards)
	checkErr(err)

	// 根据最后一个路径分隔符将路径path分隔为目录和文件名两部分（dir 和 file）
	_, realFileName := filepath.Split(file.fileName)
	// 输出目录
	dir := "./RS_Encode"
	for i, shard := range shards {
		// output为string类型，此处的output要和decoder中的input一一对应
		output := fmt.Sprintf("%d_%s", i, realFileName)
		fmt.Println("Writing to", output)
		// 将分片内容写入到文件中
		err = ioutil.WriteFile(filepath.Join(dir, output), shard, os.ModePerm)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(2)
	}
}
