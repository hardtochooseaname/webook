在很多项目里，数据库操作的目录叫 dao，是因为：

DAO 全称是 Data Access Object，意思是“数据访问对象”。

它的职责很简单明确：
👉 专门负责和数据库打交道，把数据读出来、写进去。
👉 它不关心业务逻辑，只负责“怎么从数据库取数据”、“怎么存数据”。

DAO对应的就是数据库操作，而repository是总的数据存储抽象，里面可能有cache、ElasticSearch或其他的存储实现方式