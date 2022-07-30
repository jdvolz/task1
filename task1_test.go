package main

import "testing"

func TestEveryThird(t *testing.T) {
	o := Capitalizenth("Aspiration.com", 3)
	if o != "asPirAtiOn.cOm" {
		t.Fatalf("%v", o)
	}
}

func TestEveryFourth(t *testing.T) {
	o := Capitalizenth("Aspiration.com", 4)
	if o != "aspIratIon.cOm" {
		t.Fatalf("%v", o)
	}
}
