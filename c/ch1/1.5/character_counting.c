/* #include <stdio.h>
 *
 * Count characters; 1st version
 *
 * main() {
 *     long nc;
 *
 *     nc = 0;
 *     while (getchar() != EOF)
 *         ++nc;
 *     printf("%1d\n", nc);
 * }
 * ++nc increments the variable nc
 * Alternatives include nc++, --nc, nc--
 */

#include <stdio.h>

/* Count characters; 2nd version */

main() {
    double nc;
    for (nc = 0; getchar() != EOF; nc++)
        ;
    printf("%.0f\n", nc);
}

/* Doubles are like ints that can be larger values.
 * They're double precision floats.
 * Notice the single ; in the body of our loop.
 * This is because all of the work of the loop is
 * done in the condition testing itself! But, C is
 * like people, and requires some type of
 * nonempty body. So we'll give it one.
 * This floating semicolon is called a null statement.
 */
