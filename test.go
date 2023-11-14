package main

import (
	"log/slog"
	"strconv"
	"time"
)

func main() {
	slog.Info("", strconv.Itoa(int(time.Now().UnixNano()/1e6)))
}
