# microservices-demo TarsGo部署

建议提前浏览 [Tars Go快速指南](https://github.com/TarsCloud/TarsGo/blob/master/docs/tars_go_quickstart.md#tars-go%E5%BF%AB%E9%80%9F%E6%8C%87%E5%8D%97)

##### 目录结构
MicroserviceDemo下含有三个Tarsgo微服务：**CatalogueGo**、**PaymentGo**、**UserGo**
以**CatalogueGo**为例：
```
  '    |-- tarsgo踩坑.md',
  '    |-- MicroserviceDemo',
  '        |-- CatalogueGo',                
  '        |   |-- CatalogueGo',                //编译生成的可执行
  '        |   |-- CatalogueGo.conf',           //启动运行所需的配置文件
  '        |   |-- CatalogueGo.go',             //程序入口文件
  '        |   |-- CatalogueGo.tgz',            //生成的targo所需的部署文件
  '        |   |-- DoCatalogue.tars',      
  '        |   |-- makefile',
  '        |   |-- start.sh',                   //启动运行脚本
  '        |   |-- client',                     //测试客户端
  '        |   |   |-- client.go',
  '        |   |-- data',                       //mysql数据可以直接导入
  '        |   |   |-- dump.sql',
  '        |   |-- debugtool',
  '        |   |   |-- dumpstack.go',
  '        |   |-- imp',                        //接口具体实现,mysql init()
  '        |   |   |-- DoCatalogueImp.go',
  '        |   |-- MicroserviceDemo',
  '        |   |   |-- DoCatalogue_IF.go',
  '        |   |-- vendor',
  '        |       |-- vendor.json',
  '        |       |-- MicroserviceDemo',
  '        |           |-- DoCatalogue_IF.go',
  '        |-- PaymentGo',
  '        |   |-- DoPayment.tars',
  '        |   |-- makefile',
  '        |   |-- PaymentGo',
  '        |   |-- PaymentGo.conf',
  '        |   |-- PaymentGo.go',
  '        |   |-- PaymentGo.tgz',
  '        |   |-- start.sh',
  '        |   |-- client',
  '        |   |   |-- client.go',
  '        |   |-- debugtool',
  '        |   |   |-- dumpstack.go',
  '        |   |-- imp',
  '        |   |   |-- DoPaymentImp.go',
  '        |   |-- MicroserviceDemo',
  '        |   |   |-- DoPayment_IF.go',
  '        |   |-- vendor',
  '        |       |-- vendor.json',
  '        |       |-- MicroserviceDemo',
  '        |           |-- DoPayment_IF.go',
  '        |-- UserGo',
  '            |-- DoUser.tars',
  '            |-- makefile',
  '            |-- start.sh',
  '            |-- UserGo',
  '            |-- UserGo.conf',
  '            |-- UserGo.go',
  '            |-- UserGo.tgz',
  '            |-- client',
  '            |   |-- client.go',
  '            |-- db',
  '            |   |-- db.go',                             //mongo init()
  '            |-- debugtool',
  '            |   |-- dumpstack.go',
  '            |-- docker',
  '            |   |-- user-db',
  '            |       |-- Dockerfile',
  '            |       |-- scripts',
  '            |           |-- accounts-create.js',
  '            |           |-- address-insert.js',
  '            |           |-- card-insert.js',
  '            |           |-- customer-insert.js',
  '            |           |-- mongo_create_insert.sh',
  '            |-- imp',
  '            |   |-- DoUserImp.go',
  '            |-- MicroserviceDemo',
  '            |   |-- DoUser_IF.go',
  '            |-- vendor',
  '                |-- vendor.json',
  '                |-- MicroserviceDemo',
  '                    |-- DoUser_IF.go',
 ```

**注意：CatalogueGo需要mysql，UserGo需要mongo可以运行时修改数据库地址**