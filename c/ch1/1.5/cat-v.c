/* I named it cat-v to make you mad */
#include <stdio.h>

main() {
    int c;

    c = 0;
    while ((c = getchar()) != EOF) {
        if (c == '\t') {
            putchar('\\');
            putchar('t');
        }
        if (c == '\b') {
            putchar('\\');
            putchar('b');
        }
        if (c == '\\') {
            putchar('\\');
            putchar('\\');
        }
        else putchar(c);
    }
}
