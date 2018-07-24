#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  void sortColors(vector<int> &nums) {
    int low, mid, high;
    low = mid = 0;
    high = nums.size() - 1;

    while (mid <= high) {
      if (nums[mid] == 0) {
        swap(nums[low++], nums[mid++]);
      } else if (nums[mid] == 1) {
        mid++;
      } else {
        swap(nums[mid], nums[high--]);
      }
    }
  }
};

int main() {
  vector<int> nums{0, 1, 2, 1, 2, 0};
  Solution s = Solution();
  s.sortColors(nums);
  for (auto i : nums)
    cout << i << " ";
  cout << endl;

  return 1;
}
