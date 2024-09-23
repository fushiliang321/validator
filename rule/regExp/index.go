package regExp

import "regexp"

var (
	Alpha      = regexp.MustCompile(`^[A-Za-z\\u4e00-\\u9fa5]+$`)
	AlphaNum   = regexp.MustCompile(`^[\p{L}\d]+$`)
	AlphaDash  = regexp.MustCompile(`^[\p{L}\d_-]+$`)
	Ascii      = regexp.MustCompile(`^[\x00-\x7F]{7}$`)
	MacAddress = regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)
	Email      = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$`)
	Phone      = regexp.MustCompile(`^1[3-9]\d{9}$`)
)
