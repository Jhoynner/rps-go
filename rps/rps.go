package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // Pieda. Vence a las tijeras. (tijeras + 1) % 3 = 0
	PAPER    = 1 // Papel. Vence a la piedra (piedra + 1) % 3 = 1
	SCISSORS = 2 // Tijra. Vence al papel (papel + 1) % 3 = 2
)

// Estructura para dar resultado de cada ronda
type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// Mensajes para cuando gana
var winMessages = []string{
	"Buen trabajo",
	"Estas OnFire",
	"Echale al baloto",
}

// Mensajes para cuando pierde
var loseMessages = []string{
	"Que lastima",
	"Eso fue suerte, intentalo una vez mas",
	"Mejor ya dejalo",
}

// Mensajes para cuando empata
var drawMessages = []string{
	"Te esta leyendo, ganale",
	"No te dejes, ganale",
	"John Connor no vayas a perder, ganale",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	// Mensaje dependiendo de lo que eligio la computadora
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligio TIJERAS"
	}

	// Generar un numero aleatorio de 0-2, que usamos para elegir un mensaje aleatorio
	messageInt := rand.Intn(3)

	// Declarar una var para contener el mensaje
	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana"
		message = loseMessages[messageInt]
	}

	// Retornar resultado
	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
