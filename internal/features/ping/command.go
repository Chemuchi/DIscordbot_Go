package ping

import "github.com/bwmarrin/discordgo"

func Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "핑",
		Description: "\"퐁\"을 출력합니다.",
	}
}
