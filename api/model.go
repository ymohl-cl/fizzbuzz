package api

import (
	"strconv"
)

// InputFB to request fizzbuzz endpoint
type InputFB struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// OutputStat
type OutputStat struct {
	InputFB
	Count int64 `json:"count"`
}

// Hash return a text representation of content
func (i InputFB) Hash() string {
	var hash string

	hash = strconv.Itoa(i.Limit)
	hash += strconv.Itoa(i.Int1)
	hash += strconv.Itoa(i.Int2)
	hash += i.Str1
	hash += i.Str2
	return hash
}
