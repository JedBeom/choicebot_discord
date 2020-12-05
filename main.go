package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	token   string
	version = "0.1.3"

	embed          discordgo.MessageEmbed
	spec           discordgo.MessageEmbed
	versionMessage discordgo.MessageEmbed

	color int = 0xed90ba
)

// 초기화
func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()

	// 만일 토큰이 안 들어 왔을 때
	if token == "" {
		log.Fatal("USAGE: -t [token]")
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

func vote(s *discordgo.Session, m *discordgo.MessageCreate) {
	//prefix := "!투표 "
	//content := m.Content[len(prefix):]

	// 개발 시 handler() 수정 필요

	reply(s, m, "투표 기능은 아직 개발 중이에요!")
}

// Prefix에 맞춰 함수 실행
func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!세리카" {
		sendHelp(s, m)
	} else if m.Content == "!세리카 버전" {

		versionMessage.Timestamp = time.Now().Format(time.RFC3339)

		message := discordgo.MessageSend{
			Content: "<@" + m.Author.ID + ">",
			Embed:   &versionMessage,
		}

		s.ChannelMessageSendComplex(m.ChannelID, &message)

	} else if m.Content == "!세리카 스펙" {
		sendSpec(s, m)
	} else if strings.HasPrefix(m.Content, "!선택 ") {
		choice(s, m)
	} else if strings.HasPrefix(m.Content, "!주사위 ") {
		dice(s, m)
	} else if strings.HasPrefix(m.Content, "!업다운 ") {
		updown(s, m)
	} else if strings.HasPrefix(m.Content, "!투표") {
		vote(s, m)
	} else if m.Content == "!역할놀이" {
		roleStart(s, m)
	} else if m.Content == "!역할놀이 시작" {
		roleMix(s, m)
	}
	// 해당하는 것이 없으면 아예 반응을 안함
}
