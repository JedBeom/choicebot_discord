package main

import (
	"flag"
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
	token string
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
	if token == "" {
		log.Fatal("No token")
		os.Exit(1)
	}
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error While Creating Discord Session.")
		return
	}

	dg.AddHandler(handler)

	err = dg.Open()
	if err != nil {
		log.Fatal("Error while opening connection,", err)
		return
	}

	log.Println("Bot is now running. Press CTRL-C to exit")
	dg.UpdateStatus(0, "도움말: !세리카")
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
	prefix := "!선택 "

	rand.Seed(time.Now().UnixNano())
	things := m.Content[len(prefix):]

	if strings.Contains(m.Content, " vs ") {

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

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	footer := discordgo.MessageEmbedFooter{
		Text:    "© Bombwhale",
		IconURL: "https://avatars2.githubusercontent.com/u/20675630?s=460&v=4",
	}
	fields := []*discordgo.MessageEmbedField{}

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "!선택",
		Value: "`!선택 A B` 또는 `!선택 A eats vs B dances`\n항목 중 하나를 선택해요!",
	}, &discordgo.MessageEmbedField{
		Name:  "!주사위",
		Value: "`!주사위 a-b`\n정수 a부터 b가지의 수 중 하나를 무작위로 뽑아요!",
	})

	author := discordgo.MessageEmbedAuthor{
		Name:    "하코자키 세리카",
		IconURL: "https://raw.githubusercontent.com/JedBeom/choicebot_discord/master/serika.png",
	}

	embed := discordgo.MessageEmbed{
		Author:      &author,
		Title:       "세리카 봇",
		Description: "제 사용법이에요!",
		Color:       0xed90ba,
		Fields:      fields,
		Footer:      &footer,
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!세리카" {
		help(s, m)
	} else if strings.HasPrefix(m.Content, "!선택 ") {
		choice(s, m)
	} else if strings.HasPrefix(m.Content, "!주사위 ") {
		dice(s, m)
	}
}
