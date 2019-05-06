#include<iostream>
#include<vector>
#include<string>
using namespace std;

class Lcs{
public:
    vector<vector<int>>length;
    vector<vector<string>>str;
};

Lcs LcsLength(string a,string b);
int PrintLcs(vector<vector<string>>b,string x,int i,int j);

int main()
{
    string x="ABCBDAB";
    string y="BDCABA";
    Lcs lcs;
    lcs=LcsLength(x,y);
    int m=x.size();
    int n=y.size();
    cout<<lcs.length[m][n]<<endl;
    PrintLcs(lcs.str,x,m,n);
    return 0;
}

Lcs LcsLength(string x, string y){
    int m=x.size();
    int n=y.size();
    vector<int>t(n+1, 0);
    vector<vector<int>>dp(m+1,t);
    vector<string>s(n+1,"");
    vector<vector<string>>b(m+1,s);

    for(int i=1;i<=m;i++){
        for(int j=1;j<=n;j++){
            if(x[i-1]==y[j-1]){
                dp[i][j]=dp[i-1][j-1]+1;
                b[i][j]="↖";
            }
            else if(dp[i-1][j]>=dp[i][j-1]){
                dp[i][j]=dp[i-1][j];
                b[i][j]="↑";
            }
            else{
                dp[i][j]=dp[i][j-1];
                b[i][j]="←";
            }
        }
    }
    Lcs lcs;
    lcs.length=dp;
    lcs.str=b;
    return lcs;
}

int PrintLcs(vector<vector<string>>b,string x,int i,int j)
{
    if(i==0 || j==0)
        return 0;
    if(b[i][j]=="↖"){
        //下面两个语句顺序别搞错
        //如果把这两句的顺序颠倒，会导致逆序输出LCS
        PrintLcs(b,x,i-1,j-1);
        cout<<x[i-1];
    }
    else if(b[i][j]=="↑"){
        PrintLcs(b,x,i-1,j);
    }
    else{
        PrintLcs(b,x,i,j-1);
    }
    return 0;
}

