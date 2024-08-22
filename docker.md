1. 查看在 docker ps -a 中出现但在 docker ps 中没有出现的容器的所有字段
```sh
docker ps -a --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}" | grep -v "$(docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}")"
```
2. 查看特定容器的状态
```bash
docker ps -f name=<container_name>
docker ps -f id=<container_id>

```
3. 查看Linux发行版 `cat /etc/issue`
Alpine Linux