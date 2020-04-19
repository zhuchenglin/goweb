# goweb
使用go语言封装一个简单的api框架  
author: Wisdomlin

框架组成：
1. //  数据库orm采用gorm  文档地址 https://gorm.io/zh_CN/docs/connecting_to_the_database.html



使用注意：
1. 因为 gorm里面需要 crypto库，这个库的地址被墙了，需要 在 vendor/pkg/mod/cache/download/golang.org/x 中 
    git clone  https://github.com/golang/crypto.git
    
