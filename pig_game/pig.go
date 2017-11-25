package main

import (
  "fmt"
  "math/rand"
)

const (
  win = 100 // points to win
  gamesPerSeries = 10 // number of games per series to simulate
)

type score struct {
  player, opponent, thisTurn int
}

type action func(currentScore score) (result score, turnIsOver bool)

func roll(s score) (score, bool) {
  outcome := rand.Intn(6) + 1
  if outcome == 1 {
    return score{s.opponent, s.player, 0}, true
  }
  return score{s.player, s.opponent, outcome + s.thisTurn}, false
}

// stay returns the (result, turnIsOver) outcome of staying.
func stay(s score) (score, bool) {
  return score{s.opponent, s.player + s.thisTurn, 0}, true
}

type strategy func(score) action

func stayAtK(k int) strategy {
  return func(s score) action{
    if s.thisTurn >= k {
      return stay
    }
    return roll
  }
}

// play simulates a Pig game and returns the winner (0 or 1).
func play(s1 strategy, s2 strategy) int {
  strategies := []strategy{s1, s2}
  var s score
  var turnIsOver bool
  currentPlayer := rand.Intn(2)
  for s.player + s.thisTurn < win {
    s, turnIsOver = strategies[currentPlayer](s)(s)
    if turnIsOver{
      currentPlayer = (currentPlayer + 1) % 2
    }
  }
  return currentPlayer
}

// round-robin 循环制
func roundRobin(strategies []strategy) ([]int, int) {
  wins := make([]int, len(strategies))
  for i := 0; i < len(strategies); i++ {
    for j := i + 1; j < len(strategies); j++ {
      for k := 0; k < gamesPerSeries; k++ {
        winner := play(strategies[i], strategies[j])
        if winner == 0 {
          wins[i]++
        } else {
          wins[j]++
        }
      }
    }
  }
  gamesPerStrategy := gamesPerSeries * (len(strategies) - 1)
  return wins, gamesPerStrategy
}

func ratioString(vals ...int) string {
	total := 0
	for _, val := range vals {
		total += val
	}
	s := ""
	for _, val := range vals {
		if s != "" {
			s += ", "
		}
		pct := 100 * float64(val) / float64(total)
		s += fmt.Sprintf("%d/%d (%0.1f%%)", val, total, pct)
	}
	return s
}

func main() {
	strategies := make([]strategy, win)
	for k := range strategies {
		strategies[k] = stayAtK(k + 1)
	}
	wins, games := roundRobin(strategies)

	for k := range strategies {
		fmt.Printf("Wins, losses staying at k =% 4d: %s\n",
			k+1, ratioString(wins[k], games - wins[k]))
	}
}
