#include<iostream>
#include<vector>
#include<utility>

using namespace std;

typedef vector<vector<int>> TypDouVecint;
typedef vector<vector<string>> TypDouVecstr;
typedef pair<int,TypDouVecstr> TypPairIntDouVecstr;

class Solution{
public:
    TypPairIntDouVecstr lpsLength(string s){
        int n=s.size();
        vector<int> t(n+1,0);
        TypDouVecint dp(n+1,t);
        vector<string> c(n+1,"");
        TypDouVecstr ds(n+1,c);
        for(int i=1;i<=n;i++){
            dp[i][i]=1;
            ds[i][i]="。";
        }
        for(int l=2;l<=n;l++){
            for(int i=1;i<=n-l+1;i++){
                int j=i+l-1;
                if(s[i-1]==s[j-1]){
                    dp[i][j]=dp[i+1][j-1]+2;
                    ds[i][j]="↙";
                }else if(dp[i+1][j]>=dp[i][j-1]){
                    dp[i][j]=dp[i+1][j];
                    ds[i][j]="↓";
                }else{
                    dp[i][j]=dp[i][j-1];
                    ds[i][j]="←";
                }
            }
        }
        TypPairIntDouVecstr ivs;
        ivs.first=dp[1][n];
        ivs.second=ds;
        return ivs;
    }
    void printLps(TypDouVecstr dvs, string s, int i, int j){
        if(i>j)
            return ;
        if(dvs[i][j]=="↙"){
            cout<<s[i-1];
            printLps(dvs, s, i+1,j-1);
            cout<<s[j-1];
        }else if(dvs[i][j]=="↓"){
            printLps(dvs, s, i+1, j);
        }else if(dvs[i][j]=="←"){
            printLps(dvs, s, i, j-1);
        }else{
            cout<<s[i-1];
        }
    }
};

int main()
{
    Solution solut;
    string s;
    while(cin>>s){
        TypPairIntDouVecstr ivs=solut.lpsLength(s);
        cout<<ivs.first<<endl;
        solut.printLps(ivs.second, s, 1, s.size());
        cout<<endl;
    }
    return 0;
}
