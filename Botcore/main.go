package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	token string
)

func init() {
	// Carrega variável de ambiente;
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar variável de ambiente.")
	}

	token = os.Getenv("TOKEN")
}

func main() {
	// Cria uma sessão do Discord
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Erro ao conectar bot no disc: ", err)
		return
	}

	// Define a função de tratamento para o evento "messageCreate"
	sess.AddHandler(MessageCreate)

	// Abre uma conexão com o Discord
	err = sess.Open()
	if err != nil {
		fmt.Println("Erro ao abrir sessão no disc:  ", err)
		return
	}

	fmt.Println("Bot ta milhão. Aperte CTRL + C para Parar o menino.")

	// Espera por um sinal de interrupção
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Encerra conexão com o Discord
	sess.Close()
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	for _, mention := range m.Mentions {
		if mention.ID == s.State.User.ID {
			handleCommand(s, m)
			return
		}
	}
}

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.TrimSpace(strings.TrimPrefix(m.ContentWithMentionsReplaced(), "<@!"+s.State.User.ID+">"))

	if content == "" {
		return
	}

	parts := strings.Fields(content)
	if len(parts) == 0 {
		return
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "bora?":
		s.ChannelMessageSend(m.ChannelID, "Bora tomar uma.")
	case "diz":
		if len(args) == 0 {
			s.ChannelMessageSend(m.ChannelID, "Repetir o que, meu fi?")
			return
		}

		echoText := strings.Join(args, " ")
		s.ChannelMessageSend(m.ChannelID, echoText)
	}
}
