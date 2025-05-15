package main

func Gcd(a, b uint) uint {
	for a == 0 || b == 0 {
		return 0
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
