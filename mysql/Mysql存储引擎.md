## Mysql的存储引擎种类和区别      
    
### MyIsam 
    1. 不支持数据库事务。提供高效存储和检索，以及全文搜索能力。MyIsam是默认的存储引擎。
    2. 只支持表级锁，用户在操作myisam表时，select，update，delete，insert语句都还自动给表加锁，如果加锁以后的表满足insert并发的情况下，可以在表的尾部插入新的数据   
    3. 保存有表的总行数，如果select count(*) from table,直接取出
    4. 如果执行大量select，myisam是最好的选择

### Innodb  
    1. 支持数据库事务，外部键等高级数据库功能。具有事务，回滚和崩溃修复能力
    2. 支持事务和行级锁。但是innodb的行锁，只在where主键的时候是有效的，非主键的where都会锁全表
    3. 没有保存表的总行数，如果select count(*) from table；就会遍历整个表。但是加了where条件后，myisam和innodb的处理方式是一样的
    4. 如果执行大量insert和update，处于性能方面的考虑，应该使用innodb。delete从性能上innodb更优，但是delete的时候是一行一行的删除的，如果要清除大量数据的表，最好使用truncate table命令
