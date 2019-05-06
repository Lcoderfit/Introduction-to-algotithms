#include<iostream>
#include<vector>
#include<iomanip>

using namespace std;

const int INT_MAX=0x7ffffffe;

int merge(vector<int>&a, int p, int q, int r);
int mergeSort(vector<int>&a, int p, int r);

int main()
{
    vector<int>a{6,5,4,3,2,1};
    mergeSort(a,0,5);
    cout<<setiosflags(ios::left);
    for(int i=0;i<5;i++){
        cout<<setw(4)<<a[i];
    }
    return 0;
}

int merge(vector<int>&a, int p, int q, int r)
{
    int n1=p-q+1;
    int n2=r-q;
    vector<int>L(n1+1,INT_MAX);
    vector<int>R(n2+1,INT_MAX);
    for(int i=0;i<n1;i++){
        L[i]=a[p+i];
    }
    for(int j=0;j<n2;j++){
        R[j]=a[q+j+1];
    }

    int i=0;
    int j=0;
    for(int k=p;k<=r;k++){
        if(L[i]<R[j]){
            a[k]=L[i];
            i++;
        }
        else{
            a[k]=R[j];
            j++;
        }
    }
    return 0;
}

int mergeSort(vector<int>&a, int p, int r)
{
    if(p<r){
        int q=(p+r)/2;
        mergeSort(a,p,q);
        mergeSort(a,q+1,r);
        merge(a,p,q,r);
    }
    return 0;
}
