package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Registration struct {
	Id        int64
	Name      string
	Email     string
	CreatedAt time.Time
}

var DB gorm.DB

func init() {
	// auth.InitFacebookStrategy()
	// auth.InitGoogleStrategy()
	var err error
	DB, err = gorm.Open("postgres", "user=victor dbname=ankara sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
	// DB.SingularTable(true)
	DB.AutoMigrate(Registration{})
	// DB.
}

func main() {
	r := gin.Default()
	r.Static("/public", "public")
	// r.Use(static.Serve("/public"))
	r.LoadHTMLTemplates("templates/*")
	r.GET("/", func(c *gin.Context) {
		// obj := gin.H{"title": "Main website"}
		c.HTML(200, "index.tmpl", nil)
	})
	r.POST("/", func(c *gin.Context) {
		req := c.Request
		req.ParseForm()
		name := req.FormValue("name")
		email := req.FormValue("email")
		registration := Registration{}
		if DB.Where(&Registration{Email: email}).Find(&registration).RecordNotFound() {
			DB.Save(&Registration{Name: name, Email: email})
		}
		c.HTML(200, "thank-you.tmpl", nil)
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8085")
}
