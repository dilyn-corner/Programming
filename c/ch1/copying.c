/* A program that copies its input
 * to an output one character at a time
 */

#include <stdio.h>

/* copy input to output; 1st version */
main() {

    /* If 'c' is just holding characters, why is it an int??
     * Well, because 'c' gets inputs, it has to be big enough
     * to hold an input. But what happens when input ends?
     * What about when our input finishes... How big is the end?
     * It's of size EOF, bigger than a char. So we need 'c'
     * to be able to hold a size of any input, plus the size
     * of the end of our input. Thus, c == int.
     */
    int c;

    c = getchar();
    while (c != EOF) {
        putchar(c);
        c = getchar();
    }
}

/* copy input to output, 2nd version
 *
 * main() {
 *     int c;
 *
 *     while ((c = getchar()) != EOF)
 *     putchar(c);
 * }
 *
 * The main difference here is that the value of c is tested
 * against the value of EOF. If c is not EOF, then it is printed.
 * It's more readable, more compact, simpler. The point of C.
 *
 * This points to an important idea: precedence. Without (...),
 * things are interpretted in a predefined way (like PEMDAS). In C,
 * = has higher value than !=. So:
 * c = getchar() != EOF <==> c = (getchar() != EOF)
 * and c will always be 0 or 1, depending on if the character received
 * is equal to EOF, or note equal to EOF.
 */
