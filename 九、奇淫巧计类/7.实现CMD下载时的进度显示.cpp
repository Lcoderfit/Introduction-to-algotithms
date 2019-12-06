#include<iostream>
#include<windows.h>

using namespace std;

int main()
{
    string str = "-\|/";
    int n = 45;
    while(n--){
        for(int i=0;i<str.size(); i++) {
            printf("%c", str[i]);
            Sleep(500);
            printf("\b \b");
        }
    }
}
