# KV存储系统（单节点版）

工程目录：

/bin：代码生成的.exe。

/data：数据集、测试集。

/levelDB：数据库文件。

/RS_Decode:文件恢复模块的生成文件。

/RS_Encode：文件备份模块的生成文件，也就是数据切片和校验切片。

/src/dao：对数据库的操作，包括对levelDB的增删改查等操作，以及打开数据库操作。

/src/service：service层调用dao层，实现具体业务，包括：PUT数据集、根据日志文件恢复出数据库信息。

/src/fileModule：文件模块，包括文件复制，备份，恢复。

/src/testData: 保存文件复制模块生成的复制文件。