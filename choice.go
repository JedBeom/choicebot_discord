package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	criteria = []string{
		" vs ",
		" ㄷ ",
		" ",
	}
)

// 선택 함수
func choice(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := "!선택 "
	content := m.Content[len(prefix):]

	rand.Seed(time.Now().UnixNano())
	var options []string

	for _, value := range criteria {
		if strings.Contains(content, value) {
			options = strings.Split(content, value)
			break
		}
	}

	// 선택지가 1개 이하라면 끝내기
	if len(options) < 2 {
		return
	}

	// 0부터 슬라이스의 길이 중의 정수 중에서 하나를 무작위로 뽑은 다음
	// 그 숫자를 인덱스 번호로 넣어 값을 받아옴
	truth := options[rand.Intn(len(options))]

	reply(s, m, truth)

}
