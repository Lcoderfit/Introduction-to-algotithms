#include<iostream>
#include<vector>
using namespace std;

int main()
{
    int n,m;
    cin>>n>>m;
    vector<int>need(n+1,0);
    vector<int>value(n+1,0);
    for(int i=1;i<=n;i++){
        cin>>need[i]>>value[i];
    }
    vector<int>t(m+1,0);
    vector<vector<int>>dp(n+1,t);
    for(int i=1;i<=n;i++){
        for(int j=1;j<=m;j++){
            if(j<need[i])
                dp[i][j]=dp[i-1][j];
            else
                dp[i][j]=max(dp[i-1][j],dp[i-1][j-need[i]]+value[i]);
        }
    }
    cout<<dp[n][m]<<endl;
    return 0;
}
