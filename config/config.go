package config

import (
	"log"
	"os"
	"strings"
)

var (
	Token            = mustEnv("DISCORD_TOKEN")
	RolesByMessageID = mustMap(mustEnv("ROLES_BY_MESSAGE_ID")) // <message ID>=<role name>,<message ID>=<role name>,...
)

func mustEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("%s not set\n", name)
	}
	return value
}

func mustMap(list string) map[string]string {
	pairs := strings.FieldsFunc(list, func(c rune) bool { return c == ',' || c == '=' })
	if len(pairs)%2 != 0 {
		log.Fatalf("incomplete map value '%s'", list)
	}

	m := map[string]string{}
	for i := 0; i < len(pairs); i += 2 {
		key, value := pairs[i], pairs[i+1]
		m[key] = value
	}

	return m
}
