/*************************************************
1����������
2�����ȶ���ѡ������,�÷��εķ������Ȳ�ֺ�ݹ顣
��������Ϊ2 3 2 0 6 1,��ô��jָ��0ʱ���Ὣ0���һ��2
���н������Ӷ��ı���ԭ����˳��
3��ʱ�临�Ӷȣ�ƽ��O(nlog(n)) �O(n^2) ���O(nlog(n))
4���ռ临�Ӷȣ�
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
    cout<<"����������£�";
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
