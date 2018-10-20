#include<iostream>
#include<vector>
#include<utility>
#include<iomanip>
using namespace std;

typedef vector<pair<int,string>> VecPar;

int len;
int maxHeapify(VecPar &vp, int i);
int buildMaxHeapify(VecPar &vp);
int heapSort(VecPar &vp);
int leftTree(int i);
int rightTree(int i);
int parentTree(int i);

/*
1.把元素x插入集合S中
insert(S,x);
2.返回S中具有最大键值的元素
maximum(S);
3.去掉并返回S中具有最大键值的元素
extractMax(S);
4.将元素x的关键字增加到k，这里假设k的值不小于x的值
increaseKey(S,x,k);
*/
int heapInsert(VecPar &vp, int x);
string heapMaximum(VecPar vp);
pair<int,string> heapExtractMax(VecPar &vp);
int heapIncreaseKey(VecPar &vp, int i, int k);

int printPriorityAndElement(VecPar vp, int n);

int main()
{
    int a[11]= {4,2,6,1,7,3,5,0,9,8,11};
    string str[11]= {"a","b","c","d","e",
                     "f","g","h","i","j","k"
                    };
    vector<pair<int,string>>vp;
    for(int i=0; i<11; i++) {
        vp.push_back(make_pair(a[i],str[i]));
    }
    len=vp.size();
    buildMaxHeapify(vp);
    printPriorityAndElement(vp,2);

//    string maxs=heapMaximum(vp);
//    cout<<"max string is:"<<maxs<<endl;
//
//    cout<<"删除优先级最高的元素"<<endl;
//    pair<int,string>p;
//    p=heapExtractMax(vp);
//    cout<<"优先级最高元素为："<<endl;
//    printPriorityAndElement(vp,2);
//
//    cout<<"将元素优先级为4的元素变为8"<<endl;
//    heapIncreaseKey(vp,4,8);
//    printPriorityAndElement(vp,2);

    cout<<"插入优先级为6的元素"<<endl;
    heapInsert(vp,6);
    printPriorityAndElement(vp,2);
    return 0;
}

int leftTree(int i)
{
    return 2*i;
}

int rightTree(int i)
{
    return 2*i+1;
}

int parentTree(int i)
{
    return (i-1)/2;
}

int maxHeapify(VecPar &vp, int i)
{
    int l=leftTree(i);
    int r=rightTree(i);
    int largest=i;
    if(l<len && vp[l].first>vp[largest].first) {
        largest=l;
    }
    if(r<len && vp[r].first>vp[largest].first) {
        largest=r;
    }
    if(largest!=i) {
        pair<int,string>t=vp[i];
        vp[i]=vp[largest];
        vp[largest]=t;
        maxHeapify(vp,largest);
    }
    return 0;
}

int buildMaxHeapify(VecPar &vp)
{
    int n=len-1;
    for(int i=n/2; i>=0; i--) {
        maxHeapify(vp, i);
    }
    return 0;
}

int heapSort(VecPar &vp)
{
    buildMaxHeapify(vp);
    pair<int,string> t;
    for(int i=vp.size()-1; i>0; i--) {
        t=vp[0];
        vp[0]=vp[i];
        vp[i]=t;
        //长度减一
        len--;
        maxHeapify(vp,0);
    }
    return 0;
}

string heapMaximum(VecPar vp)
{
    return vp[0].second;
}

pair<int,string> heapExtractMax(VecPar &vp)
{
    if(len<1) {
        cout<<"heap underflow"<<endl;
    }
    pair<int,string>max;
    max=vp[0];
    int n=vp.size()-1;
    vp[0]=vp[n];
    len--;
    maxHeapify(vp,0);
    return max;
}
//设置索引值为i的元素的优先级为k
int heapIncreaseKey(VecPar &vp, int i,int k)
{
    if(i<vp[i].first) {
        cout<<"new key is smaller than old key"<<endl;
        return 0;
    }
    vp[i].first=k;
    pair<int,string>t;
    while(i>0 && vp[i].first>parentTree(i)) {
        t=vp[i];
        vp[i]=vp[parentTree(i)];
        vp[parentTree(i)]=t;
        i=parentTree(i);
    }
    return 0;
}
//插入一个优先级为x的元素
int heapInsert(VecPar &vp, int x)
{
    len++;
    pair<int,string>p;
    p.first=x;
    cout<<"please input the element:";
    cin>>p.second;
    vp[len-1]=p;
    heapIncreaseKey(vp,len-1,x);
    return 0;
}

int printPriorityAndElement(VecPar vp,int n)
{
    cout<<setiosflags(ios::left);
    if(n==2) {
        cout<<"元素优先级和元素为："<<endl;
        for(int i=0; i<len; i++) {
            cout<<setw(4)<<vp[i].first;
        }
        cout<<endl;
        for(int i=0; i<len; i++) {
            cout<<setw(4)<<vp[i].second;
        }
    }
    if(n==0) {
        cout<<"元素优先级为："<<endl;
        for(int i=0; i<len; i++) {
            cout<<setw(4)<<vp[i].first;
        }
    }
    if(n==1) {
        cout<<"元素优先级对应元素为："<<endl;
        for(int i=0; i<len; i++) {
            cout<<setw(4)<<vp[i].second;
        }
    }
    cout<<endl;
    return 0;
}
