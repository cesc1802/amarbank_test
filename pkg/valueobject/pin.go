package valueobject

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type PIN struct {
	firstBlk  string
	secondBlk string
	thirdBlk  string
}

func NewPIN(f, t string) *PIN {
	return &PIN{
		firstBlk: f,
		thirdBlk: t,
	}
}

func (p *PIN) Gen(typ GenderType) string {
	current := time.Now().Format("01-02-06")
	dateString := strings.Split(current, "-")
	var secondBlk string

	if typ.IsMale() {
		secondBlk = fmt.Sprintf("%s%s%s", dateString[1], dateString[0], dateString[2])
	} else if typ.IsFeMale() {
		dayInt, _ := strconv.Atoi(dateString[1])
		secondBlk = fmt.Sprintf("%d%s%s", dayInt+40, dateString[0], dateString[2])
	} else {

	}

	return fmt.Sprintf("%s%s%s", p.firstBlk, secondBlk, p.thirdBlk)
}
