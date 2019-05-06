#include<iostream>
#include<vector>//容器
using namespace std;

int BottomUpCutRod(int p[], int n);//计算长度为n的钢条的最大收益

int main()
{
    int p[10]={1,5,8,9,10,17,17,20,24,30};//定义一个数组存储对应钢条长度的售价
    BottomUpCutRod(p,10);//
    return 0;
}

//自底向上动态规划
int BottomUpCutRod(int p[], int n)
{
    vector<int>dp(n+1,0);
    int k;
    //while(cin>>k)可以对多个输入进行运算
    //可同时输入1～10检验程序计算结果与正确值是否相等
    while(cin>>k){
        for(int i=1;i<=n;i++){
            for(int j=1;j<=i;j++){
                dp[i]=max(dp[i],p[j-1]+dp[i-j]);
            }
        }
        cout<<dp[n]<<" ";
    }

}

