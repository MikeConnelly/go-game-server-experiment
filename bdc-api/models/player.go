package models

type Player struct {
	Base
	UserID     uint
	User       User
	WorldID    uint
	World      World
	AllianceID *uint
	Alliance   Alliance
	Name       string

	ResourceWorkerCount int
	ResourceGoldCount   int
	ResourceLumberCount int
	ResourceManaCount   int

	CastleXCoord          int
	CastleYCoord          int
	CastleFarmStage       int
	CastleGoldmineStage   int
	CastleLumbermillStage int
	CastleManawellStage   int
	CastleEspionageStage  int
	CastleArmoryStage     int
	CastleStableStage     int
	CastleHeavyStage      int
}
