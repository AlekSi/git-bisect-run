package main

type Z struct {
	buf [256]byte
}

var saved interface{}

func f(x interface{}) {
	z := x.(Z)
	g(z)
}

func g(z Z) {
	allocstuff()
	h(z, z, z)
}

func h(_ Z, _ Z, x interface{}) {
	saved = x
	return

	f(x)
}

func allocstuff() {
	m := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		m[i] = i
	}
}

func main() {
	var z Z
	f(z)
	allocstuff()
}
