package task0050

// Implement pow(x, n), which calculates x raised to the power n (i.e., x^n).
//
// Constraints:
//   -100.0 < x < 100.0
//   -2^31 <= n <= 2^31-1
//   -10^4 <= x^n <= 10^4
//
// Result:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Pow(x, n).
//   Memory Usage: 2.2 MB, less than 5.88% of Go online submissions for Pow(x, n).
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n < 0 {
		return myPow(1/x, -n) // (x)^-10 => 1 / x^10
	}
	if n%2 == 0 {
		return myPow(x*x, n/2) // x^8 == (x^2)^4 == ((x^2)^2)^2
	} else {
		return x * myPow(x*x, n/2) // x^9 == x * x^8
	}
}
