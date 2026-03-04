package discord

import "github.com/bwmarrin/discordgo"

func RegisterGlobalCommands(s *discordgo.Session, appID string, cmds []*discordgo.ApplicationCommand) ([]*discordgo.ApplicationCommand, error) {
	registered := make([]*discordgo.ApplicationCommand, 0, len(cmds))
	for _, c := range cmds {
		rc, err := s.ApplicationCommandCreate(appID, "", c) // guildID = "" → global
		if err != nil {
			return nil, err
		}
		registered = append(registered, rc)
	}
	return registered, nil
}

func DeleteGlobalCommands(s *discordgo.Session, appID string, cmds []*discordgo.ApplicationCommand) {
	for _, c := range cmds {
		_ = s.ApplicationCommandDelete(appID, "", c.ID)
	}
}
