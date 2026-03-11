package server

const (
	Red   = "\033[38;5;88m"
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Blue  = "\033[38;5;27m"
	Green = "\033[38;5;46m"
	ClearScreen = "\033[2J\033[H"
)

func SetGreen(str string) string {
	return Green + str + Reset
}

func SetBlue(str string) string {
	return Blue + str + Reset
}

func SetRed(str string) string {
	return Red + str + Reset
}