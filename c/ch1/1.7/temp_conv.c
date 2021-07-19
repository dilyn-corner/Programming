#include <stdio.h>

/* Default temperatures to use for table */
#define LOWER 0
#define UPPER 300
#define STEP  20

/* Function prototype */
float conversion(float);

int main() {
    float fahr;

    fahr=LOWER;
    printf("Fahr\t Celsius\n");
    while (fahr <= UPPER) {
        printf("%3.0f\t%6.1f\n", fahr, conversion(fahr));
        fahr = fahr + STEP;
    }

    return 0;
}

float conversion(float fahr) {
    float celsius;

    celsius = (5.0/9.0)*(fahr-32.0);

    return celsius;
}
