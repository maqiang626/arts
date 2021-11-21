# Prometheus announces an Agent to address a new range of use cases		

https://www.cncf.io/blog/2021/11/16/prometheus-announces-an-agent-to-address-a-new-range-of-use-cases/



[CNCF 				Project				Blog Post](https://www.cncf.io/lf-author-category/project/)

Posted on 			November 16, 2021			 					 					

*Project post by Prometheus maintainers*

Prometheus, the leading open source monitoring solution hosted by  CNCF, has announced today a new operating mode: the Prometheus Agent.  This new way of working enables new workflows such as low-resources  environments, edge networks, and IoT. It uses significantly fewer  resources and is able to efficiently forward data to centralized remote  endpoints, all while using the stable code base millions of Prometheus  users rely upon.

Prometheus Agent is a specialized mode that focuses on three parts  that have made the success of Prometheus: service discovery, scraping,  and remote writing. Built into Prometheus itself, Prometheus Agent  behaves like a normal Prometheus server: it is a pull-based mechanism,  scraping metrics over HTTP and replicating the data to remote write  endpoints. 

Over the years, the Prometheus server has been used in many different situations. From traditional servers to huge cloud-native clusters. The default mode of data forwarding in Prometheus is federation. While this is reliable, it does not match all users’ operational needs. Prometheus introduced Prometheus Remote Write, allowing for other solutions to  aggregate into a global view. Of particular note are the CNCF sister  projects Cortex and Thanos.

Still, Prometheus itself remains a component in such setups, scraping and forwarding metrics. Our users have been successfully using  Prometheus in decentralized scenarios that report all or some of their  metrics to remote write endpoints. This works reliably and at scale, but at a cost: full Prometheus servers still have many capabilities not  required for forwarding, in particular full local storage.

In this new mode, the data is not available to query locally.  Instead, it can be forwarded to Prometheus or any other compliant remote write endpoints.

“Our agent removes the data immediately after successful writes”,  says Bartek Plotka, Principal Software Engineer at Red Hat. “This  enables Prometheus Agent to use only a fraction of resources Prometheus  would normally use in a similar situation. It is also a drop-in  replacement for Prometheus server mode, as the behavior, interfaces, and configuration are the same.”

It is worth noting that the new persistent buffering mechanism,  called Write-Ahead-Log (WAL) was heavily inspired by the existing  Prometheus TSDB WAL. It was implemented initially in Grafana Agent in  2020 and has been successfully battle-tested since then on numerous  deployments. Thanks to Robert Fratto, Grafana Agent Tech Lead at Grafana Labs, for the initial implementation and upstreaming the implementation to our main Prometheus repository and binary for native use and  upstream maintenance.

The Prometheus Agent is available in beta right now. Read more on [the Prometheus blog](https://prometheus.io/blog/).



------

Prometheus, yyds

使用学习中...

