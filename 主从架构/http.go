/**
  @author: 黄睿楠
  @since: 2022/4/30
  @desc: API接口
**/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (node *Node) write() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	// http://localhost:9000/index
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MB
	r.POST("/write", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dstA := fmt.Sprintf("./dbA/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dstA)

		// 主从复制
		dstB := fmt.Sprintf("./dbB/%s", file.Filename)
		c.SaveUploadedFile(file, dstB)

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' 上传成功!", file.Filename),
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

