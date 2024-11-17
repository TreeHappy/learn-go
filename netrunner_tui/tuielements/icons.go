package tuielements

import "strings"

const (
	ClicksIcon          = "󰔚"
	TrashIcon           = "󰩹"
	CreditIcon          = ""
	RecurringCreditIcon = ""
	VirusIcon           = "󱎷"
	ProgramIcon         = ""
	LinkIcon            = "󰿨"
	HardwareIcon        = "󰢻"
	SubroutineIcon      = "󱞩"
	InfluenceFilledChar = ""
	InfluenceEmptyChar  = ""
	ResourceIcon        = ""
	EventIcon           = "󰛕"
	IdentityIcon        = ""
	AssetIcon           = ""
	IceIcon             = ""
	OperationIcon       = ""
	AgendaIcon          = ""
	UpgradeIcon         = ""
	InterruptIcon       = ""
)

var iconDictionary = map[string]string{
	"[click]":            ClicksIcon + " ",
	"[link]":             LinkIcon + " ",
	"[credit]":           CreditIcon + " ",
	"[mu]":               ProgramIcon + " ",
	"[subroutine]":       SubroutineIcon + " ",
	"[recurring-credit]": ResourceIcon + " ",
	"[trash]":            TrashIcon + " ",
	"[interrupt]":        InterruptIcon + " ",
	"<strong>":           "",
	"</strong>":          "",
	"→":                  "",
}

func CardTextWithSymbols(input string) string {
	input = strings.ToLower(input)
	for key, value := range iconDictionary {
		input = strings.ReplaceAll(input, strings.ToLower(key), value)
	}
	return input
}
