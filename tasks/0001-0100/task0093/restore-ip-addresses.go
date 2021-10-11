package task0093

import "strconv"

// Given a string s containing only digits, return all possible valid IP addresses that can be obtained from s.
// You can return them in any order.
// A valid IP address consists of exactly four integers, each integer is between 0 and 255, separated by single
// dots and cannot have leading zeros. For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses and
// "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
//
// Constraints:
//   0 <= s.length <= 3000
//   s consists of digits only.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Restore IP Addresses.
//   Memory Usage: 2.1 MB, less than 77.94% of Go online submissions for Restore IP Addresses.
func restoreIpAddresses(s string) []string {
	res, isOk := restoreIpAddressesRecursive(s, 0)
	if isOk {
		return res
	}
	return []string{}
}

func restoreIpAddressesRecursive(s string, depth int) ([]string, bool) {
	n := len(s)
	if n < (4-depth) || n > (4-depth)*3 {
		// check the size of current part of ip address:
		// e.g. if depth=2 => string cannot contain less than 2 digits and more than 6
		return nil, false
	}
	if depth == 3 {
		// We're on the last step:
		//   1. Try to convert into integer value
		//   2. Reject values: 00,000 and >255
		val, err := strconv.Atoi(s)
		if s[0] == '0' {
			if val > 0 || n > 1 {
				return nil, false
			}
		}
		if err == nil && val < 256 {
			return []string{s}, true
		} else {
			return nil, false
		}
	}

	if s[0] == '0' {
		// if current digit starts with '0' - we can cut off only '0'
		if n == 1 {
			return nil, false
		}

		// Extract sub levels
		res, isOk := restoreIpAddressesRecursive(s[1:], depth+1)
		if !isOk {
			return nil, false
		}
		// Append '0' to the list of results and return
		for j := range res {
			res[j] = "0." + res[j]
		}
		return res, true
	}

	// extract current number:
	var result []string
	isAtLeastOk := false
	for i := 0; i < n && i < 3; i++ {
		current := s[:i+1]
		val, err := strconv.Atoi(current)
		if err != nil || val > 255 {
			break
		}

		// Extract sub levels
		res, isOk := restoreIpAddressesRecursive(s[i+1:], depth+1)
		if !isOk {
			continue
		}
		isAtLeastOk = true
		// Append current number to the list of results
		for j := range res {
			res[j] = current + "." + res[j]
			result = append(result, res[j])
		}
	}
	if !isAtLeastOk {
		return nil, false
	}
	return result, true
}
