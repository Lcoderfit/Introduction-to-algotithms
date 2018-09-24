#include<iostream>
#include<vector>
using namespace std;

int partition(vector<int>&a, int p, int r);
int quickSort(vector<int>&a, int p,int r);

int main()
{
    int n;
    cin>>n;
    vector<int>a(n,0);
    for(int i=0;i<n;i++){
        cin>>a[i];
    }
    quickSort(a,0,n-1);
    for(int i=0;i<n;i++){
        cout<<a[i]<<" ";
    }
    return 0;
}

int partition(vector<int>&a, int p, int r)
{
    int t;
    int x=a[r];
    int i=p-1;
    for(int j=p;j<=r-1;j++){
        if(a[j]<=x){
            i++;
            t=a[j];
            a[j]=a[i];
            a[i]=t;
        }
    }
    t=a[i+1];
    a[i+1]=a[r];
    a[r]=t;

    return i+1;
}

int quickSort(vector<int>&a, int p, int r)
{
    if(p<r){
        int q=partition(a,p,r);
        quickSort(a,p,q-1);
        quickSort(a,q+1,r);
    }
    return 0;
}
