/* Exercise 1.1
 * Experiment compiling these different broken versions
 * of our magical introductory function!

// ---

main() {
    printf("hello, world\n");
}

// ---

#include <stdio.h>

main() {
    printf("hello, world\n")
}

// ---

#include <stdio.h>

main() {
    printf("hello, world
    ");
}

// ---

#include <stdio.h>

main() printf("hello, world\n");
*/
