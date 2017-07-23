# manga-list-api-go
This is API server application for [manga-list-2](https://github.com/yoshi10321/manga-list-2)


Requirements
------------
* [Golang](https://golang.org/dl/) 1.7.0+

Usage
-----
You can run this application with command below.
```
go run main.go
```

SetUp DB
--------
This application uses MySQL for Database.
You need to create database names `manga`.
In [manga-list-api-go/migration/sql/](https://github.com/yoshi10321/manga-list-api-go/tree/master/migration/sql), There're sql files to create tables.
