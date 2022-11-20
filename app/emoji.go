package main

import (
	"math/rand"
	"time"
)

var (
	emojis = []string{
		"👁️‍🗨️",
		"💫",
		"💢",
		"💤",
		"🌐",
		"🌎",
		"🌆",
		"🌕",
		"🌊",
		"🎭",
		"🖼",
		"🎨",
		"🧵",
		"🪡",
		"🧶",
		"🪢",
		"💽",
		"💾",
		"💿",
		"📀",
		"😪",
		"😌",
		"🤌",
		"🏆",
		"🎶",
		"🎬",
		"🔱",
	}
)

func getRandomEmoji() string {
	rand.Seed(time.Now().Unix())

	num := rand.Intn(len(emojis))
	return emojis[num]
}
