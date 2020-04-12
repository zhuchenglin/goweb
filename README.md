# goweb
使用go语言封装一个简单的api框架  
author: Wisdomlin

框架组成：
1. 数据库orm采用gorm  文档地址 https://gorm.io/zh_CN/docs/connecting_to_the_database.html





使用注意：
1. 如需修改项目名称，除了将项目文件夹名称修改外，还需要修改文件中 import 地方对应的 名称（可以全局搜索进行修改）
    1. public/main.go
    2. route.go
    3. go.mod 文件 module 
    4. 删除 go.sum

    
