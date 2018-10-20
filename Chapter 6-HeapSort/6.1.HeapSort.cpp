#include<iostream>
#include<vector>
using namespace std;

int len;
//维持大根堆
int maxHeapify(vector<int>&a, int i);
//建立大根堆
int buildMaxHeap(vector<int>&a);
//堆排序
int heapSort(vector<int>&a);

int main()
{
    vector<int>a{2,3,4,1,6,7,9,8,0,5};
    len=a.size();
    heapSort(a);
    for(int i=0;i<a.size();i++){
        cout<<a[i]<<" ";
    }
    return 0;
}

int maxHeapify(vector<int>&a, int i)
{
    int l=2*i;
    int r=2*i+1;
    int largest=i;
    if(l<len && a[l]>a[i]){
        largest=l;
    }
    if(r<len && a[r]>a[largest]){
        largest=r;
    }
    if(largest!=i){
        int t=a[largest];
        a[largest]=a[i];
        a[i]=t;
        maxHeapify(a,largest);
    }
    return 0;
}

int buildMaxHeap(vector<int>&a)
{
    int n=a.size()-1;
    for(int i=n/2;i>=0;i--){
        maxHeapify(a,i);
    }
    return 0;
}

int heapSort(vector<int>&a)
{
    buildMaxHeap(a);
    int t;
    for(int i=a.size()-1;i>0;i--){
        t=a[0];
        a[0]=a[i];
        a[i]=t;
        len--;
        maxHeapify(a,0);
    }
    return 0;
}
