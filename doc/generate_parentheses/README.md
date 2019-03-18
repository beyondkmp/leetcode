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

## 暴力枚举

使用回溯，列举出所有的可能，然后再判断是不是合法的parentheses, 不过会做非常多的无用工作.

```go
package main

import "fmt"

func isValid(s []byte) bool {
        top := 0
        for _, v := range s {
                if v == '(' {
                        top++
                }
                if v == ')' {
                        top--
                }
                if top < 0 {
                        return false
                }
        }

        if top == 0 {
                return true
        }
        return false
}

func generateParenthesis(n int) []string {
        var result []string
        var dfs func(left, right int, sol []byte)

        dfs = func(left, right int, sol []byte) {
                if left == n && right == n {
                        if isValid(sol) {
                                result = append(result, string(sol))
                        }
                }

                if left > n || right > n {
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

## 回源加剪枝

![generateParenthesis](./generate_parenthese.png)

从上图中可以得到，如果right > left都是不合法的，所以可以直接对这些不合法的进行剪枝, 其实只要对上面的代码加上`right > left`判断就可以，具体代码如下:

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

## 参考

1. [Generate All Strings With n Matched Parentheses - Backtracking ("Generate Parentheses" on LeetCode)](https://www.youtube.com/watch?v=sz1qaKt0KGQ)
