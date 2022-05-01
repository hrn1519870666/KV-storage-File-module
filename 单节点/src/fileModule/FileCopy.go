/**
  @author: 黄睿楠
  @since: 2022/3/24
  @desc: 文件复制
**/

package fileModule

import (
	"io"
	"os"
)

// written int64:复制到的字节数
func (file *File) CopyFile() (written int64, err error) {
	src, err := os.Open(file.fileName)
	if err != nil {
		return
	}
	defer src.Close()

	// 写入的文件需要有权限
	dst, err := os.OpenFile(file.copyName, os.O_RDWR, 0777)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
