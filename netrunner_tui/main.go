package main

import (
	"fmt"
	"strconv"

	"learn-go/cardreader"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cards []cardreader.Card
}

func HardwareView(card cardreader.Card) string {
	cardWidth := 40
	greenBackground := lipgloss.Color("#005500")
	greenForeground := lipgloss.Color("#00aa00")
	costBorder := lipgloss.Border{
		Left:  "",
		Right: "",
	}
	nameBorder := lipgloss.Border{
		Left:  "",
		Right: "█",
	}
	textTabBorder := lipgloss.Border{
		Left:  "",
		Right: "",
	}
	costStyle := lipgloss.
		NewStyle().
		Background(greenBackground).
		Foreground(greenForeground).
		BorderStyle(costBorder).
		BorderLeft(true).
		BorderRight(true).
		BorderForeground(greenBackground)
	nameStyle := lipgloss.
		NewStyle().
		Foreground(greenForeground).
		Background(greenBackground).
		BorderStyle(nameBorder).
		BorderLeft(true).
		BorderRight(true).
		BorderForeground(greenBackground)
	typeTabStyle := lipgloss.
		NewStyle().
		Foreground(greenForeground).
		Background(greenBackground).
		BorderStyle(textTabBorder).
		BorderLeft(true).
		BorderRight(true).
		BorderForeground(greenBackground)
	textStyle := lipgloss.
		NewStyle().
		Foreground(greenForeground).
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(greenForeground).
		Width(cardWidth - 3).
		Height(5).
		AlignVertical(lipgloss.Center).
		AlignHorizontal(lipgloss.Center)
	influence := "\n\n\n\n\n\n"

	s := lipgloss.JoinHorizontal(lipgloss.Left, costStyle.Render(strconv.Itoa(card.Cost)), lipgloss.PlaceHorizontal(cardWidth-3, lipgloss.Right, nameStyle.Render(card.Title))) + "\n"
	s += lipgloss.PlaceHorizontal(cardWidth, lipgloss.Center, typeTabStyle.Render(card.CardTypeID)) + "\n"
	s += lipgloss.JoinHorizontal(lipgloss.Left, textStyle.Render(card.Text), influence)
	// Text
	return s
}

func View(card cardreader.Card) {
	fmt.Printf("%d, %d, %s, %d, %t, %s, %s", card.Cost, card.DeckLimit, card.FactionID, card.InfluenceCost, card.IsUnique, card.SideID, card.Text)
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Grocery List")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	return s
}

func main() {
	card := cardreader.Card{
		Title:      "Bibu",
		Cost:       4,
		Text:       "foobar",
		CardTypeID: "hardware",
	}

	fmt.Println(HardwareView(card))
	// p := tea.NewProgram(initialModel())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// }
}
