## Cgroup의 이해

> A _cgroup_ associates a set of tasks with a set of parameters for one or more subsystems.

> There are multiple efforts to provide process aggregations in the Linux kernel, mainly for resource-tracking purposes.

process aggregation이 뭘까?

[1] LWN cgroup article : https://lwn.net/Articles/524935/

[2] go cgroup api : https://github.com/containerd/cgroups

[3] linux에서 cgroup의 실체 : https://wiki.archlinux.org/index.php/cgroups

## Q & A

- cgroup에서 말하는 subsystem이란 무엇인가?

subsystem은 resource controller의 역할을 한다. 그말은 즉슨, process 그룹에 자원을 어떻게 할당할 것인가를 정하는 역할을 맡는다는 이야기이다.

- cgroup이 할 수 있는 일은 무엇인가?

일단 제대로 아는건 memory 제한, cpu 사용량 제한을 둘 수 있다는 것이다.

- 특정 프로세스에 cgroup이 할당된 상태를 확인하는 방법

`/proc/PID/cgroup` 파일에 PID에 해당하는 프로세스의 cgroup 할당 상태를 확인할 수 있다.

- cgroup의 실체

cgroup은 디렉토리들로 계층 구조를 만들어서 각 계층에 어떤 프로세스들을 할당할지 관리하는 파일 `tasks`를 두고 어떻게 자원을 할당할지 규칙들을 정의하는 파일들을 둬서 관리한다.

cgroup 파일들은 `/sys/fs/cgroup`에서 확인할 수 있다.

systemd에서도 사용하는 cgroup의 상태를 확인할 수 있다.

`systemd-cgtop` & `systemctl status`

## 신선한 녀석들

systemd도 service마다 cgroup을 할당해준다

RDMA - Remote DMA
