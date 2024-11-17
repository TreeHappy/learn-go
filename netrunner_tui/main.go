package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"learn-go/cardreader"
	"learn-go/netrunner_tui/tuielements"

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
	influenceStyle := lipgloss.
		NewStyle().
		Foreground(greenForeground)
	influenceFilledChar := "\n"
	influenceEmptyChar := "\n"
	influence := "\n" + strings.Repeat(influenceEmptyChar, 5-card.InfluenceCost) + strings.Repeat(influenceFilledChar, card.InfluenceCost)

	s := lipgloss.JoinHorizontal(lipgloss.Left, costStyle.Render(tuielements.CreditIcon+" "+strconv.Itoa(card.Cost)), lipgloss.PlaceHorizontal(cardWidth-5, lipgloss.Right, nameStyle.Render(card.Title))) + "\n"
	s += costStyle.Render(tuielements.HardwareIcon+" ") + lipgloss.PlaceHorizontal(cardWidth-4, lipgloss.Center, typeTabStyle.Render(card.CardTypeID)) + "\n"
	s += lipgloss.JoinHorizontal(lipgloss.Left, textStyle.Render(card.Text), influenceStyle.Render(influence))

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

const (
	DBPath = "/tmp/test.db"
)

func main() {
	dbExists := true
	if _, err := os.Stat(DBPath); errors.Is(err, os.ErrNotExist) {
		dbExists = false
	}
	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if !dbExists {
		cardreader.ImportCardsFromDirectory("/workspaces/learn-go/netrunner-cards-json/v2/cards/", db)
	}

	cards, _ := cardreader.ReadCardsFromDB(db)
	for _, card := range cards {
		fmt.Println(HardwareView(card))
	}
	// p := tea.NewProgram(initialModel())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// }
}
