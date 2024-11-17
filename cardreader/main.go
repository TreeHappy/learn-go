package cardreader

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Card represents a card in the database
type Card struct {
	CardTypeID    string   `json:"card_type_id"`
	Cost          int      `json:"cost"`
	DeckLimit     int      `json:"deck_limit"`
	DesignedBy    string   `json:"designed_by"`
	FactionID     string   `json:"faction_id"`
	ID            string   `json:"id"`
	InfluenceCost int      `json:"influence_cost"`
	IsUnique      bool     `json:"is_unique"`
	SideID        string   `json:"side_id"`
	StrippedText  string   `json:"stripped_text"`
	StrippedTitle string   `json:"stripped_title"`
	Subtypes      []string `json:"subtypes"`
	Text          string   `json:"text"`
	Title         string   `json:"title"`
}

func main() {
	// Define command-line flags
	sqliteFile := flag.String("sqlite", "", "SQLite database file name")
	jsonDir := flag.String("jsondir", "", "Directory of JSON files")

	fmt.Println("test")
	// Define a usage message
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("Note: Both -sqlite and -jsondir options are required.")
	}

	flag.Parse()

	// Check if both flags are provided
	if *sqliteFile == "" || *jsonDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", *sqliteFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the cards table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS cards (
            id TEXT PRIMARY KEY,
            card_type_id TEXT,
            cost INTEGER,
            deck_limit INTEGER,
            designed_by TEXT,
            faction_id TEXT,
            influence_cost INTEGER,
            is_unique INTEGER,
            side_id TEXT,
            stripped_text TEXT,
            stripped_title TEXT,
            text TEXT,
            title TEXT
        );
    `)
	if err != nil {
		log.Fatal(err)
	}

	// Read JSON files from the specified directory
	files, err := os.ReadDir(*jsonDir)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON into a Card struct
	// Insert the card into the database
	importCardsFromDirectory(files, jsonDir, db)
}

func importCardsFromDirectory(files []fs.DirEntry, jsonDir *string, db *sql.DB) {
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(*jsonDir, file.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Error reading file %s: %v\n", filePath, err)
				continue
			}

			var card Card
			err = json.Unmarshal(data, &card)
			if err != nil {
				log.Printf("Error unmarshaling JSON in file %s: %v\n", filePath, err)
				continue
			}

			_, err = db.Exec(`
                INSERT INTO cards (
                    id,
                    card_type_id,
                    cost,
                    deck_limit,
                    designed_by,
                    faction_id,
                    influence_cost,
                    is_unique,
                    side_id,
                    stripped_text,
                    stripped_title,
                    text,
                    title
                ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
            `, card.ID, card.CardTypeID, card.Cost, card.DeckLimit, card.DesignedBy, card.FactionID, card.InfluenceCost, card.IsUnique, card.SideID, card.StrippedText, card.StrippedTitle, card.Text, card.Title)
			if err != nil {
				log.Printf("Error inserting card into database: %v\n", err)
			} else {
				log.Printf("Inserted card %s into database\n", card.ID)
			}
		}
	}
}
