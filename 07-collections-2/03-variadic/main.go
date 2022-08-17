package main

func DebugLog(args ...string) []string {
	prepend := []string{"[DEBUG]"}
	args = append(prepend, args...)
	return args
}

func InfoLog(args ...string) []string {
	prepend := []string{"[INFO]"}
	args = append(prepend, args...)
	return args
}

func ErrorLog(args ...string) []string {
	prepend := []string{"[ERROR]"}
	args = append(prepend, args...)
	return args
}
