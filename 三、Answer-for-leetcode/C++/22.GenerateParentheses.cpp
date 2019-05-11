/**************************************************************
9.括号生成：generate-parenthesis
leetcode链接：https://leetcode-cn.com/problems/generate-parentheses/
时间复杂度：O(4^n/根号(n))
空间复杂度：
time:1:28   二刷:20
**************************************************************/
#include<iostream>
#include<vector>

using namespace std;

class Solution{
public:
    vector<string> generateParenthesis(int n){
        //回溯法一般内部不定义像向量这种结构，而是作为回溯函数的参数
        vector<string> ans;
        backtrack(ans, "", 0, 0, n);
        return ans;
    }
    //采用回溯法时容易混淆不清，用构造回溯树的方法滤清思路：
    //树的初始状态为""(根节点,第0层),树的第一层有一个节点(即"("),
    //然后根据回溯算法函数递归顺序，先弄一条直到叶子节点的路径(即最后一层)，然后回溯到上一层
    //以深度优先的方式搜索解空间，并且在搜索过程中用剪枝函数避免无效搜索。
    //ans:存储结果 cur:回溯的中间状态   open&close:搜索条件    max:回溯成功的判断条件
    void backtrack(vector<string>& ans, string cur,
                   int open, int close, int max){
        if(cur.size() == max*2){
            ans.push_back(cur);
            //回溯退出时候一般直接return,不返回值
            return;
        }

        if(open<max){
            backtrack(ans, cur+"(", open+1, close, max);
        }
        if(close<open){
            backtrack(ans, cur+")", open, close+1, max);
        }
    }

};

int main()
{
    int n;
    vector<string> ans;
    while(cin>>n){
        Solution s;
        ans=s.generateParenthesis(n);
        for(int i=0;i<ans.size();i++){
            cout<<ans[i]<<endl;
        }
    }
    return 0;
}
