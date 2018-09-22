#include<iostream>
#include<vector>
using namespace std;

int main()
{
    int t;
    cin>>t;
    while(t--){
        int n,m;
        cin>>n>>m;
        vector<int>p(m+1,0);
        vector<int>h(m+1,0);
        vector<int>c(m+1,0);
        for(int i=1;i<=m;i++){
            cin>>p[i]>>h[i]>>c[i];
        }
        vector<int>t(n+1,0);
        vector<vector<int>>dp(m+1,t);
        for(int i=1;i<=m;i++){
            for(int j=1;j<=n;j++){
                if(j<p[i])
                    dp[i][j]=dp[i-1][j];
                else{
                    int q=0;
                    for(int k=0;k<=j/p[i] && k<=c[i];k++)
                        q=max(q,dp[i-1][j-k*p[i]]+k*h[i]);
                    dp[i][j]=q;
                }
            }
        }
        cout<<dp[m][n]<<endl;
    }
    return 0;
}
