#include <iostream>
#include <vector>
#include <iomanip>
using namespace std;

vector<int> greedyActivitySelector(vector<int>s, vector<int>f);

int main()
{
    vector<int>s{0,1,3,0,5,3,5,6,8,8,2,12};
    vector<int>f{0,4,5,6,7,9,9,10,11,12,14,16};
    vector<int>res;
    res=greedyActivitySelector(s, f);
    for(int i=0; i<res.size();i++){
        cout<<setw(3)<<res[i];
    }
    return 0;
}

vector<int> greedyActivitySelector(vector<int>s, vector<int>f)
{
    vector<int>res;
    int length=s.size()-1;
    int k=0;
    for(int m=1; m<=length; m++){
        if(f[k]<=s[m]){
            res.push_back(m);
            k=m;
        }
    }
    return res;
}
