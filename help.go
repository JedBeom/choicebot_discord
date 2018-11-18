package main

import (
	"time"

	disgo "github.com/bwmarrin/discordgo"
)

func init() {
	// 맨 위
	author := disgo.MessageEmbedAuthor{
		Name:    "하코자키 세리카",
		IconURL: "https://raw.githubusercontent.com/JedBeom/choicebot_discord/master/serika.png",
	}

	// 중간의 제목 - 내용들
	fields := []*disgo.MessageEmbedField{}

	fields = append(fields, &disgo.MessageEmbedField{
		Name:  "!선택",
		Value: "**!선택 A eats vs B dances** 또는 **!선택 A B** 또는 **!선택 밀리 애니 ㄷ 월희 리메이크**\n항목 중 하나를 선택해요!",
	}, &disgo.MessageEmbedField{
		Name:  "!주사위",
		Value: "**!주사위 a-b**\n0을 포함한 자연수 a부터 b가지의 수 중 하나를 무작위로 뽑아요!",
	}, &disgo.MessageEmbedField{
		Name:  "!업다운",
		Value: "**!업다운 시작**으로 업다운 게임을 시작할 수 있어요.\n숫자는 1-100의 정수이며, **!업다운 23**으로 게임을 진행할 수 있어요.",
	}, &disgo.MessageEmbedField{
		Name:  "!세리카 버전",
		Value: "버전을 보여줘요.",
	}, &disgo.MessageEmbedField{
		Name:  "!세리카 스펙",
		Value: "세리카 봇의 간단한 스펙을 볼 수 있어요.",
	})

	// 우측 상단의 사진
	thumbnail := disgo.MessageEmbedThumbnail{
		URL: "https://raw.githubusercontent.com/JedBeom/choicebot_discord/master/serika.png",
	}

	// 맨 아래의 정보
	footer := disgo.MessageEmbedFooter{
		Text:    "© Bombwhale | v" + version,
		IconURL: "https://avatars2.githubusercontent.com/u/20675630?s=460&v=4",
	}

	embed = disgo.MessageEmbed{
		Color:  color,
		Author: &author,
		Footer: &footer,

		Title: "Github 레포",
		URL:   "https://github.com/JedBeom/choicebot_discord",

		Description: "아래의 행동을 할 수 있어요!",

		Fields:    fields,
		Thumbnail: &thumbnail,
	}

	versionMessage = disgo.MessageEmbed{
		Color:  color,
		Author: &author,
		Footer: &footer,

		Title: "버전 목록",
		URL:   "https://github.com/JedBeom/choicebot_discord/blob/master/versions.md",

		Description: "현재 버전은 **v" + version + "**이에요!",
	}

	specFields := []*disgo.MessageEmbedField{
		&disgo.MessageEmbedField{
			Name:   "개발 및 실행 언어",
			Value:  "Go 1.11",
			Inline: true,
		},
		&disgo.MessageEmbedField{
			Name:   "서버 스펙",
			Value:  "Raspberry Pi 2B+",
			Inline: true,
		},
		&disgo.MessageEmbedField{
			Name:   "**!선택** 명령어 구분자 우선 순위",
			Value:  "' vs ', ' ㄷ ', ' ' 순",
			Inline: true,
		},
		&disgo.MessageEmbedField{
			Name:  "개발자 및 주인",
			Value: "[Bombwhale](https://github.com/JedBeom), 하코자키 세리카 사진을 제외한 모든 저작권은 저에게 있습니다.",
		}}

	spec = disgo.MessageEmbed{
		Color:  color,
		Author: &author,
		Footer: &footer,

		Title: "Github 레포",
		URL:   "https://github.com/JedBeom/choicebot_discord",

		Description: "세리카 봇의 상세 스펙이에요!",

		Fields: specFields,
	}

}

func sendHelp(s *disgo.Session, m *disgo.MessageCreate) {

	embed.Timestamp = time.Now().Format(time.RFC3339)
	reply(s, m, "직접 보냈으니 확인해 보세요!")

	privateChannel, _ := s.UserChannelCreate(m.Author.ID)
	s.ChannelMessageSendEmbed(privateChannel.ID, &embed)

}

func sendSpec(s *disgo.Session, m *disgo.MessageCreate) {
	spec.Timestamp = time.Now().Format(time.RFC3339)
	reply(s, m, "직접 보냈으니 확인해 보세요!")

	privateChannel, _ := s.UserChannelCreate(m.Author.ID)
	s.ChannelMessageSendEmbed(privateChannel.ID, &spec)
}
