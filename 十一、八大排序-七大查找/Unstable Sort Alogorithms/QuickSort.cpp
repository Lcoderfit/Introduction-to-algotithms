/*************************************************
1、快速排序
2、不稳定的选择排序,用分治的方法，先拆分后递归。
假设序列为2 3 2 0 6 1,那么当j指向0时，会将0与第一个2
进行交换，从而改变了原来的顺序。
3、时间复杂度：平均O(nlog(n)) 最坏O(n^2) 最好O(nlog(n))
4、空间复杂度：
*************************************************/
#include<iostream>
#include<time.h>
#include<vector>

using namespace std;


class Solution{
public:
    void quickSort(vector<int>& a, int p, int r);
    int partition(vector<int>& a, int p, int r);
    void printArray(vector<int> a);
    vector<int> randForArray(int n);
};


int main(){
    int n;
    while(cin>>n){
        Solution s;
        vector<int> a = s.randForArray(n);
        s.printArray(a);

        s.quickSort(a, 0, n-1);
        s.printArray(a);
    }
}

void Solution::quickSort(vector<int>& a, int p, int r){
    if(p<r){
        int q=partition(a, p, r);
        quickSort(a, p, q-1);
        quickSort(a, q+1, r);
    }
}

int Solution::partition(vector<int>& a, int p, int r){
    int k=a[r];
    int i=p-1;
    for(int j=p;j<r;j++){
        if(a[j]<=k){
            i++;
            swap(a[i], a[j]);
        }
    }
    swap(a[i+1], a[r]);

    return i+1;
}

void Solution::printArray(vector<int> a){
    cout<<"序列输出如下：";
    for(int i=0;i<a.size();i++){
        cout<<a[i]<<" ";
    }
    cout<<endl<<endl;
}

vector<int> Solution::randForArray(int n){
    vector<int> a;
    srand((unsigned)time(NULL));
    for(int i=0;i<n;i++){
        a.push_back(rand()%100);
    }
    return a;
}
