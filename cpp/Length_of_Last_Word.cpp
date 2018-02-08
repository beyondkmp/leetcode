class Solution {
public:
    int lengthOfLastWord(string s) {
        int len=0;
        int flag=0;
        for (int i = s.length() - 1; i>=0; i--){
            if(s[i] != ' '){
                flag = 1;
                len++;
            }else if(flag == 1){
                break;
            }
        }
        return len;
    }
};
