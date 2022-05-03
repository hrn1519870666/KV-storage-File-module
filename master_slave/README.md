## 文档

### 1.克隆本地

```bash
git clone https://github.com/hrn1519870666/KV-storage-File-module.git
cd master_slave
```



### 2.编译运行

#### 主节点

在master_slave下开启cmd，执行以下命令：

```bash
// 编译
go build -o node.exe

// 执行
node.exe A
```

开始监听9000写端口，默认监听10分钟。

#### 从节点

在master_slave下开启另一个cmd，执行以下命令：

```ba
node.exe B
```

开始监听9001读端口，默认监听10分钟。



### 3.测试

#### 写端口：

postman：

<a href="https://sm.ms/image/q1SamCdcjUOpztR" target="_blank"><img src="https://s2.loli.net/2022/05/03/q1SamCdcjUOpztR.png" style="zoom: 50%;"  ></a>



写入的文件将分别保存在dbA和dbB下：

<a href="https://sm.ms/image/7F2mMiEAyQjYdV1" target="_blank"><img src="https://s2.loli.net/2022/05/03/7F2mMiEAyQjYdV1.png" style="zoom: 50%;"  ></a>



#### 读端口：

<a href="https://sm.ms/image/sPSnBlI3RkiM8QL" target="_blank"><img src="https://s2.loli.net/2022/05/03/sPSnBlI3RkiM8QL.png" style="zoom:50%;"  ></a>