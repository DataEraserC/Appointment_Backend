package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"unique"`
	Password    string
	Avatar      string
	NickName    string
	PhoneNumber string
}

type Location struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}

type Record struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	LocationID uint
	Date       string
}

type Token struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Token     string `gorm:"unique"`
	CreatedAt int64
}

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate
	db.AutoMigrate(&User{}, &Location{}, &Record{}, &Token{})

	r := gin.Default()

	// 用户注册接口
	r.POST("/register", func(c *gin.Context) {
		var request struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
			return
		}

		user := User{
			Username: request.Username,
			Password: request.Password,
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "用户名已存在"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "注册成功"})
	})

	// 用户登录接口
	r.POST("/login", func(c *gin.Context) {
		var request struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
			return
		}

		var user User
		if err := db.Where("username = ?", request.Username).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "用户名或密码错误"})
			return
		}

		if user.Password != request.Password {
			c.JSON(400, gin.H{"code": 1, "message": "用户名或密码错误"})
			return
		}

		var token Token
		if err := db.Where("user_id = ?", user.ID).First(&token).Error; err != nil {
			// If token does not exist, create a new token for the user
			token = Token{
				UserID: user.ID,
				Token:  generateToken(user.ID), // Assume there is a function generateToken to generate a unique token
			}
			if err := db.Create(&token).Error; err != nil {
				c.JSON(400, gin.H{"code": 1, "message": "无法生成token"})
				return
			}
		}

		c.JSON(200, gin.H{"code": 0, "message": "登录成功", "token": token.Token})
	})

	// 用户获取个人信息接口
	r.POST("/userinfo", func(c *gin.Context) {
		var request struct {
			Token string `json:"token"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
			return
		}

		var tokenRecord Token
		if err := db.Where("token = ?", request.Token).First(&tokenRecord).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "身份验证失败"})
			return
		}

		var user User
		if err := db.First(&user, tokenRecord.UserID).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "获取用户信息失败"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "获取个人信息成功", "data": user})

	})

	// 用户修改个人信息接口
	r.POST("/updateuserinfo", func(c *gin.Context) {
		var request struct {
			Token       string `json:"token"`
			Username    string `json:"username"`
			Password    string `json:"password"`
			Avatar      string `json:"avatar"`
			NickName    string `json:"nickname"`
			PhoneNumber string `json:"phoneNumber"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
			return
		}

		var tokenData Token
		if err := db.Model(&tokenData).Where("token = ?", request.Token).First(&tokenData).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "身份验证失败"})
			return
		}

		var user User
		if err := db.Model(&user).Where("ID = ?", tokenData.UserID).Updates(User{
			Username:    request.Username,
			Password:    request.Password,
			Avatar:      request.Avatar,
			NickName:    request.NickName,
			PhoneNumber: request.PhoneNumber,
		}).Error; err != nil {
			c.JSON(500, gin.H{"code": 2, "message": "修改个人信息失败"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "修改个人信息成功"})
	})

	// 预约地点搜索接口
	r.POST("/searchlocation", func(c *gin.Context) {
		var request struct {
			Token   string `json:"token"`
			Keyword string `json:"keyword"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
			return
		}

		var tokenData Token
		if err := db.Model(&tokenData).Where("token = ?", request.Token).First(&tokenData).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "身份验证失败"})
			return
		}

		var locations []Location
		if err := db.Where("name LIKE ?", "%"+request.Keyword+"%").Find(&locations).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "搜索失败"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "搜索成功", "data": locations})
	})

	// 用户预约接口
	r.POST("/reservation", func(c *gin.Context) {
		var request struct {
			Token      string `json:"token"`
			LocationID string `json:"location_id"`
			Date       string `json:"date"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
			return
		}

		var user User
		if err := db.Where("token = ?", request.Token).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "身份验证失败"})
			return
		}

		var record Record
		record.UserID = user.ID
		locationIDUint, err := strconv.ParseUint(request.LocationID, 10, 32)
		if err != nil {
			// Handle the error if the conversion fails
		}
		record.LocationID = uint(locationIDUint)
		record.Date = request.Date

		if err := db.Create(&record).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "预约失败"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "预约成功", "data": record})
	})

	// 预约记录搜索接口
	r.POST("/searchrecord", func(c *gin.Context) {
		var request struct {
			Token   string `json:"token"`
			Keyword string `json:"keyword"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
			return
		}

		var user User
		if err := db.Where("token = ?", request.Token).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "身份验证失败"})
			return
		}

		var records []Record
		if err := db.Where("date LIKE ? OR location_id LIKE ?", "%"+request.Keyword+"%", "%"+request.Keyword+"%").Find(&records).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "搜索失败"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "搜索成功", "data": records})
	})

	// 预约地信息查询接口
	r.POST("/locationinfo", func(c *gin.Context) {
		var request struct {
			Token      string `json:"token"`
			LocationID string `json:"location_id"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "参数错误"})
			return
		}

		var user User
		if err := db.Where("token = ?", request.Token).First(&user).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "身份验证失败"})
			return
		}

		var location Location
		if err := db.Where("id = ?", request.LocationID).First(&location).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "查询失败"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "查询成功", "data": location})
	})

	r.Run()
}

func generateToken(userID uint) string {
	token := uuid.New().String() // Generate a unique token using UUID
	return token
}
