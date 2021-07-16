#include <stdio.h>

/* print Fahrenheit-Celsius table
    for fahr = 0, 20, ..., 300    */

main() {
    /* we'll declare some variables.
     * Those variables store integers (int).
     * We declare fahr, celsius as float(ing point)
     * numbers so that we can display them as decimals.
     * This simply improves accuracy.
     */
    float fahr, celsius;
    int lower, upper, step;
    /* Other notable types: char, short, long, double */

    /* We will set some of those variables
     * initial values.
     */
    lower = 0;   /* lower limit of temperatures */
    upper = 300; /* upper limit                 */
    step = 20;   /* step size                   */

    /* Now we can set one variable to equal another */
    fahr=lower;
    /* This is a while loop; it's a form of control flow,
     * which is very useful in a program. It does
     * what it says it does!
     */

    while (fahr <= upper) {
        /* Do some operation on fahr, set celcsius'
         * value to the result of that operation.
         * Note that this works because fahr,celsius are
         * ints, not chars etc.
         * Because our numbers are floats and not ints,
         * We should probably use floating point numbers
         * in our operations on them...
         *
         * if fahr, celsius were ints:
         * celsius = 5*(fahr-32)/9;
         *
         * Integer division truncates in C,
         * so we can't use 5/9 in the previous instance
         */
        celsius = (5.0/9.0)*(fahr-32.0);
        /* We'll print the results in a table.
         * %d specifies an integer argument.
         * Because we're using floats, we'll use %f.
         * %d...%d causes fahr, celsius
         * to be printed, separated by a tab
         * character, and then a new line.
         * Formatting (of values), values.
         * printf("%d\t%d\n", fahr, celsius);
         *
         * Use %xY to get more granular with the
         * formatting; x=3, x=6 will print the first
         * item (fahr) 3 characters wide, and the
         * second item (celsius) 6 characters wide
         * Specifically, 3 characters wide with no
         * decimal, and 6 characters wide with one
         */
        printf("%3.0f\t%6.1f\n", fahr, celsius);
        /* Increment the variable controlling our loop.
         * Otherwise, it'll never end... And just print
         * 0    -17
         */
        fahr = fahr + step;
    }
}
