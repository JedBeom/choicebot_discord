package main

import (
	"fmt"
	disgo "github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"time"
)

var roleWait = make(map[string]string) // [channelID]messageID

const emojiID = "blob_wow:682930056053522452"

// const emojiID = "shiny_asahi:683651245394755593"

func roleStart(s *disgo.Session, m *disgo.MessageCreate) {
	msg, err := s.ChannelMessageSend(m.ChannelID, "@here 역할놀이를 시작할게요. 참여하려면 리액션 해주세요!\n**시작하기 전, 나와 닉네임이 같은 사람이 없는지 꼭 확인하세요!**")
	if err != nil {
		log.Println(err)
		return
	}

	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, emojiID); err != nil {
		log.Println(err)
		return
	}

	roleWait[m.ChannelID] = msg.ID
}

func roleMix(s *disgo.Session, m *disgo.MessageCreate) {
	msgID, ok := roleWait[m.ChannelID]
	if !ok {
		reply(s, m, "`!역할놀이`가 먼저 호출되어야 해요...")
		return
	}

	users, err := s.MessageReactions(m.ChannelID, msgID, emojiID, 10, "", "")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(users)

	users = users[:len(users)-1]
	if len(users) <= 1 {
		reply(s, m, "참여하는 사람이 없어요...")
		return
	}

	var nicks = make([]string, 0, 10)
	for _, u := range users {
		member, err := s.GuildMember(m.GuildID, u.ID)
		nick := "에러"
		if err == nil && member != nil {
			nick = member.User.Username
			if member.Nick != "" {
				nick = member.Nick
			}
		}

		nicks = append(nicks, nick)
	}
	oriNicks := make([]string, len(nicks))
	copy(oriNicks, nicks)
	log.Println(nicks)

	try := 0
RANDOM:
	try++
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nicks), func(i, j int) {
		nicks[i], nicks[j] = nicks[j], nicks[i]
	})

	for i := range nicks {
		if nicks[i] == oriNicks[i] && try <= 3 {
			goto RANDOM
		}
	}

	result := "@here **닉네임 변경 결과!**"

	for i := range users {
		err := s.GuildMemberNickname(m.GuildID, users[i].ID, nicks[i])
		if err != nil {
			log.Println(err)
		}
		result += fmt.Sprintf("\n%s → %s", users[i].Username, nicks[i])
	}

	_, err = s.ChannelMessageSend(m.ChannelID, result+"\n\n**이제 혼란스러운 역할놀이를 시작하세요!**")
	if err != nil {
		log.Println(err)
	}
}
