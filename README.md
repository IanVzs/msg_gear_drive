# virtual_room
事件->创建Hub->Hub中添加Client

## Client
client中包括心跳检测,可`receive`来自`Hub`的事件也可推送事件到`Hub`的`broadcast`

## Hub
Hub中包含`cron`, `broadcast`, `clients`
