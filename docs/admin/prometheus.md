# Prometheus

Coder exposes many metrics which can be consumed by a Prometheus server, and give insight into the current state of a live Coder deployment.

If you don't have an Prometheus server installed, you can follow the Prometheus [Getting started](https://prometheus.io/docs/prometheus/latest/getting_started/) guide.

## Enable Prometheus metrics

Coder server exports metrics via the HTTP endpoint, which can be enabled using either the environment variable `CODER_PROMETHEUS_ENABLE` or the flag `--prometheus-enable`.

The Prometheus endpoint address is `http://localhost:2112/` by default. You can use either the environment variable `CODER_PROMETHEUS_ADDRESS` or the flag ` --prometheus-address <network-interface>:<port>` to select a different listen address.

If `coder server --prometheus-enable` is started locally, you can preview the metrics endpoint in your browser or by using curl: <!-- markdown-link-check-disable -->http://localhost:2112/<!-- markdown-link-check-enable -->.

```shell
$ curl http://localhost:2112/
# HELP coderd_api_active_users_duration_hour The number of users that have been active within the last hour.
# TYPE coderd_api_active_users_duration_hour gauge
coderd_api_active_users_duration_hour 0
...
```

### Kubernetes deployment

The Prometheus endpoint can be enabled in the [Helm chart's](https://github.com/coder/coder/tree/main/helm) `values.yml` by setting the environment variable `CODER_PROMETHEUS_ADDRESS` to `0.0.0.0:2112`.
The environment variable `CODER_PROMETHEUS_ENABLE` will be enabled automatically.

## Available metrics

<!-- Code generated by 'make docs/admin/prometheus.md'. DO NOT EDIT -->

| Name                                         | Type      | Description                                                        | Labels                 |
| -------------------------------------------- | --------- | ------------------------------------------------------------------ | ---------------------- |
| `coderd_api_active_users_duration_hour`      | gauge     | The number of users that have been active within the last hour.    |                        |
| `coderd_api_concurrent_requests`             | gauge     | The number of concurrent API requests                              |                        |
| `coderd_api_concurrent_websockets`           | gauge     | The total number of concurrent API websockets                      |                        |
| `coderd_api_request_latencies_ms`            | histogram | Latency distribution of requests in milliseconds                   | `method` `path`        |
| `coderd_api_requests_processed_total`        | counter   | The total number of processed API requests                         | `code` `method` `path` |
| `coderd_api_websocket_durations_ms`          | histogram | Websocket duration distribution of requests in milliseconds        | `path`                 |
| `coderd_api_workspace_latest_build_total`    | gauge     | The latest workspace builds with a status.                         | `status`               |
| `coderd_provisionerd_job_timings_ms`         | histogram | The provisioner job time duration.                                 | `provisioner` `status` |
| `coderd_provisionerd_jobs_current`           | gauge     | The number of currently running provisioner jobs.                  | `provisioner`          |
| `go_gc_duration_seconds`                     | summary   | A summary of the pause duration of garbage collection cycles.      |                        |
| `go_goroutines`                              | gauge     | Number of goroutines that currently exist.                         |                        |
| `go_info`                                    | gauge     | Information about the Go environment.                              | `version`              |
| `go_memstats_alloc_bytes`                    | gauge     | Number of bytes allocated and still in use.                        |                        |
| `go_memstats_alloc_bytes_total`              | counter   | Total number of bytes allocated, even if freed.                    |                        |
| `go_memstats_buck_hash_sys_bytes`            | gauge     | Number of bytes used by the profiling bucket hash table.           |                        |
| `go_memstats_frees_total`                    | counter   | Total number of frees.                                             |                        |
| `go_memstats_gc_sys_bytes`                   | gauge     | Number of bytes used for garbage collection system metadata.       |                        |
| `go_memstats_heap_alloc_bytes`               | gauge     | Number of heap bytes allocated and still in use.                   |                        |
| `go_memstats_heap_idle_bytes`                | gauge     | Number of heap bytes waiting to be used.                           |                        |
| `go_memstats_heap_inuse_bytes`               | gauge     | Number of heap bytes that are in use.                              |                        |
| `go_memstats_heap_objects`                   | gauge     | Number of allocated objects.                                       |                        |
| `go_memstats_heap_released_bytes`            | gauge     | Number of heap bytes released to OS.                               |                        |
| `go_memstats_heap_sys_bytes`                 | gauge     | Number of heap bytes obtained from system.                         |                        |
| `go_memstats_last_gc_time_seconds`           | gauge     | Number of seconds since 1970 of last garbage collection.           |                        |
| `go_memstats_lookups_total`                  | counter   | Total number of pointer lookups.                                   |                        |
| `go_memstats_mallocs_total`                  | counter   | Total number of mallocs.                                           |                        |
| `go_memstats_mcache_inuse_bytes`             | gauge     | Number of bytes in use by mcache structures.                       |                        |
| `go_memstats_mcache_sys_bytes`               | gauge     | Number of bytes used for mcache structures obtained from system.   |                        |
| `go_memstats_mspan_inuse_bytes`              | gauge     | Number of bytes in use by mspan structures.                        |                        |
| `go_memstats_mspan_sys_bytes`                | gauge     | Number of bytes used for mspan structures obtained from system.    |                        |
| `go_memstats_next_gc_bytes`                  | gauge     | Number of heap bytes when next garbage collection will take place. |                        |
| `go_memstats_other_sys_bytes`                | gauge     | Number of bytes used for other system allocations.                 |                        |
| `go_memstats_stack_inuse_bytes`              | gauge     | Number of bytes in use by the stack allocator.                     |                        |
| `go_memstats_stack_sys_bytes`                | gauge     | Number of bytes obtained from system for stack allocator.          |                        |
| `go_memstats_sys_bytes`                      | gauge     | Number of bytes obtained from system.                              |                        |
| `go_threads`                                 | gauge     | Number of OS threads created.                                      |                        |
| `promhttp_metric_handler_requests_in_flight` | gauge     | Current number of scrapes being served.                            |                        |
| `promhttp_metric_handler_requests_total`     | counter   | Total number of scrapes by HTTP status code.                       | `code`                 |

<!-- End generated by 'make docs/admin/prometheus.md'. -->