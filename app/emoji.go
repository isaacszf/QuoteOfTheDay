package main

import (
	"math/rand"
	"time"
)

var (
	emojis = []string{
		"๐๏ธโ๐จ๏ธ",
		"๐ซ",
		"๐ข",
		"๐ค",
		"๐",
		"๐",
		"๐",
		"๐",
		"๐",
		"๐ญ",
		"๐ผ",
		"๐จ",
		"๐งต",
		"๐ชก",
		"๐งถ",
		"๐ชข",
		"๐ฝ",
		"๐พ",
		"๐ฟ",
		"๐",
		"๐ช",
		"๐",
		"๐ค",
		"๐",
		"๐ถ",
		"๐ฌ",
		"๐ฑ",
	}
)

func getRandomEmoji() string {
	rand.Seed(time.Now().Unix())

	num := rand.Intn(len(emojis))
	return emojis[num]
}
