class Solution:
    def removeKdigits(self, num, k):
        """
        :type num: str
        :type k: int
        :rtype: str
        """
        while k > 0:
            for i, x in enumerate(num):
                if i == len(num) - 1:
                    num = num[:-1]
                    break
                if num[i] > num[i + 1]:
                    num = num[:i] + num[i + 1:]
                    break
            k -= 1

        for i, x in enumerate(num):
            if int(x) != 0:
                break
        if num[i:]:
            return num[i:]
        return "0"


def main():
    s = Solution()
    r = s.removeKdigits("1232143292", 4)
    print(r)


if __name__ == "__main__":
    main()
