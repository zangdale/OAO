package oao

import (
	"os/exec"
)

func Parent(cmd *exec.Cmd) (and *And, err error) {
	and.stdin, err = cmd.StdinPipe()
	if err != nil {
		return
	}
	and.stdout, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	and.stderr, err = cmd.StderrPipe()
	return
}
