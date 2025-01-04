/*
####### sdk.jw (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package jw

import (
	"strconv"
	"strings"
)

const (
	_defaultName   = "without"
	_defaultOrigin = "without"
)

type (
	Status string
)

const (
	StatusTodo      Status = "todo"
	StatusRunning   Status = "running"
	StatusPending   Status = "pending"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
)

func (s Status) String() string {
	return string(s)
}

type (
	Priority int
)

const (
	PriorityNone     Priority = 0
	PriorityLow      Priority = 20
	PriorityMedium   Priority = 50
	PriorityHigh     Priority = 80
	PriorityCritical Priority = 100
)

func StringToPriority(priority string) Priority {
	switch strings.ToLower(priority) {
	case "low":
		return PriorityLow
	case "medium":
		return PriorityMedium
	case "high":
		return PriorityHigh
	case "critical":
		return PriorityCritical
	default:
		return PriorityNone
	}
}

func (p *Priority) Fix() {
	if *p < PriorityNone {
		*p = PriorityNone
	} else if *p > PriorityCritical {
		*p = PriorityCritical
	}
}

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "low"
	case PriorityMedium:
		return "medium"
	case PriorityHigh:
		return "high"
	case PriorityCritical:
		return "critical"
	default:
		return strconv.Itoa(int(p))
	}
}

/*
####### END ############################################################################################################
*/
