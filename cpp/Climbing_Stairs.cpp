#include<iostream>

class Solution
{
public:
    int climbStairs(int n)
    {
        int stepOne = 1, stepTwo = 1;
        int ret = 1;
        for (int i = 2; i <= n; i++)
        {
            ret = stepOne + stepTwo;
            stepOne = stepTwo;
            stepTwo = ret;
        }
        return ret;
    }
};

int main()
{
    Solution solution = Solution();
    std::cout << solution.climbStairs(5)<<std::endl;
    return 0;
}
