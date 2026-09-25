package hopscotch

import (
	"fmt"
	"strings"
)

const envPrefix = "HOPSCOTCH_REDIRECT_"

func parseEnvString(env string) (string, string, string, error) {
	trimmed := strings.TrimPrefix(strings.TrimSpace(env), envPrefix)

	parts := strings.SplitN(trimmed, "=", 2)
	if len(parts) != 2 {
		return "", "", "", fmt.Errorf("invalid env format: %s", env)
	}

	k, v := strings.ToLower(strings.TrimSpace(parts[0])), strings.TrimSpace(parts[1])

	lastIndex := strings.LastIndex(k, "_")
	if lastIndex == -1 {
		return "", "", "", fmt.Errorf("invalid name format: %s expected like %s_{NAME}_FIELD", env, envPrefix)
	}

	name := k[:lastIndex]
	field := k[lastIndex+1:]

	return name, field, v, nil
}
