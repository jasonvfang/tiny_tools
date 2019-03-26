#include <stdio.h>
#include <string.h>
#include <stdlib.h>





/*
** Escape special characters such as ", \ etc.
** outString buffer size must large enough
*/
int escape_special_characters(char *inString, int inStringLen, char *outString, int outStringBufferSize)
{
    int rc = 0;
    char *ptr, *newPtr, *data ;
    char espTagList[] = {'\\',  '\"'};/* special character list */

    if (!inString || inStringLen <= 0)
        return -1;

    data = inString;
    newPtr = NULL;

    int i = 0; 

    for(; i < sizeof(espTagList); i ++)
    {
        if (newPtr != NULL)
        {
            /* reset ptr */
            data = newPtr;
        }
        
        while (*data != '\0')
        {
            ptr = strchr(data, espTagList[i]);

            if (ptr != NULL)
            {
                if (newPtr == NULL)
                {
                    newPtr = calloc(1, (inStringLen + 2));
                    if (newPtr)
                    {
                        int escape_len = ptr - data;    
                        strncpy(newPtr, data, ptr - data);

                        *(newPtr + escape_len) = '\\';/* add escape char at the right pos */
                        strncat(newPtr, ptr, (inStringLen + 2));

                        data = newPtr + escape_len + 2;  
                    }
                    else
                    {
                        rc = -1;
                        goto end;
                    }
                }
                else
                {
                    int tmpSize = strlen(newPtr) + 2; /* 2 bytes allocated for reserved space */   
                    char *tmpPtr = calloc(1, strlen(newPtr) + 2);

                    if (tmpPtr)
                    {   
                        int escape_len = ptr - data; 
                        int remain_len = ptr - newPtr; 

                        strncpy(tmpPtr, newPtr, remain_len);  
                        
                        *(tmpPtr + remain_len) = '\\';/* add escape char at the right pos */

                        strncat(tmpPtr, ptr, strlen(ptr));

                        data = tmpPtr + remain_len + 2 ;

                        free(newPtr);
                        newPtr = tmpPtr;
                    }
                    else
                    {
                        rc = -1;
                        goto end;
                    }
                }
            }
            else
            {
                break;
            }
        }
    }
       
end:
    
    memset(outString, 0, outStringBufferSize);

    if (rc == 0 && newPtr)
    {
        strncpy(outString, newPtr, outStringBufferSize - 1);
    }
    else
    {
        //not escaped, just return the original string
        strncpy(outString, inString, outStringBufferSize - 1);
    }

    if (newPtr)
       free(newPtr);

    return rc;
}




int main(void)
{
        char tags[][1024] =
        {
                "asdfasd\\\\\\\\\\\\fasdf\"",
        };
        
        int i = 0; 
        for(;i < sizeof(tags) / sizeof(tags[i]); i ++)
        {
                char out[1024] = {0};
                
                printf("In =[ %s ] \n", tags[i]);  

                escape_special_characters(tags[i], strlen(tags[i]), out, sizeof(out));

                printf("Out=[ %s ]\n", out);  
        }

        return 0;
}