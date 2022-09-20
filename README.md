# nuget-downloader

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
nuget-downloader -n PackageName -v PackageVersion -r Nexus-repository
````
-n - указывает какой пакет необходимо найти

-v - указывает какой версии пакет должен соответствовать 

-r (необязательный) - указывает репозиторий в Action-Nexus ( по умолчанию nuget-freeze )