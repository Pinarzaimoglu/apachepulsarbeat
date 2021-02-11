package metrics

import (
	"github.com/elastic/beats/v7/metricbeat/helper/prometheus"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

func init(){
    mapping := &prometheus.MetricsMapping{
        Metrics: map[string]prometheus.MetricMap{
	    //namespace
	    "pulsar_entry_size_le_128":				  prometheus.Metric("entry_size_le_128"),
	    "pulsar_entry_size_le_512":				  prometheus.Metric("entry_size_le_512"),
	    "pulsar_entry_size_le_1_kb":			  prometheus.Metric("entry_size_le_1_kb"),
	    "pulsar_entry_size_le_2_kb":			  prometheus.Metric("entry_size_le_2_kb"),
	    "pulsar_entry_size_le_4_kb":			  prometheus.Metric("entry_size_le_4_kb"),
	    "pulsar_entry_size_le_16_kb":			  prometheus.Metric("entry_size_le_16_kb"),
	    "pulsar_entry_size_le_100_kb":			  prometheus.Metric("entry_size_le_100_kb"),
	    "pulsar_entry_size_le_1_mb":			  prometheus.Metric("entry_size_le_1_mb"),
	    "pulsar_entry_size_le_overflow":			  prometheus.Metric("entry_size_le_overflow"),
	    "pulsar_producers_count":				  prometheus.Metric("producers_count"),
	    "pulsar_rate_in":					  prometheus.Metric("rate_in"),
	    "pulsar_rate_out":					  prometheus.Metric("rate_out"),
	    "pulsar_replication_backlog":			  prometheus.Metric("replication_backlog"),
	    "pulsar_replication_rate_in":			  prometheus.Metric("replication_rate_in"),
	    "pulsar_replication_rate_out":			  prometheus.Metric("replication_rate_out"),
	    "pulsar_replication_throughput_in":		          prometheus.Metric("replication_throughput_in"),
	    "pulsar_replication_throughput_out":		  prometheus.Metric("replication_throughput_out"),
	    "pulsar_subscriptions_count":			  prometheus.Metric("subscriptions_count"),
	    "pulsar_throughput_in":				  prometheus.Metric("throughput_in"),
	    "pulsar_throughput_out":				  prometheus.Metric("throughput_out"),
	    "pulsar_topics_count":				  prometheus.Metric("topics_count"),
	    "pulsar_consumers_count":				  prometheus.Metric("consumers_count"),

	    //topics
	    "pulsar_in_bytes_total":				  prometheus.Metric("in_bytes_total"),
	    "pulsar_in_messages_total":			          prometheus.Metric("in_messages_total"),
	    "pulsar_out_messages_total":			  prometheus.Metric("out_messages_total"),
	    "pulsar_out_bytes_total":				  prometheus.Metric("out_bytes_total"),
	    "pulsar_storage_backlog_quota_limit":		  prometheus.Metric("storage_backlog_quota_limit"),
	    "pulsar_storage_backlog_size":			  prometheus.Metric("storage_backlog_size"),
	    "pulsar_storage_read_rate":				  prometheus.Metric("storage_read_rate"),
	    "pulsar_storage_offloaded_size":			  prometheus.Metric("storage_offloaded_size"),
	    "pulsar_storage_size":				  prometheus.Metric("storage_size"),
	    "pulsar_storage_write_latency_le_0_5":		  prometheus.Metric("storage_write_latency_le_0_5"),
	    "pulsar_storage_write_latency_le_1":		  prometheus.Metric("storage_write_latency_le_1"),
	    "pulsar_storage_write_latency_le_5":		  prometheus.Metric("storage_write_latency_le_5"),
	    "pulsar_storage_write_latency_le_10":		  prometheus.Metric("storage_write_latency_le_10"),
	    "pulsar_storage_write_latency_le_20":		  prometheus.Metric("storage_write_latency_le_20"),
	    "pulsar_storage_write_latency_le_50":		  prometheus.Metric("storage_write_latency_le_50"),
	    "pulsar_storage_write_latency_le_100":		  prometheus.Metric("storage_write_latency_le_100"),
	    "pulsar_storage_write_latency_le_200":		  prometheus.Metric("storage_write_latency_le_200"),
	    "pulsar_storage_write_latency_le_1000":		  prometheus.Metric("storage_write_latency_le_1000"),
	    "pulsar_storage_write_latency_le_overflow":		  prometheus.Metric("storage_write_latency_le_overflow"),
	    "pulsar_storage_write_rate":			  prometheus.Metric("storage_write_rate"),

	    //subscription
	    "pulsar_subscription_back_log":			  prometheus.Metric("subscription_back_log"),
	    "pulsar_subscription_blocked_on_unacked_messages":    prometheus.Metric("subscription_blocked_on_unacked_messages"),
	    "pulsar_subscription_delayed":			  prometheus.Metric("subscription_delayed"),
	    "pulsar_subscription_msg_rate_out":			  prometheus.Metric("subscription_msg_rate_out"),
	    "pulsar_subscription_msg_rate_redeliver":		  prometheus.Metric("subscription_msg_rate_redeliver"),
	    "pulsar_subscription_msg_throughput_out":		  prometheus.Metric("subscription_msg_throughput_out"),
	    "pulsar_subscription_unacked_message":		  prometheus.Metric("subscription_unacked_message"),

	    //consumer
	    "pulsar_consumer_available_permits":		  prometheus.Metric("consumer_available_permits"),
	    "pulsar_consumer_msg_throughput_out":		  prometheus.Metric("consumer_msg_throughput_out"),
	    "pulsar_consumer_msg_rate_out":			  prometheus.Metric("consumer_msg_rate_out"),
	    "pular_consumer_blocked_on_unacked_messages":	  prometheus.Metric("consumer_blocked_on_unacked_messages"),
	    "pulsar_consumer_unacked_messages":			  prometheus.Metric("consumer_unacked_messages"),
	    "pulsar_consumer_msg_rate_redeliver":		  prometheus.Metric("consumer_msg_rate_redeliver"),
         },
	 ExtraFields: map[string]string{"api_version": "2"},
	 Namespace:   "apachepulsar",
    }

    mb.Registry.MustAddMetricSet("apachepulsar", "metrics", prometheus.MetricSetBuilder(mapping), mb.WithHostParser(prometheus.HostParser))
}
