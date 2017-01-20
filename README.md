# date-filter

用于筛选当前时间到之前指定的范围内的日志内容的小工具，支持对php慢查询日志等多行日志文件的筛选。

    tail -n 20000 /var/log/nginx/access.log \
         | date-filter -layout '2006-01-02 15:04:05' -offset 1 -valid-time 1m

上例中会输出发生时间在当前时间1分钟以内的日志。

## 参数说明

### -layout 日期格式模板

日期格式使用Go语言标准的日期时间表示方法。

| 符号 | 说明
|-----|--------
| 2006 | 年
| 01 | 月
| 02 | 日
| 15 | 时（24小时制）
| 04 | 分
| 05 | 秒

例如 

- `-layout '2006-01-02 15:04:05'`

        [2017-01-20 10:27:17] production.DEBUG: request-consuming: 370.5661 ms [] {"process_id":10412}

- `-layout '2006-01-02T15:04:05'`

        time[2017-01-20T10:25:54+08:00] ip[125.84.236.232] ...

- `-layout '02-Jan-2006 15:04:05'`
    
        [20-Jan-2017 10:24:20]  [pool www] pid 10409

### -offset 时间截取位置

目前只支持固定时间偏移位置的方式获取时间。`offset`指定了日期时间在每一行中的开始位置，`date-filter`将会从`offset`位置开始，截取`layout`长度的内容作为当前行的时间。

### -valid-time 有效时间

只有在当前时间前`valid-time`时间范围内的行会被输出。时间单位支持`ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`。比如 `300ms`，`1.5h`，`2h45m`。

## 使用范例

    LOG_FILE=/directory/log/xxx.log
    tail -n 20000 $LOG_FILE | date-filter -layout '2006-01-02 15:04:05' -offset 1 -valid-time 1m | awk -v log_file=$LOG_FILE  '{ printf "%s %s\n", log_file, $0 }' | dos2unix

