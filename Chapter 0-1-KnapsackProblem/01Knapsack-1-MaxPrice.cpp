#include<iostream>
#include<vector>

using namespace std;

class Solution{
public:
    vector<vector<int> > maxPriceArray(vector<int>a, int value);
    int maxPrice(vector<vector<int> >ans, vector<int>a);
};

int main()
{
    int kind;
    int value;
    while(cin>>value>>kind){
        vector<int> a(kind, 0);
        for(int i=0;i<kind;i++){
            cin>>a[i];
        }

        Solution s;
        vector<vector<int> > ans = s.maxPriceArray(a, value);
        int price = s.maxPrice(ans, a);
        cout<<price<<endl;
    }
    return 0;
}

vector<vector<int> > Solution::maxPriceArray(vector<int> a, int value){
    int kind=a.size();
    vector<int>t(value+1, 0);
    vector<vector<int> > ans(kind+1, t);

    for(int i=1;i<kind;i++){
        for(int j=1;j<=value;j++){
            if(j<a[i-1]){
                ans[i][j] = ans[i-1][j];
            }else{
                ans[i][j] = max(ans[i-1][j], ans[i-1][j-a[i-1]]+1);
            }
        }
    }

    return ans;
}

//倒着来判断是否取了
int Solution::maxPrice(vector<vector<int> >ans, vector<int>a){
    int kind = ans.size()-1;
    int value = ans[0].size()-1;

    int res=0;
    for(int i=kind;i>=1;i--){
        if(ans[i][value]!=ans[i-1][value]){
            res +=a[i-1];
            value-=a[i-1];
        }
    }

    return res;
}
