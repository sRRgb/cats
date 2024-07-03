package models

import (
	"github.com/jinzhu/gorm"
)

type Mission struct {
	gorm.Model
	CatID    uint     `json:"cat_id"`
	Targets  []Target `gorm:"foreignkey:MissionID" json:"targets"`
	Complete bool     `json:"complete"`
}

type Target struct {
	gorm.Model
	MissionID uint   `json:"-"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Notes     string `json:"notes"`
	Complete  bool   `json:"complete"`
}

type CreateMissionInput struct {
	CatID    uint          `json:"cat_id" binding:"required"`
	Targets  []TargetInput `json:"targets" binding:"required"`
	Complete bool          `json:"complete"`
}

type TargetInput struct {
	Name     string `json:"name" binding:"required"`
	Country  string `json:"country" binding:"required"`
	Notes    string `json:"notes"`
	Complete bool   `json:"complete"`
}
