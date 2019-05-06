#include<iostream>
#include<vector>
#include<tuple>
#include<iomanip>

using namespace std;


class Solution{
public:
    const int MIN_INT=-0x7ffffffe;
    typedef tuple<int,int,int> TupleType;
    //求跨分隔点的最大子数组
    TupleType findMaxCrossingSubarray(vector<int>&a, int low, int mid, int high)
    {
        int leftsum=MIN_INT;
        int leftmark=0;
        int sum=0;
        for(int i=mid;i>=low;i--){
            sum+=a[i];
            if(sum>leftsum){
                leftsum=sum;
                leftmark=i;
            }
        }
        int rightsum=MIN_INT;
        int rightmark=0;
        sum=0;
        for(int j=mid+1;j<=high;j++){
            sum+=a[j];
            if(sum>rightsum){
                rightsum=sum;
                rightmark=j;
            }
        }
        int res=leftsum+rightsum;
        return make_tuple(leftmark,rightmark,res);
    }
    //求整个数组的最大子数组，先求分隔点左边的最大子数组，然后求分隔点右边的，最后跨分隔点的
    TupleType findMaximumSubarray(vector<int>&a, int low, int high)
    {
        if(low==high){
            return make_tuple(low,high,a[low]);
        }
        else{
            int mid=(low+high)/2;
            int leftlow,lefthigh,leftsum;
            tie(leftlow,lefthigh,leftsum)=findMaximumSubarray(a,low,mid);
            int rightlow,righthigh,rightsum;
            tie(rightlow,righthigh,rightsum)=findMaximumSubarray(a,mid+1,high);
            int crosslow,crosshigh,crosssum;
            tie(crosslow,crosshigh,crosssum)=findMaxCrossingSubarray(a,low,mid,high);
            if(leftsum>=rightsum && leftsum>=crosssum){
                return make_tuple(leftlow,lefthigh,leftsum);
            }
            else if(rightsum>=leftsum && rightsum>=crosssum){
                return make_tuple(rightlow,righthigh,rightsum);
            }
            else{
                return make_tuple(crosslow,crosshigh,crosssum);
            }
        }
    }
};

int main()
{
    vector<int>a{13,-3,-25,20,-3,-16,-23,
                18,20,-7,12,-5,-22,15,-4,7};
    int i,j,sum;
    Solution solution;
    tie(i,j,sum)=solution.findMaximumSubarray(a,0,a.size()-1);
    cout<<setiosflags(ios::left);
    cout<<setw(4)<<i<<setw(4)<<j<<setw(4)<<sum<<endl;
    return 0;
}
