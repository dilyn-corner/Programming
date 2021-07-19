#include <stdio.h>

main() {
    /* We're going to do basically the same thing here as we did for
     * length_hist.c, except we're going to do something very cute.
     *
     * Each spot in the array stands for a letter, a-z (1-26).
     * Because ANSI characters like A are just placeholders for
     * values like 65, we can just ask if c happens to be in
     * some ANSI range for English letters. We'll treat A and a
     * as the same because... Well, why distinguish between the two?
     *
     * We also don't count special characters, though it won't be hard
     * to extend this to include them.
     */
    int c, i;
    c = 0;

    int charfreq[26];
    for (i = 0; i < 26; ++i)
        charfreq[i] = 0;

    while ((c = getchar()) != EOF) {
        if (c != ' ' && c != '\t' && c != '\n') {
            if (65 <= c && c <= 90)
                ++charfreq[c-65];
            else if (97 <= c && c <= 122)
                ++charfreq[c-97];
        }
    }

    for (i = 0; i < 26; ++i) {
        putchar(i+97);
        printf("|");
        while (charfreq[i] > 0) {
            printf("*");
            charfreq[i] = --charfreq[i];
        }
        printf("\n");
    }
}
