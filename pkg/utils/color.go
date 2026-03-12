package utils

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Red   = "\033[38;5;196m"
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