package steps

import (
	"testing"
)

func TestStepsGood(t *testing.T) {
	want := `#   
##  
### 
####`
	if got := steps(4); got != want {
		t.Errorf("steps(4) = %q, want %q", got, want)
	}
}

func TestStepsRecursion(t *testing.T) {
	want := `#   
##  
### 
####`
	if got := stepsRecursion(4); got != want {
		t.Errorf("stepsRecursion(4) = %q, want %q", got, want)
	}
}
