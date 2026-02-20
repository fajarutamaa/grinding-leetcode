package sumprimes

func SumPrimes(limit int) int {
	if limit < 2 {
		return 0
	}
	result := 0

	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			result += i
		}
	}
	return result
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
