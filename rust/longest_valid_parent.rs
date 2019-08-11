use std::cmp;
fn longest_valid_parentheses(s: String) -> i32 {
    if s.len() < 2 {
        return 0;
    }
    let mut result = 0;
    let mut longest = vec![0; s.len()];
    let mut count = 0;

    for (i, c) in s.chars().enumerate() {
        if c == '(' {
            count += 1;
        } else if count > 0 {
            longest[i] = longest[i - 1] + 2;
            if i > longest[i] {
                longest[i] += longest[i - longest[i]]
            }
            result = cmp::max(result, longest[i]);
            count -= 1;
        }
    }
    result as i32
}

fn main() {
    let s = String::from("()");
    println!("the result is {}", longest_valid_parentheses(s))
}
