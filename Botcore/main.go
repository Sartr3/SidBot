package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

var (
	token string
)

func init() {

	// carrega variavel de ambiente;
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}

	token = os.Getenv("TOKEN")
}

func main() {

	// cria uma sessão do discord;
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// define uma função de tratamento para o evento "ready";
	dg.AddHandler(ready)

	// abre uma conexão com o discord;
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	// espera por um sinal de interrupção;
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// encerra conexão com o discord;
	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// função chamada quando o bot estiver pronto;
	fmt.Println("Bot pronto. Nome: ", s.State.User.Username)
}

