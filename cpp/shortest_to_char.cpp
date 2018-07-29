#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
  vector<int> shortestToChar(string S, char C) {
    int n = S.size();
    vector<int> val(n, n);
    int pos = -n;

    for (int i = 0; i < n; i++) {
      if (S[i] == C) {
        pos = i;
      }
      val[i] = min(val[i], i - pos);
    }

    for (int i = n - 1; i >= 0; i--) {
      if (S[i] == C) {
        pos = i;
      }
      val[i] = min(val[i], abs(i - pos));
    }

    return val;
  }
};

int main() {
  string str = "loveleetcode";
  char c = 'e';
  Solution s = Solution();
  for (auto i : s.shortestToChar(str, c)) {
    cout << i << " ";
  }
  cout << endl;

  return 1;
}
