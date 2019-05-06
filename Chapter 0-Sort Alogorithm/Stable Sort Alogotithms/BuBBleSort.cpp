/***********************************
1、冒泡排序
2、稳定的交换排序
3、时间复杂度：平均O(n^2)、最坏O(n^2)、最好O(n)
4、空间复杂度：O(1)
***********************************/
#include<iostream>
#include<time.h>
#include<vector>

using namespace std;

class Solution {
  public:
    void bubbleSort(vector<int>& a, int rev);
    void printArray(vector<int> a);
    vector<int> randForArray(int n);
};

int main() {
    int n;
    while(cin>>n){
        Solution s;
        vector<int> a = s.randForArray(n);
        s.printArray(a);

        s.bubbleSort(a, 0);
        s.printArray(a);
    }
    return 0;
}

void Solution::bubbleSort(vector<int>& a, int rev) {
    for(int i=0; i<a.size()-1; i++) {
        for(int j=0; j<a.size()-i-1; j++) {
            if(rev==0) {
                if(a[j]>a[j+1]) {
                    int t=a[j];
                    a[j]=a[j+1];
                    a[j+1]=t;
                }
            }else{
                for(int j=0;j<a.size()-i-1;j++){
                    if(a[j]<a[j+1]){
                        int t=a[j];
                        a[j]=a[j+1];
                        a[j+1]=t;
                    }
                }
            }
        }
    }
}

void Solution::printArray(vector<int> a){
    cout<<"序列输出如下："<<endl;
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
