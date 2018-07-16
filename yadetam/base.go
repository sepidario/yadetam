package yadetam

import "time"

// Base is the common parts of all entities
type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
