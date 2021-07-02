package interaction

import (
	"facom-bot/internal/logger"

	dg "github.com/bwmarrin/discordgo"
)

type interactionFunc func(s *dg.Session, i *dg.InteractionCreate)

const internalErrorMessage = "Ocorreu um erro interno do bot."

var interactionHandlersMap = map[string]interactionFunc{
	"materia1-4": materiaCommand,
	"materia5-8": materiaCommand,
	"ufu":        ufuCommand,
}

// InteractionHandler calls the adequate function for a given Slash Command
func InteractionHandler(s *dg.Session, i *dg.InteractionCreate) {
	logger.Info.Printf("Command %s called", i.Data.Name)
	if handler, ok := interactionHandlersMap[i.Data.Name]; ok {
		handler(s, i)
		return
	}

	interactionReplyEphemeral(s, i, "𐄂 Este comando não deveria existir")
}

// interactionReplyEphemeral replies to the interaction with the given message and only allows the caller of the
// to see the reply
func interactionReplyEphemeral(s *dg.Session, i *dg.InteractionCreate, msg string) error {
	return s.InteractionRespond(
		i.Interaction,
		&dg.InteractionResponse{
			Type: dg.InteractionResponseChannelMessageWithSource,
			Data: &dg.InteractionApplicationCommandResponseData{
				Content: msg,
				Flags:   64,
			},
		})
}

// interactionReply replies to the interaction with a regular message.
func interactionReply(s *dg.Session, i *dg.InteractionCreate, msg string) error {
	return s.InteractionRespond(
		i.Interaction,
		&dg.InteractionResponse{
			Type: dg.InteractionResponseChannelMessageWithSource,
			Data: &dg.InteractionApplicationCommandResponseData{
				Content: msg,
			},
		})
}

// RegisterCommands registers all commands supported by the program
func RegisterCommands(s *dg.Session) {
	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "564863853331218433", &cmd)
		if err != nil {
			logger.Error.Fatalf("Failed to register command %s: %s", cmd.Name, err.Error())
		}
	}
}
