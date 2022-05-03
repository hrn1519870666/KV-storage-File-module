# KV存储系统-文件模块

工程目录：

/data：数据集、测试集。

/levelDB：levelDB数据库中的数据。

/RS_Decode:文件恢复模块的生成文件。

/RS_Encode：文件备份模块的生成文件，也就是数据切片和校验切片。

/src/dao：对levelDB的增删改查等操作，以及打开数据库操作。

/src/service：PUT数据集、根据日志文件恢复出数据库信息。

/src/fileModule：文件模块，包括文件复制，备份，恢复。

/src/testData: 保存文件复制模块生成的复制文件。