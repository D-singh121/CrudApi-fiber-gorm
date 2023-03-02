package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// const urlDSN = "root:Dee154595@@tcp(localhost:3306)/goDB?parseTime=true"
const urlDSN = "postgres://testuser:test123@localhost:5432/mydb?sslmode=disable"

type User struct {
	gorm.Model
	FirstName string `JSON:"firstname"`
	LastName  string `JSON:"lastname"`
	Email     string `JSON:"email"`
}

// -------connection establised with database using gorm
func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to database ")
	}
	DB.AutoMigrate(&User{})

}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&user)
	return c.JSON(&user)
}

func Getusers(c *fiber.Ctx) error {
	var users []User
	DB.Find(&users)
	return c.JSON(&users)

}

func Getuser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.Find(&user, id)
	return c.JSON(&user)

}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("user not available")
	}

	DB.Delete(&user)
	return c.SendString("user Deleted from database")

}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)

	if user.Email == "" {
		return c.Status(500).SendString("user not found")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&user)
	return c.JSON(&user)
}
