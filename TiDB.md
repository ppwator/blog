## 1. TiDB
在使用MySQL数据库时，为了方便，我们都习惯使用自增ID来作为表的主键。因此，将数据从MySQL迁移到TiDB之后，原来的表结构都保持不变，仍然是以自增ID作为表的主键。这样就造成了批量导入数据时出现TiDB写入热点的问题，导致Region分裂不断进行，消耗大量资源。

对此，在进行TiDB优化时，我们从表结构入手，对以自增ID作为主键的表进行重建，删除自增ID，使用TiDB隐式的_tidb_rowid列作为主键，将

create table t (a int primary key auto_increment, b int)；
改为：

create table t (a int, b int)SHARD_ROW_ID_BITS=4 PRE_SPLIT_REGIONS=2


通过设置SHARD_ROW_ID_BITS，将RowID打散写入多个不同的Region，从而缓解写入热点问题。

此处需要注意，SHARD_ROW_ID_BITS值决定分片数量：

SHARD_ROW_ID_BITS = 0 表示 1 个分片
SHARD_ROW_ID_BITS = 4 表示 16 个分片
SHARD_ROW_ID_BITS = 6 表示 64 个分片
SHARD_ROW_ID_BITS值设置的过大会造成RPC请求数放大，增加CPU和网络开销，这里我们将SHARD_ROW_ID_BITS设置为4。

PRE_SPLIT_REGIONS指的是建表成功后的预均匀切分，我们通过设置PRE_SPLIT_REGIONS=2，实现建表成功后预均匀切分2^(PRE_SPLIT_REGIONS)个Region。