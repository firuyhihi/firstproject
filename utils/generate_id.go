package utils

import "math/rand"

func GenerateId(dept string) string {
	if dept == "ENGINEERING" {
		return "EN" + RandStringBytes(6)
	} else if dept == "BUSINESS" {
		return "BU" + RandStringBytes(6)
	} else if dept == "OTHERS" {
		return "NA" + RandStringBytes(6)
	}
	return ""
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
