// main.c
#include <stdio.h>

int main() {
	int sum = 0;
	for (int i = 1; i <= 10; i++){
		sum += i;
	}
	printf("Sum number from 1 to 10 equal: %d\n", sum);
	return 0;
}
