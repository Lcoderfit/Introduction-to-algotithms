#include<iostream>
#include<vector>
#include<list>
#include<iomanip>
using namespace std;

bool visit[100]={0};
vector<list<int>>graph;

void dfs(int v)
{
    list<int>::iterator it;
    cout<<setw(4)<<v;
    visit[v]=true;
    for(it=graph[v].begin();it!=graph[v].end();it++){
        if(!visit[*it]){
            dfs(*it);
        }
    }
}

int main()
{
    int n;
    cin>>n;
    int t;
    for(int i=0;i<n;i++){
        list<int>l;
        while(cin>>t && t!=n){
            l.push_back(t);
        }
        graph.push_back(l);
    }
    dfs(0);
    return 0;
}
