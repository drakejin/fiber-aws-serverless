package model

// tag validate format: https://github.com/go-playground/validator
// tag gorm format: https://gorm.io
type Todo struct {
	DefaultModel

	Type string `gorm:"type:VARCHAR(100);index;" validate:"required"`
	Sort int32  `validate:"required"`
	Name string `validate:"required"`
	Note string
}
