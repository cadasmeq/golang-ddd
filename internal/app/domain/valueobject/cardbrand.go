package valueobject

import "strings"

type CardBrand string

const (
	VISA       CardBrand = "VISA"
	MASTERCARD CardBrand = "MASTERCARD"
)

func ConvertToCardBrand(s string) CardBrand {
	switch strings.ToUpper(s) {
	case "VISA":
		return VISA
	case "MASTERCARD":
		return MASTERCARD
	default:
		panic("Invalid CardBrand")
	}

}
