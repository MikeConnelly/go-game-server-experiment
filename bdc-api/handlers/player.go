package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	. "go_server_test/database"
	. "go_server_test/models"
)

// request body format for creating a new player
type CreatePlayerRequestBody struct {
	UserID  uint   `json:"userID"`
	WorldID uint   `json:"worldID"`
	XCoord  int    `json:"xCoord"`
	YCoord  int    `json:"yCoord"`
	Name    string `json:"name"`
}

// returns a new player with default values
// could also set these defaults in the database
func newPlayerBase() Player {
	return Player{
		ResourceWorkerCount:   30,
		ResourceGoldCount:     500,
		ResourceLumberCount:   500,
		ResourceManaCount:     100,
		CastleFarmStage:       1,
		CastleGoldmineStage:   1,
		CastleLumbermillStage: 1,
		CastleManawellStage:   1,
		CastleEspionageStage:  0,
		CastleArmoryStage:     1,
		CastleStableStage:     0,
		CastleHeavyStage:      0,
	}
}

// Checks that user does not have a player in this world
func findPlayerWithUserIDAndWorldIDExists(UserID uint, WorldID uint) string {
	return fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM players WHERE user_id = %d AND world_id = %d) AS found", UserID, WorldID)
}

// Checks that a user of the same name does not exist in this world
func findPlayerWithWorldIDAndNameExists(WorldID uint, Name string) string {
	return fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM players WHERE world_id = %d AND name = %s) AS found", WorldID, Name)
}

// Checks that a castle does not exist at the same position in this world
func findPlayerWithWorldIDAndCoordsExists(WorldID uint, XCoord int, YCoord int) string {
	return fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM players WHERE world_id = %d AND castle_x_coord = %d AND castle_y_coord = %d) AS found", WorldID, XCoord, YCoord)
}

func CreatePlayerAndJoinWorld(c *fiber.Ctx) error {
	// user should be retrieve from some middleware jwt
	var req CreatePlayerRequestBody
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("Error parsing request body")
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// make checks that this player is allowed to be created
	result := struct {
		Found bool
	}{}
	DB.Raw(findPlayerWithUserIDAndWorldIDExists(req.UserID, req.WorldID)).First(&result)
	if result.Found {
		fmt.Println("Player already exists")
		return c.Status(fiber.StatusBadRequest).SendString("Player already exists")
	}
	DB.Raw(findPlayerWithWorldIDAndNameExists(req.WorldID, req.Name)).First(&result)
	if result.Found {
		fmt.Println("Player already exists")
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request. Name not unique")
	}
	DB.Raw(findPlayerWithWorldIDAndCoordsExists(req.WorldID, req.XCoord, req.YCoord)).First(&result)
	if result.Found {
		fmt.Println("Player already exists")
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request. Location taken")
	}

	// create new player
	player := newPlayerBase()
	player.UserID = req.UserID
	player.WorldID = req.WorldID
	player.CastleXCoord = req.XCoord
	player.CastleYCoord = req.YCoord
	player.Name = req.Name

	DB.Create(&player)

	fmt.Println(player)
	c.JSON(player)

	return nil
}
