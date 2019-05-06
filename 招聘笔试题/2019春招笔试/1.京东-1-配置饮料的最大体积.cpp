/*
��һ��������������n(0<n<=1000)�������������ɵ�������v��
�ڶ�������n��������ʾ�����������ȣ�a1:a2:a3...an��
����������n��������ʾÿ�����ϵ�����b1 b2 b3 ... bn��
a, b(0<=a, b<=1000000, a+b>1)
����������������ɶ������úõ����ϣ��������4λС��
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
