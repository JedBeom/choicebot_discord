package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// 선택 함수
func choice(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := "!선택 "
	content := m.Content[len(prefix):]

	rand.Seed(time.Now().UnixNano())

	// ' vs '가 내용 속에 있을 경우
	if strings.Contains(content, " vs ") {

		// ' vs '를 기준으로 나눠 슬라이스로
		options := strings.Split(content, " vs ")
		// 슬라이스가 1개 이하일 경우
		if len(options) < 2 {
			return
		}

		// 0부터 슬라이스의 길이 중의 정수 중에서 하나를 무작위로 뽑은 다음
		// 그 숫자를 인덱스 번호로 넣어 값을 받아옴
		truth := options[rand.Intn(len(options))]

		reply(s, m, truth)

	} else if strings.Contains(content, " ㄷ ") {
		options := strings.Split(content, " ㄷ ")
		if len(options) < 2 {
			return
		}

		truth := options[rand.Intn(len(options))]
		reply(s, m, truth)

	} else {
		// 띄어쓰기를 기준으로 슬라이스로 나눔
		options := strings.Split(content, " ")
		// 요소가 1개 이하일 경우
		if len(options) < 2 {
			return
		}

		// 0부터 슬라이스의 길이 중의 정수 중에서 하나를 무작위로 뽑은 다음
		// 그 숫자를 인덱스 번호로 넣어 값을 받아옴
		truth := options[rand.Intn(len(options))]

		reply(s, m, truth)

	}

}
