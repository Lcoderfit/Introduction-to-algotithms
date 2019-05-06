/**************************************************************
39.??????????????????????????????????????????
leetcode链接：
时间复杂度：
空间复杂度：
time:
**************************************************************/
#include<iostream>
#include<random>
#include<time.h>
#include<vector>

using namespace std;

typedef vector<vector<int> > IntVecDouble;

class Solution {
public:
    vector<vector<int>> combinationSum(vector<int> candidates, int target) {
        IntVecDouble ans;
        vector<int>cur;
        for(int i=0;i<candidates.size();i++){
            backTrack(ans, cur, candidates, 0, 0, target);
        }
        return ans;
    }
    void backTrack(IntVecDouble& ans, vector<int>cur,
        vector<int> candidates, int num, int i, int target){

        if(i>candidates.size() || num>target)
            return ;
        if(num == target){
            cur.push_back(candidates[i]);
            ans.push_back(cur);
            return ;
        }
        //invalid use of expression:可能原因是调用了返回值为void的函数
        cur.push_back(candidates[i]);
        backTrack(ans, cur, candidates, num+candidates[i], i, target);
        cur.push_back(candidates[i+1]);
        backTrack(ans, cur, candidates, num+candidates[i], i+1, target);
        return ;
    }
    vector<int> randForCandidates(int n){
        srand((unsigned)time(NULL));
        vector<int>candidates;
        for(int i=0;i<n;i++){
            int t=rand()%100+1;
            int j=0;
            for(;j<candidates.size();j++){
                if(t == candidates[j])
                    break;
            }
            if(candidates.size()==0 || t!=candidates[i])
                candidates.push_back(t);
        }
        return candidates;
    }
};

int main()
{
    int n;
    while(cin>>n){
        Solution s;
        vector<int> candidates=s.randForCandidates(n);
        int target=rand()%100;

        cout<<"candidates:";
        for(int i=0;i<n;i++){
            cout<<candidates[i]<<" ";
        }
        cout<<endl;
        cout<<"target:"<<target<<endl;

        IntVecDouble ans=s.combinationSum(candidates, target);
        for(int i=0;i<ans.size();i++){
            cout<<"[";
            for(int j=0;j<ans[0].size();j++){
                if(j==0){
                    cout<<ans[i][j];
                }else{
                    cout<<ans[i][j]<<" ";
                }
            }
            cout<<"]"<<endl;
        }
    }
    return 0;
}
