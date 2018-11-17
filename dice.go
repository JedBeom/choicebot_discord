package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// 주사위 함수
func dice(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := "!주사위 "
	content := m.Content[len(prefix):]

	rand.Seed(time.Now().UnixNano())

	if strings.Contains(content, "-") {

		numberWithString := strings.Split(content, "-")
		if len(numberWithString) != 2 {
			return
		}

		number := [2]int{0, 0}
		var err error
		number[0], err = strconv.Atoi(numberWithString[0])
		if err != nil {
			return
		}

		number[1], err = strconv.Atoi(numberWithString[1])
		if err != nil {
			return
		}

		if number[0] >= number[1] {
			return
		}

		magicNumber := rand.Intn(number[1] - number[0] + 1)
		magicNumber += number[0]

		reply(s, m, strconv.Itoa(magicNumber))
	}

}
