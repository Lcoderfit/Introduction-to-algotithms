#include<iostream>
#include<set>
#include<vector>
#include<iomanip>
using namespace std;

vector<int> recursiveActivitySelector(
    vector<int>s, vector<int>f, int k, int n);

int main()
{
    vector<int>s{0,1,3,0,5,3,5,6,8,8,2,12};
    vector<int>f{0,4,5,6,7,9,9,10,11,12,14,16};
    vector<int>res;
    res=recursiveActivitySelector(s, f, 0, 11);
    for(int i=0; i<res.size(); i++){
        cout<<setw(3)<<res[i];
    }
    return 0;
}

vector<int> recursiveActivitySelector(
        vector<int>s, vector<int>f, int k, int n)
{
    int m=k+1;
    //f[0]==s[0]==0
    while(m<=n && f[k]>s[m]){
        m++;
    }
    vector<int>v;
    if(m<=n){
        v=recursiveActivitySelector(s, f, m, n);
        v.insert(v.begin(), m);
    }
    return v;
}

//#include <iostream>
//#include <vector>
//
//using namespace std;
//
//int main()
//{
//    vector<int>a;
//    int t=1;
//    a.push_back(t);
//    a.insert(a.begin(), 3);//在开头插入元素3
//    a.insert(a.end(), 4, 1);//在末尾插入4个1
//    cout<<a[0]<<endl;
//    return 0;
//}
