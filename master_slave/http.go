/**
  @author: 黄睿楠
  @since: 2022/4/30
  @desc: API接口
**/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func (node *Node) write() {
	r := gin.Default()

	// http://localhost:9000/write?filename=xxx
	r.POST("/write", func(c *gin.Context) {
		// 从URL中获取文件名
		filename := c.Query("filename")

		// 从原有Request.Body读取
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)

		fileA,_ := os.Create("./dbA/"+filename)
		defer fileA.Close()
		// 将文件存入dbA
		_ = ioutil.WriteFile("./dbA/"+filename,bodyBytes,0777)

		_,err := CopyFile("./dbB/"+filename,"./dbA/"+filename)
		if err != nil {
			fmt.Println("主从复制失败:", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"filename": filename,
		})
	})
	r.Run(":9000")
}



func (node *Node) read() {
	r := gin.Default()
	//http://localhost:9001/read?filename=xxx
	r.GET("/read", func(c *gin.Context) {
		filename := c.Query("filename")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"filename": filename,
		})
		// TODO   读取操作

	})
	r.Run(":9001")
}

