package main

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
    db *gorm.DB
    sqlConnection = "root:ylx25122@tcp(127.0.0.1:3306)/chapter8?" +
		"charset=utf8&parseTime=true"
)

func init()  {
    var err error
    db,err = gorm.Open("mysql",sqlConnection)
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&User{})
}

func main()  {
    router := gin.Default()
    v2 := router.Group("/api/v2/user")
    {
        v2.POST("/", createUser)      //POST方法，创建新用户
		v2.GET("/", fetchAllUser)     //GET方法，获取所有用户
		v2.GET("/:id", fetchUser)     //GET方法，获取某一个用户，形如：/api/v2/user/1
		v2.PUT("/:id", updateUser)    //PUT方法，更新用户，形如：/api/v2/user/1
		v2.DELETE("/:id", deleteUser) //DELETE方法，删除用户，形如：/api/v2/user/1
	}
    router.Run("127.0.0.1:8086")
} 

type (
    User struct {
        ID  uint    `json:"id"`
        Phone   string  `json:"phone`
        Name    string  `json:"name"`
        Password    string  `json:"password"`
    }

    UserRes struct {
        ID    uint   `json:"id"`
		Phone string `json:"phone"`
		Name  string `json:"name"`
    }

)

func md5Password(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func createUser(c *gin.Context)  {
    phone := c.PostForm("phone")
    name := c.PostForm("name")
    user := User {
        Phone: phone,
        Name: name,
        Password: md5Password("666666"),
    }
    db.Save(&user)
    c.JSON(
        http.StatusCreated,
        gin.H{
            "status":  http.StatusCreated,
            "message":"User Created successfully",
            "ID":   user.ID,
        })

}

func fetchAllUser(c *gin.Context)  {
    var user []User
    var _userRes []UserRes

    db.Find(&user)
    if len(user) <= 0 {
        c.JSON(
            http.StatusNotFound,
            gin.H{
                "status": http.StatusNotFound,
                "message":"No user found",
            })
            return
    }
    
    for _,item := range user {
        _userRes = append(_userRes, 
            UserRes{
                ID: item.ID,
                Phone: item.Phone,
                Name: item.Name,
            })
    }

    c.JSON(http.StatusOK,
        gin.H{
            "status":http.StatusOK,
            "data":_userRes,
        })
}

// 获取单个用户
func fetchUser(c *gin.Context) {
	var user User       //定义User结构体
	ID := c.Param("id") //获取参数id

	db.First(&user, ID)

	if user.ID == 0 { //如果用户不存在，则返回
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	//返回响应结构体
	res := UserRes{ID: user.ID, Phone: user.Phone, Name: user.Name}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
}

func updateUser(c *gin.Context) {
	var user User           //定义User结构体
	userID := c.Param("id") //获取参数id
	db.First(&user, userID) //查找数据库

	if user.ID == 0 { //如果数据库不存在，则返回
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	//更新对应的字段值
	db.Model(&user).Update("phone", c.PostForm("phone"))
	db.Model(&user).Update("name", c.PostForm("name"))
	c.JSON(http.StatusOK,
		gin.H{"status": http.StatusOK, "message": "Updated User successfully!"})
}

// 删除用户
func deleteUser(c *gin.Context) {
	var user User           //定义User结构体
	userID := c.Param("id") //获取参数id

	db.First(&user, userID) //查找数据库

	if user.ID == 0 { //如果数据库不存在，则返回
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	//删除用户
	db.Delete(&user)
	c.JSON(http.StatusOK,
		gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}