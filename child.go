package oao

import (
	"os"
)

func Child() (and *And, err error) {
	and = new(And)
	// out = bufio.NewWriter(os.Stdout)
	// input := bufio.NewReader(os.Stdin)
	and.stdin = os.Stdout
	and.stdout = os.Stdin
	and.stderr = os.Stderr
	return
}
