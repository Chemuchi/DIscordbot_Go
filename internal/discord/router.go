package discord

import "github.com/bwmarrin/discordgo"

type Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)

type Router struct {
	handlers map[string]Handler
}

func NewRouter() *Router {
	return &Router{handlers: map[string]Handler{}}
}

func (r *Router) Register(name string, h Handler) {
	r.handlers[name] = h
}

func (r *Router) OnInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	name := i.ApplicationCommandData().Name
	h, ok := r.handlers[name]
	if !ok {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Unknown command",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	h(s, i)
}
