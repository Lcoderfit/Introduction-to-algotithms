/*************************************************
1���鲢����
2���ȶ��ıȽ������ù鲢�ķ������ȷֲ��ݹ��鲢��
3��ʱ�临�Ӷȣ����������ΪO(nlog(n))
4���ռ临�Ӷȣ���ҪO(n)�Ķ���洢�ռ�
**************************************************/
#include<iostream>
#include<time.h>
#include<vector>

using namespace std;


class Solution{
public:
    //������õ�a[p...q]��a[q+1...r]��������
    void merge(vector<int>& a, int p, int q, int r);
    //��δ�����a[p...r]��������
    void mergeSort(vector<int>& a, int p, int r);
    //��ӡ����
    void printArray(vector<int> a);
    //���ɰ���n�������������
    vector<int> randForArray(int n);
};


int main()
{
    int n;
    while(cin>>n){
        Solution s;
        vector<int> a = s.randForArray(n);
        s.printArray(a);

        s.mergeSort(a, 0, n-1);
        s.printArray(a);
    }
    return 0;
}

void Solution::merge(vector<int>& a, int p, int q, int r){
    //������䣬����������������ִ�У��������򱨴�
    if(!(a.size()>0 && p>=0 && r<a.size() && p<=q && q<=r))
        return ;

    int n1 = q-p+1; //a[p...q]��Ԫ�ظ���
    int n2 = r-q;   //a[q+1...r]��Ԫ�ظ���

    vector<int> L(n1+1, INT_MAX);   //���һ��Ԫ���趨Ϊ�Ƚ�ʱ���˳�����
    vector<int> R(n2+1, INT_MAX);

    for(int i=0;i<n1;i++){
        L[i] = a[i+p];
    }
    for(int j=0;j<n2;j++){
        R[j] = a[j+q+1];
    }

    int i=0;
    int j=0;
    //��L��R�е�Ԫ�ش�С�������a��
    for(int k=p; k<=r; k++){
        if(L[i]<R[j]){
            a[k]=L[i];
            i++;
        }else{
            a[k]=R[j];
            j++;
        }
    }
}

void Solution::mergeSort(vector<int>& a, int p, int r){
    if(!(a.size()>0 && p>=0 && r<a.size() && p<r))
        return ;

    int q=(p+r)/2;
    mergeSort(a, p, q);
    mergeSort(a, q+1, r);
    merge(a, p, q, r);
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
    for(int i=0;i<n; i++){
        a.push_back(rand()%100);
    }
    return a;
}
