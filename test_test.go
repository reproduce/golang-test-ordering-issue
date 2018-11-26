package mypkg

import "testing"

func TestSuccess(t *testing.T) {

}

func TestSkip(t *testing.T) {
	t.Skip("skipping test")
}

func TestFail(t *testing.T) {
	t.Fatal("failed test")
}

func TestMain(t *testing.T) {
	t.Run("success", func(t *testing.T) {

	})

	t.Run("skip", func(t *testing.T) {
		t.Skip("skipping test")
	})

	t.Run("fail", func(t *testing.T) {
		t.Fatal("failed test")
	})
}
