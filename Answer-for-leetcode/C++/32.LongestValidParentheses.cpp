#include<iostream>
#include<vector>
#include<stack>

using namespace std;

class Solution {
public:
    int longestValidParentheses(string s) {
        int slen = s.size();
        stack<int> ts;
        vector<int>res(slen, 0);
        for(int i=0;i<slen;i++){
            if(s[i]=='('){
                ts.push(i);
            }else{
                if(!ts.empty()){
                    res[i] = 1;
                    res[ts.top()] = 1;
                    ts.pop();
                }
            }
        }
        int max=0;
        int count=0;
        for(int i=0;i<slen;i++){
            if(res[i]==1){
                count++;
            }else{
                if(max<count){
                    max = count;
                }
                count = 0;
            }
        }
        if(max < count){
            max = count;
        }
        return max;
    }
};

int main()
{
    string s;
    while(cin>>s){
        Solution ans;
        int res =  ans.longestValidParentheses(s);
        cout<<res<<endl;
    }
    return 0;
}
