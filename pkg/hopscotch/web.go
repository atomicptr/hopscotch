package hopscotch

import (
	"fmt"
	"log"
	"net/http"

	"atomicptr.dev/bits"
)

const repoUrl = "https://github.com/atomicptr/hopscotch"

func Run() error {
	host := bits.GetEnv("HOPSCOTCH_HOST", "127.0.0.1")
	port := bits.GetEnvInt("HOPSCOTCH_PORT", 3000)

	handler := NewRedirectHandler()

	err := handler.LoadRulesFromEnv()
	if err != nil {
		return fmt.Errorf("could not load entries from env: %w", err)
	}

	if len(handler.Rules) == 0 {
		return fmt.Errorf("no rules were configured, please read the documentation at: %s", repoUrl)
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("hopscoptch running on %s\n", addr)
	return http.ListenAndServe(addr, handler)
}
