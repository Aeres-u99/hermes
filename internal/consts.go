package internal

type Call struct {
	Caller string
	Callee string
}

type Dependency struct {
	Source string
	Target string
}
