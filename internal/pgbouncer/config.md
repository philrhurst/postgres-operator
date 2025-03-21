<!--
# Copyright 2021 - 2025 Crunchy Data Solutions, Inc.
#
# SPDX-License-Identifier: Apache-2.0
-->

PgBouncer is configured through INI files. It will reload these files when
receiving a `HUP` signal or [`RELOAD` command][RELOAD] in the admin console.

There is a [`SET` command][SET] available in the admin console, but it is not
clear when those changes take affect.

- https://www.pgbouncer.org/config.html

[RELOAD]: https://www.pgbouncer.org/usage.html#process-controlling-commands
[SET]: https://www.pgbouncer.org/usage.html#other-commands

The [`%include` directive](https://www.pgbouncer.org/config.html#include-directive)
allows one file to refer other existing files.

There are three sections in the files:

 - `[pgbouncer]` is for settings that apply to the PgBouncer process.
 - `[databases]` is a list of databases to which clients can connect.
 - `[users]` changes a few database settings based on the client user.

```
psql (12.6, server 1.15.0/bouncer)

pgbouncer=# SHOW CONFIG;
            key            |                         value                          |                        default                         | changeable 
---------------------------+--------------------------------------------------------+--------------------------------------------------------+------------
 admin_users               |                                                        |                                                        | yes
 application_name_add_host | 0                                                      | 0                                                      | yes
 auth_file                 |                                                        |                                                        | yes
 auth_hba_file             |                                                        |                                                        | yes
 auth_query                | SELECT usename, passwd FROM pg_shadow WHERE usename=$1 | SELECT usename, passwd FROM pg_shadow WHERE usename=$1 | yes
 auth_type                 | md5                                                    | md5                                                    | yes
 auth_user                 |                                                        |                                                        | yes
 autodb_idle_timeout       | 3600                                                   | 3600                                                   | yes
 client_idle_timeout       | 0                                                      | 0                                                      | yes
 client_login_timeout      | 60                                                     | 60                                                     | yes
 client_tls_ca_file        |                                                        |                                                        | no
 client_tls_cert_file      |                                                        |                                                        | no
 client_tls_ciphers        | fast                                                   | fast                                                   | no
 client_tls_dheparams      | auto                                                   | auto                                                   | no
 client_tls_ecdhcurve      | auto                                                   | auto                                                   | no
 client_tls_key_file       |                                                        |                                                        | no
 client_tls_protocols      | secure                                                 | secure                                                 | no
 client_tls_sslmode        | disable                                                | disable                                                | no
 conffile                  | /tmp/pgbouncer.ini                                     |                                                        | yes
 default_pool_size         | 20                                                     | 20                                                     | yes
 disable_pqexec            | 0                                                      | 0                                                      | no
 dns_max_ttl               | 15                                                     | 15                                                     | yes
 dns_nxdomain_ttl          | 15                                                     | 15                                                     | yes
 dns_zone_check_period     | 0                                                      | 0                                                      | yes
 idle_transaction_timeout  | 0                                                      | 0                                                      | yes
 ignore_startup_parameters |                                                        |                                                        | yes
 job_name                  | pgbouncer                                              | pgbouncer                                              | no
 listen_addr               | *                                                      |                                                        | no
 listen_backlog            | 128                                                    | 128                                                    | no
 listen_port               | 6432                                                   | 6432                                                   | no
 log_connections           | 1                                                      | 1                                                      | yes
 log_disconnections        | 1                                                      | 1                                                      | yes
 log_pooler_errors         | 1                                                      | 1                                                      | yes
 log_stats                 | 1                                                      | 1                                                      | yes
 logfile                   |                                                        |                                                        | yes
 max_client_conn           | 100                                                    | 100                                                    | yes
 max_db_connections        | 0                                                      | 0                                                      | yes
 max_packet_size           | 2147483647                                             | 2147483647                                             | yes
 max_user_connections      | 0                                                      | 0                                                      | yes
 min_pool_size             | 0                                                      | 0                                                      | yes
 pidfile                   |                                                        |                                                        | no
 pkt_buf                   | 4096                                                   | 4096                                                   | no
 pool_mode                 | session                                                | session                                                | yes
 query_timeout             | 0                                                      | 0                                                      | yes
 query_wait_timeout        | 120                                                    | 120                                                    | yes
 reserve_pool_size         | 0                                                      | 0                                                      | yes
 reserve_pool_timeout      | 5                                                      | 5                                                      | yes
 resolv_conf               |                                                        |                                                        | no
 sbuf_loopcnt              | 5                                                      | 5                                                      | yes
 server_check_delay        | 30                                                     | 30                                                     | yes
 server_check_query        | select 1                                               | select 1                                               | yes
 server_connect_timeout    | 15                                                     | 15                                                     | yes
 server_fast_close         | 0                                                      | 0                                                      | yes
 server_idle_timeout       | 600                                                    | 600                                                    | yes
 server_lifetime           | 3600                                                   | 3600                                                   | yes
 server_login_retry        | 15                                                     | 15                                                     | yes
 server_reset_query        | DISCARD ALL                                            | DISCARD ALL                                            | yes
 server_reset_query_always | 0                                                      | 0                                                      | yes
 server_round_robin        | 0                                                      | 0                                                      | yes
 server_tls_ca_file        |                                                        |                                                        | no
 server_tls_cert_file      |                                                        |                                                        | no
 server_tls_ciphers        | fast                                                   | fast                                                   | no
 server_tls_key_file       |                                                        |                                                        | no
 server_tls_protocols      | secure                                                 | secure                                                 | no
 server_tls_sslmode        | disable                                                | disable                                                | no
 so_reuseport              | 0                                                      | 0                                                      | no
 stats_period              | 60                                                     | 60                                                     | yes
 stats_users               |                                                        |                                                        | yes
 suspend_timeout           | 10                                                     | 10                                                     | yes
 syslog                    | 0                                                      | 0                                                      | yes
 syslog_facility           | daemon                                                 | daemon                                                 | yes
 syslog_ident              | pgbouncer                                              | pgbouncer                                              | yes
 tcp_defer_accept          | 1                                                      |                                                        | yes
 tcp_keepalive             | 1                                                      | 1                                                      | yes
 tcp_keepcnt               | 0                                                      | 0                                                      | yes
 tcp_keepidle              | 0                                                      | 0                                                      | yes
 tcp_keepintvl             | 0                                                      | 0                                                      | yes
 tcp_socket_buffer         | 0                                                      | 0                                                      | yes
 tcp_user_timeout          | 0                                                      | 0                                                      | yes
 unix_socket_dir           | /tmp                                                   | /tmp                                                   | no
 unix_socket_group         |                                                        |                                                        | no
 unix_socket_mode          | 511                                                    | 0777                                                   | no
 user                      |                                                        |                                                        | no
 verbose                   | 0                                                      |                                                        | yes
(84 rows)
```
