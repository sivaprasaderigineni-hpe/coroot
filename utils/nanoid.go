package utils

import "github.com/matoous/go-nanoid"

func NanoId(size int) string {
	id, _ := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", size)
	return id
}
