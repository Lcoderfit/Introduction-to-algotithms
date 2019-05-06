/**************************************************************
46.??????????????????????????????????????????
leetcode链接：
时间复杂度：
空间复杂度：
time:20:20
**************************************************************/
#include<iostream>
#include<time.h>
#include<vector>

using namespace std;

class Solution{
public:
    vector<vector<int> > permutations(vector<int> nums){
        vector<vector<int> >ans;
        vector<int>cur;
        backTrack(ans, cur, nums, 0);
        return ans;
    }
    void backTrack(vector<vector<int> >&ans, vector<int>cur, vector<int> nums, int i){
        if(i>nums.size())
            return ;
        if(cur.size()==nums.size()){
            ans.push_back(cur);
            return ;
        }
        vector<int> new1=cur;
        new1.push_back(nums[i]);
        backTrack(ans, new1, nums, i+1);

        vector<int> new2=cur;
        new2.push_back(nums[i+1]);
        backTrack(ans, new2, nums, i);
    }
    vector<int> randDigitalForNums(int n){
        srand((unsigned)time(NULL));
        vector<int>nums;
        nums.push_back(rand()%10);
        for(int i=0;i<n;i++){
            int t;
            while(1){
                t=rand()%10;
                int j=0;
                for(; j<nums.size();j++){
                    if(t == nums[j])
                        break;
                }
                if(t != nums[j])
                    break;
            }
            nums.push_back(t);
        }
        return nums;
    }
};

int main()
{
    int n;
    while(cin>>n){
        Solution s;
        vector<int> nums=s.randDigitalForNums(n);
        vector<vector<int> >ans=s.permutations(nums);
        cout<<"["<<endl;
        for(int i=0;i<ans.size();i++){
            cout<<"[";
            for(int j=0;j<ans[0].size();j++){
                cout<<ans[i][j];
            }
            cout<<"]";
        }
        cout<<"]"<<endl;
    }
    return 0;
}

