#include <stdio.h>

/* count digits, white space, others */
main() {
    int c, i, nwhite, nother;
    /* This is how we declare an array! It's ten spots long */
    int ndigit[10];

    nwhite = nother = 0;
    for (i = 0; i < 10; ++i)
        /* We refer to the individual positions of places
         * within an array by using [i], and we're just
         * iterating over those positions
         */
        ndigit[i] = 0;

    /* This is the real meat! Similar to previous programs */
    while ((c = getchar()) != EOF)
        if (c >= '0' && c <= '9')
            ++ndigit[c-'0'];
        else if (c == ' ' || c == '\n' || c == '\t')
            ++nwhite;
        else
            ++nother;

    printf("digits =");
    for (i = 0; i < 10; ++i)
        printf(" %d", ndigit[i]);
    printf(", white space = %d, other = %d\n", nwhite, nother);
}
