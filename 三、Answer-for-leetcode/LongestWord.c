#include<stdio.h>

void main()
{
    char a[20];
    char maxStr[20];
    memset(a, '\0', sizeof(a));
    int count = 0;
    int maxLen = 0;
    char c;
    while((c=getchar()) != '\n')
    {
        if (c == ' ' && a[0] == '\0') {
            continue;
        }
        if (c == ' ' && a[0] != '\0') {
            printf("%-20s%d\n", a, count);
            if (count > maxLen) {
                maxLen = count;
                strcpy(maxStr, a);
            }
            count = 0;
            memset(a, '\0', sizeof(a));
        } else {
            a[count++] = c;
        }
    }
    if(a[0] != '\0') {
        printf("%-20s%d\n", a, count);
    }
    if(count > maxLen) {
        strcpy(maxStr, a);
    }
    printf("the longest word is: %s", maxStr);
}
