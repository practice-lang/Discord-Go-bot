package bot

import (
	"../config"
	"fmt"
  "github.com/bwmarrin/discordgo"
)

var BotID string 
var gotBat *discordgo.Session

func Start() {
  gotBat, err := discordgo.New("Bot " + config.Token)

  if err != nil {
	  fmt.Println(err.Error())
	  return
  }

  u, err := gotBat.User("@me")

  if err != nil {
	 fmt.Println(err.Error())
	 return
  }

  BotID = u.ID

  gotBat.AddHandler(messageHandler)

  err = gotBat.Open()

  if err != nil {
	fmt.Println(err.Error())
	return
  }
  fmt.Println("Bot ready")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
  fmt.Println(m.Author.Username + "#" + m.Author.Discriminator + "("+ m.Author.ID + ")"+ ": " + m.Content)

  if m.Author.ID == BotID {
	  return
  }

  if m.Content == "ping" {
	  _, _ = s.ChannelMessageSend(m.ChannelID, "pong")
  }
  if m.Content == "ㅇㅇ" {
    _, _ = s.ChannelMessageSend(m.ChannelID, m.Author.Username)
  }
}
