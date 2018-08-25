#include<iostream>
#include<vector>//容器
#include<ctime>//包含srand()和rand()，用于生成随机数
#include<iomanip>//包含输出格式控制函数
using namespace std;

//如果觉得二维容器书写太麻烦可以用typedef取别名替换。
//typedef vector<vector<int>> tdvec;

//计算两矩阵相乘的函数，Matrix：矩阵，multiply：相乘
vector<vector<int>> TwoMatrixMultiply(vector<vector<int>>a, vector<vector<int>>b);
//容器随机数初始化函数
int VectorRandomInit(vector<vector<int>>& v);

int main()
{
    vector<int>t1(3,0);
    vector<int>t2(4,0);
    vector<vector<int>>a(2,t1);
    vector<vector<int>>b(3,t2);

    VectorRandomInit(a);//用随机数初始化矩阵a
    VectorRandomInit(b);//用随机数初始化矩阵b

    vector<vector<int>>c;
    c=TwoMatrixMultiply(a,b);

    cout<<"两矩阵相乘的结果矩阵为："<<endl;
    //控制输出格式为左对齐
    cout<<setiosflags(ios::left);
    for(int i=1;i<=a.size();i++){
        for(int j=1;j<=b[0].size();j++){
            cout<<setw(4)<<c[i-1][j-1];
        }
        cout<<endl;
    }

    return 0;
}

//容器随机数初始化函数
int VectorRandomInit(vector<vector<int>>& v)
{
    cout<<"随机数生成的"<<v.size()<<"X";
    cout<<v[0].size()<<"矩阵如下："<<endl;
    cout<<setiosflags(ios::left);
    srand((unsigned)time(NULL));//利用系统时间获取随机数种子
    //v.size()是容器v对应的行数，v[0].size()计算容器的列数
    for(int i=0;i<v.size();i++){
        for(int j=0;j<v[0].size();j++){
            v[i][j]=rand()%10;//产生[0,10)的随机数
            cout<<setw(4)<<v[i][j];
        }
        cout<<endl;
    }
    cout<<endl;
    return 0;
}

//两矩阵相乘函数
vector<vector<int>> TwoMatrixMultiply(vector<vector<int>>a, vector<vector<int>>b)
{
    if(a[0].size() != b.size()){
        printf("error");
        vector<vector<int>>c;
        return c;
    }
    else{
        vector<int>t(b[0].size(),0);
        vector<vector<int>>c(a.size(),t);
        for(int i=1;i<=a.size();i++){
            for(int j=1;j<=b[0].size();j++){
                int s=0;
                for(int k=1;k<=a[0].size();k++){
                    s+=a[i-1][k-1]*b[k-1][j-1];
                }
                c[i-1][j-1]=s;
            }
        }
        return c;
    }
}
