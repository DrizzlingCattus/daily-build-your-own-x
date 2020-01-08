
#define _GNU_SOURCE
// fcntl.h - file control options
#include <fcntl.h>
// for setns
#include <sched.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>

#define errExit(msg)    do { perror(msg); exit(EXIT_FAILURE); \
                               } while (0)

int main(int argc, char *argv[]) {
        if (argc < 3) {
                fprintf(stderr, "arg is %s", argv[0]);
                exit(EXIT_FAILURE);
        }

        int fd = open(argv[1], O_RDONLY);
        if (fd == -1) {
                errExit("open: there is no file");
        }
        const int ANY_NS = 0;
        int is_ns_err = setns(fd, ANY_NS);
        if (is_ns_err == -1) {
                errExit("setns: fail to open");
        }
        execvp(argv[2], &argv[2]);
        errExit("execvp");
}

