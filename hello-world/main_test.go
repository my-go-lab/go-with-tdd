package main

/**
 * 실행방법
 * $ go test
 *
 * 참고
 * https://y0c.github.io/2019/08/11/beginning-go/
 * https://github.com/stretchr/testify
 */

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	expect := "Hello, World!"

	if got != expect {
		t.Errorf("got %q expect %q", got, expect)
	}
}

func TestHello(t *testing.T) {
	got := Hello("Burt.K")
	expect := "Hello, Burt.K!"

	if got != expect {
		t.Errorf("got %q expect %q", got, expect)
	}
}
