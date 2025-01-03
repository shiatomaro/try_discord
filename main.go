package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"golang.org/x/exp/rand"
)

// command prefix to idetify gotrybot bot commands
const prefix string = "!gotrybot"

func main() {
	// get environmental variables from a .env file
	godotenv.Load()

	// retrieve the token from the environment.
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("No discord bot token provided. Set DISCORD_BOT_TOKEN environment")
	}

	// create a new discord session using bot token
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	// add message handler for processing commands
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		log.Println("Message received: ", m.Content)
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")

		if args[0] != prefix {
			return
		}

		if args[1] == "Commands" {
			s.ChannelMessageSend(m.ChannelID, "gohappy: jokes")
		}
		if args[1] == "gohappy" {
			gohappy := []string{
				"How do trees get on the Internet? They log in.",
				"What do computers like to eat? Chips.",
				"What do you call a space magician? A flying saucerer.",
				"What is a computer’s first sign of old age? Loss of memory.",
				"What does a baby computer call his father? Instead of Da-da, it says 'Da-ta.'",
				"What is an astronaut’s favorite control on the computer keyboard? The space bar.",
				"What happened when the computer fell on the floor? It slipped a disk.",
				"How does a boy cell phone propose to his girlfriend? He gives her a ring, of course.",
				"Why was there a bug in the computer? It was looking for a byte to eat.",
				"What is a computer virus? A terminal illness.",
				"How did the mouse get out of the Roman Cathedral? He clicked on an icon and opened a window.",
				"What kind of doctor fixes broken websites? A URLologist.",
				"Have you heard about the Disney virus? It makes everything on your computer go Goofy.",
				"What happened when a dragon breathed on several Macintosh computers? He wound up with baked Apples!",
				"Why did the chicken cross the Web? To get to the other site.",
				"Why did the computer go to a doctor? It thought it had a terminal illness.",
				"Knock-knock! Who’s there? I-M. I-M, who? I-M on the computer, and I can’t answer the door.",
			}

			selection := rand.Intn(len(gohappy))

			author := discordgo.MessageEmbedAuthor{
				Name: "Uknown",
				URL:  "https://www.elon.edu/u/imagining/about/kidzone/jokes-laughs/",
			}

			embed := discordgo.MessageEmbed{
				Title:  gohappy[selection],
				Author: &author,
			}

			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		}
	})

	// set the bot's intents to listen for all non-priveleged events
	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// start session to establish a connection with discord
	err = sess.Open()
	if err != nil {
		log.Fatalf("Error in creating a discord bot session: %v", err)
	}
	defer sess.Close()

	fmt.Println("Bot is currently online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	log.Println("Bot shutting down..")

}
