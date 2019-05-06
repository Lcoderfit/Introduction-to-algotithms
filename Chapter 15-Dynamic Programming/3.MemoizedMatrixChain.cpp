#include<iostream>
#include<vector>
using namespace std;

//带备忘录的自顶向下动态规划矩阵链乘法
int memoizedChain(vector<int>p);
//备忘录查表函数
int lookupChain(vector<vector<int>>m, vector<int>p, int i, int j);

int main()
{
    //定义一个容器保存矩阵链的规模信息
    vector<int>p{30,35,15,5,10,20,25};
    int res=memoizedChain(p);
    cout<<res<<endl;
    return 0;
}

//带备忘录的自顶向下动态规划矩阵链乘法
int memoizedChain(vector<int>p)
{
    int n=p.size()-1;
    //容器内的全部元素赋初值为int类型最大值
    vector<int>t(n+1,0x7ffffffe);
    vector<vector<int>>m(n+1,t);

    return lookupChain(m, p, 1, n);
}

//备忘录查表函数
int lookupChain(vector<vector<int>>m, vector<int>p, int i, int j)
{
    if(m[i][j]<0x7ffffffe)
        return m[i][j];
    if(i==j)
        m[i][j]=0;
    else{
        for(int k=i;k<=j-1;k++){
            int q=lookupChain(m, p, i,k) + \
                    lookupChain(m, p, k+1, j) + \
                    p[i-1]*p[k]*p[j];
            if(q<m[i][j])
                m[i][j] = q;
        }
    }
    return m[i][j];
}
