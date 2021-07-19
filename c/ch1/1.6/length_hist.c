#include <stdio.h>

main() {
/* print a histogram
 *
 *
 *
 * 1 | ****
 * 2 | ***
 * 3 | ******
 * . | ***
 * . |...
 * . |...
 * . |...
 * . |...
 * n | **
 *
 * The number of *'s is the number of words of the length on the axis.
 *
 * So, we need to count the length of each word in our input.
 * Then, after we know how long it is, we need to increment the number that
 * occupies that length's position.
 *
 * Words are separated, canonically, by spaces. We will assume sane language
 * speakers here for this program.
 *
 * So we need to get our input, we need to count the length of a word, and we
 * need to reset that count when we encounter a space. When we reset our
 * counter, we should increment the position.
 *
 * We shouldn't make any assumptions about how big of a word we care about, but
 * we should have a limit.
 */

#define MAX_SIZE 30

    int c, i, wordLength;
    c = wordLength = 0;

    int words[MAX_SIZE];
    for (i = 1; i < MAX_SIZE; ++i)
        words[i] = 0;

    while ((c = getchar()) != EOF) {
        if (c != ' ' && c != '\t' && c != '\n')
            ++wordLength;
        else {
            ++words[wordLength];
            wordLength = 0;
        }
    }

    for (i = 1; i < MAX_SIZE; ++i) {
        printf("%3d", i);
        printf("|");
        while (words[i] > 0) {
            printf("*");
            words[i] = words[i]-1;
        }
        printf("\n");
    }
}
