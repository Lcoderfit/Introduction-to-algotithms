/*
输入两个字符串，判断字符串1转换成字符串2所需要的最少操作数
允许的操作如下：
1.每删除一个字符算一次操作
2.每添加一个字符算一次操作
3.每替换一个字符算一次操作
*/

#include<iostream>
#include<algorithm>
#include<vector>

using namespace std;

int main()
{
    string s1, s2;
    while(cin>>s1>>s2){
        int len1 = s1.size();
        int len2 = s2.size();
        int res;
        if(len1==0){
            cout<<len2<<endl;
            continue;
        }
        if(len2==0){
            cout<<len1<<endl;
            continue;
        }
        vector<int> t(len2+1, 0);
        vector<vector<int> >dp(len1+1, t);
        dp[0][1] = 1;
        dp[1][0] = 1;
        for(int i=1;i<=len1;i++){
            for(int j=1;j<=len2;j++){
                if(s1[i-1]==s2[j-1]){
                    dp[i][j] = dp[i-1][j-1];
                }else{
                    dp[i][j]=min(dp[i-1][j]+1, dp[i][j-1]+1);
                    dp[i][j]=min(dp[i][j], dp[i-1][j-1]+1);
                }
            }
        }
        cout<<dp[len1][len2]<<endl;
    }
}
