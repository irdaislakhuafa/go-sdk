package codes

import "slices"

type CodeInterface interface {
	IsBetween(startCode, endCode Code) bool
	IsOneOf(listCodes ...Code) bool
	IsNotOneOf(listCodes ...Code) bool
}

var _ CodeInterface = new(Code)

func (c *Code) IsBetween(startCode Code, endCode Code) bool {
	return *c >= startCode && *c <= endCode
}

func (c *Code) IsOneOf(listCodes ...Code) bool {
	return slices.Contains(listCodes, *c)
}

func (c *Code) IsNotOneOf(listCodes ...Code) bool {
	return !c.IsOneOf(listCodes...)
}
