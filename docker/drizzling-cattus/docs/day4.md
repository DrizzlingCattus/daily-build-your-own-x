## Namespace API - setns

`int setns(int fd, int nstype);`

> Given a file descriptor referring to a namespace, reassociate the calling thread with that namespace.

현재 프로세스에 namespace를 할당한다.

> setns() disassociates the calling process from one instance of a particular namespace type and reassociates the process with another instance of the same namespace type.

```
fd = open(argv[1], O_RDONLY);   /* Get descriptor for namespace */

setns(fd, 0);                   /* Join that namespace */

execvp(argv[2], &argv[2]);      /* Execute a command in namespace */
```

[1] setns man page : http://man7.org/linux/man-pages/man2/setns.2.html

## Namespace API - unshare

`int unshare(int flags)`

> it creates the new namespaces specified by the CLONE_NEW\* bits in its flags argument and makes the caller a member of the namespaces.

> unshare() allows a process (or thread) to disassociate parts of its execution context that are currently being shared with other processes (or threads).

> The main use of unshare() is to allow a process to control its shared execution context without creating a new process.

커널은 process가 특정 자원들을 공유할 수 있도록 허가하고 있다.
`clone`을 통해서 thread를 생성할때, thread 사이에서 공유할 자원을 선택할 수 있게 한다.
이 시점에 unshare를 통해서 thread들 사이에서 공유하지않을 자원을 선택할 수 있다.

`exec` system call은 같은 행동을 하는 다양한 이름을 가진 형제같은 명령들이 있다. (execvp, ...)
이 형제들에 붙은 이름에는 규칙이 있는데 다음과 같다.

- l, v : argv인자를 넘겨줄 때 사용한다. (l일 경우는 char _로 하나씩 v일 경우에는 char _[]로 배열로 한번에 넘겨준다)

- e : 환경변수를 넘겨줄 때 사용한다. (e는 위에서 v와 같이 char \*[]로 배열로 넘겨준다.)

- p : p가 있는 경우에는 환경변수 PATH를 참조하기 때문에 절대경로를 입력하지 않아도 된다.

[1] unshare man page : http://man7.org/linux/man-pages/man2/unshare.2.html

[2] bbolmin : https://bbolmin.tistory.com/35

[3] unshare가 나타난 문맥을 완벽히 설명해주는 글 : https://www.kernel.org/doc/html/latest/userspace-api/unshare.html

## Q & A

- `unshare`와 `clone`의 차이는?

`unshare`는 호출자의 execution context를 분리 & 변경시키는 반면에 `clone`은 새롭게 프로세스를 생성하고 해당 프로세스의 namespace나 자원을 할당하는 동작을 보인다.

- `exec`류의 system call 함수들은 어떤 동작을 하는가?

호출하는 process의 image를 중단시키고(exec가 호출하는 시점에 코드가 더 진행이 되질않고 종료된다), exec를 통해서 실행하는 process image로 교체한다.

- `execvp`는 `exec`와 같은 동작을 하고 매개변수만 다른 형태로 주입받는다. 여기에서 vp는 뭘 의미하는가?

v는 argv인자를 넘겨준다고 명시

p는 환경변수를 참조하겠다고 명시
