/*************************************************
1、直接选择排序
2、不稳定的选择排序，例如(7) 3 4 6 [7] 1,用直接选
择排序会把第一个(7)调到[7]后面，改变了原来的顺序。
3、时间复杂度：三种情况均为O(n^2)
4、空间复杂度：O(1)
*************************************************/
#include<iostream>
#include<time.h>
#include<vector>

using namespace std;


class Solution{
public:
    void selectionSort(vector<int>& a);
    void printArray(vector<int> a);
    vector<int> randForArray(int n);
};

int main(){
    int n;
    while(cin>>n){
        Solution s;
        vector<int> a = s.randForArray(n);
        s.printArray(a);

        s.selectionSort(a);
        s.printArray(a);
    }
}

void Solution::selectionSort(vector<int>& a){
    for(int i=0;i<a.size();i++){
        for(int j=i+1;j<a.size();j++){
            if(a[i]>a[j]){
                int t=a[i];
                a[i]=a[j];
                a[j]=t;
            }
        }
    }
}

void Solution::printArray(vector<int> a){
    cout<<"输出序列如下：";
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

