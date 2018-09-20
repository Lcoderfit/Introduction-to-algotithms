#include<iostream>
#include<vector>
#include<list>
#include<iomanip>
#include<queue>
using namespace std;

vector<int>visit(100,0);
vector<list<int>>graph;

void bfs(int v)
{
    list<int>::iterator it;
    cout<<setw(4)<<v;
    visit[v]=true;
    queue<int>q;
    q.push(v);
    while(!q.empty()){
        v=q.front();
        q.pop();
        for(it=graph[v].begin();it!=graph[v].end();it++){
            if(!visit[*it]){
                q.push(*it);
                cout<<setw(4)<<*it;
                visit[*it]=true;
            }
        }
    }
    cout<<endl;
}

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
    cout<<"please input:"<<endl;
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
    cout<<"bsf algorithm:"<<endl;
    bfs(0);

    for(int i=0;i<visit.size();i++)
        visit[i]=0;
    cout<<"dfs algorithm"<<endl;
    dfs(0);

    return 0;
}
