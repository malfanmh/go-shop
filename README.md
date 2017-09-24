# go-Shop
----------
### Requerment Summary:

 * GO v1.8 or later varsion
 * [Iris](iris-go.com) framework
 * Vendoring Using [Govendor](github.com/kardianos/govendor)


### Instalation :
* Read Go installation steps from [hire](golang.org/doc/install).
* After Go installed and the environment variables are set, install govendor

```sh
go get -u github.com/kardianos/govendor
```


### Project
* Clone the project
```sh
git clone https://github.com/malfanmh/go-shop.git
```
* Restore vendor source , (full documentation [hire](github.com/kardianos/govendor/blob/master/doc/dev-guide.md))
```sh
govendor sync
```
* Setup Environment Variable to config file ([./files/etc/go-shop/config.yml](github.com/malfanmh/go-shop/blob/master/files/etc/go-shop/config.yml)) or move file to any directory as you will. for example :
```sh
GOSHOP_CONNFIG=/etc/go-ship/config.yml
```
* Type below command To run application
```sh
go build && ./go-shop
```
or you can direct output binary , for example :
```sh
go build -o /usr/local/bin/go-shop
```
----
### api documentation

* [Postman Collection](https://www.getpostman.com/collections/26aaf74af977dfb4c580)
* [swagger](swagger.io) (comming soon)

