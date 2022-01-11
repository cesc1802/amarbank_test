package valueobject

import "strings"

type GenderType int

const (
	MALE GenderType = iota
	FEMALE
)

func GenderFromString(gender string) GenderType {
	genderMap := map[string]GenderType{
		"MALE":   MALE,
		"FEMALE": FEMALE,
	}

	return genderMap[strings.ToUpper(gender)]
}

func (gender GenderType) String() string {
	names := [...]string{
		"MALE",
		"FEMALE"}

	if gender < MALE || gender > FEMALE {
		return "Unknown"
	}
	return names[gender]
}

func (gender GenderType) IsMale() bool {
	switch gender {
	case MALE:
		return true
	default:
		return false
	}
}

func (gender GenderType) IsFeMale() bool {
	switch gender {
	case FEMALE:
		return true
	default:
		return false
	}
}
