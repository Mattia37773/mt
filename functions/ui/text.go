/*
Copyright © 2026 Matze
*/
package ui

func Bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}

func Lime(s string) string {
	return "\033[38;5;154m" + s + "\033[0m"
}

func Pink(s string) string {
	return "\033[38;5;198m" + s + "\033[0m"
}

func GlowPink(s string) string {
	return Bold("\033[38;2;255;20;147m" + s + "\033[0m")
}

func GlowPurple(s string) string {
	return Bold("\033[38;2;180;0;255m" + s + "\033[0m")
}

func Purple(s string) string {
	return "\033[38;5;93m" + s + "\033[0m"
}

func Blue(s string) string {
	return "\033[34m" + s + "\033[0m" /// Terminal
}

func Red(s string) string {
	return "\033[38;5;196m" + s + "\033[0m"
}

func Green(s string) string {
	return "\033[32m" + s + "\033[0m" // Terminal
}
