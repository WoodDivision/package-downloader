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