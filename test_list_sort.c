#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "wm_list.h"

typedef struct _pda_control_context {
//    sem_t g_pda_msg_queue;
    struct list_head msg_pda_head;
    pthread_mutex_t mutex_pda_msg;
} pda_control_context;

typedef struct _pda_msg_cont {
    struct list_head entry;
    char msg[32];
    long long msgTimeStamp;
} pda_msg_cont;

struct list_head MsgListHeader = {NULL, NULL};

void debug_print_list(struct list_head *msgListHeader)
{
    pda_msg_cont *tmpEntry = NULL;
    struct list_head *pos = NULL;
    
    list_for_each(pos, msgListHeader) 
    {
        tmpEntry = list_entry(pos, pda_msg_cont, entry);

        printf("[D] MsgName [%s], timestmp [%lld] \n", tmpEntry->msg, tmpEntry->msgTimeStamp);
    }
}



void msg_sort_by_timestamp(struct list_head *msgListHeader)
{
    if (!msgListHeader)
        return;

    printf("Before Sort:\n");
    debug_print_list(msgListHeader);
    
    static struct list_head sortedMsgHeader = {NULL, NULL};
    pda_msg_cont *currEntry = NULL, *NextEntry = NULL, *LastEntry = NULL;

    /* sort list if necessary */
    if (list_number(msgListHeader) >= 2)
    {
        INIT_LIST_HEAD(&sortedMsgHeader);

        list_for_each_entry_safe(currEntry, NextEntry, msgListHeader, entry) 
        {
            list_del(&currEntry->entry);

            if (currEntry == NULL)
                break;
            
            if (list_empty(&sortedMsgHeader))
            {
                list_add_tail(&currEntry->entry, &sortedMsgHeader);
                continue;
            }
            
            pda_msg_cont *currNewEntry = NULL, *NextNewEntry = NULL, *rightPos = NULL;
            int i = 0;
            int newListEntryNum = list_number(&sortedMsgHeader);

            list_for_each_entry_safe(currNewEntry, NextNewEntry, &sortedMsgHeader, entry) 
            {
                if (currEntry->msgTimeStamp > currNewEntry->msgTimeStamp)
                {
                    i ++;
                    continue;
                }
                else
                {
                    rightPos = currNewEntry;
                    break;
                }
            }

            if (i == newListEntryNum)
            {
                list_add_tail(&currEntry->entry, &sortedMsgHeader);
            }
            else
            {
                list_add_tail(&currEntry->entry, &rightPos->entry);
            }
        }    
    }

    printf("After Sort:\n");
    debug_print_list(&sortedMsgHeader);

    return; 
}


void msg_sort_by_timestamp2(struct list_head **msgListHeader)
{
    if (!msgListHeader)
        return;

    //printf("Before Sort:\n");
    //debug_print_list(*msgListHeader);
    
    static struct list_head sortedMsgHeader = {NULL, NULL};
    pda_msg_cont *currEntry = NULL, *NextEntry = NULL, *LastEntry = NULL;

    /* sort list if necessary */
    if (list_number(*msgListHeader) >= 2)
    {
        INIT_LIST_HEAD(&sortedMsgHeader);

        list_for_each_entry_safe(currEntry, NextEntry, *msgListHeader, entry) 
        {
            list_del(&currEntry->entry);

            
            if (list_empty(&sortedMsgHeader))
            {
                list_add_tail(&currEntry->entry, &sortedMsgHeader);
                continue;
            }
            
            pda_msg_cont *currNewEntry = NULL, *NextNewEntry = NULL, *rightPos = NULL, *TailEntry = NULL;
            int i = 0;
            int newListEntryNum = list_number(&sortedMsgHeader);

            list_for_each_entry_safe(currNewEntry, NextNewEntry, &sortedMsgHeader, entry) 
            {
                if (currEntry->msgTimeStamp > currNewEntry->msgTimeStamp)
                {                   
                    i ++;
                    continue;
                }
                else
                {
                    rightPos = currNewEntry;
                    break;
                }
            }
            
            if (i == newListEntryNum)
            {
                list_add_tail(&currEntry->entry, &sortedMsgHeader);
            }
            else
            {
                list_add_tail(&currEntry->entry, &rightPos->entry);
            }
        }    
    }

    //printf("After Sort:\n");
    //debug_print_list(&sortedMsgHeader);

    *msgListHeader = &sortedMsgHeader;
    
    return; 
}

int main()
{    
    struct list_head *pos = NULL;
    pda_msg_cont *tmpEntry = NULL, *NextEntry = NULL;

    pda_msg_cont msg1 = {{NULL,NULL}, "msg 1", 1};
    pda_msg_cont msg2 = {{NULL,NULL}, "msg 2", 2};
    pda_msg_cont msg3 = {{NULL,NULL}, "msg 3", 3};
    pda_msg_cont msg4 = {{NULL,NULL}, "msg 4", 4};
    pda_msg_cont msg5 = {{NULL,NULL}, "msg 5", 5};
    pda_msg_cont msg6 = {{NULL,NULL}, "msg 6", 6};

    pda_msg_cont msg11 = {{NULL,NULL}, "msg 11", 11};
    pda_msg_cont msg22 = {{NULL,NULL}, "msg 22", 22};
    pda_msg_cont msg33 = {{NULL,NULL}, "msg 33", 33};
    pda_msg_cont msg44 = {{NULL,NULL}, "msg 44", 44};
    pda_msg_cont msg55 = {{NULL,NULL}, "msg 55", 55};
    pda_msg_cont msg66 = {{NULL,NULL}, "msg 66", 66};


    INIT_LIST_HEAD(&MsgListHeader);

    list_add_tail(&msg5.entry, &MsgListHeader);
    list_add_tail(&msg3.entry, &MsgListHeader);
    list_add_tail(&msg1.entry, &MsgListHeader);
    list_add_tail(&msg4.entry, &MsgListHeader);
    list_add_tail(&msg2.entry, &MsgListHeader);
    list_add_tail(&msg6.entry, &MsgListHeader);

    list_add_tail(&msg55.entry, &MsgListHeader);
    list_add_tail(&msg33.entry, &MsgListHeader);
    list_add_tail(&msg11.entry, &MsgListHeader);
    list_add_tail(&msg44.entry, &MsgListHeader);
    list_add_tail(&msg66.entry, &MsgListHeader);
    list_add_tail(&msg22.entry, &MsgListHeader);

    printf("List entry num is [%d]\n", list_number(&MsgListHeader));
    struct list_head *MsgListHeaderTmp = &MsgListHeader;
    msg_sort_by_timestamp2(&MsgListHeaderTmp);
    
    return 0;
}

