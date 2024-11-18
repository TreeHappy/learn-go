package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"learn-go/cardreader"
	"learn-go/imagechopper"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cards []cardreader.Card
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

func RenderImage(imagePath string) string {
	imageBytes, _ := os.ReadFile(imagePath)
	base64Image := base64.StdEncoding.EncodeToString(imageBytes)

	return fmt.Sprintf("\033]1337;File=height=10;inline=1:%s\007", base64Image)
}

func main() {
	imagechopper.Dodido()
	fmt.Println(RenderImage("/tmp/11106.jpg"))
	// dbExists := true
	// if _, err := os.Stat(DBPath); errors.Is(err, os.ErrNotExist) {
	// 	dbExists = false
	// }
	// // Connect to the SQLite database
	// db, err := sql.Open("sqlite3", DBPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// if !dbExists {
	// 	cardreader.ImportCardsFromDirectory("/workspaces/learn-go/netrunner-cards-json/v2/cards/", db)
	// }
	//
	// cards, _ := cardreader.ReadCardsFromDB(db)
	// for _, card := range cards {
	// 	fmt.Println(tuielements.HardwareView(card))
	// }

	// p := tea.NewProgram(initialModel())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// }
}
