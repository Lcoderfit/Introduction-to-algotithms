#include<iostream>
#include<vector>
#include<iomanip>

using namespace std;

class Solution{
public:
    int randomizeInPlace(vector<int> &a){
        int n=a.size();
        cout<<setiosflags(ios::left);
        for(int i=0;i<n;i++){
            cout<<setw(4)<<i;
        }
        cout<<endl;
        for(int i=0;i<n;i++){
            cout<<setw(4)<<a[i];
        }
        cout<<endl;
        for(int i=0;i<n-1;i++){
            int q=i+rand()%(n-i);
            cout<<setw(4)<<q;
            swap(a[i],a[q]);
        }
        cout<<endl;
        for(int i=0;i<n;i++){
            cout<<setw(4)<<a[i];
        }
    }
};

int main()
{
    Solution s;
    vector<int>a{0,1,2,3,4,5,6,7,8,9};
    s.randomizeInPlace(a);
    return 0;
}
