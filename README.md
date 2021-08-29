# leetcode-go

Here I'm solving leetcode problems using Go lang. This is my first Go lang experience, so I'll be very appreciative if you send PR or a message if you have something to improve

## Performance table

Performance results of algorithms table

#task id | task name | cpu (ms) | cpu (better than) | memory (mBytes) | memory (better than)
--- | --- | --- | --- | --- | ---
[1](./tasks/task0001/two-sum.go) | [Two Sum](https://leetcode.com/problems/two-sum/) | 4ms | <span style="color:green">> 97.43%</span> | 4.2mB | <span style="color:orange">> 61.84%</span>
[2](./tasks/task0002/add-two-numbers.go) | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | 4ms | <span style="color:green">> 98.33%</span> | 4.8mB | <span style="color:green">> 92.75%</span>
[3](./tasks/task0003/longest-substring-without-repeating-characters.go) | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | 4ms | <span style="color:green">> 90.69%</span> | 3.2mB | <span style="color:#f00">> 44.02%</span>
[4](./tasks/task0004/median-of-two-sorted-arrays.go) | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | 8ms | <span style="color:green">> 97.82%</span> | 5.3mB | <span style="color:orange">> 78.69%</span>
[5](./tasks/task0005/longest-palindromic-substring.go) | [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) | 4ms | <span style="color:green">> 96.51%</span> | 2.7mB | <span style="color:green">> 100%</span>
[6](./tasks/task0006/zigzag-conversion.go) | [ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/) | 4ms | <span style="color:green">> 92.93%</span> | 4mB | <span style="color:green">> 100%</span>
[7](./tasks/task0007/reverse-integer.go) | [Reverse Integer](https://leetcode.com/problems/reverse-integer/) | 0ms | <span style="color:green">> 100%</span> | 2.2mB | <span style="color:#f00">> 26.85%</span>
[8](./tasks/task0008/string-to-integer-atoi.go) | [String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/) | 0ms | <span style="color:green">> 100%</span> | 2.3mB | <span style="color:green">> 100%</span>
[9](./tasks/task0009/palindrome-number.go) | [Palindrome Number](https://leetcode.com/problems/palindrome-number/) | 4ms | <span style="color:green">> 98.3%</span> | 5.1mB | <span style="color:yellow">> 81.56%</span>
[10](./tasks/task0010/regular-expression-matching.go) | [Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/) | 0ms | <span style="color:green">> 100%</span> | 2.4mB | <span style="color:orange">> 64.37%</span>
[11](./tasks/task0011/container-with-most-water.go) | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | 80ms | <span style="color:yellow">> 83.62%</span> | 8.8mB | <span style="color:#f00">> 14.86%</span>
[12](./tasks/task0012/integer-to-roman.go) | [Integer to Roman](https://leetcode.com/problems/integer-to-roman/) | 4ms | <span style="color:green">> 92.66%</span> | 3.3mB | <span style="color:yellow">> 86.31%</span>
[13](./tasks/task0013/roman-to-integer.go) | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) | 0ms | <span style="color:green">> 100%</span> | 3mB | <span style="color:#f00">> 40.5%</span>
[14](./tasks/task0014/longest-common-prefix.go) | [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) | 0ms | <span style="color:green">> 100%</span> | 2.4mB | <span style="color:green">> 100%</span>
[15](./tasks/task0015/3sum.go) | [3Sum](https://leetcode.com/problems/3sum/) | 28ms | <span style="color:green">> 92.04%</span> | 7.7mB | <span style="color:orange">> 52.96%</span>
[16](./tasks/task0016/3sum-closest.go) | [3Sum Closest](https://leetcode.com/problems/3sum-closest/) | 4ms | <span style="color:green">> 96.37%</span> | 2.8mB | <span style="color:green">> 100%</span>
[17](./tasks/task0017/letter-combinations-of-a-phone-number.go) | [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | 0ms | <span style="color:green">> 100%</span> | 2.2mB | <span style="color:#f00">> 23.93%</span>
[18](./tasks/task0018/4sum.go) | [4Sum](https://leetcode.com/problems/4sum/) | 4ms | <span style="color:green">> 96.75%</span> | 2.8mB | <span style="color:orange">> 68.18%</span>
[19](./tasks/task0019/remove-nth-node-from-end-of-list.go) | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | 0ms | <span style="color:green">> 100%</span> | 2.5mB | <span style="color:#f00">> 6.36%</span>
[20](./tasks/task0020/valid-parentheses.go) | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | 0ms | <span style="color:green">> 100%</span> | 2mB | <span style="color:orange">> 77.59%</span>
[21](./tasks/task0021/merge-two-sorted-lists.go) | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | 0ms | <span style="color:green">> 100%</span> | 2.6mB | <span style="color:green">> 100%</span>
[22](./tasks/task0022/generate-parentheses.go) | [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/) | 0ms | <span style="color:green">> 100%</span> | 3.7mB | <span style="color:#f00">> 16.51%</span>
[23](./tasks/task0023/merge-k-sorted-lists.go) | [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) | 32ms | <span style="color:#f00">> 36.97%</span> | 5.8mB | <span style="color:#f00">> 40.72%</span>
[24](./tasks/task0024/swap-nodes-in-pairs.go) | [Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/) | 0ms | <span style="color:green">> 100%</span> | 2.1mB | <span style="color:#f00">> 16.14%</span>
[25](./tasks/task0025/reverse-nodes-in-k-group.go) | [Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/) | 0ms | <span style="color:green">> 100%</span> | 3.7mB | <span style="color:green">> 100%</span>
[26](./tasks/task0026/remove-duplicates-from-sorted-array.go) | [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | 4ms | <span style="color:green">> 99.44%</span> | 4.4mB | <span style="color:#f00">> 34.64%</span>
[27](./tasks/task0027/remove-element.go) | [Remove Element](https://leetcode.com/problems/remove-element/) | 0ms | <span style="color:green">> 100%</span> | 2.2mB | <span style="color:green">> 99.76%</span>
[28](./tasks/task0028/implement-strstr.go) | [Implement strStr()](https://leetcode.com/problems/implement-strstr/) | 4ms | <span style="color:orange">> 66.29%</span> | 2.5mB | <span style="color:#f00">> 49.09%</span>
[30](./tasks/task0030/substring-with-concatenation-of-all-words.go) | [Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words/) | 4ms | <span style="color:green">> 100%</span> | 4mB | <span style="color:yellow">> 84.38%</span>
[31](./tasks/task0031/next-permutation.go) | [Next Permutation](https://leetcode.com/problems/next-permutation/) | 0ms | <span style="color:green">> 100%</span> | 2.5mB | <span style="color:#f00">> 12.32%</span>
[32](./tasks/task0032/longest-valid-parentheses.go) | [Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/) | 0ms | <span style="color:green">> 100%</span> | 3.7mB | <span style="color:#f00">> 22.75%</span>
[33](./tasks/task0033/search-in-rotated-sorted-array.go) | [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) | 0ms | <span style="color:green">> 100%</span> | 2.6mB | <span style="color:#f00">> 14.6%</span>
[34](./tasks/task0034/find-first-and-last-position-of-element-in-sorted-array.go) | [Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | 4ms | <span style="color:green">> 98.77%</span> | 4mB | <span style="color:#f00">> 19.7%</span>
[35](./tasks/task0035/search-insert-position.go) | [Search Insert Position](https://leetcode.com/problems/search-insert-position/) | 0ms | <span style="color:green">> 100%</span> | 3mB | <span style="color:#f00">> 13.94%</span>
[36](./tasks/task0036/valid-sudoku.go) | [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/) | 0ms | <span style="color:green">> 100%</span> | 3.2mB | <span style="color:#f00">> 45.77%</span>
[37](./tasks/task0037/sudoku-solver.go) | [Sudoku Solver](https://leetcode.com/problems/sudoku-solver/) | 0ms | <span style="color:green">> 100%</span> | 2.4mB | <span style="color:#f00">> 34.65%</span>
[38](./tasks/task0038/count-and-say.go) | [Count and Say](https://leetcode.com/problems/count-and-say/) | 0ms | <span style="color:green">> 100%</span> | 2.5mB | <span style="color:yellow">> 86.55%</span>
[39](./tasks/task0039/combination-sum.go) | [Combination Sum](https://leetcode.com/problems/combination-sum/) | 0ms | <span style="color:green">> 100%</span> | 3mB | <span style="color:yellow">> 84.8%</span>
[40](./tasks/task0040/combination-sum-ii.go) | [Combination Sum II](https://leetcode.com/problems/combination-sum-ii/) | 0ms | <span style="color:green">> 100%</span> | 2.7mB | <span style="color:yellow">> 80%</span>
[41](./tasks/task0041/first-missing-positive.go) | [First Missing Positive](https://leetcode.com/problems/first-missing-positive/) | 116ms | <span style="color:yellow">> 84.7%</span> | 25.2mB | <span style="color:orange">> 50.18%</span>
[42](./tasks/task0042/trapping-rain-water.go) | [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) | 4ms | <span style="color:yellow">> 84.15%</span> | 3.6mB | <span style="color:#f00">> 14.52%</span>
[43](./tasks/task0043/multiply-strings.go) | [Multiply Strings](https://leetcode.com/problems/multiply-strings/) | 0ms | <span style="color:green">> 100%</span> | 5.9mB | <span style="color:#f00">> 31.49%</span>
[45](./tasks/task0045/jump-game-ii.go) | [Jump Game II](https://leetcode.com/problems/jump-game-ii/) | 8ms | <span style="color:green">> 99.42%</span> | 6.1mB | <span style="color:#f00">> 49.85%</span>
[46](./tasks/task0046/permutations.go) | [Permutations](https://leetcode.com/problems/permutations/) | 0ms | <span style="color:green">> 100%</span> | 3.5mB | <span style="color:#f00">> 6.58%</span>
[47](./tasks/task0047/permutations-ii.go) | [Permutations II](https://leetcode.com/problems/permutations-ii/) | 4ms | <span style="color:yellow">> 89.41%</span> [*](#performanceFootnote) | 6.4mB | <span style="color:#f00">> 19.41%</span>
[55](./tasks/task0055/jump-game.go) | [Jump Game](https://leetcode.com/problems/jump-game/) | 56ms | <span style="color:green">> 92.67%</span> [*](#performanceFootnote) | 7.1mB | <span style="color:yellow">> 85.71%</span>
[60](./tasks/task0060/permutation-sequence.go) | [Permutation Sequence](https://leetcode.com/problems/permutation-sequence/) | 0ms | <span style="color:green">> 100%</span> | 2.2mB | <span style="color:#f00">> 27.27%</span>
[61](./tasks/task0061/rotate-list.go) | [Rotate List](https://leetcode.com/problems/rotate-list/) | 0ms | <span style="color:green">> 100%</span> | 2.6mB | <span style="color:#f00">> 29%</span>
[82](./tasks/task0082/remove-duplicates-from-sorted-list-ii.go) | [Remove Duplicates from Sorted List II](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/) | 0ms | <span style="color:green">> 100%</span> | 3.1mB | <span style="color:green">> 100%</span>
[83](./tasks/task0083/remove-duplicates-from-sorted-list.go) | [Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/) | 4ms | <span style="color:yellow">> 80.59%</span> | 3.2mB | <span style="color:green">> 100%</span>
[86](./tasks/task0086/partition-list.go) | [Partition List](https://leetcode.com/problems/partition-list/) | 0ms | <span style="color:green">> 100%</span> | 2.5mB | <span style="color:green">> 100%</span>
[92](./tasks/task0092/reverse-linked-list-ii.go) | [Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/) | 0ms | <span style="color:green">> 100%</span> | 2.1mB | <span style="color:#f00">> 25.06%</span>
[94](./tasks/task0094/binary-tree-inorder-traversal.go) | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) | 0ms | <span style="color:green">> 100%</span> | 2.1mB | <span style="color:green">> 99.63%</span>
[100](./tasks/task0100/same-tree.go) | [Same Tree](https://leetcode.com/problems/same-tree/) | 0ms | <span style="color:green">> 100%</span> | 2.1mB | <span style="color:#f00">> 20.08%</span>
[101](./tasks/task0101/symmetric-tree.go) | [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) | 0ms | <span style="color:green">> 100%</span> | 2.9mB | <span style="color:#f00">> 38.33%</span>
[108](./tasks/task0108/convert-sorted-array-to-binary-search-tree.go) | [Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/) | 0ms | <span style="color:green">> 100%</span> | 4.1mB | <span style="color:#f00">> 28.33%</span>
[109](./tasks/task0109/convert-sorted-list-to-binary-search-tree.go) | [Convert Sorted List to Binary Search Tree](https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/) | 4ms | <span style="color:green">> 99.07%</span> | 5.9mB | <span style="color:green">> 100%</span>
[141](./tasks/task0141/linked-list-cycle.go) | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) | 8ms | <span style="color:yellow">> 88.4%</span> | 4.4mB | <span style="color:green">> 100%</span>
[142](./tasks/task0142/linked-list-cycle-ii.go) | [Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/) | 4ms | <span style="color:green">> 97.89%</span> | 3.8mB | <span style="color:green">> 100%</span>
[189](./tasks/task0189/rotate-array.go) | [Rotate Array](https://leetcode.com/problems/rotate-array/) | 28ms | <span style="color:green">> 96.3%</span> | 8.1mB | <span style="color:orange">> 62.59%</span>
[203](./tasks/task0203/remove-linked-list-elements.go) | [Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/) | 8ms | <span style="color:yellow">> 86.15%</span> | 4.7mB | <span style="color:green">> 100%</span>
[206](./tasks/task0206/reverse-linked-list.go) | [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) | 0ms | <span style="color:green">> 100%</span> | 2.6mB | <span style="color:green">> 100%</span>
[226](./tasks/task0226/invert-binary-tree.go) | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) | 0ms | <span style="color:green">> 100%</span> | 2.1mB | <span style="color:orange">> 50.25%</span>
[234](./tasks/task0234/palindrome-linked-list.go) | [Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/) | 148ms | <span style="color:yellow">> 87.03%</span> | 11.4mB | <span style="color:#f00">> 19.92%</span>
[278](./tasks/task0278/first-bad-version.go) | [First Bad Version](https://leetcode.com/problems/first-bad-version/) | 0ms | <span style="color:green">> 100%</span> | 1.9mB | <span style="color:#f00">> 30.5%</span>
[331](./tasks/task0331/verify-preorder-serialization-of-a-binary-tree.go) | [Verify Preorder Serialization of a Binary Tree](https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/) | 0ms | <span style="color:green">> 100%</span> | 2.8mB | <span style="color:yellow">> 84%</span>
[522](./tasks/task0522/longest-uncommon-subsequence-ii.go) | [Longest Uncommon Subsequence II](https://leetcode.com/problems/longest-uncommon-subsequence-ii/) | 0ms | <span style="color:green">> 100%</span> | 2.3mB | <span style="color:#f00">> 14.29%</span>
[633](./tasks/task0633/sum-of-square-numbers.go) | [Sum of Square Numbers](https://leetcode.com/problems/sum-of-square-numbers/) | 0ms | <span style="color:green">> 100%</span> | 2mB | <span style="color:orange">> 57.14%</span>
[876](./tasks/task0876/middle-of-the-linked-list.go) | [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) | 0ms | <span style="color:green">> 100%</span> | 2mB | <span style="color:#f00">> 11.59%</span>
[988](./tasks/task0988/smallest-string-starting-from-leaf.go) | [Smallest String Starting From Leaf](https://leetcode.com/problems/smallest-string-starting-from-leaf/) | 4ms | <span style="color:green">> 96.15%</span> | 6.3mB | <span style="color:#f00">> 7.69%</span>
[996](./tasks/task0996/number-of-squareful-arrays.go) | [Number of Squareful Arrays](https://leetcode.com/problems/number-of-squareful-arrays/) | 0ms | <span style="color:green">> 100%</span> | 2.1mB | <span style="color:orange">> 71.43%</span>
[1448](./tasks/task1448/count-good-nodes-in-binary-tree.go) | [Count Good Nodes in Binary Tree](https://leetcode.com/problems/count-good-nodes-in-binary-tree/) | 88ms | <span style="color:orange">> 66.04%</span> | 10.9mB | <span style="color:#f00">> 32.08%</span>
[1734](./tasks/task1734/decode-xored-permutation.go) | [Decode XORed Permutation](https://leetcode.com/problems/decode-xored-permutation/) | 156ms | <span style="color:green">> 100%</span> | 10mB | <span style="color:#f00">> 37.5%</span>

---

<p id="performanceFootnote">
    * Result is unstable (shows different result for each run).
</p>
