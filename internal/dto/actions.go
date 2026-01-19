package dto

type userStates int

const (
	StateIdle userStates = iota
	StateWaitingFace
	StateWaitingTarget
)

// tracking actions
type Action struct {
	states map[int64]userStates
	face   map[int64]string
	target map[int64]string
}

// initialize actions
func NewAction() Action {
	return Action{
		states: make(map[int64]userStates),
		face:   make(map[int64]string),
		target: make(map[int64]string),
	}
}

// get State
func (a *Action) State(chatID int64) userStates {
	return a.states[chatID]
}

// set for waiting face state
func (a *Action) SetWaitingFaceState(chatID int64) {
	a.states[chatID] = StateWaitingFace
}

// set for waiting target state
func (a *Action) SetWaitingTargetState(chatID int64) {
	a.states[chatID] = StateWaitingTarget
}

// set for idle state
func (a *Action) SetIdleState(chatID int64) {
	a.states[chatID] = StateIdle
}

// add Face
func (a *Action) AddFace(chatID int64, face string) {
	a.face[chatID] = face
}

// add target
func (a *Action) AddTarget(chatID int64, target string) {
	a.target[chatID] = target
}

func (a *Action) Get(chatID int64) (string, string) {
	return a.face[chatID], a.target[chatID]
}
