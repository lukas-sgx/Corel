package cmd

import "testing"

func TestCli(t *testing.T) {
	result := Cli()
	
	if result != nil {
		t.Fatal("Cli launch fail")
	}
}