package rolada

import (
	"strconv"
	"strings"
	"sync"
)

type Rolada struct {
	Number    [6]int
	sum       *int
	sumMutex  *sync.Mutex
	str       string
	strMuxtex *sync.Mutex
}

func New(n int) *Rolada {
	r := [6]int{}
	c := len(r) - 1
	for n > 0 {
		r[c] = n % 10
		n = n / 10
		c--
	}

	return &Rolada{
		Number:    r,
		sumMutex:  &sync.Mutex{},
		strMuxtex: &sync.Mutex{},
	}
}

func (r *Rolada) Sum() int {
	if r.sum == nil {
		r.sumMutex.Lock()
		defer r.sumMutex.Unlock()
		if r.sum == nil {
			sum := 0
			for _, i := range r.Number {
				sum += i
			}
			r.sum = &sum
			return sum
		}
	}

	return *r.sum
}

func (r *Rolada) String() string {
	builder := strings.Builder{}

	for _, v := range r.Number {
		if v == 0 {
			continue
		}
		builder.WriteString(strconv.Itoa(v))
	}

	return builder.String()
}

func (r *Rolada) Int() int {
	if r.IsZero() {
		return 0
	}
	i, _ := strconv.Atoi(r.String())

	return i
}

func (r *Rolada) IsZero() bool {
	if r.Sum() == 0 {
		return true
	}

	return false
}

func (r *Rolada) HarshadCheck() bool {
	if r.IsZero() {
		return false
	}

	return r.Int()%r.Sum() == 0
}
