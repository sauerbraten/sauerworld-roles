package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/sauerbraten/sauerworld-roles/config"
)

func main() {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v\n", err)
	}

	session.AddHandler(func(_ *discordgo.Session, _ *discordgo.Connect) {
		log.Println("connected to Discord")
	})
	session.AddHandler(func(_ *discordgo.Session, _ *discordgo.Disconnect) {
		log.Println("disconnected from Discord")
	})
	session.AddHandler(messageReactionAdded)
	session.AddHandler(messageReactionRemoved)

	session.Identify.Intents = discordgo.IntentsGuildMessageReactions

	err = session.Open()
	if err != nil {
		log.Fatalf("error opening connection: %v\n", err)
	}

	// wait for kill signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Println("received interrupt, shutting down")

	err = session.Close()
	if err != nil {
		log.Printf("error closing connection: %v\n", err)
	}
}

func messageReactionAdded(s *discordgo.Session, a *discordgo.MessageReactionAdd) {
	toggleRole(s, a.MessageReaction, true)
}

func messageReactionRemoved(s *discordgo.Session, a *discordgo.MessageReactionRemove) {
	toggleRole(s, a.MessageReaction, false)
}
