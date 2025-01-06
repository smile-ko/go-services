package global

import (
	"go-services/pkg/setting"

	"gorm.io/gorm"
)

var (
	PDB    *gorm.DB
	Config setting.Config
)
