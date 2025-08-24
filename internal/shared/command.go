package shared

type Command struct {
	Name string
	Args []string
}


type CommandCallbackRegistry struct {
	Reg map[string]func(*State, Command) error
}
