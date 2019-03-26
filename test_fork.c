#include <stdio.h>
#include <unistd.h>

int main()
{
	int i = 0;
	
	for(i = 0; i < 2 ; i ++)
	{
		fork();
		printf("%d,-\n", i);
	}
	
	wait(NULL);
	wait(NULL);
	return 0;
}
