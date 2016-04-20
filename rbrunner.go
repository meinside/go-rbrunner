package rbrunner

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

const (
	tag = "rbrunner"

	ExitStatusNotSet = -1
	TempDir          = "/tmp"
	RubyExec         = "ruby"
)

type RunResult struct {
	ExitStatus   int
	StdoutOutput string
	StderrOutput string
}

func genTempFile(code string) (filepath string) {
	filepath = ""

	if file, err := ioutil.TempFile(TempDir, "rbrunner_"); err == nil {
		defer file.Close()

		if _, err := file.WriteString(code); err == nil {
			filepath = file.Name()
		}
	}

	return filepath
}

// Run code with given parameters
func Run(code string, params ...string) (result RunResult) {
	result.ExitStatus = ExitStatusNotSet

	// generate tempfile
	if rbFilepath := genTempFile(code); rbFilepath != "" {
		execParams := []string{rbFilepath}
		execParams = append(execParams, params...)
		cmd := exec.Command(RubyExec, execParams...)

		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		// referenced: http://stackoverflow.com/questions/10385551/get-exit-code-go
		if err := cmd.Start(); err == nil {
			// read from stdout
			if bytes, err := ioutil.ReadAll(stdout); err == nil {
				result.StdoutOutput = string(bytes)
			} else {
				result.StdoutOutput = err.Error()
			}

			// read from stderr
			if bytes, err := ioutil.ReadAll(stderr); err == nil {
				result.StderrOutput = string(bytes)
			} else {
				result.StderrOutput = err.Error()
			}

			// get exit status
			if err := cmd.Wait(); err != nil {
				if exitErr, ok := err.(*exec.ExitError); ok {
					if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
						result.ExitStatus = status.ExitStatus()
					}
				} else {
					fmt.Printf("%s: Failed to get exit status - %s\n", tag, err)
				}
			} else {
				result.ExitStatus = 0
			}
		}

		// delete tempfile
		if err := os.Remove(rbFilepath); err != nil {
			fmt.Printf("%s: Failed to delete tempfile - %s\n", tag, rbFilepath)
		}
	} else {
		fmt.Printf("%s: %s\n", tag, "Failed to generate tempfile")
	}

	return result
}
