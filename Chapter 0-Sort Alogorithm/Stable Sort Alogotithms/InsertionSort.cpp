/*************************************************
1、直接插入排序
2、稳定的插入排序
3、时间复杂度：平均O(n^2)、最坏O(n^2)、最好O(n)
4、空间复杂度：O(1)
*************************************************/
#include<iostream>
#include<time.h>
#include<vector>

using namespace std;

class Solution {
  public:
    //插入排序,rev==0时从小到大排序,rev==1时从大到小排序
    void insertionSort1(vector<int>& a, int rev);
    //for循环嵌套的方法实现插入排序
    void insertionSort2(vector<int>& a, int rev);
    //产生包含n个100以内随机数的向量
    vector<int> randForArray(int n);
    //打印序列
    void printArray(vector<int> a);
};

int main() {
    int n;
    while(cin>>n) {
        Solution s;
        vector<int>a = s.randForArray(n);
        s.printArray(a);

        s.insertionSort1(a, 0);
        s.printArray(a);
    }
    return 0;
}

void Solution::insertionSort1(vector<int>&a, int rev) {
    //while循环方法
    for(int i=1; i<a.size(); i++) {
        int k=a[i];
        int j=i-1;
        //从小到大排序
        if(rev==0) {
            while(j>=0 && a[j]>k) {
                a[j+1]=a[j];
                j--;
            }
            a[j+1]=k;
        } else { //从大到小排序
            while(j>=0 && a[j]<k) {
                a[j+1]=a[j];
                j--;
            }
            a[j+1]=k;
        }
    }
}

void Solution::insertionSort2(vector<int>&a, int rev) {
    //for循环方法
    for(int i=1; i<a.size(); i++) {
        int tp=a[i];
        int j;
        if(rev==0) {
            for(j=i-1; j>=0 && a[j]>tp; j--) {
                a[j+1]=a[j];
            }
            a[j+1]=tp;
        } else {
            for(j=i-1; j>=0 && a[j]<tp; j--) {
                a[j+1]=a[j];
            }
            a[j+1]=tp;
        }
    }
}

vector<int> Solution::randForArray(int n) {
    vector<int>a;
    srand((unsigned)time(NULL));
    for(int i=0; i<n; i++) {
        a.push_back(rand()%100);
    }
    return a;
}

void Solution::printArray(vector<int> a) {
    cout<<"序列输出如下:"<<endl;
    for(int i=0; i<a.size(); i++) {
        cout<<a[i]<<" ";
    }
    cout<<endl<<endl ;
}
