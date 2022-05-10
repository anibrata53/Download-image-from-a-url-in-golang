package test

import (
	"imgdownload/controllers"
	"testing"
)

func TestFindImages1(t *testing.T) {

	input := "https://www.geeksforgeeks.org/"
	expectedOutput := "done"
	output := controllers.FindImages(input)

	if output != expectedOutput {
		t.Errorf("got %q, wanted %q", output, expectedOutput)
	}

}
