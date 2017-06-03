package utils

func NewTrueBoolPtr() *bool {
	b := true
	return &b
}

func NewFalseBoolPtr() *bool {
	b := false
	return &b
}
