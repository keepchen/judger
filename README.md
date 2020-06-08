#### Judger

- What's this

> **A pure judge tool based on QingdaoU Judger**

- Environment dependency
  * cmake
  * make
  * g++
  * golang (optional)

- How to use

```shell
# build judger

sudo sh build.sh
```

```shell
# run test program

sudo sh run_test.sh
```

```shell
# run test program with cgo

LD_LIBRARY_PATH=$(pwd)/shared-lib/output go run main.go
```

- Big thanks

[QingdaoU/Judger](https://github.com/QingdaoU/Judger)
