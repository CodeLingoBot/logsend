package logsend

import (
	logpkg "log"
	"os"
)

var (
	log        = logpkg.New(os.Stderr, "", logpkg.Lmicroseconds)
	Debug      = true
	SendBuffer = 50
)

var Conf = struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	UDP        bool
}{
	"localhost:8086",
	"root",
	"root",
	"test1",
	false,
}

func debug(msg ...interface{}) {
	if !Debug {
		return
	}
	log.Printf("debug: %+v", msg)
}
