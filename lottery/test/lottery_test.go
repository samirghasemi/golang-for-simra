package test

import (
	"testing"

	"github.com/samirghasemi/golang-for-simra/lottery/pkg/lottery"
)

func TestDraw(t *testing.T) {
	prizes := []lottery.Prize{
		{Name: "A", Weight: 0.1},
		{Name: "B", Weight: 0.3},
		{Name: "C", Weight: 0.2},
		{Name: "D", Weight: 0.15},
		{Name: "E", Weight: 0.25},
	}

	result := lottery.Draw(prizes)
	if result == "" {
		t.Errorf("Expected a prize, but got nothing")
	}
}
