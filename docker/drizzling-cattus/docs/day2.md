## Cgroup과 Namespace

namespace는 6가지 종류가 존재합니다.

각 namespace마다 os가 프로세스에게 제공할 수 있는 자원들을 정의하고있는데, namespace를 사용함으로서 프로세스에게 사용자가 원하는 자원을 외부 환경(호스트 os)에 영향을 받지않도록 `고립`시킬 수 있는 방법을 제공합니다.

> Currently, Linux implements six different types of namespaces. The purpose of each namespace is to wrap a particular global system resource in an abstraction that makes it appear to the processes within the namespace that they have their own isolated instance of the global resource.

cgroup은 control group의 약자로서, namespace에 비해 더 큰 개념을 이야기합니다.

namespace가 **단일 프로세스**에 대해서 이야기하는데 비해서 cgroup은 **프로세스 그룹**에게 자원을 할당하는 방법을 제공합니다.

[1] 컨트롤 그룹이란? : https://access.redhat.com/documentation/ko-kr/red_hat_enterprise_linux/6/html/resource_management_guide/ch01

## UTS namespace

UTS namespace는 system identifier인 nodename과 domainname을 고립시킬 수 있는 namespace입니다.

즉 다시말해 UTS namespace를 통해서 container는 다음과 같은 값을 호스트 환경에 관계없이 할당할 수 있습니다.

- hostname
- NIS(YP) domain

이는 hostname에 의존하고있는 스크립트나 모듈에 영향을 최소화시킬 수 있는 방법을 제공합니다.

## Namespace API

namespace를 위한 system call API가 있습니다.

- clone
- setns
- unshare

이 중에서 `clone`은 프로세스를 생성하도록 os에게 요청할 수 있습니다.

그리고 생성하면서 동시에 namespace를 할당할 수 있는 방법을 제공하고 있습니다.

그렇다면 각 프로세스가 namespace를 어떤 형태로 관리하고 있을까요?

```
[root@goldilocks-web workspace]# ls -al /proc/10/ns/
total 0
dr-x--x--x 2 root root 0 Jan  4 13:44 .
dr-xr-xr-x 9 root root 0 Dec 27 16:19 ..
lrwxrwxrwx 1 root root 0 Jan  4 13:45 ipc -> ipc:[4026531839]
lrwxrwxrwx 1 root root 0 Jan  4 13:45 mnt -> mnt:[4026531840]
lrwxrwxrwx 1 root root 0 Jan  4 13:45 net -> net:[4026531992]
lrwxrwxrwx 1 root root 0 Jan  4 13:45 pid -> pid:[4026531836]
lrwxrwxrwx 1 root root 0 Jan  4 13:45 user -> user:[4026531837]
lrwxrwxrwx 1 root root 0 Jan  4 13:45 uts -> uts:[4026531838]
```

리눅스에서는 모든게 파일입니다.

마찬가지로 namespace도 파일로 관리되고 있고, 프로세스에 할당된 namespace들은 /proc/PID/ns에서 확인할 수 있습니다.

위에서 보듯 namespace는 symbolic link되어있고 이는 약간 특별한 형태를 사용합니다.

process에 할당되어있는 namespace를 확인하는 방법을 알았으니, 이제 실험을 통해서 namespace가 제대로 할당되는지 확인해봅시다.

간단하게 `clone` system call을 통해서 child process를 독립적인 UTS namespace를 할당하는 프로그램을 만들었습니다.

핵심은 다음 코드입니다.

```
/* Create a child that has its own UTS namespace;
*        the child commences execution in childFunc() */

child_pid = clone(
    /* child process callback */
    childFunc,

    /* stack 의 top을 가리키는 포인터 주소 */
¦   child_stack + STACK_SIZE,

    /* UTS namespace flag */
¦   CLONE_NEWUTS | SIGCHLD,

    /* child's UTS hostname */
    argv[1]
);
```

실행하면 다음과 같은 결과를 얻을 수 있습니다.

```
[root@goldilocks-web namespace]# ./uts helo
PID of child created by clone() is 11120
uts.nodename in child:  helo
uts.nodename in parent: goldilocks-web
^Z
[1]+  Stopped                 ./uts helo

[root@goldilocks-web namespace]# jobs -l
[1]+ 11119 Stopped                 ./uts helo
```

parent process의 pid는 11119 이고 child process의 pid는 11120 임을 알 수 있습니다.

이를 토대로 각자가 가지고 있는 UTS namespace를 확인해봅시다.

```
[root@goldilocks-web namespace]# readlink /proc/11119/ns/uts
uts:[4026531838]
[root@goldilocks-web namespace]# readlink /proc/11120/ns/uts
uts:[4026532182]
```

child와 parent의 UTS id가 다른것을 확인할 수 있습니다.

또다른 관계없는 process의 UTS는 어떨까요?

```
[root@goldilocks-web namespace]# readlink /proc/10/ns/uts
uts:[4026531838]
```

당연하게도 host os의 namespace를 그대로 같이 사용하기 때문에 parent의 UTS namespace와 같음을 확인할 수 있습니다.

[1] 리눅스 뉴스 namespace : https://lwn.net/Articles/531114/

[2] namespace와 cgroup overview : https://itnext.io/chroot-cgroups-and-namespaces-an-overview-37124d995e3d

[3] 지금 읽고 있는 부분 https://lwn.net/Articles/531381/
