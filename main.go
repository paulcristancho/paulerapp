package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)


func suma(a int64, b int64) int64 {
	return a + b
}

// boolean true, false
// int     numero
// float   numero flotante -> 0.8 80.8 106.7
// strings "Hola"

type User struct {
	gorm.Model
	Name         string
	Age          int64
	Email        string  `gorm:"type:varchar(100);unique_index"`
}

type CREACION struct{
	Name string `json:"name"`
	Age int64 `json:"age"`
	Email string `json:"email"`
}


func main() {
	//log.Println("Hola Pauler")
	//log.Println( suma(100, 24) )
	// TCP
	// UDP
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/pauler?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(User{})


	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/pauler", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ivant",
		})
	})

	r.POST("/pauler", func(c *gin.Context) {
		var mensaje CREACION
		c.BindJSON(&mensaje)

		user := User{
			Name: mensaje.Name,
			Age: mensaje.Age,
			Email: mensaje.Email,
		}

		err := db.Create(&user).Error
		respuesta := "Ok"
		if err != nil {
			respuesta = "Ya existe el correo"
		}

		c.JSON(200, gin.H{
			"message": respuesta,
			//{
			//	"message": "ivantrips1@gmail.com"
			//}
		})
	})


	r.Run()
	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

}



