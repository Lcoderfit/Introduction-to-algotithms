/**************************************************************
7.������ת��reverse-integer
leetcode���ӣ�https://leetcode-cn.com/problems/reverse-integer/submissions/
ʱ�临�Ӷȣ�O(log10(n)),nΪ���������
�ռ临�Ӷȣ�O(1)
�ĵã�
1.�����ǰһ��״̬�ж���һ��״̬�Ƿ�����
2.�������ݷ�Χ��ʹ�����ܽ���ԭ����������ݣ��Ӷ������ж��Ƿ����
**************************************************************/
#include<iostream>
//����pow()����
#include<math.h>
#include<vector>
using namespace std;

/*
//int��������ȡ�õ����ֵ
//2147483647,IMAX+1==-2147483648
//IMAX+1+a==-2147483648+a
int const IMAX=pow(2,31)-1;
//int��������ȡ�õ���Сֵ
//-2147483648��IMIN-1==2147483647
//IMIN-1-a==2147483647-a
int const IMIN=pow(-2,31);
*/

/*******�ⷨһ******/
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

/******�ⷨ��******/
//-2^31
static const int IMIN=0x80000000;
//2^31-1
static const int IMAX=0x7FFFFFFF;

class Solution2{
public:
    int reverse(int x){
        //sizeof(long)>sizeof(int)
        //��Linux64ϵͳ�ϣ�sizeof(long)==64
        //long long �����л����¶�ռ64λ
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
