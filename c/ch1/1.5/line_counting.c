/* count lines in input */

/* #include <stdio.h>
*
* main() {
*     int c, nl;
*
*     nl = 0;
*     while ((c = getchar()) != EOF)
*         if (c == '\n')
*             nl++;
*     printf("%d\n", nl);
* }

* '\n' is the representation of the integer value of
* the quoted constant, aptly named a character constant.
* For instance, 'A' is a placeholder for 65, the ASCII
* value of the character A. Depending on any given context,
* 'A' or 65 could be more informative. \n has value 10.
* '\n' is a single character, holding an integer value.
* "\n" is a string constant, containing a single character.
* These differences are meaningful.
* Watch your "'s!
*
* Suppose we want a program that counts blanks and tabs as
* well, and if the input is multiple blanks, they are reduced
* to single blanks!
*/

#include <stdio.h>

main() {
    int c, nl;

    nl=0;
    while ((c=getchar()) != EOF) {
        if (c == '\n' || c == '\t' || c == ' ')
            nl++;
    }
    printf("%d\n", nl);
}
