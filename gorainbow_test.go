package gorainbow

import (
	"testing"
)

func TestGorainbow(t *testing.T) {
	expected := "\x1b[38;2;128;237;18mh\x1b[0m\x1b[38;2;140;231;12me\x1b[0m\x1b[38;2;153;223;7ml\x1b[0m\x1b[38;2;165;214;4ml\x1b[0m\x1b[38;2;177;204;1mo\x1b[0m\x1b[38;2;188;194;1m \x1b[0m\x1b[38;2;199;182;1mw\x1b[0m\x1b[38;2;209;171;2mo\x1b[0m\x1b[38;2;219;159;5mr\x1b[0m\x1b[38;2;227;146;9ml\x1b[0m\x1b[38;2;234;133;15md\x1b[0m"
	output := Rainbow("hello world")

	if expected != output {
		t.Errorf("The text should match the expected input.\nExpected: %s\nOutput: %s", expected, output)
	}
}
