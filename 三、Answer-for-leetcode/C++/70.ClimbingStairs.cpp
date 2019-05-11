/**************************************************************
70.��¥��:climbing-stairs
leetcode���ӣ�https://leetcode-cn.com/problems/climbing-stairs/
**************************************************************/

#include<iostream>
//STL����
#include<vector>
//�����ռ�
using namespace std;

//����һ������������
class Solution {
public:
    //����һ����¥�ݷ���
    int climbStairs(int n) {
        //����һ��һά������Ԫ�ظ���Ϊn+1��
        //��������Ԫ��ȫ����ֵΪ0
        vector<int>dp(n+1,0);
        //��ʼ��dp[0]��dp[1]
        dp[0]=1;
        dp[1]=1;
        //��̬�滮�㷨��dp[i]��ʾ�����i��̨�׵ķ�����
        //�����i��̨�׵ķ�����=�����i-1��̨�׵ķ�����
        //+�����i-2��̨�׵ķ�����(i>=2)
        for(int i=2;i<=n;i++){
            dp[i]=dp[i-1]+dp[i-2];
        }
        //����dp[n]:�����n��̨�׵ķ�����
        return dp[n];
    }
};

//�����������в�������
int main()
{
    int n;
    int res;
    //������Զ�������������
    while(cin>>n){
        Solution s;
        //res��������n��̨�׵ķ�����
        res=s.climbStairs(n);
        cout<<res<<endl;
    }

    return 0;
}
