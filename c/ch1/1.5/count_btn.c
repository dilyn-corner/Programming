#include <stdio.h>

main() {

    /* Declare the variables we care about */
    int c, nl, tab, blank;

    /* Set their initial values */
    nl    = 0;
    tab   = 0;
    blank = 0;

    /* A single loop is fine enough */
    while ((c = getchar()) != EOF) {
        if (c == '\n')
            ++nl;
        if (c == '\t')
            ++tab;
        if (c == ' ')
            ++blank;
    }
    /* Headings so we know what each value means */
    printf("newline\t tab\t blank\t\n");
    printf("%d\t %d\t %d\t\n", nl, tab, blank);
}
