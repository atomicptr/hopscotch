package hopscotch

import "testing"

func assertParseEnvOk(t *testing.T, env, name, field, value string) {
	n, f, v, err := parseEnvString(env)
	if err != nil {
		t.Errorf("expected `parseEnvString(\"%s\")` to return no error, returned: %s", env, err)
		return
	}

	if n != name {
		t.Errorf("expected `parseEnvString(\"%s\")` name to be: %s, got: %s instead", env, name, n)
		return
	}

	if f != field {
		t.Errorf("expected `parseEnvString(\"%s\")` field to be: %s, got: %s instead", env, field, f)
		return
	}

	if v != value {
		t.Errorf("expected `parseEnvString(\"%s\")` value to be: %s, got: %s instead", env, value, v)
		return
	}
}

func assertParseEnvFail(t *testing.T, env string) {
	_, _, _, err := parseEnvString(env)
	if err == nil {
		t.Errorf("expected `parseEnvString(\"%s\")` to return an error", env)
		return
	}
}

func TestParseEnvString(t *testing.T) {
	assertParseEnvOk(t, "HOPSCOTCH_REDIRECT_EXAMPLE_FROM=example.com", "example", "from", "example.com")
	assertParseEnvOk(t, "HOPSCOTCH_REDIRECT_EXAMPLE_WITH_UNDERSCORES_FROM=example.com", "example_with_underscores", "from", "example.com")

	// failure cases
	assertParseEnvFail(t, "HOPSCOTCH_REDIRECT_EXAMPLE=example.com")
	assertParseEnvFail(t, "HOPSCOTCH_REDIRECT_EXAMPLE_FROM")
}
