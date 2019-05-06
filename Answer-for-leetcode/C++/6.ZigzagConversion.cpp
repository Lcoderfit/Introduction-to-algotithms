/**************************************************************
9.Z字形变换：zigzag-conversion
leetcode链接：https://leetcode-cn.com/problems/zigzag-conversion/solution/
时间复杂度：O(n),n==len(s)
空间复杂度：O(n)
**************************************************************/

//#include<iostream>
//#include<vector>
//using namespace std;
//
//class Solution {
//  public:
//    string convert(string s, int numRows) {
//        if(numRows==1) {
//            return s;
//        }
//        int p=numRows*2-2;
//        int num=s.size()/p;
//        int numCols=num*(numRows-1);
//        int q=s.size()%p;
//        if(q>0 && q<=numRows) {
//            numCols++;
//        } else if(q>numRows) {
//            numCols=numCols+q-3;
//        }
//        vector<char>t(numCols, '\0');
//        vector<vector<char> >chVec(numRows, t);
//        int i=0;
//        int j=0;
//        int count=0;
//        for(int k=0; k<s.size(); k++) {
//            count=0;
//            chVec[i][j]=s[k];
//            if(i==0) {
//                i++;
//                count++;
//            }
//            if(count==0 && i>0 && i<numRows) {
//                if(chVec[i-1][j]!='\0' && i!=numRows-1) {
//                    i++;
//                    count++;
//                } else if(i==numRows-1) {
//                    i--;
//                    j++;
//                    count++;
//                } else {
//                    i--;
//                    j++;
//                    count++;
//                }
//            }
//        }
//        string str="";
//        for(int i=0; i<numRows; i++) {
//            for(int j=0; j<numCols; j++) {
//                if(chVec[i][j]!='\0') {
//                    str=str+chVec[i][j];
//                }
//            }
//        }
//        return str;
//    }
//};
//
//int main()
//{
//    Solution s;
//    string str;
//    int numRows;
//    string res;
//    while(cin>>str>>numRows){
//        res=s.convert(str, numRows);
//        cout<<res<<endl;
//    }
//    return 0;
//}

#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

class Solution{
public:
    string convert(string s, int numRows){
        if(numRows==1)
            return s;

        vector<string> rows(min(numRows, int(s.size())));

        int curRow=0;
        int goDown=false;

        for(char c:s){
            rows[curRow]+=c;
            if(curRow==0 || curRow==numRows-1){
                goDown=!goDown;
                curRow+=goDown?1:-1;
            }
        }

        string ret;
        for(string row:rows){
            ret+=row;
        }
        return ret;
    }
};

int main()
{
    Solution s;
    string str;
    int numRows;
    while(cin>>str>>numRows){
        string ret=s.convert(str, numRows);
        cout<<ret<<endl;
    }
    return 0;
}

