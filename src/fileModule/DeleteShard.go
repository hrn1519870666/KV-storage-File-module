/**
  @author: 黄睿楠
  @since: 2022/3/16
  @desc： 模拟文件分片丢失
**/

package fileModule

import "os"

func (file *File) DeleteShard(fileName string) {
	err1 := os.Remove(fileName)
	checkErr(err1)
}
