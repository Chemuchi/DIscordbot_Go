package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"discord-bot/internal/discord"
	"discord-bot/internal/features/ping"
)

func mustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing %s", key)
	}
	return v
}

func main() {
	token := mustGetenv("BOT_TOKEN")
	appID := mustGetenv("APPLICATION_ID")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
	dg.Identify.Intents = discordgo.IntentsGuilds

	router := discord.NewRouter()
	router.Register("핑", ping.Handler)
	dg.AddHandler(router.OnInteractionCreate)

	if err := dg.Open(); err != nil {
		log.Fatal(err)
	}
	defer dg.Close()

	// 🔥 Global 커맨드 등록
	registered, err := discord.RegisterGlobalCommands(dg, appID, []*discordgo.ApplicationCommand{
		ping.Command(),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 운영에서는 보통 삭제 안함 (지금은 개발용이니까 유지)
	defer discord.DeleteGlobalCommands(dg, appID, registered)

	log.Println("bot is running. Ctrl+C to stop.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
