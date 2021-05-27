package prime

import "testing"

func TestSmallPrimes(t *testing.T) {
	testPrimes := []int{1, 3, 5, 7}
	for _, i := range (testPrimes) {
		prime := isPrime(int32(i))
		if !prime {
			t.Errorf("%d should be prime, but got not prime", i)
		}
	}
}

func TestSmallNotPrimes(t *testing.T) {
	testNotPrimes := []int{2,4,9,12}
	for _, i := range(testNotPrimes) {
		prime := isPrime(int32(i))
		if prime {
			t.Errorf("%d should not be prime, but got prime", i)
		}
	}
}