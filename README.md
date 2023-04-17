# package-downloader

## Установка
1. Скачать репозиторий

2. Выполнить поманду
```shell
go instal .
```
3. Экспортировать бинарник в переменную PATH
```shell
export PATH=$PATH:/Users/user/go/bin
```
## Использование

````shell
package-downloader -n PackageName -v PackageVersion -t (npm/nuget)
````
-n - указывает какой пакет необходимо найти

-v - указывает какой версии пакет должен соответствовать 

-t - указать тип пакета


Примеры без компиляции
```shell
go run main.go -n @babel/core -v 7.0.0 -t npm
go run main.go -n api -v 5.0.8 -t npm

go run main.go -n HotChocolate -v 13.0.5 -t nuget
 go run main.go -n System.Diagnostics.PerformanceCounter -v 8.0.0-preview.3.23174.8 -t nuget
```
