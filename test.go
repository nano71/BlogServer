package main

import "log/slog"

var readerMap = make(map[string][]int)

func main() {
	slog.Info("", len(readerMap["11"]) == 0)
	readerMap["11"] = append(readerMap["11"], 0)
	slog.Info("", readerMap["11"])
}
