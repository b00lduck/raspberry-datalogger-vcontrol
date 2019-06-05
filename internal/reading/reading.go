package reading

import "math"

type Reading interface {
    Process() error
}

type reading struct {
    code string
    command string
}

func round(x float64) float64 {
    return math.Floor(x + 0.5)
}