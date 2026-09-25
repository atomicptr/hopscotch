package hopscotch

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type RedirectRule struct {
	Name      string
	From      string
	To        string
	Permanent bool
}

type RedirectHandler struct {
	Rules map[string]RedirectRule
}

func NewRedirectHandler() RedirectHandler {
	return RedirectHandler{
		Rules: make(map[string]RedirectRule),
	}
}

func (handler *RedirectHandler) LoadRulesFromEnv() error {
	tmpRules := make(map[string]RedirectRule)

	for _, env := range os.Environ() {
		if !strings.HasPrefix(env, envPrefix) {
			continue
		}

		name, field, value, err := parseEnvString(env)
		if err != nil {
			log.Printf("could not parse %s: %s\n", env, err)
			continue
		}

		err = applyRule(tmpRules, name, field, value)
		if err != nil {
			log.Printf("could not apply %s: %s\n", env, err)
			continue
		}
	}

	for name, rule := range tmpRules {
		if rule.From == "" {
			return fmt.Errorf("rule: '%s' does not have 'from' field, did you forget to set '%s%s_FROM'?", name, envPrefix, strings.ToUpper(name))
		}

		if rule.To == "" {
			return fmt.Errorf("rule: '%s' does not have 'to' field, did you forget to set '%s%s_TO'?", name, envPrefix, strings.ToUpper(name))
		}

		handler.Rules[rule.From] = rule

		log.Printf("Rule: %s | From: %s | To: %s | Permanent: %t", name, rule.From, rule.To, rule.Permanent)
	}

	return nil
}

func applyRule(m map[string]RedirectRule, name, field, value string) error {
	rule, ok := m[name]
	if !ok {
		m[name] = RedirectRule{}
		rule = m[name]
	}

	rule.Name = name

	switch field {
	case "from":
		rule.From = cleanUrl(value)
	case "to":
		rule.To = cleanUrl(value)
	case "permanent":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("%s permanent: could not parse value '%s' to bool", name, value)
		}

		rule.Permanent = val
	default:
		return fmt.Errorf("unknown field '%s'", field)
	}

	m[name] = rule

	return nil
}

func cleanUrl(url string) string {
	u := url
	u = strings.ToLower(u)
	u = strings.TrimPrefix(u, "https://")
	u = strings.TrimPrefix(u, "http://")

	parts := strings.Split(u, "/")
	return parts[0]
}

func (handler RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	entry, ok := handler.Rules[r.Host]
	if !ok {
		http.NotFound(w, r)
		return
	}

	code := http.StatusTemporaryRedirect
	if entry.Permanent {
		code = http.StatusPermanentRedirect
	}

	scheme := "https"
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		scheme = proto
	}

	path := r.URL.RequestURI()

	target := fmt.Sprintf("%s://%s%s", scheme, entry.To, path)

	log.Printf("Hit rule %s: redirecting %s -> %s", entry.Name, fmt.Sprintf("%s%s", r.Host, path), target)
	http.Redirect(w, r, target, code)
}
