package goes

import "testing"

type testClass struct{}

func TestTypeOf(t *testing.T) {
	const ex = "goes.testClass"
	res := TypeOf(testClass{})
	if res != ex {
		t.Errorf("expected %s but was %s", ex, res)
	}
}

func TestMessageType(t *testing.T) {
	const ex = "goes.testClass"
	res := MessageType(&testClass{})
	if res != ex {
		t.Errorf("expected %s but was %s", ex, res)
	}
}

func TestRegex(t *testing.T) {
	expected := []string{"a7Z.d9F", "a2F.2Fs"}
	res := Regex("(\\w+)\\.([0-9a-zA-Z]+)", expected[0]+" 2eW.?d2 "+expected[1])
	if len(res) != 2 {
		t.Errorf("expected 2 results but got %d", len(res))
	}
	for i, item := range res {
		if expected[i] != item {
			t.Errorf("expected the value #%d to be %s but was %s", i, expected[i], item)
		}
	}
}
