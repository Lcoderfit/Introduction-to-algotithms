#include<iostream>
#include<algorithm>

using namespace std;

int main()
{
    int v, h;
    while(cin>>v>>h){
        int times;
        cin>>times;
        vector<int>hv;
        vector<int>vv;
        char s;
        int t;
        for(int i=0;i<times;i++){
            cin>>s>>t;
            if(s == 'H'){
                hv.push_back(t);
            } else {
                vv.push_back(t);
            }

            int hmax = 0;
            int vmax = 0;
            hv.push_back(h);
            hv.push_back(0);
            vv.push_back(v);
            vv.push_back(0);

            sort(hv.begin(), hv.end());
            for(int i=1;i<hv.size();i++){
                int k = hv[i] - hv[i - 1];
                if(hmax < k){
                    hmax = k;
                }
            }
            sort(vv.begin(), vv.end());
            for(int i=1;i<vv.size();i++){
                int k = vv[i] - vv[i - 1];
                if(vmax < k){
                    vmax = k;
                }
            }

            printf("%d\n", hmax*vmax);
        }
    }
}

/*
8 6
4
H 4
V 4
V 6
V 2

8


7 6
5
H 4
V 3
V 5
H 2
V 1

4
*/
