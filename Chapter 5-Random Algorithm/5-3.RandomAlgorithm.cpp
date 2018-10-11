#include<iostream>
#include<vector>

using namespace std;

int main()
{
    vector<int>a{1,2,3,4,5,6,7,8,9};
    for(int i=0;i<a.size();i++){
        cout<<a[i]<<" ";
    }
    cout<<endl;
    srand(time(NULL));
    vector<int>t(a.size(),0);
    int q=0;
    for(int i=0;i<a.size();i++){
        while(t[q]){
            q=rand()%10;
        }
        t[q]=a[i];
        cout<<q<<" ";
    }
    cout<<endl;
    for(int i=0;i<a.size();i++){
        cout<<t[i]<<" ";
    }
    return 0;
}
