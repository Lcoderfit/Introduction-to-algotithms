#include<iostream>
//容器
#include<vector>
using namespace std;

//用于存放最优代价和最优分割方案
class MatrixAndPrint{
public:
    vector<vector<int>>m;
    vector<vector<int>>s;
};
//求解矩阵链最优分割方案和最少代价
MatrixAndPrint MatrixChainOrder(vector<int>p);
//打印最佳括号分割方案
int PrintOptimalParens(vector<vector<int>>s, int i, int j);

int main()
{
    vector<int>p(7,0);
    int a[7]={30,35,15,5,10,20,25};
    for(int i=0;i<p.size();i++){
        p[i]=a[i];
    }
    MatrixAndPrint ms;
    ms=MatrixChainOrder(p);
    cout<<ms.m[1][6]<<endl;
    PrintOptimalParens(ms.s, 1, 6);
    return 0;
}

//求解矩阵链最优分割方案和最少代价
MatrixAndPrint MatrixChainOrder(vector<int>p)
{
    int n=p.size()-1;
    vector<int>t(n+1,0);
    vector<vector<int>>m(n+1,t);
    vector<vector<int>>s(n+1,t);
    for(int l=2;l<=n;l++){
        for(int i=1;i<=n-l+1;i++){
            int j=i+l-1;
            m[i][j]=0x7ffffffe;
            for(int k=i;k<=j-1;k++){
                int q=m[i][k]+m[k+1][j]+p[i-1]*p[k]*p[j];
                if(m[i][j]>q){
                    m[i][j]=q;
                    s[i][j]=k;
                }
            }
        }
    }
    MatrixAndPrint ms;
    ms.m=m;
    ms.s=s;
    return ms;
}

//打印最优括号分割方案
int PrintOptimalParens(vector<vector<int>>s, int i, int j)
{
    if(i==j){
        cout<<"A"<<i;
    }
    else{
        cout<<"(";
        PrintOptimalParens(s, i, s[i][j]);
        PrintOptimalParens(s, s[i][j]+1, j);
        cout<<")";
    }
    return 0;
}
