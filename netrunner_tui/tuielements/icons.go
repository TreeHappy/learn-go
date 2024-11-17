package tuielements

import "strings"

const (
	ClicksIcon     = "󰔚"
	TrashIcon      = "󰩹"
	CreditIcon     = ""
	VirusIcon      = "󱎷"
	ProgramIcon    = ""
	LinkIcon       = "󰿨"
	HardwareIcon   = "󰢻"
	SubroutineIcon = "↳"
)

var dictionary = map[string]string{
	"[click]":      ClicksIcon + " ",
	"[link]":       LinkIcon + " ",
	"[credit]":     CreditIcon + " ",
	"[mu]":         ProgramIcon + " ",
	"[subroutine]": SubroutineIcon + " ",
}

func CardTextWithSymbols(input string) string {
	input = strings.ToLower(input)
	for key, value := range dictionary {
		input = strings.ReplaceAll(input, strings.ToLower(key), value)
	}
	return input
}
