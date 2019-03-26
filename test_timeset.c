#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <time.h>
#include <errno.h>
#include <sys/time.h>

void showCurrentTime()
{
    struct tm *pTm = NULL;

    time_t _seconds = time(NULL);

    //pTm = gmtime(&_seconds);
    pTm = localtime(&_seconds);
    printf("local time:%s\n", asctime(pTm));

    int isdaylighttime = pTm->tm_isdst;
    printf("isdaylighttime:%d\n", isdaylighttime);

}

int main(int argc, char **argv)
{
    struct tm *curTm = NULL;
    time_t _nowtimeseconds = 0;
    
    showCurrentTime();

    setenv("TZ", argv[1], 1);
    tzset();
    
    system("echo \"export TZ=America/Los_Angeles\" >> /etc/profile");
    system("source /etc/profile");

    char *x = getenv("TZ");    
    printf("TimeZone=%s\n", x);

    //system("hwclock -s");

    printf("After change:\n");
    showCurrentTime();

    return 0;
}


