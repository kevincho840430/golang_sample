# How to use migration
---
~~~
migration 是為了確保開發人員db schema一致
若db有更新(增加欄位、減少欄位)，再下一次simple migrate command就好

migration只會比對model中新增或有更動的table 跟column 並加入或修改
並不會自動刪除你在model中刪除的column
---
若已migrate 過但想刪除欄位 請先刪除`models`中的欄位
再新增刪除指令新增於migration.go中

~~~

**simple migrate command**
```go run ./migration/migration.go```

**add table**
~~~
建立新 `models` 
將以下程式碼新增進 `migration.go`
--db.AutoMigrate(新創的model)--
重新simple migrate一次
~~~

**add column**
~~~
建立新欄位在 `model.go`中
重新simple migrate一次
~~~

**update column**
~~~
更新在`model.go`中的欄位
重新simple migrate一次
~~~

**drop column**
~~~
刪除在`model.go`中的欄位
在`migration.go`中新增判斷欄位是否存在並刪除的code 如下

---
genderExisted := db.Migrator().HasColumn(&models.User{} , "Gender")
if genderExisted{
    db.Migrator().DropColumn(&models.User{},"Gender")
}
---

重新simple migrate一次
~~~
