#include <stdio.h>

#define MAXLINE 1000 /* maximum input line size */

/*
 * Name it something other than getline() because that function was added
 * to stdio.h in 2010 rip K&R https://c-for-dummies.com/blog/?p=1112
 */
int getLine(char line[], int maxline);
void copy(char to[], char from[]);

/* print longest input line */
int main() {
    int len;                /* current line length        */
    int max;                /* maximum length seen so far */
    char line[MAXLINE];     /* current input line         */
    char longest[MAXLINE];  /* longest line saved here    */

    max = 0;
    while ((len = getLine(line, MAXLINE)) > 0)
        if (len > max) {
            max = len;
            copy(longest, line);
        }
    if (max < 80)
        printf("No line long enough womp womp\n");
    else
        printf("%s", longest);
    return 0;
}

/* getline: read a line into s, return length */
int getLine(char s[], int lim) {
    int c, i;

    for (i = 0; i < lim-1 && (c = getchar()) != EOF && c != '\n'; ++i)
        s[i] = c;
    if ( c == '\n') {
        s[i] = c;
        ++i;
    }
    s[i] = '\0';
    return i;
}

/* copy: copy 'from' into 'to'; assume to is big enough */
void copy(char to[], char from[]) {
    int i;

    i = 0;
    while ((to[i] = from[i]) != '\0')
        ++i;
}

//aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
//aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
//aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
//aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
