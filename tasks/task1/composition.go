// Package task1 :
// Реализовать композицию структуры Action от
// родительской структуры Human
package task1

import (
	"fmt"
	"strings"
	"time"
)

// Human struct - base parent struct
type Human struct {
	name     string // name of human
	lastName string // lastName of human
	age      int    // age of human
	danger   string // dange is collision field

}

// describeHuman ...
func (h *Human) describeHuman() string {
	return fmt.Sprintf(
		"name : %q, lastname : %q, age : %d, dangerHuman : %q",
		h.name,
		h.lastName,
		h.age,
		h.danger,
	)
}

// Action struct is child struct for human struct
type Action struct {
	Human                    // parent composition through embedding
	actionName string        // action name
	duration   time.Duration // duration of action
	isComplete bool          // isComplete for action
	danger     string        // collision field
}

// New is constructor for Action struct
func New(humanName string, humanLastName string, humanAge int, humanDanger string,
	actionName string, actionDuration time.Duration, actionComplete bool, actionDanger string,
) *Action {
	return &Action{
		Human: Human{
			name:     humanName,
			lastName: humanLastName,
			age:      humanAge,
			danger:   humanDanger,
		},
		actionName: actionName,
		duration:   actionDuration,
		isComplete: actionComplete,
		danger:     actionDanger,
	}
}

// Describe is public method
func (a *Action) Describe() string {
	b := new(strings.Builder)
	b.WriteString(a.describeHuman())
	b.WriteString(", ")
	b.WriteString(fmt.Sprintf(
		"duration : %q, action : %q, complete : %v dangerAction: %q",
		a.duration,
		a.actionName,
		a.isComplete,
		a.danger,
	))
	return b.String()
}

// GetParent returns copy oh parent embedded entity
func (a *Action) GetParent() Human {
	return a.Human
}
