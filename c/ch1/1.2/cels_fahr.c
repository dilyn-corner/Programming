#include <stdio.h>

main() {
    float fahr, celsius;
    int lower, upper, step;

    /* We could change these conditions if we wanted a
     * literal reversal of our other table. But I'm lazy */
    lower = 0;
    upper = 300;
    step  = 20;

    celsius = lower;
    printf("Celsius\t Fahrenheit\n");
    while (celsius <= upper) {
        fahr = ((9.0/5.0) * celsius) + 32;
        printf("%3.0f\t%6.1f\n", celsius, fahr);
        celsius = celsius + step;
    }
}
