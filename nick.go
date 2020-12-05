package main

import (
	disgo "github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"time"
)

var roleWait map[string]string // [channelID]messageID

const emojiID = "682930056053522452"

func roleStart(s *disgo.Session, m *disgo.MessageCreate) {
	msg, err := s.ChannelMessageSend(m.ChannelID, "@here 역할놀이를 시작할게요. 참여하려면 리액션 해주세요!")
	if err != nil {
		log.Println(err)
		return
	}

	if err = s.MessageReactionAdd(m.ChannelID, msg.ID, emojiID); err != nil {
		log.Println(err)
		return
	}
}

func roleMix(s *disgo.Session, m *disgo.MessageCreate) {
	msgID, ok := roleWait[m.ChannelID]
	if !ok {
		reply(s, m, "`!역할놀이`가 먼저 호출되어야 해요...")
		return
	}

	users, err := s.MessageReactions(m.ChannelID, msgID, emojiID, 10, "", "")
	if err != nil {
		return
	}

	var nicks = make([]string, 0, 10)
	for _, u := range users {
		member, err := s.GuildMember(m.GuildID, u.ID)
		nick := "에러"
		if err == nil && member != nil {
			nick = member.Nick
		}

		nicks = append(nicks, nick)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nicks), func(i, j int) {
		nicks[i], nicks[j] = nicks[j], nicks[i]
	})

	for i := range users {
		err := s.GuildMemberNickname(m.GuildID, users[i].ID, nicks[i])
		if err != nil {
			log.Println(err)
		}
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "@here 닉네임 변경을 완료했어요! 역할극 시작!")
	if err != nil {
		log.Println(err)
	}
}
