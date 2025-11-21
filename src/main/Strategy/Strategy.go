package Strategy

type Action int

const (
	Hit Action = iota
	Stand
)

type Strategy interface {
	DecideAction(ctx GameContext) Action
	DecideBetting() int
}
