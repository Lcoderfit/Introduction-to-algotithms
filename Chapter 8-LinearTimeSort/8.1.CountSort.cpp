#include<iostream>
#include<vector>
#include<iomanip>
#include<iomanip>
using namespace std;

typedef vector<int> TypedVec;
int MIN_INT=-0x7ffffffe;

class Solution{
public:
    int countSort(TypedVec a, TypedVec &b){
        int max=MIN_INT;
        for(int i=0;i<a.size();i++){
            if(max<a[i]){
                max=a[i];
            }
        }
        TypedVec c(max+1,0);
        for(int i=0;i<a.size();i++){
            c[a[i]]++;
        }
        for(int i=1;i<c.size();i++){
            c[i]=c[i]+c[i-1];
        }
        for(int i=a.size()-1;i>=0;i--){
            b[c[a[i]]-1]=a[i];
            c[a[i]]--;
        }
        return 0;
    }
};

int main()
{
    TypedVec a{3,2,4,6,1,1,7,9,0,5,6,2};
    TypedVec b(a.size(),0);
    Solution s;
    s.countSort(a, b);
    cout<<setiosflags(ios::left);
    for(int i=0;i<a.size();i++){
        cout<<setw(4)<<b[i];
    }
    return 0;
}
