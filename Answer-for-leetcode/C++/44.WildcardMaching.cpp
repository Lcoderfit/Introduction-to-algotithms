#include<iostream>
#include<vector>

using namespace std;

class Solution {
public:
    bool isMatch(string s, string p) {
        int len1 = s.size();
        int len2 = p.size();
        vector<int> t(len2+1, 0);
        vector<vector<int> > dp(len1+1, t);
        dp[0][0]=true;
        for(int i=0;i<len2;i++){
            if(dp[0][i] && p[i]=='*')
                dp[0][i+1] = true;
        }

        for(int i=1;i<=len1;i++){
            for(int j=1;j<=len2;j++){
                if(p[j-1]=='?' || s[i-1] == p[j-1]){
                    dp[i][j] = dp[i-1][j-1];
                }else if(p[j-1]=='*'){
                    dp[i][j] = dp[i-1][j] || dp[i][j-1];
                }
            }
        }
        return dp[len1][len2];
    }
};

int main()
{
    string s, p;
    while(cin>>s>>p){
        Solution ans;
        bool res = ans.isMatch(s, p);
        if(res){
            cout<<"true"<<endl;
        }else{
            cout<<"false"<<endl;
        }
    }
    return 0;
}
