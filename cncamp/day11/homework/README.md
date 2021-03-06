### yaml文件修改成了仅deploy，增加了annotations及相关监控参数配置
### 源代码增加基于prometheus的监控指标获取，添加一个测试用延迟处理函数
### 要求4 获取指标数据
```
http://127.0.0.1:8000/metrics/返回结果
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 9
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.17.3"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 1.918912e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 1.918912e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 6986
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 1284
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 2.34468e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 1.918912e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 4.202496e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.8912e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 14168
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 4.202496e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 8.093696e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 15452
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 9344
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 64056
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 65536
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 999278
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 294912
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 294912
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.1821472e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 8
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.046875
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.6777216e+07
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 110
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.2746752e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.638935685e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.8014208e+07
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP service_http_request_count_total Total number of HTTP requests made.
# TYPE service_http_request_count_total counter
service_http_request_count_total{endpoint="/",method="GET",status="404"} 1
service_http_request_count_total{endpoint="/delay/",method="GET",status="200"} 5
service_http_request_count_total{endpoint="/favicon.ico",method="GET",status="404"} 1
service_http_request_count_total{endpoint="/healthz/",method="GET",status="200"} 1
# HELP service_http_request_duration_seconds HTTP request latencies in seconds.
# TYPE service_http_request_duration_seconds histogram
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="0.005"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="0.01"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="0.025"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="0.05"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="0.1"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="0.25"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="0.5"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="1"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="2.5"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="5"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="10"} 1
service_http_request_duration_seconds_bucket{endpoint="/",method="GET",status="404",le="+Inf"} 1
service_http_request_duration_seconds_sum{endpoint="/",method="GET",status="404"} 0
service_http_request_duration_seconds_count{endpoint="/",method="GET",status="404"} 1
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="0.005"} 0
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="0.01"} 0
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="0.025"} 0
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="0.05"} 0
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="0.1"} 0
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="0.25"} 0
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="0.5"} 0
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="1"} 1
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="2.5"} 5
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="5"} 5
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="10"} 5
service_http_request_duration_seconds_bucket{endpoint="/delay/",method="GET",status="200",le="+Inf"} 5
service_http_request_duration_seconds_sum{endpoint="/delay/",method="GET",status="200"} 6.8790857
service_http_request_duration_seconds_count{endpoint="/delay/",method="GET",status="200"} 5
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="0.005"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="0.01"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="0.025"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="0.05"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="0.1"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="0.25"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="0.5"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="1"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="2.5"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="5"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="10"} 1
service_http_request_duration_seconds_bucket{endpoint="/favicon.ico",method="GET",status="404",le="+Inf"} 1
service_http_request_duration_seconds_sum{endpoint="/favicon.ico",method="GET",status="404"} 0.0009962
service_http_request_duration_seconds_count{endpoint="/favicon.ico",method="GET",status="404"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="0.005"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="0.01"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="0.025"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="0.05"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="0.1"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="0.25"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="0.5"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="1"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="2.5"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="5"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="10"} 1
service_http_request_duration_seconds_bucket{endpoint="/healthz/",method="GET",status="200",le="+Inf"} 1
service_http_request_duration_seconds_sum{endpoint="/healthz/",method="GET",status="200"} 5.43e-05
service_http_request_duration_seconds_count{endpoint="/healthz/",method="GET",status="200"} 1
# HELP service_http_request_size_bytes HTTP request sizes in bytes.
# TYPE service_http_request_size_bytes summary
service_http_request_size_bytes_sum{endpoint="/",method="GET",status="404"} 589
service_http_request_size_bytes_count{endpoint="/",method="GET",status="404"} 1
service_http_request_size_bytes_sum{endpoint="/delay/",method="GET",status="200"} 3063
service_http_request_size_bytes_count{endpoint="/delay/",method="GET",status="200"} 5
service_http_request_size_bytes_sum{endpoint="/favicon.ico",method="GET",status="404"} 519
service_http_request_size_bytes_count{endpoint="/favicon.ico",method="GET",status="404"} 1
service_http_request_size_bytes_sum{endpoint="/healthz/",method="GET",status="200"} 597
service_http_request_size_bytes_count{endpoint="/healthz/",method="GET",status="200"} 1
# HELP service_http_response_size_bytes HTTP response sizes in bytes.
# TYPE service_http_response_size_bytes summary
service_http_response_size_bytes_sum{endpoint="/",method="GET",status="404"} 0
service_http_response_size_bytes_count{endpoint="/",method="GET",status="404"} 1
service_http_response_size_bytes_sum{endpoint="/delay/",method="GET",status="200"} 144
service_http_response_size_bytes_count{endpoint="/delay/",method="GET",status="200"} 5
service_http_response_size_bytes_sum{endpoint="/favicon.ico",method="GET",status="404"} 0
service_http_response_size_bytes_count{endpoint="/favicon.ico",method="GET",status="404"} 1
service_http_response_size_bytes_sum{endpoint="/healthz/",method="GET",status="200"} 12
service_http_response_size_bytes_count{endpoint="/healthz/",method="GET",status="200"} 1
# HELP service_uptime HTTP service uptime.
# TYPE service_uptime counter
service_uptime 121
```