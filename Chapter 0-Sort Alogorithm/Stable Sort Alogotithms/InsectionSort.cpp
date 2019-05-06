/*************************************************
1��ֱ�Ӳ�������
2���ȶ��Ĳ�������
3��ʱ�临�Ӷȣ�ƽ��O(n^2)���O(n^2)�����O(n)
4���ռ临�Ӷȣ�O(1)
*************************************************/
#include<iostream>
#include<time.h>
#include<vector>

using namespace std;

class Solution {
  public:
    //��������,rev==0ʱ��С��������,rev==1ʱ�Ӵ�С����
    void insertionSort1(vector<int>& a, int rev);
    //forѭ��Ƕ�׵ķ���ʵ�ֲ�������
    void insertionSort2(vector<int>& a, int rev);
    //��������n��100���������������
    vector<int> randForArray(int n);
    //��ӡ����
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
    //whileѭ������
    for(int i=1; i<a.size(); i++) {
        int k=a[i];
        int j=i-1;
        //��С��������
        if(rev==0) {
            while(j>=0 && a[j]>k) {
                a[j+1]=a[j];
                j--;
            }
            a[j+1]=k;
        } else { //�Ӵ�С����
            while(j>=0 && a[j]<k) {
                a[j+1]=a[j];
                j--;
            }
            a[j+1]=k;
        }
    }
}

void Solution::insertionSort2(vector<int>&a, int rev) {
    //forѭ������
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
    cout<<"�����������:"<<endl;
    for(int i=0; i<a.size(); i++) {
        cout<<a[i]<<" ";
    }
    cout<<endl<<endl ;
}
