/*
输入一个数n，求n的阶乘的末尾的非零数
示例：
输入：6
输出：2
解释：6！=720，末尾非0数为2
*/



#include<iostream>
#include<vector>

using namespace std;

int main()
{
    int n;
    while(cin>>n){
        int res = 1;
        for(int i=2;i<=n;i++){
            res*=i;
            while(res%10==0){
                res = res/10;
            }
            //这个需要调参，100~1000000
            res = res % 100;
        }
        cout<<res%10<<endl;
    }
    return 0;
}
