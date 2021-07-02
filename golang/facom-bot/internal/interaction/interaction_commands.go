package interaction

import (
	"facom-bot/internal/logger"

	dg "github.com/bwmarrin/discordgo"
)

func materiaCommand(s *dg.Session, i *dg.InteractionCreate) {
	if i.GuildID == "" {
		interactionReplyEphemeral(s, i, "Esse comando não funciona em DMs.")
		return
	}

	if i.Data.Name != "materia1-4" && i.Data.Name != "materia5-8" {
		logger.Error.Printf("materiaCommand received command with name %s", i.Data.Name)
		interactionReplyEphemeral(s, i, internalErrorMessage)
		return
	}

	if len(i.Data.Options) != 1 {
		logger.Error.Printf("materiaCommand received an Options slice with length %d", len(i.Data.Options))
		interactionReplyEphemeral(s, i, internalErrorMessage)
		return
	}

	roleID := i.Data.Options[0].StringValue()

	for _, memberRoleID := range i.Member.Roles {
		if memberRoleID == roleID {
			err := s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, roleID)
			if err != nil {
				logger.Error.Printf("materiaCommand failed to remove a user's role: %s", err.Error())
				interactionReplyEphemeral(s, i, internalErrorMessage)
				return
			}
			interactionReplyEphemeral(s, i, "❌ Cargo removido com sucesso")
			return
		}
	}

	err := s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, roleID)
	if err != nil {
		logger.Error.Printf("materiaCommand failed to assign a role to a user: %s", err.Error())
		interactionReplyEphemeral(s, i, internalErrorMessage)
		return
	}
	interactionReplyEphemeral(s, i, "✅ Cargo adquirido com sucesso")
}

var commands = []dg.ApplicationCommand{
	{
		Name:        "materia1-4",
		Description: "Adiciona ou remove um cargo de matéria na sua conta",
		Options: []*dg.ApplicationCommandOption{
			{
				Type:        dg.ApplicationCommandOptionString,
				Name:        "materia",
				Description: "A matéria cujo cargo você deseja adicionar ou remover",
				Required:    true,
				Choices: []*dg.ApplicationCommandOptionChoice{
					{
						Name:  "P1 - Empreendedorismo em Informática",
						Value: "742549539051012126",
					},
					{
						Name:  "P1 - Cálculo Diferencial e Integral I",
						Value: "742549533489233960",
					},
					{
						Name:  "P1 - Geometria Analítica e Álgebra Linear",
						Value: "742549536131514469",
					},
					{
						Name:  "P1 - Programação Procedimental",
						Value: "742549538249768981",
					},
					{
						Name:  "P1 - Introdução a Ciência da Computação",
						Value: "742549528703664130",
					},
					{
						Name:  "P1 - Lógica para a Computação",
						Value: "742549537322958858",
					},
					{
						Name:  "P2 - Profissão em Comput. e Informática",
						Value: "742549539394814026",
					},
					{
						Name:  "P2 - Cálculo Diferencial e Integral 2",
						Value: "742549540493852723",
					},
					{
						Name:  "P2 - Matemática para Ciência da Computação",
						Value: "742549541043306557",
					},
					{
						Name:  "P2 - Algoritmos e Estrutura de Dados I",
						Value: "742549541622120473",
					},
					{
						Name:  "P2 - Programação Lógica",
						Value: "742549542934806579",
					},
					{
						Name:  "P2 - Sistemas Digitais",
						Value: "742549543299842089",
					},
					{
						Name:  "P3 - Estatística",
						Value: "742549546587914250",
					},
					{
						Name:  "P3 - Cálculo Diferencial e Integral 3",
						Value: "742549547171184751",
					},
					{
						Name:  "P3 - Programação Funcional",
						Value: "742549544121925752",
					},
					{
						Name:  "P3 - Algoritmos e Estruturas de Dados 2",
						Value: "742549544310538311",
					},
					{
						Name:  "P3 - Programação Orientada a Objetos 1",
						Value: "742549544923037849",
					},
					{
						Name:  "P3 - Arquitetura e Organização de Computadores 1",
						Value: "742549546147774584",
					},
					{
						Name:  "P4 - Estatística Computacional",
						Value: "742549550752989224",
					},
					{
						Name:  "P4 - Teoria dos Grafos",
						Value: "742549547921703052",
					},
					{
						Name:  "P4 - Sistema de Banco de Dados",
						Value: "742549548270092370",
					},
					{
						Name:  "P4 - Linguagens Formais e Autômatos",
						Value: "742549549322731520",
					},
					{
						Name:  "P4 - Sistemas Operacionais",
						Value: "742552758129000460",
					},
					{
						Name:  "P4 - Arquitetura e Organização de Computadores 2",
						Value: "742549550035632198",
					},
				},
			},
		},
	},
	{
		Name:        "materia5-8",
		Description: "Adiciona ou remove um cargo de matéria na sua conta",
		Options: []*dg.ApplicationCommandOption{
			{
				Type:        dg.ApplicationCommandOptionString,
				Name:        "materia",
				Description: "A matéria cujo cargo você deseja adicionar ou remover",
				Required:    true,
				Choices: []*dg.ApplicationCommandOptionChoice{
					{
						Name:  "P5 - Computação Científica e Otimização",
						Value: "742549551449374800",
					},
					{
						Name:  "P5 - Análise de Algoritmos",
						Value: "742549551751364722",
					},
					{
						Name:  "P5 - Gerenciamento de Banco de Dados",
						Value: "742549552577642586",
					},
					{
						Name:  "P5 - Modelagem de Software",
						Value: "742549552875176017",
					},
					{
						Name:  "P5 - Programação Orientada a Objetos 2",
						Value: "742549553990991892",
					},
					{
						Name:  "P5 - Arquitetura de Redes de Computadores",
						Value: "742549554209095752",
					},
					{
						Name:  "P6 - Gestão Empresarial",
						Value: "742549555366854757",
					},
					{
						Name:  "P6 - Teoria da Computação",
						Value: "742549555693879428",
					},
					{
						Name:  "P6 - Inteligência Artificial",
						Value: "742549556541259787",
					},
					{
						Name:  "P6 - Engenharia de Software",
						Value: "742550986400202923",
					},
					{
						Name:  "P6 - Modelagem e Simulação",
						Value: "742551002900856893",
					},
					{
						Name:  "P6 - Arquitetura de Redes TCP/IP",
						Value: "742551003777466530",
					},
					{
						Name:  "P7 - Construção de Compiladores",
						Value: "742551005165518888",
					},
					{
						Name:  "P7 - Projeto de Graduação 1",
						Value: "742551004389572759",
					},
					{
						Name:  "P7 - Inteligência Computacional",
						Value: "742551005773692949",
					},
					{
						Name:  "P7 - Sistemas Distribuídos",
						Value: "742551006520279151",
					},
					{
						Name:  "P8 - Direito e Legislação",
						Value: "742551007116132432",
					},
					{
						Name:  "P8 - Projeto de Graduação 2",
						Value: "742551008030359642",
					},
					{
						Name:  "P8 - Segurança da Informação",
						Value: "742551008584007802",
					},
					{
						Name:  "P8 - Programação para Internet",
						Value: "742551009380794402",
					},
				},
			},
		},
	},
}
