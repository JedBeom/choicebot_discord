package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	token string
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error While Creating Discord Session.")
		return
	}

	dg.AddHandler(choice)

	err = dg.Open()
	if err != nil {
		log.Fatal("Error while opening connection,", err)
		return
	}

	log.Println("Bot is now running. Press CTRL-C to exit")
	dg.UpdateStatus(0, "!선택 도움말")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func reply(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	message := "<@" + m.Author.ID + "> " + content
	s.ChannelMessageSend(m.ChannelID, message)
}

func choice(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!선택 ") {
		rand.Seed(time.Now().UnixNano())
		things := m.Content[8:]

		if things == "도움말" {
			help := `'!선택'을 앞에 붙인 뒤, 항목들을 vs나 띄어쓰기로 구분해서 보내주세요.
vs가 있을 경우 띄어쓰기로 항목이 구별되지 않습니다.
알아들을 수 없는 경우 대답하지 않아요.`

			reply(s, m, help)

		} else if strings.Contains(m.Content, " vs ") {

			options := strings.Split(things, " vs ")
			if len(options) < 2 {
				return
			}
			truth := options[rand.Intn(len(options))]

			reply(s, m, truth)

		} else {
			options := strings.Split(things, " ")
			if len(options) < 2 {
				return
			}

			truth := options[rand.Intn(len(options))]

			reply(s, m, truth)

		}

	}
}
