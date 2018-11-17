package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

// 업다운 게임을 위한 구조체
type Updown struct {
	Number   int
	TryCount int
}

// 여러 개 있어야겠지
var games = make(map[string]Updown)

// 업다운 함수
func updown(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := "!업다운 "
	content := m.Content[len(prefix):]

	// !업다운 시작
	if content == "시작" {

		// 이미 진행 중인 게임이 있을 때
		if game, ok := games[m.Author.ID]; ok {
			reply(s, m, "이미 게임이 진행 중이에요!")
		} else {
			rand.Seed(time.Now().UnixNano())

			game.Number = rand.Intn(100) + 1
			games[m.Author.ID] = game
			reply(s, m, "자, 게임을 시작 할게요!")
		}

		// !업다운 [숫자] 일 때
	} else if number, err := strconv.Atoi(content); err == nil {

		if game, ok := games[m.Author.ID]; ok {

			// 정답일 때
			if number == game.Number {
				delete(games, m.Author.ID)
				answer := fmt.Sprintf("우와! 정답이에요! %d번만에 맞췄어요!", game.TryCount)
				reply(s, m, answer)

				// 정답보다 클 때
			} else if number > game.Number {
				game.TryCount += 1
				games[m.Author.ID] = game

				answer := fmt.Sprintf("다운! %d번째 시도였어요!", game.TryCount)
				reply(s, m, answer)

				// 정답보다 작을 때
			} else if number < game.Number {
				game.TryCount += 1
				games[m.Author.ID] = game

				answer := fmt.Sprintf("업! %d번째 시도였어요!", game.TryCount)
				reply(s, m, answer)
			}

		} else {
			reply(s, m, "게임을 안하고 있는 걸요? `!업다운 시작`으로 게임을 시작 하시는 것이 어때요?")
		}

		// !업다운 이긴 한데 알아들을 수 없음
	} else {
		reply(s, m, "알아 들을 수 없어요! `!세리카`로 제 사용법을 읽어 보실래요?")
	}
}
