package startup

import (
	"log"

	"github.com/joqd/EchoNotes/internal/adapter/config"
)

func Go() {
	conf := config.NewConfig()
	log.Printf("server running on %d", conf.Server.Port)
}