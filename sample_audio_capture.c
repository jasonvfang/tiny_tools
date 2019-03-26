#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/select.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#include <unistd.h>
#include <sys/ioctl.h>
#include <errno.h>

#define AUDIO_CAPTURE_FIFO_NAME    "/tmp/audio_capture_fifo"
#define OUTPUT_CAPTURED_PCM_FILE "/tmp/output_captured.pcm"
#define AUDIO_BUFF_SIZE 4096

int main(int argc, char *argv[])
{
    fd_set readfds;
    int keepLooping = 1;
    int audio_cap_fifo_fd = -1, ret = -1, output_fd = -1;

    char audio_buff[AUDIO_BUFF_SIZE];

    if ((audio_cap_fifo_fd = open(AUDIO_CAPTURE_FIFO_NAME, O_RDONLY | O_NONBLOCK)) < 0)
    {
        perror("fifo open failed\n");
        return -1;
    }

	/* remove output file firstly because we write data in appending mode */
	remove(OUTPUT_CAPTURED_PCM_FILE);

    if ((output_fd = open(OUTPUT_CAPTURED_PCM_FILE, O_CREAT | O_RDWR | O_APPEND)) < 0)
    {
        perror("output fd open failed\n");
        return -1;
    }

    while (keepLooping)
    {
        FD_ZERO(&readfds);
        FD_SET(audio_cap_fifo_fd, &readfds);

        ret = select(audio_cap_fifo_fd + 1, &readfds, NULL, NULL, NULL);

        if (ret < 0)
        {
            continue;
        }

        if (FD_ISSET(audio_cap_fifo_fd, &readfds)) /* FD Msg */
        {
            memset(audio_buff, 0, AUDIO_BUFF_SIZE);

            int ret_size = read(audio_cap_fifo_fd, audio_buff, AUDIO_BUFF_SIZE);

            if (ret_size > 0)
            {
                write(output_fd, audio_buff, ret_size);
                fsync(output_fd);
            }
        }
    }

    close(audio_cap_fifo_fd);
    close(output_fd);

    return 0;
}

