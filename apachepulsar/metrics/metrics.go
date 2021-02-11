package metrics

import (
	"github.com/elastic/beats/v7/metricbeat/helper/prometheus"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

func init(){
    mapping := &prometheus.MetricsMapping{
        Metrics: map[string]prometheus.MetricMap{
	    //namespace
	    "pulsar_entry_size_le_128":				  prometheus.Metric("namespace.entry_size_le_128"),
	    "pulsar_entry_size_le_512":				  prometheus.Metric("namespace.entry_size_le_512"),
	    "pulsar_entry_size_le_1_kb":			  prometheus.Metric("namespace.entry_size_le_1_kb"),
	    "pulsar_entry_size_le_2_kb":			  prometheus.Metric("namespace.entry_size_le_2_kb"),
	    "pulsar_entry_size_le_4_kb":			  prometheus.Metric("namespace.entry_size_le_4_kb"),
	    "pulsar_entry_size_le_16_kb":			  prometheus.Metric("namespace.entry_size_le_16_kb"),
	    "pulsar_entry_size_le_100_kb":			  prometheus.Metric("namespace.entry_size_le_100_kb"),
	    "pulsar_entry_size_le_1_mb":			  prometheus.Metric("namespace.entry_size_le_1_mb"),
	    "pulsar_entry_size_le_overflow":			  prometheus.Metric("namespace.entry_size_le_overflow"),
	    "pulsar_producers_count":				  prometheus.Metric("namespace.producers_count"),
	    "pulsar_rate_in":					  prometheus.Metric("namespace.rate_in"),
	    "pulsar_rate_out":					  prometheus.Metric("namespace.rate_out"),
	    "pulsar_replication_backlog":			  prometheus.Metric("namespace.replication_backlog"),
	    "pulsar_replication_rate_in":			  prometheus.Metric("namespace.replication_rate_in"),
	    "pulsar_replication_rate_out":			  prometheus.Metric("namespace.replication_rate_out"),
	    "pulsar_replication_throughput_in":		          prometheus.Metric("namespace.replication_throughput_in"),
	    "pulsar_replication_throughput_out":		  prometheus.Metric("namespace.replication_throughput_out"),
	    "pulsar_subscriptions_count":			  prometheus.Metric("namespace.subscriptions_count"),
	    "pulsar_throughput_in":				  prometheus.Metric("namespace.throughput_in"),
	    "pulsar_throughput_out":				  prometheus.Metric("namespace.throughput_out"),
	    "pulsar_topics_count":				  prometheus.Metric("namespace.topics_count"),
	    "pulsar_consumers_count":				  prometheus.Metric("namespace.consumers_count"),

	    //topic
	    "pulsar_in_bytes_total":				  prometheus.Metric("topic.in_bytes_total"),
	    "pulsar_in_messages_total":			          prometheus.Metric("topic.in_messages_total"),
	    "pulsar_out_messages_total":			  prometheus.Metric("topic.out_messages_total"),
	    "pulsar_out_bytes_total":				  prometheus.Metric("topic.out_bytes_total"),
	    "pulsar_storage_backlog_quota_limit":		  prometheus.Metric("topic.storage_backlog_quota_limit"),
	    "pulsar_storage_backlog_size":			  prometheus.Metric("topic.storage_backlog_size"),
	    "pulsar_storage_read_rate":				  prometheus.Metric("topic.storage_read_rate"),
	    "pulsar_storage_offloaded_size":			  prometheus.Metric("topic.storage_offloaded_size"),
	    "pulsar_storage_size":				  prometheus.Metric("topic.storage_size"),
	    "pulsar_storage_write_latency_le_0_5":		  prometheus.Metric("topic.storage_write_latency_le_0_5"),
	    "pulsar_storage_write_latency_le_1":		  prometheus.Metric("topic.storage_write_latency_le_1"),
	    "pulsar_storage_write_latency_le_5":		  prometheus.Metric("topic.storage_write_latency_le_5"),
	    "pulsar_storage_write_latency_le_10":		  prometheus.Metric("topic.storage_write_latency_le_10"),
	    "pulsar_storage_write_latency_le_20":		  prometheus.Metric("topic.storage_write_latency_le_20"),
	    "pulsar_storage_write_latency_le_50":		  prometheus.Metric("topic.storage_write_latency_le_50"),
	    "pulsar_storage_write_latency_le_100":		  prometheus.Metric("topic.storage_write_latency_le_100"),
	    "pulsar_storage_write_latency_le_200":		  prometheus.Metric("topic.storage_write_latency_le_200"),
	    "pulsar_storage_write_latency_le_1000":		  prometheus.Metric("topic.storage_write_latency_le_1000"),
	    "pulsar_storage_write_latency_le_overflow":		  prometheus.Metric("topic.storage_write_latency_le_overflow"),
	    "pulsar_storage_write_rate":			  prometheus.Metric("topic.storage_write_rate"),

	    //subscription
	    "pulsar_subscription_back_log":			  prometheus.Metric("subscription.back_log"),
	    "pulsar_subscription_blocked_on_unacked_messages":    prometheus.Metric("subscription.blocked_on_unacked_messages"),
	    "pulsar_subscription_delayed":			  prometheus.Metric("subscription.delayed"),
	    "pulsar_subscription_msg_rate_out":			  prometheus.Metric("subscription.msg_rate_out"),
	    "pulsar_subscription_msg_rate_redeliver":		  prometheus.Metric("subscription.msg_rate_redeliver"),
	    "pulsar_subscription_msg_throughput_out":		  prometheus.Metric("subscription.msg_throughput_out"),
	    "pulsar_subscription_unacked_message":		  prometheus.Metric("subscription.unacked_message"),

	    //consumer
	    "pulsar_consumer_available_permits":		  prometheus.Metric("consumer.available_permits"),
	    "pulsar_consumer_msg_throughput_out":		  prometheus.Metric("consumer.msg_throughput_out"),
	    "pulsar_consumer_msg_rate_out":			  prometheus.Metric("consumer.msg_rate_out"),
	    "pular_consumer_blocked_on_unacked_messages":	  prometheus.Metric("consumer.blocked_on_unacked_messages"),
	    "pulsar_consumer_unacked_messages":			  prometheus.Metric("consumer.unacked_messages"),
	    "pulsar_consumer_msg_rate_redeliver":		  prometheus.Metric("consumer.msg_rate_redeliver"),
         },
	 ExtraFields: map[string]string{"api_version": "2"},
	 Namespace:   "apachepulsar",
    }

    mb.Registry.MustAddMetricSet("apachepulsar", "metrics", prometheus.MetricSetBuilder(mapping), mb.WithHostParser(prometheus.HostParser))
}
