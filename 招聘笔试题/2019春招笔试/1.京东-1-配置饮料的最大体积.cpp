/*
第一行输入饮料种数n(0<n<=1000)，和容器能容纳的最大体积v。
第二行输入n个数，表示饮料种类的配比（a1:a2:a3...an）
第三行输入n个数，表示每种饮料的量（b1 b2 b3 ... bn）
a, b(0<=a, b<=1000000, a+b>1)
求用容器最多能容纳多少配置好的饮料，输出保留4位小数
*/


#include<iostream>
#include<vector>

using namespace std;

int main()
{
    int n;
    double v;
    while(cin>>n>>v){
        vector<double> a(n, 0);
        vector<double> b(n, 0);
        double sum=0;
        for(int i=0;i<n;i++){
            cin>>a[i];
            sum+=a[i];
        }
        for(int i=0;i<n;i++){
            cin>>b[i];
        }
        int i,j;
        double maxres=0;
        int k=0;
        for(i=0;i<n;i++){
            int flag = 0;
            double zb = a[i]/sum;
            double num = b[i]/zb;
            for(j=0;j<n;j++){
                if(num*(a[j]/sum)>b[j]){
                    flag=1;
                    break;
                }
            }
            if(flag==0){
                if(maxres<num){
                    maxres=num;
                }
            }
        }
        if(maxres>v){
            printf("%.4f", v);
        }else{
            printf("%.4f", maxres);
        }
    }
    return 0;
}
