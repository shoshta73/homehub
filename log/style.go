package log

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	charmLogger "github.com/charmbracelet/log"
)

var style *charmLogger.Styles = &charmLogger.Styles{
	Timestamp: lipgloss.NewStyle(),
	Caller:    lipgloss.NewStyle().Faint(true),
	Prefix:    lipgloss.NewStyle().Bold(true).Faint(true),
	Message:   lipgloss.NewStyle(),
	Key:       lipgloss.NewStyle().Faint(true),
	Value:     lipgloss.NewStyle(),
	Separator: lipgloss.NewStyle().Faint(true),
	Levels: map[charmLogger.Level]lipgloss.Style{
		charmLogger.DebugLevel: lipgloss.NewStyle().
			SetString(strings.ToUpper("DEBUG")).
			Bold(true).
			MaxWidth(5).
			Foreground(lipgloss.Color("63")),
		charmLogger.InfoLevel: lipgloss.NewStyle().
			SetString(strings.ToUpper("INFO")).
			Bold(true).
			MaxWidth(4).
			Foreground(lipgloss.Color("86")),
		charmLogger.WarnLevel: lipgloss.NewStyle().
			SetString(strings.ToUpper("WARN")).
			Bold(true).
			MaxWidth(4).
			Foreground(lipgloss.Color("192")),
		charmLogger.ErrorLevel: lipgloss.NewStyle().
			SetString(strings.ToUpper("ERROR")).
			Bold(true).
			MaxWidth(5).
			Foreground(lipgloss.Color("204")),
		charmLogger.FatalLevel: lipgloss.NewStyle().
			SetString(strings.ToUpper("FATAL")).
			Bold(true).
			MaxWidth(5).
			Foreground(lipgloss.Color("134")),
	},
	Keys:   map[string]lipgloss.Style{},
	Values: map[string]lipgloss.Style{},
}
