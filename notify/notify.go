package notify

import (
	"runtime"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Severity int

const (
	SeverityLow Severity = iota
	SeverityNormal
	SeverityUrgent
)

type Notify struct {
	title    string
	message  string
	severity Severity
}

func New(title, message string, severity Severity) *Notify {
	return &Notify{
		title:    title,
		message:  message,
		severity: severity,
	}
}

func (s Severity) String() string {
	sev := "low"

	switch s {
	case SeverityLow:
		sev = "low"
	case SeverityNormal:
		sev = "normal"
	case SeverityUrgent:
		sev = "critical"
	}

	if runtime.GOOS == "darwin" {
		caser := cases.Title(language.English)
		sev = caser.String(sev)
	}

	if runtime.GOOS == "windows" {
		switch s {
		case SeverityLow:
			sev = "Info"
		case SeverityNormal:
			sev = "Warning"
		case SeverityUrgent:
			sev = "Error"
		}
	}

	return sev
}
