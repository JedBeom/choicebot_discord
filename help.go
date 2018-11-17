package main

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	// 맨 위
	author := discordgo.MessageEmbedAuthor{
		Name:    "하코자키 세리카",
		IconURL: "https://raw.githubusercontent.com/JedBeom/choicebot_discord/master/serika.png",
	}

	// 중간의 제목 - 내용들
	fields := []*discordgo.MessageEmbedField{}

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "!선택",
		Value: "**!선택 A B** 또는 **!선택 밀리 애니 ㄷ 월희 리메이크** 또는 **!선택 A eats vs B dances**\n항목 중 하나를 선택해요!",
	}, &discordgo.MessageEmbedField{
		Name:  "!주사위",
		Value: "**!주사위 a-b**\n0을 포함한 자연수 a부터 b가지의 수 중 하나를 무작위로 뽑아요!",
	}, &discordgo.MessageEmbedField{
		Name:  "!업다운",
		Value: "**!업다운 시작**으로 업다운 게임을 시작할 수 있어요.\n숫자는 1-100의 정수이며, **!업다운 23**으로 게임을 진행할 수 있어요.",
	}, &discordgo.MessageEmbedField{
		Name:  "!세리카 버전",
		Value: "버전을 보여줘요.",
	})

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: "https://raw.githubusercontent.com/JedBeom/choicebot_discord/master/serika.png",
	}

	// 맨 아래의 정보
	footer := discordgo.MessageEmbedFooter{
		Text:    "© Bombwhale | v" + version,
		IconURL: "https://avatars2.githubusercontent.com/u/20675630?s=460&v=4",
	}

	embed = discordgo.MessageEmbed{
		Color:  0xed90ba,
		Author: &author,

		Title:       "Github 레포",
		URL:         "https://github.com/JedBeom/choicebot_discord",
		Description: "아래의 행동을 할 수 있어요!",

		Fields:    fields,
		Thumbnail: &thumbnail,

		Footer: &footer,
	}

	versionMessage = discordgo.MessageEmbed{
		Color:  0xed90ba,
		Author: &author,
		Footer: &footer,

		Title: "버전 목록",
		URL:   "https://github.com/JedBeom/blob/master/versions.md",

		Description: "현재 버전은 **v" + version + "**이에요!",
	}

}
