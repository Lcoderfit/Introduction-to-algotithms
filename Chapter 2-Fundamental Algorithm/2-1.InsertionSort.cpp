#include<iostream>
#include<vector>
using namespace std;

int insertionSort(vector<int>&v);

int main()
{
    vector<int>v{9,5,7,2,6,8,1,3,4,0};
    insertionSort(v);
    for(int i=0;i<v.size();i++){
        cout<<v[i]<<" ";
    }
    cout<<endl;
    return 0;
}

int insertionSort(vector<int>&v)
{
    for(int j=1;j<v.size();j++){
        int k=v[j];
        int i=j-1;
        while(i>=0 && v[i]>k){
            v[i+1]=v[i];
            i=i-1;
        }
        v[i+1]=k;
    }
    return 0;
}
