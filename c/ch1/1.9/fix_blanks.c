#include <stdio.h>

#define MAXLINE 1000000000

int getLine(char line[], int maxline);

int main() {
    int c, i, len;
    char line[MAXLINE];

    while ((len = getLine(line, MAXLINE)) > 0)
        if (line[1] == '\n')
            line[1] = '\0';
    for (i = MAXLINE; i > 0; --i)
        else
            printf("%s\n", line);
}

int getLine(char s[], int lim) {
    int c, i;

    for (i = 0; i < lim-1 && (c = getchar()) != EOF && c != '\n'; ++i)
        s[i] = c;
    if (c == '\n') {
        s[i] = c;
        ++i;
    }

    s[i] = '\0';
    return i;
}
