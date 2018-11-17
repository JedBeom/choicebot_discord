package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	token   string
	version = "18.11.17.2"
	embed   discordgo.MessageEmbed
)

// 초기화
func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()

	// 만일 토큰이 안 들어 왔을 때
	if token == "" {
		log.Fatal("USAGE: -t [token]")
		os.Exit(1) // 프로그램 종료
	}

	// 맨 아래의 정보
	footer := discordgo.MessageEmbedFooter{
		Text:    "© Bombwhale | v" + version,
		IconURL: "https://avatars2.githubusercontent.com/u/20675630?s=460&v=4",
	}

	// 중간의 제목 - 내용들
	fields := []*discordgo.MessageEmbedField{}

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "!선택",
		Value: "`!선택 A B` 또는 `!선택 A eats vs B dances`\n항목 중 하나를 선택해요!",
	}, &discordgo.MessageEmbedField{
		Name:  "!주사위",
		Value: "`!주사위 a-b`\n0을 포함한 자연수 a부터 b가지의 수 중 하나를 무작위로 뽑아요!",
	}, &discordgo.MessageEmbedField{
		Name:  "!업다운",
		Value: "`!업다운 시작`으로 업다운 게임을 시작할 수 있어요.\n숫자는 1-100의 정수이며, `!업다운 23`으로 게임을 진행할 수 있어요.",
	})

	// 맨 위
	author := discordgo.MessageEmbedAuthor{
		Name:    "하코자키 세리카",
		IconURL: "https://raw.githubusercontent.com/JedBeom/choicebot_discord/master/serika.png",
	}

	embed = discordgo.MessageEmbed{
		Author: &author,

		Title:       "세리카 봇",
		Description: "아래의 행동을 할 수 있어요!",

		Color:     0xed90ba,
		Timestamp: time.Now().Format(time.RFC3339),
		Fields:    fields,
		Footer:    &footer,
	}
}

func main() {
	// 토큰으로 디스코드 로그인
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error While Creating Discord Session.")
		return
	}

	dg.AddHandler(handler)

	// 커넥션 열기
	err = dg.Open()
	if err != nil {
		log.Fatal("Error while opening connection,", err)
		return
	}

	log.Println("Bot is now running. Press CTRL-C to exit")
	// 디스코드 내에서 플레이 중인 게임 이름 지정
	dg.UpdateStatus(0, "도움말: !세리카")

	// C-c 들어오면
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// 디스코드 로그아웃
	dg.Close()
}

// 부른 사람에게 멘션하기
func reply(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	message := "<@" + m.Author.ID + "> " + content
	s.ChannelMessageSend(m.ChannelID, message)
}

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

func vote(s *discordgo.Session, m *discordgo.MessageCreate) {
	//prefix := "!투표 "
	//content := m.Content[len(prefix):]

	reply(s, m, "투표 기능은 아직 개발 중이에요!")
}

// Prefix에 맞춰 함수 실행
func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!세리카" {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		if err != nil {
			log.Println(err)
		}
	} else if strings.HasPrefix(m.Content, "!선택 ") {
		choice(s, m)
	} else if strings.HasPrefix(m.Content, "!주사위 ") {
		dice(s, m)
	} else if strings.HasPrefix(m.Content, "!업다운 ") {
		updown(s, m)
	} else if strings.HasPrefix(m.Content, "!투표 ") {
		vote(s, m)
	}
	// 해당하는 것이 없으면 아예 반응을 안함
}
