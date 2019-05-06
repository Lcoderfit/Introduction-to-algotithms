/**************************************************************
9.回文数：palindrome-number
leetcode链接：https://leetcode-cn.com/problems/palindrome-number/
时间复杂度：O(log10(n))
空间复杂度：O(1)
**************************************************************/

#include<iostream>
using namespace std;

class Solution {
  public:
    bool palindromeNumber(int n) {
        //首先排除特殊情况
        if(n<0 || n%10==0 && n!=0) {
            return false;
        }
        //用于存储反转一半的数字
        int revertedNumber=0;
        //如果n有偶数位，则恰好反转一半后即n==revertedNumber时退出循环

        //如果n为奇数位，则反转至n<revertedNumber时退出循环
        //此时原数的中位数为revertedNumber的个位数
        while(n>revertedNumber) {
            revertedNumber = revertedNumber*10 + n%10;
            n/=10;
        }

        return n == revertedNumber || n == revertedNumber/10;
    }
};

int main() {
    int n;
    while(cin>>n) {
        Solution s;
        bool res=s.palindromeNumber(n);
        cout<<res<<endl;
    }
    return 0;
}
