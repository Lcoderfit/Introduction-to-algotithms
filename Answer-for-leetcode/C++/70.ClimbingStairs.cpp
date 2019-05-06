/**************************************************************
70.爬楼梯:climbing-stairs
leetcode链接：https://leetcode-cn.com/problems/climbing-stairs/
**************************************************************/

#include<iostream>
//STL向量
#include<vector>
//命名空间
using namespace std;

//定义一个解决问题的类
class Solution {
public:
    //定义一个爬楼梯方法
    int climbStairs(int n) {
        //定义一个一维向量，元素个数为n+1个
        //将向量内元素全部赋值为0
        vector<int>dp(n+1,0);
        //初始化dp[0]和dp[1]
        dp[0]=1;
        dp[1]=1;
        //动态规划算法，dp[i]表示到达第i个台阶的方法数
        //到达第i个台阶的方法数=到达第i-1个台阶的方法数
        //+到达第i-2个台阶的方法数(i>=2)
        for(int i=2;i<=n;i++){
            dp[i]=dp[i-1]+dp[i-2];
        }
        //返回dp[n]:到达第n个台阶的方法数
        return dp[n];
    }
};

//主函数，运行测试用例
int main()
{
    int n;
    int res;
    //这个可以多次输入测试用例
    while(cin>>n){
        Solution s;
        //res保存爬到n阶台阶的方法数
        res=s.climbStairs(n);
        cout<<res<<endl;
    }

    return 0;
}
