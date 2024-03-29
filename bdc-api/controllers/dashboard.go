package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	. "go_server_test/database"
)

// Raw SQL Query for Dashboard Data:
//   SELECT worlds.id AS world_id, worlds.name AS world_name, p.name AS player_name,
//   p.resource_worker_count, p.resource_gold_count, p.resource_lumber_count, p.resource_mana_count
//   FROM "worlds" LEFT JOIN players AS p ON p.world_id = worlds.id AND p.user_id = ?
//   WHERE worlds.active = true

type DashboardData struct {
	WorldID             uint
	WorldName           string
	PlayerName          string
	ResourceWorkerCount int
	ResourceGoldCount   int
	ResourceLumberCount int
	ResourceManaCount   int
}

func GetDashboardData(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(sess)
	userID := sess.Get("userId")
	if userID == nil {
		return c.SendStatus(401)
	}

	var data []DashboardData
	DB.Select("w.id AS world_id, w.name AS world_name, p.name AS player_name, "+
		"p.resource_worker_count, p.resource_gold_count, "+
		"p.resource_lumber_count, p.resource_mana_count").
		Where("w.active = true").
		Table("worlds AS w").
		Joins("LEFT JOIN players AS p ON p.world_id = w.id AND p.user_id = ?", userID).
		Find(&data)

	fmt.Println(data)
	c.JSON(data)

	return nil
}
