package main

import (
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
	Time       string
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
		token = Token{
			UserID: user.ID,
			Token:  generateToken(user.ID), // Assume there is a function generateToken to generate a unique token
		}
		if err := db.Create(&token).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "无法生成token"})
			return
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
		user.Password = ""

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

	// 地点添加接口
	r.POST("/addlocation", func(c *gin.Context) {
		var request struct {
			Token       string `json:"token"`
			Name        string `json:"name"`
			Description string `json:"description"`
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

		location := Location{
			Name:        request.Name,
			Description: request.Description,
		}

		if err := db.Create(&location).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "添加地点失败"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "添加地点成功", "data": location})
	})

	// 预约地点修改接口
	r.POST("/updatelocation", func(c *gin.Context) {
		var request struct {
			Token       string `json:"token"`
			LocationID  int    `json:"location_id"`
			Name        string `json:"name"`
			Description string `json:"description"`
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

		var location Location
		if err := db.Model(&location).Where("id = ?", request.LocationID).First(&location).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "查询预约地点失败"})
			return
		}

		// 更新预约地点信息
		if err := db.Model(&location).Updates(Location{Name: request.Name, Description: request.Description}).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "更新预约地点失败"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "更新预约地点成功"})
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
			LocationID int    `json:"location_id"`
			Date       string `json:"date"`
			Time       string `json:"time"`
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

		var record Record
		record.UserID = tokenData.UserID
		if err != nil {
			// Handle the error if the conversion fails
		}
		record.LocationID = uint(request.LocationID)
		record.Date = request.Date
		record.Time = request.Time

		if err := db.Create(&record).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "message": "预约失败"})
			return
		}

		c.JSON(200, gin.H{"code": 0, "message": "预约成功", "data": record})
	})

	// 用户预约记录列表接口
	r.POST("/listrecord", func(c *gin.Context) {
		var request struct {
			Token string `json:"token"`
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

		var records []Record

		db.Where("user_id = ?", tokenData.UserID).Find(&records)

		c.JSON(200, gin.H{"code": 0, "message": "搜索成功", "data": records})
	})

	// 预约地信息查询接口
	r.POST("/locationinfo", func(c *gin.Context) {
		var request struct {
			Token      string `json:"token"`
			LocationID int    `json:"location_id"`
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
