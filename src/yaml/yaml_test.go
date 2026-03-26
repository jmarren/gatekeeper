package yaml

import (
	"os"
	"testing"
)

func TestYaml(t *testing.T) {
	content, err := os.ReadFile("example.yaml")

	if err != nil {
		t.Error(err)
	}
	t.Logf("content = %s\n", string(content))

	p := NewParser(string(content))

	p.Tokens()

}
