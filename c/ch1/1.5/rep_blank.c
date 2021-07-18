#include <stdio.h>

main() {
    int c;
    int space;

    c = 0;
    space = 0;
    /* We keep track of a second value to determine
     * if the last character we printed was a space or not.
     * If the caracter read is a space, and our previous
     * character was not a space, we print the character
     * and set our 'space' count to one. If our space counter
     * is nonzero, we skip it. If our character read isn't a space,
     * we simply print it and set space count to zero.
     */
    while ((c = getchar()) != EOF) {
        if (c == ' ') {
            if (space == 0) {
            putchar(c);
            space = 1;
            }
        }
        else {
            putchar(c);
            space = 0;
        }
    }
    printf("%d\n", c);
}
