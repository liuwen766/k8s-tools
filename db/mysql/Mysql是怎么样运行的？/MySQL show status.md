MySQL show global status;

```
root@localhost : (none) 09:15:06> show global status;
+-----------------------------------------------+---------------------------------------------------------------------------------------------------+
| Variable_name                                 | Value                                                                                             |
+-----------------------------------------------+---------------------------------------------------------------------------------------------------+
| Aborted_clients                               | 4                                                                                                 |
| Aborted_connects                              | 0                                                                                                 |
| Audit_log_buffer_size_overflow                | 0                                                                                                 |
| Binlog_cache_disk_use                         | 0                                                                                                 |
| Binlog_cache_use                              | 1578216                                                                                           |
| Binlog_stmt_cache_disk_use                    | 0                                                                                                 |
| Binlog_stmt_cache_use                         | 0                                                                                                 |
| Bytes_received                                | 1563450533                                                                                        |
| Bytes_sent                                    | 4027169172                                                                                        |
| Com_admin_commands                            | 392165                                                                                            |
| Com_assign_to_keycache                        | 0                                                                                                 |
| Com_alter_db                                  | 0                                                                                                 |
| Com_alter_db_upgrade                          | 0                                                                                                 |
| Com_alter_event                               | 0                                                                                                 |
| Com_alter_function                            | 0                                                                                                 |
| Com_alter_instance                            | 0                                                                                                 |
| Com_alter_procedure                           | 0                                                                                                 |
| Com_alter_server                              | 0                                                                                                 |
| Com_alter_table                               | 0                                                                                                 |
| Com_alter_tablespace                          | 0                                                                                                 |
| Com_alter_user                                | 0                                                                                                 |
| Com_analyze                                   | 0                                                                                                 |
| Com_begin                                     | 1578216                                                                                           |
| Com_binlog                                    | 0                                                                                                 |
| Com_call_procedure                            | 0                                                                                                 |
| Com_change_db                                 | 1                                                                                                 |
| Com_change_master                             | 1                                                                                                 |
| Com_change_repl_filter                        | 0                                                                                                 |
| Com_check                                     | 0                                                                                                 |
| Com_checksum                                  | 0                                                                                                 |
| Com_commit                                    | 1578216                                                                                           |
| Com_create_db                                 | 0                                                                                                 |
| Com_create_event                              | 0                                                                                                 |
| Com_create_function                           | 0                                                                                                 |
| Com_create_index                              | 0                                                                                                 |
| Com_create_procedure                          | 0                                                                                                 |
| Com_create_server                             | 0                                                                                                 |
| Com_create_table                              | 0                                                                                                 |
| Com_create_trigger                            | 0                                                                                                 |
| Com_create_udf                                | 0                                                                                                 |
| Com_create_user                               | 0                                                                                                 |
| Com_create_view                               | 0                                                                                                 |
| Com_dealloc_sql                               | 0                                                                                                 |
| Com_delete                                    | 0                                                                                                 |
| Com_delete_multi                              | 0                                                                                                 |
| Com_do                                        | 0                                                                                                 |
| Com_drop_db                                   | 0                                                                                                 |
| Com_drop_event                                | 0                                                                                                 |
| Com_drop_function                             | 0                                                                                                 |
| Com_drop_index                                | 0                                                                                                 |
| Com_drop_procedure                            | 0                                                                                                 |
| Com_drop_server                               | 0                                                                                                 |
| Com_drop_table                                | 0                                                                                                 |
| Com_drop_trigger                              | 0                                                                                                 |
| Com_drop_user                                 | 0                                                                                                 |
| Com_drop_view                                 | 0                                                                                                 |
| Com_empty_query                               | 0                                                                                                 |
| Com_execute_sql                               | 0                                                                                                 |
| Com_explain_other                             | 0                                                                                                 |
| Com_flush                                     | 0                                                                                                 |
| Com_get_diagnostics                           | 0                                                                                                 |
| Com_grant                                     | 0                                                                                                 |
| Com_ha_close                                  | 0                                                                                                 |
| Com_ha_open                                   | 0                                                                                                 |
| Com_ha_read                                   | 0                                                                                                 |
| Com_help                                      | 0                                                                                                 |
| Com_insert                                    | 1                                                                                                 |
| Com_insert_select                             | 0                                                                                                 |
| Com_install_plugin                            | 0                                                                                                 |
| Com_kill                                      | 0                                                                                                 |
| Com_load                                      | 0                                                                                                 |
| Com_lock_tables                               | 0                                                                                                 |
| Com_optimize                                  | 0                                                                                                 |
| Com_preload_keys                              | 0                                                                                                 |
| Com_prepare_sql                               | 0                                                                                                 |
| Com_purge                                     | 0                                                                                                 |
| Com_purge_before_date                         | 0                                                                                                 |
| Com_release_savepoint                         | 0                                                                                                 |
| Com_rename_table                              | 0                                                                                                 |
| Com_rename_user                               | 0                                                                                                 |
| Com_repair                                    | 0                                                                                                 |
| Com_replace                                   | 0                                                                                                 |
| Com_replace_select                            | 0                                                                                                 |
| Com_reset                                     | 1                                                                                                 |
| Com_resignal                                  | 0                                                                                                 |
| Com_revoke                                    | 0                                                                                                 |
| Com_revoke_all                                | 0                                                                                                 |
| Com_rollback                                  | 0                                                                                                 |
| Com_rollback_to_savepoint                     | 0                                                                                                 |
| Com_savepoint                                 | 0                                                                                                 |
| Com_select                                    | 6630806                                                                                           |
| Com_set_option                                | 392166                                                                                            |
| Com_signal                                    | 0                                                                                                 |
| Com_show_binlog_events                        | 0                                                                                                 |
| Com_show_binlogs                              | 30906                                                                                             |
| Com_show_charsets                             | 0                                                                                                 |
| Com_show_collations                           | 0                                                                                                 |
| Com_show_create_db                            | 0                                                                                                 |
| Com_show_create_event                         | 0                                                                                                 |
| Com_show_create_func                          | 0                                                                                                 |
| Com_show_create_proc                          | 0                                                                                                 |
| Com_show_create_table                         | 0                                                                                                 |
| Com_show_create_trigger                       | 0                                                                                                 |
| Com_show_databases                            | 0                                                                                                 |
| Com_show_engine_logs                          | 0                                                                                                 |
| Com_show_engine_mutex                         | 0                                                                                                 |
| Com_show_engine_status                        | 0                                                                                                 |
| Com_show_events                               | 0                                                                                                 |
| Com_show_errors                               | 0                                                                                                 |
| Com_show_fields                               | 0                                                                                                 |
| Com_show_function_code                        | 0                                                                                                 |
| Com_show_function_status                      | 0                                                                                                 |
| Com_show_grants                               | 0                                                                                                 |
| Com_show_keys                                 | 0                                                                                                 |
| Com_show_master_status                        | 4                                                                                                 |
| Com_show_open_tables                          | 0                                                                                                 |
| Com_show_plugins                              | 0                                                                                                 |
| Com_show_privileges                           | 0                                                                                                 |
| Com_show_procedure_code                       | 0                                                                                                 |
| Com_show_procedure_status                     | 0                                                                                                 |
| Com_show_processlist                          | 0                                                                                                 |
| Com_show_profile                              | 0                                                                                                 |
| Com_show_profiles                             | 0                                                                                                 |
| Com_show_relaylog_events                      | 0                                                                                                 |
| Com_show_slave_hosts                          | 0                                                                                                 |
| Com_show_slave_status                         | 654597                                                                                            |
| Com_show_status                               | 30917                                                                                             |
| Com_show_storage_engines                      | 0                                                                                                 |
| Com_show_table_status                         | 0                                                                                                 |
| Com_show_tables                               | 0                                                                                                 |
| Com_show_triggers                             | 0                                                                                                 |
| Com_show_variables                            | 30912                                                                                             |
| Com_show_warnings                             | 3                                                                                                 |
| Com_show_create_user                          | 0                                                                                                 |
| Com_shutdown                                  | 0                                                                                                 |
| Com_slave_start                               | 1                                                                                                 |
| Com_slave_stop                                | 1                                                                                                 |
| Com_group_replication_start                   | 0                                                                                                 |
| Com_group_replication_stop                    | 0                                                                                                 |
| Com_stmt_execute                              | 0                                                                                                 |
| Com_stmt_close                                | 0                                                                                                 |
| Com_stmt_fetch                                | 0                                                                                                 |
| Com_stmt_prepare                              | 0                                                                                                 |
| Com_stmt_reset                                | 0                                                                                                 |
| Com_stmt_send_long_data                       | 0                                                                                                 |
| Com_truncate                                  | 0                                                                                                 |
| Com_uninstall_plugin                          | 0                                                                                                 |
| Com_unlock_tables                             | 0                                                                                                 |
| Com_update                                    | 1578215                                                                                           |
| Com_update_multi                              | 0                                                                                                 |
| Com_xa_commit                                 | 0                                                                                                 |
| Com_xa_end                                    | 0                                                                                                 |
| Com_xa_prepare                                | 0                                                                                                 |
| Com_xa_recover                                | 0                                                                                                 |
| Com_xa_rollback                               | 0                                                                                                 |
| Com_xa_start                                  | 0                                                                                                 |
| Com_stmt_reprepare                            | 0                                                                                                 |
| Connection_errors_accept                      | 0                                                                                                 |
| Connection_errors_internal                    | 0                                                                                                 |
| Connection_errors_max_connections             | 0                                                                                                 |
| Connection_errors_peer_address                | 0                                                                                                 |
| Connection_errors_select                      | 0                                                                                                 |
| Connection_errors_tcpwrap                     | 0                                                                                                 |
| Connections                                   | 2886930                                                                                           |
| Created_tmp_disk_tables                       | 5                                                                                                 |
| Created_tmp_files                             | 6                                                                                                 |
| Created_tmp_tables                            | 123646                                                                                            |
| Delayed_errors                                | 0                                                                                                 |
| Delayed_insert_threads                        | 0                                                                                                 |
| Delayed_writes                                | 0                                                                                                 |
| Flush_commands                                | 1                                                                                                 |
| Handler_commit                                | 9743578                                                                                           |
| Handler_delete                                | 1                                                                                                 |
| Handler_discover                              | 0                                                                                                 |
| Handler_external_lock                         | 14424469                                                                                          |
| Handler_mrr_init                              | 0                                                                                                 |
| Handler_prepare                               | 6312864                                                                                           |
| Handler_read_first                            | 37                                                                                                |
| Handler_read_key                              | 5008972                                                                                           |
| Handler_read_last                             | 0                                                                                                 |
| Handler_read_next                             | 2                                                                                                 |
| Handler_read_prev                             | 0                                                                                                 |
| Handler_read_rnd                              | 1578269                                                                                           |
| Handler_read_rnd_next                         | 1184796657                                                                                        |
| Handler_rollback                              | 0                                                                                                 |
| Handler_savepoint                             | 0                                                                                                 |
| Handler_savepoint_rollback                    | 0                                                                                                 |
| Handler_update                                | 3761451                                                                                           |
| Handler_write                                 | 28440491                                                                                          |
| Innodb_buffer_pool_dump_status                | Dumping of buffer pool not started                                                                |
| Innodb_buffer_pool_load_status                | Cannot open '/var/lib/mysql/data/innodb_ts/ib_buffer_pool' for reading: No such file or directory |
| Innodb_buffer_pool_resize_status              |                                                                                                   |
| Innodb_buffer_pool_pages_data                 | 429                                                                                               |
| Innodb_buffer_pool_bytes_data                 | 7028736                                                                                           |
| Innodb_buffer_pool_pages_dirty                | 0                                                                                                 |
| Innodb_buffer_pool_bytes_dirty                | 0                                                                                                 |
| Innodb_buffer_pool_pages_flushed              | 7075851                                                                                           |
| Innodb_buffer_pool_pages_free                 | 359975                                                                                            |
| Innodb_buffer_pool_pages_misc                 | 0                                                                                                 |
| Innodb_buffer_pool_pages_total                | 360404                                                                                            |
| Innodb_buffer_pool_read_ahead_rnd             | 0                                                                                                 |
| Innodb_buffer_pool_read_ahead                 | 0                                                                                                 |
| Innodb_buffer_pool_read_ahead_evicted         | 0                                                                                                 |
| Innodb_buffer_pool_read_requests              | 36133223                                                                                          |
| Innodb_buffer_pool_reads                      | 376                                                                                               |
| Innodb_buffer_pool_wait_free                  | 0                                                                                                 |
| Innodb_buffer_pool_write_requests             | 22378377                                                                                          |
| Innodb_data_fsyncs                            | 10928279                                                                                          |
| Innodb_data_pending_fsyncs                    | 0                                                                                                 |
| Innodb_data_pending_reads                     | 0                                                                                                 |
| Innodb_data_pending_writes                    | 0                                                                                                 |
| Innodb_data_read                              | 7852544                                                                                           |
| Innodb_data_reads                             | 545                                                                                               |
| Innodb_data_writes                            | 13991289                                                                                          |
| Innodb_data_written                           | 237165706240                                                                                      |
| Innodb_dblwr_pages_written                    | 7075802                                                                                           |
| Innodb_dblwr_writes                           | 2432518                                                                                           |
| Innodb_log_waits                              | 0                                                                                                 |
| Innodb_log_write_requests                     | 3106881                                                                                           |
| Innodb_log_writes                             | 3443412                                                                                           |
| Innodb_os_log_fsyncs                          | 4482911                                                                                           |
| Innodb_os_log_pending_fsyncs                  | 0                                                                                                 |
| Innodb_os_log_pending_writes                  | 0                                                                                                 |
| Innodb_os_log_written                         | 4772833792                                                                                        |
| Innodb_page_size                              | 16384                                                                                             |
| Innodb_pages_created                          | 53                                                                                                |
| Innodb_pages_read                             | 376                                                                                               |
| Innodb_pages_written                          | 7075851                                                                                           |
| Innodb_row_lock_current_waits                 | 0                                                                                                 |
| Innodb_row_lock_time                          | 0                                                                                                 |
| Innodb_row_lock_time_avg                      | 0                                                                                                 |
| Innodb_row_lock_time_max                      | 0                                                                                                 |
| Innodb_row_lock_waits                         | 0                                                                                                 |
| Innodb_rows_deleted                           | 1                                                                                                 |
| Innodb_rows_inserted                          | 101                                                                                               |
| Innodb_rows_read                              | 5008973                                                                                           |
| Innodb_rows_updated                           | 3761451                                                                                           |
| Innodb_num_open_files                         | 37                                                                                                |
| Innodb_truncated_status_writes                | 0                                                                                                 |
| Innodb_available_undo_logs                    | 128                                                                                               |
| Key_blocks_not_flushed                        | 0                                                                                                 |
| Key_blocks_unused                             | 6695                                                                                              |
| Key_blocks_used                               | 3                                                                                                 |
| Key_read_requests                             | 6                                                                                                 |
| Key_reads                                     | 3                                                                                                 |
| Key_write_requests                            | 0                                                                                                 |
| Key_writes                                    | 0                                                                                                 |
| Locked_connects                               | 0                                                                                                 |
| Max_execution_time_exceeded                   | 0                                                                                                 |
| Max_execution_time_set                        | 0                                                                                                 |
| Max_execution_time_set_failed                 | 0                                                                                                 |
| Max_used_connections                          | 3                                                                                                 |
| Max_used_connections_time                     | 2021-10-09 09:41:27                                                                               |
| Not_flushed_delayed_rows                      | 0                                                                                                 |
| Ongoing_anonymous_transaction_count           | 0                                                                                                 |
| Open_files                                    | 20                                                                                                |
| Open_streams                                  | 0                                                                                                 |
| Open_table_definitions                        | 110                                                                                               |
| Open_tables                                   | 112                                                                                               |
| Opened_files                                  | 216791                                                                                            |
| Opened_table_definitions                      | 110                                                                                               |
| Opened_tables                                 | 119                                                                                               |
| Performance_schema_accounts_lost              | 0                                                                                                 |
| Performance_schema_cond_classes_lost          | 0                                                                                                 |
| Performance_schema_cond_instances_lost        | 0                                                                                                 |
| Performance_schema_digest_lost                | 0                                                                                                 |
| Performance_schema_file_classes_lost          | 0                                                                                                 |
| Performance_schema_file_handles_lost          | 0                                                                                                 |
| Performance_schema_file_instances_lost        | 0                                                                                                 |
| Performance_schema_hosts_lost                 | 0                                                                                                 |
| Performance_schema_index_stat_lost            | 0                                                                                                 |
| Performance_schema_locker_lost                | 0                                                                                                 |
| Performance_schema_memory_classes_lost        | 0                                                                                                 |
| Performance_schema_metadata_lock_lost         | 0                                                                                                 |
| Performance_schema_mutex_classes_lost         | 0                                                                                                 |
| Performance_schema_mutex_instances_lost       | 0                                                                                                 |
| Performance_schema_nested_statement_lost      | 0                                                                                                 |
| Performance_schema_prepared_statements_lost   | 0                                                                                                 |
| Performance_schema_program_lost               | 0                                                                                                 |
| Performance_schema_rwlock_classes_lost        | 0                                                                                                 |
| Performance_schema_rwlock_instances_lost      | 0                                                                                                 |
| Performance_schema_session_connect_attrs_lost | 0                                                                                                 |
| Performance_schema_socket_classes_lost        | 0                                                                                                 |
| Performance_schema_socket_instances_lost      | 0                                                                                                 |
| Performance_schema_stage_classes_lost         | 0                                                                                                 |
| Performance_schema_statement_classes_lost     | 0                                                                                                 |
| Performance_schema_table_handles_lost         | 0                                                                                                 |
| Performance_schema_table_instances_lost       | 0                                                                                                 |
| Performance_schema_table_lock_stat_lost       | 0                                                                                                 |
| Performance_schema_thread_classes_lost        | 0                                                                                                 |
| Performance_schema_thread_instances_lost      | 0                                                                                                 |
| Performance_schema_users_lost                 | 0                                                                                                 |
| Prepared_stmt_count                           | 0                                                                                                 |
| Qcache_free_blocks                            | 0                                                                                                 |
| Qcache_free_memory                            | 0                                                                                                 |
| Qcache_hits                                   | 0                                                                                                 |
| Qcache_inserts                                | 0                                                                                                 |
| Qcache_lowmem_prunes                          | 0                                                                                                 |
| Qcache_not_cached                             | 0                                                                                                 |
| Qcache_queries_in_cache                       | 0                                                                                                 |
| Qcache_total_blocks                           | 0                                                                                                 |
| Queries                                       | 14360350                                                                                          |
| Questions                                     | 10811752                                                                                          |
| Rpl_semi_sync_master_clients                  | 0                                                                                                 |
| Rpl_semi_sync_master_net_avg_wait_time        | 0                                                                                                 |
| Rpl_semi_sync_master_net_wait_time            | 0                                                                                                 |
| Rpl_semi_sync_master_net_waits                | 0                                                                                                 |
| Rpl_semi_sync_master_no_times                 | 0                                                                                                 |
| Rpl_semi_sync_master_no_tx                    | 1578216                                                                                           |
| Rpl_semi_sync_master_status                   | OFF                                                                                               |
| Rpl_semi_sync_master_timefunc_failures        | 0                                                                                                 |
| Rpl_semi_sync_master_tx_avg_wait_time         | 0                                                                                                 |
| Rpl_semi_sync_master_tx_wait_time             | 0                                                                                                 |
| Rpl_semi_sync_master_tx_waits                 | 0                                                                                                 |
| Rpl_semi_sync_master_wait_pos_backtraverse    | 0                                                                                                 |
| Rpl_semi_sync_master_wait_sessions            | 0                                                                                                 |
| Rpl_semi_sync_master_yes_tx                   | 0                                                                                                 |
| Rpl_semi_sync_slave_status                    | ON                                                                                                |
| Select_full_join                              | 0                                                                                                 |
| Select_full_range_join                        | 0                                                                                                 |
| Select_range                                  | 0                                                                                                 |
| Select_range_check                            | 0                                                                                                 |
| Select_scan                                   | 2326848                                                                                           |
| Slave_open_temp_tables                        | 0                                                                                                 |
| Slow_launch_threads                           | 0                                                                                                 |
| Slow_queries                                  | 1                                                                                                 |
| Sort_merge_passes                             | 0                                                                                                 |
| Sort_range                                    | 0                                                                                                 |
| Sort_rows                                     | 54                                                                                                |
| Sort_scan                                     | 3                                                                                                 |
| Ssl_accept_renegotiates                       | 0                                                                                                 |
| Ssl_accepts                                   | 0                                                                                                 |
| Ssl_callback_cache_hits                       | 0                                                                                                 |
| Ssl_cipher                                    |                                                                                                   |
| Ssl_cipher_list                               |                                                                                                   |
| Ssl_client_connects                           | 0                                                                                                 |
| Ssl_connect_renegotiates                      | 0                                                                                                 |
| Ssl_ctx_verify_depth                          | 0                                                                                                 |
| Ssl_ctx_verify_mode                           | 0                                                                                                 |
| Ssl_default_timeout                           | 0                                                                                                 |
| Ssl_finished_accepts                          | 0                                                                                                 |
| Ssl_finished_connects                         | 0                                                                                                 |
| Ssl_server_not_after                          |                                                                                                   |
| Ssl_server_not_before                         |                                                                                                   |
| Ssl_session_cache_hits                        | 0                                                                                                 |
| Ssl_session_cache_misses                      | 0                                                                                                 |
| Ssl_session_cache_mode                        | NONE                                                                                              |
| Ssl_session_cache_overflows                   | 0                                                                                                 |
| Ssl_session_cache_size                        | 0                                                                                                 |
| Ssl_session_cache_timeouts                    | 0                                                                                                 |
| Ssl_sessions_reused                           | 0                                                                                                 |
| Ssl_used_session_cache_entries                | 0                                                                                                 |
| Ssl_verify_depth                              | 0                                                                                                 |
| Ssl_verify_mode                               | 0                                                                                                 |
| Ssl_version                                   |                                                                                                   |
| Table_locks_immediate                         | 2203300                                                                                           |
| Table_locks_waited                            | 0                                                                                                 |
| Table_open_cache_hits                         | 7212116                                                                                           |
| Table_open_cache_misses                       | 119                                                                                               |
| Table_open_cache_overflows                    | 0                                                                                                 |
| Tc_log_max_pages_used                         | 0                                                                                                 |
| Tc_log_page_size                              | 0                                                                                                 |
| Tc_log_page_waits                             | 0                                                                                                 |
| Threads_cached                                | 1                                                                                                 |
| Threads_connected                             | 2                                                                                                 |
| Threads_created                               | 3                                                                                                 |
| Threads_running                               | 1                                                                                                 |
| Uptime                                        | 1871068                                                                                           |
| Uptime_since_flush_status                     | 1871068                                                                                           |
+-----------------------------------------------+---------------------------------------------------------------------------------------------------+
369 rows in set (0.00 sec)
```

### Aborted_clients

 Aborted Clients表示由于客户端没有正确关闭连接而中止的连接数。 

当Aborted Clients增大的时候意味着有客户端成功建立连接，但是由于某些原因断开连接或者被终止了，这种情况一般发生在网络不稳定的环境中。主要的可能性有：

1. 客户端程序在退出之前未调用mysql_close（）正确关闭MySQL连接。
2. 客户端休眠的时间超过了系统变量wait_timeout和interactive_timeout的值，导致连接被MySQL进程终止。
3. 客户端程序在数据传输过程中突然结束。

###  Aborted Connect

Aborted Connect表示尝试连接到MySQL服务器失败的次数。这个状态变量可以结合host_cache表和其错误日志一起来分析问题。 引起这个状态变量激增的原因如下：

1. 客户端没有权限但是尝试访问MySQL数据库。
2. 客户端输入的密码有误。
3. A connection packet does not contain the right information.
4. 超过连接时间限制，主要是这个系统变量connect_timeout控制（mysql默认是10s，基本上，除非网络环境极端不好，一般不会超时。）

- **Audit_log_buffer_size_overflow**

### Binlog_cache_disk_use & Binlog_cache_use

如果开启binlog，那么binlog_cache_size用来在事务运行期间在内存中缓存binlog event。如果经常使用大事务应该加大这个缓存，避免过多的磁盘使用影响性能。

当binlog_cache_size不足以容纳所有的binlog event时，便转而使用临时文件来缓存binlog event。从Binlog_cache_use和Binlog_cache_disk_use可以看出是否使用了binlog cache或binlog 临时文件用于保存binlog event。

- **Binlog_stmt_cache_disk_use & Binlog_stmt_cache_use**

Binlog_cache_disk_use：（事务类）二进志日志缓存的已经存在硬盘的条数。

Binlog_cache_use：（事务类）二进制日志已缓存的条数（内存中）。 注意，这个不是容量，而是事务个数。每次有一条事务提交，都会有一次增加。

Binlog_stmt_cache_disk_use：（非事务类）二进志日志缓存的已经存在硬盘的条数。

Binlog_stmt_cache_use：（非事务类）二进制日志已缓存的条数（内存中）。非事务型的语句，都存在这儿，比如MYISAM引擎的表，插入记录就存在这儿。

### Bytes_received & Bytes_sent

Bytes_received：从所有客户端接收到的字节数。结合bytes sent, 可以作为数据库网卡吞吐量的评测指标,单位字节。

Bytes_sent：发送给所有客户端的字节数。结合bytes received,可以作为数据库网卡吞吐量的评测指标,单位字节。

### Com系列

Com_xxx 语句计数变量表示每个xxx 语句执行的次数。每类语句有一个状态变量。
例如：Com_delete和Com_insert分别统计DELETE 和INSERT语句执行的次数。

Com_xxx包括：

Com_admin_commands
Com_alter_db
Com_alter_db_upgrade
Com_alter_event
Com_alter_function
Com_alter_procedure
Com_alter_server
Com_alter_table
Com_alter_tablespace
Com_analyze
Com_assign_to_keycache
Com_begin
Com_binlog
Com_call_procedure
Com_change_db
Com_change_master
Com_check
Com_checksum
**Com_commit**：MySQL提交的事务数量,可以用来统计TPS(每秒事务数),计算公式：Com_commit/S+Com_rollback/S
Com_create_db
Com_create_event
Com_create_function
Com_create_index
Com_create_procedure
Com_create_server
Com_create_table
Com_create_trigger
Com_create_udf
Com_create_user
Com_create_view
Com_dealloc_sql
**Com_delete**：MySQL删除的数量,可以用来统计qps,计算公式：questions / uptime 或者基于com %计算：Com_select/s + Com_insert/s + Com_update/s + Com_delete/s
Com_delete_multi
Com_do
Com_drop_db
Com_drop_event
Com_drop_function
Com_drop_index
Com_drop_procedure
Com_drop_server
Com_drop_table
Com_drop_trigger
Com_drop_user
Com_drop_view
Com_empty_query
Com_execute_sql
Com_flush
Com_grant
Com_ha_close
Com_ha_open
Com_ha_read
Com_help
**Com_insert**：MySQL插入的数量,可以用来统计qps,qps计算公式：questions / uptime 或者基于com %计算：Com_select/s + Com_insert/s + Com_update/s + Com_delete/s
Com_insert_select
Com_install_plugin
Com_kill
Com_load
Com_lock_tables
Com_optimize
Com_preload_keys
Com_prepare_sql
Com_purge
Com_purge_before_date
Com_release_savepoint
Com_rename_table
Com_rename_user
Com_repair
Com_replace
Com_replace_select
Com_reset
Com_resignal
Com_revoke
Com_revoke_all
**Com_rollback**：MySQL回滚的事务数量,可以用来统计TPS(每秒事务数),计算公式：Com_commit/S+Com_rollback/S
Com_rollback_to_savepoint
Com_savepoint
**Com_select**：MySQL查询的数量,可以用来统计qps,qps计算公式：questions / uptime 或者基于com%计算：Com_select/s + Com_insert/s + Com_update/s + Com_delete/s
Com_set_option
Com_show_authors
Com_show_binlog_events
Com_show_binlogs
Com_show_charsets
Com_show_collations
Com_show_contributors
Com_show_create_db
Com_show_create_event
Com_show_create_func
Com_show_create_proc
Com_show_create_table
Com_show_create_trigger
Com_show_databases
Com_show_engine_logs
Com_show_engine_mutex
Com_show_engine_status
Com_show_errors
Com_show_events
Com_show_fields
Com_show_function_code
Com_show_function_status
Com_show_grants
Com_show_keys
Com_show_logs
Com_show_master_status
Com_show_new_master
Com_show_open_tables
Com_show_plugins
Com_show_privileges
Com_show_procedure_code
Com_show_procedure_status
Com_show_processlist
Com_show_profile
Com_show_profiles
Com_show_relaylog_events
Com_show_slave_hosts
Com_show_slave_status
Com_show_status
Com_show_storage_engines
Com_show_table_status
Com_show_tables
Com_show_triggers
Com_show_variables
Com_show_warnings
Com_signal
Com_slave_start
Com_slave_stop
Com_stmt_close
Com_stmt_execute
Com_stmt_fetch
Com_stmt_prepare
Com_stmt_reprepare
Com_stmt_reset
Com_stmt_send_long_data
Com_truncate
Com_uninstall_plugin
Com_unlock_tables
**Com_update**：MySQL更新的数量,可以用来统计qps,qps计算公式：questions / uptime 或者基于com_%计算：Com_select/s + Com_insert/s + Com_update/s + Com_delete/s
Com_update_multi
Com_xa_commit
Com_xa_end
Com_xa_prepare
Com_xa_recover
Com_xa_rollback
Com_xa_start

### Connections

试图连接到（不管是否成功）MySQL服务器的连接数。

Connection_errors_accept
Connection_errors_internal
Connection_errors_max_connections
Connection_errors_peer_address
Connection_errors_select
Connection_errors_tcpwrap

### Created_tmp_disk_tables

 MySQL服务在磁盘上创建的临时表数。 

### Created_tmp_files

 MySQL服务创建的临时文件文件数。

### Created_tmp_tables

 MySQL服务创建的临时表数。

比较理想的状态是：Created_tmp_disk_tables / Created_tmp_tables * 100% <= 25%

- max_heap_table_size

 只有大小为max_heap_table_size以下的临时表才能全部放内存，超过的就会用到硬盘临时表。 



### Innodb系列

- Innodb_buffer_pool_dump_status
- Innodb_buffer_pool_load_status
- Innodb_buffer_pool_resize_status
- Innodb_buffer_pool_pages_data
- Innodb_buffer_pool_bytes_data
- Innodb_buffer_pool_pages_dirty
- Innodb_buffer_pool_bytes_dirty
- Innodb_buffer_pool_pages_flushed
- Innodb_buffer_pool_pages_free
- Innodb_buffer_pool_pages_misc
- Innodb_buffer_pool_pages_total