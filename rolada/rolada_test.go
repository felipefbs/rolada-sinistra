package rolada_test

import (
	"slices"
	"testing"

	"github.com/felipefbs/rolada/rolada"
)

func TestNew(t *testing.T) {
	want := &rolada.Rolada{
		Number: [6]int{1, 2, 3, 4, 5, 6},
	}

	got := rolada.New(123456)
	if !slices.Equal(want.Number[:], got.Number[:]) {
		t.Errorf("rolada.New(123456) want %s, but got %s", want, got)
	}
}

func TestString(t *testing.T) {
	r := rolada.Rolada{
		Number: [6]int{1, 2, 3, 4, 5, 6},
	}

	want := "123456"
	got := r.String()

	if want != got {
		t.Errorf("rolada.String() want %s, but got %s", want, got)
	}
}

func TestHarshad(t *testing.T) {
	tt := []struct {
		Name      string
		R         rolada.Rolada
		IsHarshad bool
	}{
		{
			Name:      "906072",
			R:         *rolada.New(906072),
			IsHarshad: true,
		},
		{
			Name:      "123456",
			R:         *rolada.New(123456),
			IsHarshad: false,
		},
	}

	for _, test := range tt {
		t.Run(test.Name, func(t *testing.T) {
			if test.R.HarshadCheck() != test.IsHarshad {
				t.Errorf("r.HarshadCheck(%v) wanted %v, but got %v", test.R.Number, test.IsHarshad, test.R.HarshadCheck())
			}
		})
	}
}
