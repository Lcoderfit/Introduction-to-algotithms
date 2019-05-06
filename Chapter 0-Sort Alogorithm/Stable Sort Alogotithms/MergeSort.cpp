/*************************************************
1、归并排序
2、稳定的比较排序，用归并的方法，先分部递归后归并。
3、时间复杂度：三种情况均为O(nlog(n))
4、空间复杂度：需要O(n)的额外存储空间
**************************************************/
#include<iostream>
#include<time.h>
#include<vector>

using namespace std;


class Solution{
public:
    //将排序好的a[p...q]和a[q+1...r]进行排序
    void merge(vector<int>& a, int p, int q, int r);
    //对未排序的a[p...r]进行排序
    void mergeSort(vector<int>& a, int p, int r);
    //打印序列
    void printArray(vector<int> a);
    //生成包含n个随机数的序列
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
    //断言语句，满足条件继续往下执行，不满足则报错
    if(!(a.size()>0 && p>=0 && r<a.size() && p<=q && q<=r))
        return ;

    int n1 = q-p+1; //a[p...q]的元素个数
    int n2 = r-q;   //a[q+1...r]的元素个数

    vector<int> L(n1+1, INT_MAX);   //最后一个元素设定为比较时的退出条件
    vector<int> R(n2+1, INT_MAX);

    for(int i=0;i<n1;i++){
        L[i] = a[i+p];
    }
    for(int j=0;j<n2;j++){
        R[j] = a[j+q+1];
    }

    int i=0;
    int j=0;
    //将L和R中的元素从小到大存入a中
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
    cout<<"序列输出如下：";
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
