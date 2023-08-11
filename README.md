# Node exporter

[![CircleCI](https://circleci.com/gh/prometheus/node_exporter/tree/master.svg?style=shield)][circleci]
[![Buildkite status](https://badge.buildkite.com/94a0c1fb00b1f46883219c256efe9ce01d63b6505f3a942f9b.svg)](https://buildkite.com/prometheus/node-exporter)
[![Docker Repository on Quay](https://quay.io/repository/prometheus/node-exporter/status)][quay]
[![Docker Pulls](https://img.shields.io/docker/pulls/prom/node-exporter.svg?maxAge=604800)][hub]
[![Go Report Card](https://goreportcard.com/badge/github.com/prometheus/node_exporter)][goreportcard]

**1.4.1-5 / 2023-07-25**

**[BUGfix] 修复采集filebeat大日志慢的问题**


**1.4.1-4 / 2023-07-22**

**[ADD] 增加对filebeat打开文件状态的监控**

- HELP node_filebeat_openfiles Filebeat monitoring log harvester openfiles running.
- TYPE node_filebeat_openfiles gauge
- node_filebeat_openfiles 2



**1.4.1-3 / 2023-07-10**

**[ADD] 过滤对rootfs|tmpfs文件系统的采集**



**1.4.1-2 / 2023-07-04**

**[ADD] 增加对filebeat进程状态的监控**

- HELP node_filebeat_up Value is 1 if filebeat process is 'up', 0 otherwise.
- TYPE node_filebeat_up gauge
- node_filebeat_up{process_name="filebeat"}


**[ADD] 增加对rsyslog进程状态的监控**

- HELP node_rsyslog_up Value is 1 if rsyslog process is 'up', 0 otherwise.
- TYPE node_rsyslog_up gauge
- node_rsyslog_up{process_name="rsyslog"}


**[DEL] 过滤网络相关监控信息中对虚拟网卡的采集（docker0、vethxxx、lo）**

- node_network_xxxx{device="docker0"}
- node_network_xxxx{device="lo"}
- node_network_xxxx{device="veth261ba77"}


**[DEL] 移除以go_开头的33个监控项,具体如下：**

- go_gc_duration_seconds{quantile="0"}
- go_gc_duration_seconds{quantile="0.25"}
- go_gc_duration_seconds{quantile="0.5"}
- go_gc_duration_seconds{quantile="0.75"}
- go_gc_duration_seconds{quantile="1"}
- go_gc_duration_seconds_sum
- go_gc_duration_seconds_count
- go_goroutines
- go_info{version="go1.19.1"}
- go_memstats_alloc_bytes
- go_memstats_alloc_bytes_total
- go_memstats_buck_hash_sys_bytes
- go_memstats_frees_total
- go_memstats_gc_sys_bytes
- go_memstats_heap_alloc_bytes
- go_memstats_heap_idle_bytes
- go_memstats_heap_inuse_bytes
- go_memstats_heap_objects
- go_memstats_heap_released_bytes
- go_memstats_heap_sys_bytes
- go_memstats_last_gc_time_seconds
- go_memstats_lookups_total
- go_memstats_mallocs_total
- go_memstats_mcache_inuse_bytes
- go_memstats_mcache_sys_bytes
- go_memstats_mspan_inuse_bytes
- go_memstats_mspan_sys_bytes
- go_memstats_next_gc_bytes
- go_memstats_other_sys_bytes
- go_memstats_stack_inuse_bytes
- go_memstats_stack_sys_bytes
- go_memstats_sys_bytes
- go_threads




## Installation and Usage

If you are new to Prometheus and `node_exporter` there is a [simple step-by-step guide](https://prometheus.io/docs/guides/node-exporter/).

The `node_exporter` listens on HTTP port 9100 by default. See the `--help` output for more options.

### Ansible

For automated installs with [Ansible](https://www.ansible.com/), there is the [Cloud Alchemy role](https://github.com/cloudalchemy/ansible-node-exporter).

### RHEL/CentOS/Fedora

There is a [community-supplied COPR repository](https://copr.fedorainfracloud.org/coprs/ibotty/prometheus-exporters/) which closely follows upstream releases.

To enable only some specific collector(s), use `--collector.disable-defaults --collector.<name> ...`.

### Enabled by default

Name     | Description | OS
---------|-------------|----
bcache | Exposes bcache statistics from `/sys/fs/bcache/`. | Linux
bonding | Exposes the number of configured and active slaves of Linux bonding interfaces. | Linux
btrfs | Exposes btrfs statistics | Linux
boottime | Exposes system boot time derived from the `kern.boottime` sysctl. | Darwin, Dragonfly, FreeBSD, NetBSD, OpenBSD, Solaris
conntrack | Shows conntrack statistics (does nothing if no `/proc/sys/net/netfilter/` present). | Linux
cpu | Exposes CPU statistics | Darwin, Dragonfly, FreeBSD, Linux, Solaris, OpenBSD
cpufreq | Exposes CPU frequency statistics | Linux, Solaris
diskstats | Exposes disk I/O statistics. | Darwin, Linux, OpenBSD
dmi | Expose Desktop Management Interface (DMI) info from `/sys/class/dmi/id/` | Linux
edac | Exposes error detection and correction statistics. | Linux
entropy | Exposes available entropy. | Linux
exec | Exposes execution statistics. | Dragonfly, FreeBSD
fibrechannel | Exposes fibre channel information and statistics from `/sys/class/fc_host/`. | Linux
filefd | Exposes file descriptor statistics from `/proc/sys/fs/file-nr`. | Linux
filesystem | Exposes filesystem statistics, such as disk space used. | Darwin, Dragonfly, FreeBSD, Linux, OpenBSD
filebeat | filebeat process check | Linux
hwmon | Expose hardware monitoring and sensor data from `/sys/class/hwmon/`. | Linux
infiniband | Exposes network statistics specific to InfiniBand and Intel OmniPath configurations. | Linux
ipvs | Exposes IPVS status from `/proc/net/ip_vs` and stats from `/proc/net/ip_vs_stats`. | Linux
loadavg | Exposes load average. | Darwin, Dragonfly, FreeBSD, Linux, NetBSD, OpenBSD, Solaris
mdadm | Exposes statistics about devices in `/proc/mdstat` (does nothing if no `/proc/mdstat` present). | Linux
meminfo | Exposes memory statistics. | Darwin, Dragonfly, FreeBSD, Linux, OpenBSD
netclass | Exposes network interface info from `/sys/class/net/` | Linux
netdev | Exposes network interface statistics such as bytes transferred. | Darwin, Dragonfly, FreeBSD, Linux, OpenBSD
netstat | Exposes network statistics from `/proc/net/netstat`. This is the same information as `netstat -s`. | Linux
nfs | Exposes NFS client statistics from `/proc/net/rpc/nfs`. This is the same information as `nfsstat -c`. | Linux
nfsd | Exposes NFS kernel server statistics from `/proc/net/rpc/nfsd`. This is the same information as `nfsstat -s`. | Linux
nvme | Exposes NVMe info from `/sys/class/nvme/` | Linux
os | Expose OS release info from `/etc/os-release` or `/usr/lib/os-release` | _any_
powersupplyclass | Exposes Power Supply statistics from `/sys/class/power_supply` | Linux
pressure | Exposes pressure stall statistics from `/proc/pressure/`. | Linux (kernel 4.20+ and/or [CONFIG\_PSI](https://www.kernel.org/doc/html/latest/accounting/psi.html))
rapl | Exposes various statistics from `/sys/class/powercap`. | Linux
schedstat | Exposes task scheduler statistics from `/proc/schedstat`. | Linux
selinux | Exposes SELinux statistics. | Linux
sockstat | Exposes various statistics from `/proc/net/sockstat`. | Linux
softnet | Exposes statistics from `/proc/net/softnet_stat`. | Linux
stat | Exposes various statistics from `/proc/stat`. This includes boot time, forks and interrupts. | Linux
tapestats | Exposes statistics from `/sys/class/scsi_tape`. | Linux
textfile | Exposes statistics read from local disk. The `--collector.textfile.directory` flag must be set. | _any_
thermal | Exposes thermal statistics like `pmset -g therm`. | Darwin
thermal\_zone | Exposes thermal zone & cooling device statistics from `/sys/class/thermal`. | Linux
time | Exposes the current system time. | _any_
timex | Exposes selected adjtimex(2) system call stats. | Linux
udp_queues | Exposes UDP total lengths of the rx_queue and tx_queue from `/proc/net/udp` and `/proc/net/udp6`. | Linux
uname | Exposes system information as provided by the uname system call. | Darwin, FreeBSD, Linux, OpenBSD
vmstat | Exposes statistics from `/proc/vmstat`. | Linux
xfs | Exposes XFS runtime statistics. | Linux (kernel 4.4+)
zfs | Exposes [ZFS](http://open-zfs.org/) performance statistics. | FreeBSD, [Linux](http://zfsonlinux.org/), Solaris

### Disabled by default

arp | Exposes ARP statistics from `/proc/net/arp`. | Linux
`node_exporter` also implements a number of collectors that are disabled by default.  Reasons for this vary by
collector, and may include:
* High cardinality
* Prolonged runtime that exceeds the Prometheus `scrape_interval` or `scrape_timeout`
* Significant resource demands on the host

You can enable additional collectors as desired by adding them to your
init system's or service supervisor's startup configuration for
`node_exporter` but caution is advised.  Enable at most one at a time,
testing first on a non-production system, then by hand on a single
production node.  When enabling additional collectors, you should
carefully monitor the change by observing the `
scrape_duration_seconds` metric to ensure that collection completes
and does not time out.  In addition, monitor the
`scrape_samples_post_metric_relabeling` metric to see the changes in
cardinality.

Name     | Description | OS
---------|-------------|----
buddyinfo | Exposes statistics of memory fragments as reported by /proc/buddyinfo. | Linux
cgroups | A summary of the number of active and enabled cgroups | Linux
devstat | Exposes device statistics | Dragonfly, FreeBSD
drbd | Exposes Distributed Replicated Block Device statistics (to version 8.4) | Linux
ethtool | Exposes network interface information and network driver statistics equivalent to `ethtool`, `ethtool -S`, and `ethtool -i`. | Linux
interrupts | Exposes detailed interrupts statistics. | Linux, OpenBSD
ksmd | Exposes kernel and system statistics from `/sys/kernel/mm/ksm`. | Linux
lnstat | Exposes stats from `/proc/net/stat/`. | Linux
logind | Exposes session counts from [logind](http://www.freedesktop.org/wiki/Software/systemd/logind/). | Linux
meminfo\_numa | Exposes memory statistics from `/proc/meminfo_numa`. | Linux
mountstats | Exposes filesystem statistics from `/proc/self/mountstats`. Exposes detailed NFS client statistics. | Linux
network_route | Exposes the routing table as metrics | Linux
ntp | Exposes local NTP daemon health to check [time](./docs/TIME.md) | _any_
perf | Exposes perf based metrics (Warning: Metrics are dependent on kernel configuration and settings). | Linux
processes | Exposes aggregate process statistics from `/proc`. | Linux
qdisc | Exposes [queuing discipline](https://en.wikipedia.org/wiki/Network_scheduler#Linux_kernel) statistics | Linux
runit | Exposes service status from [runit](http://smarden.org/runit/). | _any_
slabinfo | Exposes slab statistics from `/proc/slabinfo`. Note that permission of `/proc/slabinfo` is usually 0400, so set it appropriately. | Linux
supervisord | Exposes service status from [supervisord](http://supervisord.org/). | _any_
sysctl | Expose sysctl values from `/proc/sys`. Use `--collector.sysctl.include(-info)` to configure. | Linux
systemd | Exposes service and system status from [systemd](http://www.freedesktop.org/wiki/Software/systemd/). | Linux
tcpstat | Exposes TCP connection status information from `/proc/net/tcp` and `/proc/net/tcp6`. (Warning: the current version has potential performance issues in high load situations.) | Linux
wifi | Exposes WiFi device and station statistics. | Linux
zoneinfo | Exposes NUMA memory zone metrics. | Linux

### Perf Collector

The `perf` collector may not work out of the box on some Linux systems due to kernel
configuration and security settings. To allow access, set the following `sysctl`
parameter:

```
sysctl -w kernel.perf_event_paranoid=X
```

- 2 allow only user-space measurements (default since Linux 4.6).
- 1 allow both kernel and user measurements (default before Linux 4.6).
- 0 allow access to CPU-specific data but not raw tracepoint samples.
- -1 no restrictions.

Depending on the configured value different metrics will be available, for most
cases `0` will provide the most complete set. For more information see [`man 2
perf_event_open`](http://man7.org/linux/man-pages/man2/perf_event_open.2.html).

By default, the `perf` collector will only collect metrics of the CPUs that
`node_exporter` is running on (ie
[`runtime.NumCPU`](https://golang.org/pkg/runtime/#NumCPU). If this is
insufficient (e.g. if you run `node_exporter` with its CPU affinity set to
specific CPUs), you can specify a list of alternate CPUs by using the
`--collector.perf.cpus` flag. For example, to collect metrics on CPUs 2-6, you
would specify: `--collector.perf --collector.perf.cpus=2-6`. The CPU
configuration is zero indexed and can also take a stride value; e.g.
`--collector.perf --collector.perf.cpus=1-10:5` would collect on CPUs
1, 5, and 10.

The `perf` collector is also able to collect
[tracepoint](https://www.kernel.org/doc/html/latest/core-api/tracepoint.html)
counts when using the `--collector.perf.tracepoint` flag. Tracepoints can be
found using [`perf list`](http://man7.org/linux/man-pages/man1/perf.1.html) or
from debugfs. And example usage of this would be
`--collector.perf.tracepoint="sched:sched_process_exec"`.

### Sysctl Collector

The `sysctl` collector can be enabled with `--collector.sysctl`. It supports exposing numeric sysctl values
as metrics using the `--collector.sysctl.include` flag and string values as info metrics by using the
`--collector.sysctl.include-info` flag. The flags can be repeated. For sysctl with multiple numeric values,
an optional mapping can be given to expose each value as its own metric. Otherwise an `index` label is used
to identify the different fields.

#### Examples
##### Numeric values
###### Single values
Using `--collector.sysctl.include=vm.user_reserve_kbytes`:
`vm.user_reserve_kbytes = 131072` -> `node_sysctl_vm_user_reserve_kbytes 131072`

###### Multiple values
A sysctl can contain multiple values, for example:
```
net.ipv4.tcp_rmem = 4096	131072	6291456
```
Using `--collector.sysctl.include=net.ipv4.tcp_rmem` the collector will expose:
```
node_sysctl_net_ipv4_tcp_rmem{index="0"} 4096
node_sysctl_net_ipv4_tcp_rmem{index="1"} 131072
node_sysctl_net_ipv4_tcp_rmem{index="2"} 6291456
```
If the indexes have defined meaning like in this case, the values can be mapped to multiple metrics by appending the mapping to the --collector.sysctl.include flag:
Using `--collector.sysctl.include=net.ipv4.tcp_rmem:min,default,max` the collector will expose:
```
node_sysctl_net_ipv4_tcp_rmem_min 4096
node_sysctl_net_ipv4_tcp_rmem_default 131072
node_sysctl_net_ipv4_tcp_rmem_max 6291456
```

##### String values
String values need to be exposed as info metric. The user selects them by using the `--collector.sysctl.include-info` flag.

###### Single values
`kernel.core_pattern = core` -> `node_sysctl_info{key="kernel.core_pattern_info", value="core"} 1`

###### Multiple values
Given the following sysctl:
```
kernel.seccomp.actions_avail = kill_process kill_thread trap errno trace log allow
```
Setting `--collector.sysctl.include-info=kernel.seccomp.actions_avail` will yield:
```
node_sysctl_info{key="kernel.seccomp.actions_avail", index="0", value="kill_process"} 1
node_sysctl_info{key="kernel.seccomp.actions_avail", index="1", value="kill_thread"} 1
...
```

### Textfile Collector

The `textfile` collector is similar to the [Pushgateway](https://github.com/prometheus/pushgateway),
in that it allows exporting of statistics from batch jobs. It can also be used
to export static metrics, such as what role a machine has. The Pushgateway
should be used for service-level metrics. The `textfile` module is for metrics
that are tied to a machine.

To use it, set the `--collector.textfile.directory` flag on the `node_exporter` commandline. The
collector will parse all files in that directory matching the glob `*.prom`
using the [text
format](http://prometheus.io/docs/instrumenting/exposition_formats/). **Note:** Timestamps are not supported.

To atomically push completion time for a cron job:
```
echo my_batch_job_completion_time $(date +%s) > /path/to/directory/my_batch_job.prom.$$
mv /path/to/directory/my_batch_job.prom.$$ /path/to/directory/my_batch_job.prom
```

To statically set roles for a machine using labels:
```
echo 'role{role="application_server"} 1' > /path/to/directory/role.prom.$$
mv /path/to/directory/role.prom.$$ /path/to/directory/role.prom
```

### Filtering enabled collectors

The `node_exporter` will expose all metrics from enabled collectors by default.  This is the recommended way to collect metrics to avoid errors when comparing metrics of different families.

For advanced use the `node_exporter` can be passed an optional list of collectors to filter metrics. The `collect[]` parameter may be used multiple times.  In Prometheus configuration you can use this syntax under the [scrape config](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#<scrape_config>).

```
  params:
    collect[]:
      - foo
      - bar
```

This can be useful for having different Prometheus servers collect specific metrics from nodes.

## Development building and running

Prerequisites:

* [Go compiler](https://golang.org/dl/)
* RHEL/CentOS: `glibc-static` package.

Building:

    cd node_exporter
    make build  
    ./node_exporter <flags>

To see all available configuration flags:

    ./node_exporter -h

## Running tests

    make test

## TLS endpoint

** EXPERIMENTAL **

The exporter supports TLS via a new web configuration file.

```console
./node_exporter --web.config=web-config.yml
```

See the [exporter-toolkit https package](https://github.com/prometheus/exporter-toolkit/blob/v0.1.0/https/README.md) for more details.

[travis]: https://travis-ci.org/prometheus/node_exporter
[hub]: https://hub.docker.com/r/prom/node-exporter/
[circleci]: https://circleci.com/gh/prometheus/node_exporter
[quay]: https://quay.io/repository/prometheus/node-exporter
[goreportcard]: https://goreportcard.com/report/github.com/prometheus/node_exporter
