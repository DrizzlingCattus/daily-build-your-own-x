우리의 앱은 컨테이너를 관리하는 서버와, 사용자의 명령을 받는 클라이언트로 구성된다.

### Feature list

1. 쉘에서 특정한 컨테이너 환경을 만들 수 있다 (기본은 백그라운드)

```
$ myapp create -port -host ...
```

2. 현재 만들어진 컨테이너 리스트를 확인할 수 있다.

```
$ myapp list
> container1. ...
> container2. ...
```

3. 만들어진 컨테이너를 forground로 만들 수 있다.

```
myapp attach container1
```

