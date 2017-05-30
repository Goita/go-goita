package commands

// Command holds command name and activity
type Command struct {
	Name string
	Run  func(args []string)
}
