#  Given an encoded string, return it's decoded string.

#  The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

#  You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

#  Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].

#  Examples:

#  s = "3[a]2[bc]", return "aaabcbc".
#  s = "3[a2[c]]", return "accaccacc".
#  s = "2[abc]3[cd]ef", return "abcabccdcdcdef".


class Solution(object):
    def decodeString(self, s):
        """
        :type s: str
        :rtype: str
        """
        nums = ''
        strs = ''
        stack_nums = []
        stack_strs = []
        result = ''
        for w in s:
            if '0' <= w <= '9':
                nums += w
            elif w == '[':
                stack_nums.append(int(nums))
                nums = ''
                if strs:
                    stack_strs.append(strs)
                    strs = ''
                stack_strs.append('[')
            elif 'a' <= w <= 'z':
                strs += w
            elif w == ']':
                if strs:
                    stack_strs.append(strs)
                    strs = ''
                n = stack_nums.pop()
                s = stack_strs.pop()
                result = ''
                while s != '[':
                    result = s + result
                    s = stack_strs.pop()
                stack_strs.append(n * result)
        return ''.join(stack_strs) + strs

class Solution2(object):
    def decodeString(self, s):
        """
        :type s: str
        :rtype: str
        """
        nums = ''
        stack = []
        stack.append(['',1])
        for ch in s:
            if ch.isdigit():
                nums += ch
            elif ch == '[':
                stack.append(['',int(nums)])
                nums = ''
            elif ch == ']':
                strs, k = stack.pop()
                stack[-1][0] += strs * k
            else:
                stack[-1][0] += ch
        return stack[0][0]
