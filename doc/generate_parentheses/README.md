# generate parentheses

## Detail

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

```
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```

## Solution 1

```go
package main

import "fmt"

func generateParenthesis(n int) []string {
        var result []string
        var dfs func(left, right int, sol []byte)

        dfs = func(left, right int, sol []byte) {
                if left == n && right == n {
                        result = append(result, string(sol))
                }

                if left > n || right > n || right > left {
                        return
                }

                dfs(left+1, right, append(sol, '('))
                dfs(left, right+1, append(sol, ')'))

        }
        dfs(0, 0, nil)
        return result
}

func main() {
        fmt.Println(generateParenthesis(3))
}
```

