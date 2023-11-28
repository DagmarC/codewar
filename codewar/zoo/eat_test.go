package zoo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWhoEatsWho(t *testing.T) {
	tc := []struct {
		input    string
		exoected []string
	}{
		{
			input: "fox,bug,chicken,grass,sheep",
			exoected: []string{"fox,bug,chicken,grass,sheep",
				"chicken eats bug",
				"fox eats chicken",
				"sheep eats grass",
				"fox eats sheep",
				"fox"},
		},
		{
			input: "bear,bear,sheep,sheep,cow,grass,antelope,lion",
			exoected: []string{"bear,bear,sheep,sheep,cow,grass,antelope,lion",
				"bear eats sheep",
				"bear eats sheep",
				"bear eats cow",
				"antelope eats grass",
				"lion eats antelope",
				"bear,bear,lion"},
		},
		{
			input: "cow,bear,leaves,chicken,busker,big-fish,antelope,leaves",
			exoected: []string{"cow,bear,leaves,chicken,busker,big-fish,antelope,leaves",
				"bear eats cow",
				"bear eats leaves",
				"bear eats chicken",
				"bear,busker,big-fish,antelope,leaves"},
		},
	}
	for _, test := range tc {
		if got := WhoEatsWho(test.input); !cmp.Equal(got, test.exoected) {
			t.Errorf("got %v, expected %v", got, test.exoected)
		}
	}

}
