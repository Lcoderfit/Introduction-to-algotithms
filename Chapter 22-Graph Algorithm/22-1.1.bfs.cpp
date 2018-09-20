#include<iostream>
#include<vector>
#include<list>
#include<queue>
#include<iomanip>
using namespace std;

bool visit[100]={0};
vector<list<int>>graph;//定义一个邻接表

void bfs(int v)
{
    cout<<setw(4)<<v;
    queue<int>q;
    list<int>::iterator it;
    q.push(v);
    visit[v]=true;
    while(!q.empty()){
        v=q.front();
        q.pop();
        for(it=graph[v].begin();it!=graph[v].end();it++){
            if(!visit[*it]){
                cout<<setw(4)<<*it;
                q.push(*it);
                visit[*it]=true;
            }
        }
    }
    cout<<endl;
}

int main()
{
    int n;
    cin>>n;
    for(int i=0;i<n;i++){
        list<int>l;
        int t;
        while(cin>>t && t!=n){
            l.push_back(t);
        }
        graph.push_back(l);
    }
    bfs(0);

    return 0;
}
