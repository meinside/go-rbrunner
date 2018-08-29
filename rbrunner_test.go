package rbrunner

import (
	"fmt"
	"strconv"
	"testing"
)

const code = `#!/usr/bin/env ruby
# The answer to the life, the universe and everything:

print "The answer is #{ARGV[0].to_i * ARGV[1].to_i}."
`

const resultFormat = `The answer is %d.`

func TestRun(t *testing.T) {
	arg1, arg2 := 6, 7
	expected := fmt.Sprintf(resultFormat, arg1*arg2)

	result := Run(code, strconv.Itoa(arg1), strconv.Itoa(arg2))

	if result.ExitStatus != 0 {
		t.Errorf("Failed to run script (%d): %s", result.ExitStatus, result.StderrOutput)
	} else if result.StdoutOutput != expected {
		t.Errorf("Result is not as expected: '%s' (should be: '%s')", result.StdoutOutput, expected)
	}
}
