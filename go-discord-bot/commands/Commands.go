package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Response struct {
	Name    string    `json:"name"`
	Id      int       `json:"id"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"id"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

func HelloWorld(s *discordgo.Session, m *discordgo.MessageCreate) {
	//ignore bots messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".helloworld" {
		s.ChannelMessageSend(m.ChannelID, "Hello World")
	}
}

func Time(s *discordgo.Session, m *discordgo.MessageCreate) {
	//ignore bots messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".time" {
		s.ChannelMessageSend(m.ChannelID, time.Now().Format(time.RFC822))
	}
}

func RandomPokemon(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == ".pokemon" {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1008)
		response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(n))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var data Response
		json.Unmarshal(responseData, &data)

		out := "Pokemon #" + strconv.Itoa(data.Id) + " " + data.Name
		s.ChannelMessageSend(m.ChannelID, out)
	}
}
