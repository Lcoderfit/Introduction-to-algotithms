#include<iostream>
#include<queue>

using namespace std;

typedef struct BinaryTreeNode {
    int data;
    BinaryTreeNode* leftPtr;
    BinaryTreeNode* rightPtr;
}*pTreeNode;

int CreateBiTree()

void FloorPrint_QUEUE(pTreeNode &Tree)  //≤„–Ú±È¿˙
{
    if (!Tree) {
        return;
    }

    queue<pTreeNode> q;
    q.push(Tree);


    while(!q.empty()) {
        cout<<q.front()->data<<"->";
        if (q.front()->leftPtr!=NULL){
            q.push(q.front()->leftPtr);
        }
        if (q.front()->rightPtr!=NULL) {
            q.push(q.front()->rightPtr);
        }
        q.pop();
    }
}


int main()
{

}
