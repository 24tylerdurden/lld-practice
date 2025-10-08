package services

import (
	"LLD-PRACTICE/text_editor/models"
	"sync"
)

type HistoryManager struct {
	mu         sync.Mutex
	UndoStack  []models.Command
	RedoStack  []models.Command
	maxHistory int
}

func NewHistoryManager(maxHistory int) *HistoryManager {
	return &HistoryManager{
		maxHistory: maxHistory,
	}
}

func (hm *HistoryManager) ExecuteCommand(cmd models.Command) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	cmd.Execute()

	hm.UndoStack = append(hm.UndoStack, cmd)

	if len(hm.UndoStack) > hm.maxHistory {
		hm.UndoStack = hm.UndoStack[1:]
	}

	hm.RedoStack = nil
}

func (hm *HistoryManager) Redo() bool {

	hm.mu.Lock()

	defer hm.mu.Unlock()

	if len(hm.RedoStack) == 0 {
		return false
	}

	//  Pop last Cmd
	lastCmd := hm.RedoStack[len(hm.RedoStack)-1]
	hm.RedoStack = hm.RedoStack[:len(hm.RedoStack)-1]

	lastCmd.Redo()
	return true

}

func (hm *HistoryManager) Undo() bool {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	if len(hm.UndoStack) == 0 {
		return false
	}

	lastCmd := hm.UndoStack[len(hm.UndoStack)-1]
	hm.UndoStack = hm.UndoStack[:len(hm.UndoStack)-1]

	lastCmd.Undo()
	hm.RedoStack = append(hm.RedoStack, lastCmd)

	return true
}

// there are two stack

//  undo stack - Undo
// [cmd, cmd, cmd]
// [cmd, cmd, cmd]
