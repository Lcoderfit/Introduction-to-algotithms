#include<iostream>
#include<vector>
using namespace std;

int main()
{
    int n,m;
    cin>>n>>m;
    vector<int>w(n+1,0);
    vector<int>v(n+1,0);
    for(int i=1;i<=n;i++)
        cin>>w[i]>>v[i];
    vector<int>t(m+1,0);
    vector<vector<int>>dp(n+1,t);
    for(int i=1;i<=n;i++){
        for(int j=1;j<=m;j++){
            if(j<w[i])
                dp[i][j]=dp[i-1][j];
            else{
                int p=0;
                for(int k=0;k<=j/w[i];k++){
                    p=max(p,dp[i-1][j-k*w[i]]+k*v[i]);
                }
                dp[i][j]=p;
            }
        }
    }
    cout<<dp[n][m]<<endl;
    return 0;
}
