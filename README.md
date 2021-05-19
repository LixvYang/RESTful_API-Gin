# Go-REST

仅供学习使用，适合初学者

---

### 使用Mysql

```go
//使用mysql
go get github.com/go-sql-driver/mysql
```
#### 数据库连接 

yourname:password@tcp(127.0.0.1:3306)/users

名称为users
### Mysql配置
```go
//创建表
CREATE TABLE `users` (
    -> `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    -> `phone` varchar(255) DEFAULT NULL,
    -> `name` varchar(255) DEFAULT NULL,
    -> `password` varchar(255) DEFAULT NULL,
    -> PRIMARY KEY (`id`)
    -> ) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8;
```

下载

```
git clone https://github.com/YLX621/RESTful_API-Gin.git

cd RESTful_API-Gin

go run main.go
```

监听 http://127.0.0.1:8086/api/v2/user/