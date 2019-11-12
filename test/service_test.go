package test

import (
	"testing"
)

func TestError(t *testing.T) {
	t.Error("error")
}

func TestOK(t *testing.T) {
	t.Log("ok")
}
