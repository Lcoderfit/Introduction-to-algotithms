/**************************************************************
9.�������ɣ�generate-parenthesis
leetcode���ӣ�https://leetcode-cn.com/problems/generate-parentheses/
ʱ�临�Ӷȣ�O(4^n/����(n))
�ռ临�Ӷȣ�
time:1:28   ��ˢ:20
**************************************************************/
#include<iostream>
#include<vector>

using namespace std;

class Solution{
public:
    vector<string> generateParenthesis(int n){
        //���ݷ�һ���ڲ����������������ֽṹ��������Ϊ���ݺ����Ĳ���
        vector<string> ans;
        backtrack(ans, "", 0, 0, n);
        return ans;
    }
    //���û��ݷ�ʱ���׻������壬�ù���������ķ�������˼·��
    //���ĳ�ʼ״̬Ϊ""(���ڵ�,��0��),���ĵ�һ����һ���ڵ�(��"("),
    //Ȼ����ݻ����㷨�����ݹ�˳����Ūһ��ֱ��Ҷ�ӽڵ��·��(�����һ��)��Ȼ����ݵ���һ��
    //��������ȵķ�ʽ������ռ䣬�����������������ü�֦����������Ч������
    //ans:�洢��� cur:���ݵ��м�״̬   open&close:��������    max:���ݳɹ����ж�����
    void backtrack(vector<string>& ans, string cur,
                   int open, int close, int max){
        if(cur.size() == max*2){
            ans.push_back(cur);
            //�����˳�ʱ��һ��ֱ��return,������ֵ
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
