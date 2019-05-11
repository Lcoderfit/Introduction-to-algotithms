#include<iostream>
#include<vector>
using namespace std;

int gcd(int a, int b){
    int t=0;
    while(b!=0){
        t=a;
        a=b;
        b=t%b;
    }
    return a;
}

int main()
{
    int n;
    while(cin>>n){
        vector<int>a(10005, 0);
        int t;
        for(int i=1;i<=n;i++){
            cin>>t;
            a[t] += 1;
        }
        int j=1;
        while(a[j]==0){
            j++;
        }
        int b = a[j];
        for(int i=j+1;i<a.size();i++){
            if(a[i])
                b = gcd(b, i);
        }
        if(b<2){
            cout<<0<<endl;
            continue;
        }

        int sum = 0;
        for(int i=j;i<a.size();i++){
            if(a[i]){
                sum += a[i]/b;
            }
        }
        cout<<sum<<endl;
    }
    return 0;
}
