package model

type CardBrand struct {
	Brand string `gorm:"type:text;primarKey"`
}
