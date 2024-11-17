package tuielements

import (
	"strconv"
	"strings"

	"learn-go/cardreader"

	"github.com/charmbracelet/lipgloss"
)

const (
	ShaperBackground        = lipgloss.Color("#005500")
	ShaperForeground        = lipgloss.Color("#aaFF11")
	CriminalBackground      = lipgloss.Color("#000055")
	CriminalForeground      = lipgloss.Color("#1166FF")
	AnarchBackground        = lipgloss.Color("#550000")
	AnarchForeground        = lipgloss.Color("#FF6611")
	NeutralRunnerBackground = lipgloss.Color("#666666")
	NeutralRunnerForeground = lipgloss.Color("#aaaaaa")
)

var (
	CostBorder = lipgloss.Border{
		Left:  "",
		Right: "",
	}
	NameBorder = lipgloss.Border{
		Left:  "",
		Right: "█",
	}
	TextTabBorder = lipgloss.Border{
		Left:  "",
		Right: "",
	}
)

func GetFactionTheme(faction string) (lipgloss.Color, lipgloss.Color) {
	switch faction {
	case "shaper":
		return ShaperForeground, ShaperBackground
	case "criminal":
		return CriminalForeground, CriminalBackground
	case "anarch":
		return AnarchForeground, AnarchBackground
	case "neutral_runner":
		return NeutralRunnerForeground, NeutralRunnerBackground
	}

	return lipgloss.Color("#000000"), lipgloss.Color("#FF0000")
}

func TypeToIcon(cardtype string) string {
	switch cardtype {
	case "hardware":
		return HardwareIcon + " "

	case "resource":
		return ResourceIcon + " "

	case "event":
		return EventIcon + " "

	case "corp_identity":
		return IdentityIcon + " "

	case "runner_identity":
		return IdentityIcon + " "

	case "program":
		return ProgramIcon + " "

	case "agenda":
		return AgendaIcon + " "

	case "ice":
		return IceIcon + " "

	case "asset":
		return AssetIcon + " "

	case "upgrade":
		return UpgradeIcon + " "

	case "operation":
		return OperationIcon + " "
	}

	return "💣"
}

func HardwareView(card cardreader.Card) string {
	cardWidth := 60
	factionForeground, factionBackground := GetFactionTheme(card.FactionID)

	costStyle := lipgloss.
		NewStyle().
		Background(factionBackground).
		Foreground(factionForeground).
		BorderStyle(CostBorder).
		BorderLeft(true).
		BorderRight(true).
		BorderForeground(factionBackground)
	nameStyle := lipgloss.
		NewStyle().
		Foreground(factionForeground).
		Background(factionBackground).
		BorderStyle(NameBorder).
		BorderLeft(true).
		BorderRight(true).
		BorderForeground(factionBackground)
	typeTabStyle := lipgloss.
		NewStyle().
		Foreground(factionForeground).
		Background(factionBackground).
		BorderStyle(TextTabBorder).
		BorderLeft(true).
		BorderRight(true).
		BorderForeground(factionBackground)
	textStyle := lipgloss.
		NewStyle().
		Foreground(factionForeground).
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(factionForeground).
		Width(cardWidth - 3).
		Height(5).
		AlignVertical(lipgloss.Center).
		AlignHorizontal(lipgloss.Center)
	influenceStyle := lipgloss.
		NewStyle().
		Foreground(factionForeground)
	influence := "\n" + strings.Repeat(InfluenceEmptyChar+"\n", 5-card.InfluenceCost) + strings.Repeat(InfluenceFilledChar+"\n", card.InfluenceCost)

	s := lipgloss.JoinHorizontal(lipgloss.Left, costStyle.Render(CreditIcon+" "+strconv.Itoa(card.Cost)), lipgloss.PlaceHorizontal(cardWidth-5, lipgloss.Right, nameStyle.Render(card.Title))) + "\n"
	s += costStyle.Render(TypeToIcon(card.CardTypeID)) + lipgloss.PlaceHorizontal(cardWidth-4, lipgloss.Center, typeTabStyle.Render(card.CardTypeID)) + "\n"
	s += lipgloss.JoinHorizontal(lipgloss.Left, textStyle.Render(CardTextWithSymbols(card.Text)), influenceStyle.Render(influence))

	return s
}
