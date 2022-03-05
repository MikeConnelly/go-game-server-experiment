package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"golang.org/x/crypto/bcrypt"

	. "go_server_test/database"
	"go_server_test/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		log.Fatal("error parsing request")
		return err
	}

	sess, err := store.Get(c)
	if err != nil {
		log.Fatal("error getting context")
		return err
	}

	var user models.User

	//DB.Where("email = ?", data["email"]).First(&user)
	// for testing
	user.ID = 1
	user.Name = "name"
	user.Email = "email"
	user.Password, _ = bcrypt.GenerateFromPassword([]byte("password"), 12)

	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	sess.Set("email", user.Email)
	sess.Set("userId", user.ID)
	err = sess.Save()
	if err != nil {
		log.Fatal("error saving session")
		return err
	}

	return c.SendStatus(200)
}
