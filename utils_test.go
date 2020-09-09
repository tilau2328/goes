package goes

import "testing"

type testClass struct {}

func TestTypeOf(t *testing.T) {
	const ex = "goes.testClass"
	res := TypeOf(testClass{})
	if res != ex {
		t.Errorf("expected %s but was %s", ex, res)
	}
}
