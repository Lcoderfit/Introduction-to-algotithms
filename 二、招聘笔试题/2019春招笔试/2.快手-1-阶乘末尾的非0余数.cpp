/*
����һ����n����n�Ľ׳˵�ĩβ�ķ�����
ʾ����
���룺6
�����2
���ͣ�6��=720��ĩβ��0��Ϊ2
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
            //�����Ҫ���Σ�100~1000000
            res = res % 100;
        }
        cout<<res%10<<endl;
    }
    return 0;
}
