package rolada_test

import (
	"slices"
	"testing"

	"github.com/felipefbs/rolada/rolada"
)

func TestNew(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  [6]int
	}{
		{
			name:  "default",
			input: 123456,
			want:  [6]int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "false",
			input: 42,
			want:  [6]int{0, 0, 0, 0, 4, 2},
		},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := rolada.New(test.input)
			if !slices.Equal(test.want[:], got.Number[:]) {
				t.Errorf("rolada.New(%v) want %v, but got %v", test.input, test.want, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	tt := []struct {
		name  string
		input rolada.Rolada
		want  string
	}{
		{
			name:  "default",
			input: rolada.Rolada{Number: [6]int{1, 2, 3, 4, 5, 6}},
			want:  "123456",
		},
		{
			name:  "starts with 0",
			input: rolada.Rolada{Number: [6]int{0, 2, 3, 4, 5, 6}},
			want:  "23456",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.String()
			if test.want != got {
				t.Errorf("rolada.String() want %s, but got %s", test.want, got)
			}
		})
	}
}

func TestIsZero(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "true",
			input: 0,
			want:  true,
		},
		{
			name:  "false",
			input: 123456,
			want:  false,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			r := rolada.New(test.input)
			got := r.IsZero()
			if test.want != got {
				t.Errorf("rolada.String() want %v, but got %v", test.want, got)
			}
		})
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
		{
			Name:      "000000",
			R:         *rolada.New(000000),
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
