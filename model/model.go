package model

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Load() []interface{} {
	return []interface{}{
		&Todo{},
	}
}

type DefaultModel struct {
	ID string `gorm:"primaryKey;type:CHAR(36);not null"`

	Status    EnumStatus     `gorm:"type:ENUM('WAIT','IDLE','INVALID','DELETED');default:'WAIT'"`
	CreatedBy string         `gorm:"type:CHAR(36);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedBy string         `gorm:"type:CHAR(36);not null"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;index"`
}

type EnumStatus string

const (
	StatusWait    EnumStatus = "WAIT"
	StatusIdle    EnumStatus = "IDLE"
	StatusInvalid EnumStatus = "INVALID"
	StatusRemoved EnumStatus = "DELETED"
)

func (e *EnumStatus) Scan(value interface{}) error {
	*e = EnumStatus(value.([]byte))
	return nil
}

func (e EnumStatus) Value() (driver.Value, error) {
	return string(e), nil
}

func (m *DefaultModel) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *DefaultModel) BeforeUpdate(_ *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

func (m *DefaultModel) BeforeDelete(_ *gorm.DB) error {
	m.DeletedAt = gorm.DeletedAt{
		Time:  time.Now(),
		Valid: true,
	}
	return nil
}
