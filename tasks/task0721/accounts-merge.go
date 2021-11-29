package task0721

import "sort"

// Given a list of accounts where each element accounts[i] is a list of strings,
// where the first element accounts[i][0] is a name, and the rest of the elements
// are emails representing emails of the account.
// Now, we would like to merge these accounts. Two accounts definitely belong to
// the same person if there is some common email to both accounts. Note that even
// if two accounts have the same name, they may belong to different people as
// people could have the same name. A person can have any number of accounts
// initially, but all of their accounts definitely have the same name.
// After merging the accounts, return the accounts in the following format:
// the first element of each account is the name, and the rest of the elements
// are emails in sorted order. The accounts themselves can be returned in any order.
//
// Constraints:
//   1 <= accounts.length <= 1000
//   2 <= accounts[i].length <= 10
//   1 <= accounts[i][j] <= 30
//   accounts[i][0] consists of English letters.
//   accounts[i][j] (for j > 0) is a valid email.
//
// Results:
//   Runtime: 40 ms, faster than 99.22% of Go online submissions for Accounts Merge.
//   Memory Usage: 8 MB, less than 70.54% of Go online submissions for Accounts Merge.
func accountsMerge(accounts [][]string) [][]string {
	// Set that contains the links to the parent accounts
	// the account parentSet[i] == i is supposed to be root account (has no parent)
	var parentSet = make([]int, len(accounts))

	// Pointer for emails to the parent index
	var emailToIdx = make(map[string]*int)

	// Build the parentSet
	for i, account := range accounts {
		parentSet[i] = i
		for j := 1; j < len(account); j++ {
			// Check current email has already been in some group
			ptrSet := emailToIdx[account[j]]
			if ptrSet == nil {
				emailToIdx[account[j]] = &parentSet[i]
			} else {
				// Get the root and set up the root link
				root_i := *ptrSet
				for parentSet[root_i] != root_i {
					root_i = parentSet[root_i]
				}
				if parentSet[i] == i {
					parentSet[i] = root_i
				} else {
					parentSet[root_i] = parentSet[i]
				}
			}
		}
	}

	result := make([][]string, len(accounts))
	for email, ptrSet := range emailToIdx {
		root_i := *ptrSet
		for root_i != parentSet[root_i] {
			root_i = parentSet[root_i]
		}
		if len(result[root_i]) == 0 {
			// We don't have result for root_i: create it
			result[root_i] = append(result[root_i], accounts[root_i][0])
		}
		// append the email to the result
		result[root_i] = append(result[root_i], email)
	}

	// Extract the result
	upperIdx := len(result)
	for i := range result {
		if len(result[i]) == 0 {
			if len(accounts[i]) == 1 {
				result[i] = append(result[i], accounts[i][0])
			} else {
				if upperIdx == len(result) {
					upperIdx = i
				}
				continue
			}
		} else {
			// Sort result item
			tmp := result[i][1:]
			sort.Slice(result[i][1:], func(i, j int) bool { return tmp[i] < tmp[j] })
		}
		if upperIdx != len(result) {
			result[upperIdx] = result[i]
			upperIdx++
		}
	}
	return result[:upperIdx]
}
