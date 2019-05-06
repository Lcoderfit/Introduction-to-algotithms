/**************************************************************
509.ì³²¨ÄÇÆõÊý£ºfibonacci-number
https://leetcode-cn.com/problems/fibonacci-number/
**************************************************************/

#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int fib(int N) {
        vector<int>dp(N+1, 0);
        dp[0]=0;
        dp[1]=1;
        for(int i=2;i<=N;i++){
          dp[i]=dp[i-1]+dp[i-2];
        }

        return dp[N];
    }
};

int main()
{
    int N;
    int res;
    while(cin>>N){
        Solution s;
        res=s.fib(N);
        cout<<res<<endl;
    }
    return 0;
}
