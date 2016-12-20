package reading

type Reading interface {
    Process() error
}

type reading struct {
    oldValue float64
    code string
    command string
}