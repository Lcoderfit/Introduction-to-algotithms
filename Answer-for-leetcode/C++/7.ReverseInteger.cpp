/**************************************************************
7.整数反转：reverse-integer
leetcode链接：https://leetcode-cn.com/problems/reverse-integer/submissions/
时间复杂度：O(log10(n)),n为输入的数据
空间复杂度：O(1)
心得：
1.在溢出前一个状态判断下一个状态是否会溢出
2.扩大数据范围，使变量能接收原本溢出的数据，从而加以判断是否溢出
**************************************************************/
#include<iostream>
//包含pow()函数
#include<math.h>
#include<vector>
using namespace std;

/*
//int类型所能取得的最大值
//2147483647,IMAX+1==-2147483648
//IMAX+1+a==-2147483648+a
int const IMAX=pow(2,31)-1;
//int类型所能取得的最小值
//-2147483648，IMIN-1==2147483647
//IMIN-1-a==2147483647-a
int const IMIN=pow(-2,31);
*/

/*******解法一******/
class Solution1{
public:
    int reverse(int x){
        int rev=0;
        while(x!=0){
            int pop=x%10;
            x/=10;

            if(rev>INT_MAX/10 || (rev==
                INT_MAX/10 && pop>7)){
                return 0;
            }
            if(rev<INT_MIN/10 || (rev==
                INT_MIN/10 && pop<-8)){
                return 0;
            }

            rev=rev*10+pop;
        }
        return rev;
    }
};

/******解法二******/
//-2^31
static const int IMIN=0x80000000;
//2^31-1
static const int IMAX=0x7FFFFFFF;

class Solution2{
public:
    int reverse(int x){
        //sizeof(long)>sizeof(int)
        //再Linux64系统上，sizeof(long)==64
        //long long 在所有环境下都占64位
        long long rev=0;
        while(x!=0){
            rev=rev*10+x%10;
            x/=10;
        }

        if(rev>IMAX || rev<IMIN){
            return 0;
        }
        return rev;
    }
};

int main()
{
    int n;

    while(cin>>n){
        Solution2 s;
        int ret=s.reverse(n);
        cout<<ret<<endl;
    }
    return 0;
}
