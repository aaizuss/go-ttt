package game

import (
	"testing"
)

func TestHumanHuman(t *testing.T) {
	players := humanHuman()

	if players[0].isHuman == false || players[1].isHuman == false {
		t.Errorf("Expected both players to be human, got %v", players)
	}
}

func TestAiHuman(t *testing.T) {
	players := aiHuman()

	if players[0].isHuman {
		t.Errorf("Expected first player to be ai, got %v", players[0])
	}

	if !players[1].isHuman {
		t.Errorf("Expected second player to be human, got %v", players[0])
	}
}
