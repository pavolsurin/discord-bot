package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	dotenv "github.com/pavolsurin/discord-bot/configs"
)

var (
	BotID string
	URL   []string
	Voice *discordgo.VoiceConnection
)

func Bot() {
	session, err := discordgo.New("Bot " + dotenv.GetEnv(dotenv.DiscordToken))
	if err != nil {
		log.Fatal("Failed to create a bot: " + err.Error())
	}
	user, err := session.User("@me")
	if err != nil {
		log.Fatal("Failed to get user: " + err.Error())
	}
	BotID = user.ID
	session.AddHandler(messageHandler)
	err = session.Open()
	if err != nil {
		log.Fatal("Could not open session: " + err.Error())
	}
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	prefix := dotenv.GetEnv(dotenv.BotPrefix)

	log.Printf("Message from %s: %s", m.Author.Username, m.Content)

	if m.Content == prefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	}
	if m.Content == prefix+"join" {
		Voice = joinVoiceChannel(s, m)
		return
	}
	if m.Content == prefix+"disconnect" {
		if Voice == nil {
			s.ChannelMessageSend(m.ChannelID, "Not connected to a voice channel.")
			return
		}
		Voice.Disconnect()
		Voice = nil
		URL = nil
		return
	}
	if strings.HasPrefix(m.Content, prefix+"play") {
		Voice = joinVoiceChannel(s, m)
		if Voice == nil {
			return
		}
		content := strings.TrimPrefix(m.Content, prefix+"play ")
		URL = append(URL, content)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Added to queue: %s", content))
		return
	}
}

func joinVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate) *discordgo.VoiceConnection {
	s.ChannelMessageSend(m.ChannelID, "Trying to join the voice channel...")
	state, err := s.State.VoiceState(m.GuildID, m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "You are not in a voice channel!")
		return nil
	}
	voice, err := s.ChannelVoiceJoin(m.GuildID, state.ChannelID, false, false)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Could not join the voice channel")
		return nil
	}
	return voice
}
