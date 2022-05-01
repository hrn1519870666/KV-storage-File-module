/**
  @author: 黄睿楠
  @since: 2022/3/16
  @desc: 文件恢复模块
**/

package fileModule

import (
	"fmt"
	"github.com/klauspost/reedsolomon"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (file *File) Decoder() {
	enc, err := reedsolomon.New(file.dataShards, file.parShards)
	checkErr(err)
	shards := make([][]byte, file.dataShards+file.parShards)
	_, realFileName := filepath.Split(file.fileName)

	for i := range shards {
		input := fmt.Sprintf("%d_%s", i, realFileName)
		fmt.Println("Opening", input)
		// 将i_data中的内容读入shards[i]中
		shards[i], err = ioutil.ReadFile("./RS_Encode/" + input)
		if err != nil {
			fmt.Println("Error reading file", err)
			shards[i] = nil
		}
	}

	ok, err := enc.Verify(shards)
	if ok {
		fmt.Println("No reconstruction needed")
	} else {
		fmt.Println("Verification failed. Reconstructing data")
		err = enc.Reconstruct(shards)
		if err != nil {
			fmt.Println("Reconstruct failed -", err)
			os.Exit(1)
		}
		ok, err = enc.Verify(shards)
		if !ok {
			fmt.Println("Verification failed after reconstruction, data likely corrupted.")
			os.Exit(1)
		}
		checkErr(err)
	}

	output := "./RS_Decode/reconstructData"
	fmt.Println("Writing data to", output)
	f, err := os.Create(output)
	checkErr(err)

	// We don't know the exact filesize.
	err = enc.Join(f, shards, len(shards[0])*file.dataShards)
	checkErr(err)
}
