#include <stdio.h>

main() {
    float fahr, celsius;
    int lower, upper, step;
    lower = 0;
    upper = 300;
    step = 20;

    fahr=lower;

    /* We place this outside the loop
     * so that it only prints once!  */
    printf("Fahr\t Celcius\n");

    while (fahr <= upper) {
        celsius = (5.0/9.0)*(fahr-32.0);
        printf("%3.0f\t%6.1f\n", fahr, celsius);
        fahr = fahr + step;
    }
}
