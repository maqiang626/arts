# Service mesh 101: the role of Envoy		

https://www.cncf.io/blog/2021/10/22/service-mesh-101-the-role-of-envoy/



[CNCF 				Member				Blog Post](https://www.cncf.io/lf-author-category/member/)

Posted on 			October 22, 2021			 		By Scott Lowe			 					

*Guest post originally published on [Kong’s blog](https://konghq.com/blog/envoy-service-mesh/) by Scott Lowe*

If you’ve done any reading about service meshes, you’ve probably come across mentions of an open source project named Envoy. And if you’ve  done any reading about Envoy, you’ve probably seen references to service meshes. How are these two technologies related? How are they different? Do they work together? I’ll attempt to answer all those questions in  this blog post’s first and [second](https://konghq.com/blog/envoy-service-mesh-configuration/) parts, plus possibly a few more.

## **What Is a Service Mesh?**

As companies are increasingly re-architecting their applications and embracing a [microservices](https://konghq.com/learning-center/microservices/)-based approach, the need for solutions to traffic management, observability, security and reliability features increases. A *service mesh* is one approach to adding these features to the underlying platform instead of individual applications or services.

What kinds of traffic management, observability, security and  reliability features are we talking about here? A few examples include:

- *Enforcing mutual TLS (*[*mTLS*](https://kuma.io/docs/1.2.0/policies/mutual-tls/)*) between services on the mesh:* This isn’t just about encryption, which lower-level network constructs could also handle. It’s also about authentication (ensuring that the service  claiming to be Service A *is* Service A) and authorization (ensuring that Service A is *allowed* to connect to some other service).
- *Enabling sophisticated traffic patterns:* A service mesh can perform more advanced traffic routing and steering, simplifying things like [canary deployments](https://konghq.com/blog/canary-deployment-5-minutes-service-mesh/) or rerouting around failed instances of a particular service by watching for upstream errors.
- *Enabling traffic tracing:* Because of how a service mesh operates (more on that in a moment) it can expose more advanced metrics about traffic patterns between services on the  mesh than might have otherwise been possible. This enables organizations to more easily identify, and possibly correct, issues affecting  application performance.

A service mesh can do these things without modifications to the  underlying applications or services. This is an important  point—offloading this functionality to the service mesh means that  organizations can focus on their applications instead.

## **What Makes Up a Service Mesh?**

While the details of a specific implementation may vary, at a high level, all service meshes consist of two basic components:

- The *control plane* interfaces over a communications  channel with one or more data planes. The control plane is where  policies are defined; these policies control things like mTLS, traffic  routing or [rate limiting](https://kuma.io/docs/1.2.0/policies/rate-limit/).

The *data planes* sit in front of services or workloads that  communicate through those data planes with other services or workloads.  It’s the responsibility of the data planes to enforce the policies  defined on and passed down by the control plane.

### ***The control plane and data plane together make up a service mesh.\***

![envoy-service-mesh-architecture](https://2tjosk2rxzc21medji3nfn1g-wpengine.netdna-ssl.com/wp-content/uploads/2021/08/control-panel.png.webp)

Figure 1: A generic service mesh architecture

### **More About the Control Plane**

A service mesh’s control plane is responsible for “command and  control” functions. Some typical functions of the control plane include:

- The control plane integrates with other systems, like [Kubernetes](https://konghq.com/learning-center/kubernetes/what-is-kubernetes), for service discovery (figuring out what services are on the mesh) and gathering configuration details.
- It manages the service mesh configuration, including all the policies for  how traffic should move between and among services on the mesh.
- The control plane programs or configures the data planes to enforce policies.

It’s important to note that the control plane is out-of-band, so it  doesn’t sit in-line between two communicating workloads. This ensures  that the control plane isn’t a bottleneck and that an outage of the  control plane doesn’t disrupt traffic.

The service mesh doesn’t redirect all the traffic up through the  control plane. Instead, data planes are distributed across the mesh and  placed close to each of the workloads. They’re only responsible for  managing the traffic moving in and out of that particular workload and  directly scale with the number of workloads deployed in the environment.

### **Digging into Data Planes**

Unlike the control plane, a data plane is in-band; it is responsible  for actually passing traffic into or out of a workload. This is the  piece that actually moves packets to and from the workloads. Because it  is in-band, it must be lightweight (not consume a lot of computing  resources) and be very performant (not introduce significant latency or  delays in the traffic it processes).

The data plane’s in-band position allows it to expose more detailed  metrics about a workload’s network traffic; the data planes see *all* the traffic moving into or out of a workload. This enables observability  functionality that might otherwise be difficult to achieve.

As an aside, some folks use “the data plane” (singular) to refer to  the entity where traffic passes back and forth between and among  services or workloads. Here we use the term “data planes” (plural) to  refer to the collection of smaller entities responsible for managing the traffic in and out of a single workload only.

## **Where Does Envoy Fit In?**

Originally created at Lyft, Envoy is a high-performance data plane  designed for service mesh architectures. Lyft open sourced it and  donated it to the CNCF, where it is now one of the CNCF’s graduated open source projects. 

What makes Envoy so well-suited for service mesh use cases? It offers several features that are very useful as a service mesh data plane.

- *Envoy can be dynamically configured.* You don’t just  pass it as a configuration file and start it up (although you can do  that, too). You can pass its configuration values dynamically over an  API.
- *Envoy has multiple load balancing algorithms.* This allows it to support a variety of traffic patterns and a wider range of applications.
- *Envoy supports retries and circuit breaking.* Envoy can retry requests, and if the upstream service returns enough errors,  Envoy can “break the circuit.” This can help keep cascading failures  from affecting the rest of your distributed, microservices-based  application or system.
- *Envoy also has first-class support for the latest protocols.* This includes protocols like HTTP/2, HTTP/3 and gRPC.

Many different service meshes use Envoy. For example, Kong’s open source project [Kuma](https://kuma.io/)—and its enterprise counterpart [Kong Mesh](https://konghq.com/kong-mesh)—use Envoy for the data planes.

## **More Resources**

Now that you understand the basics, look at these helpful resources  for implementing an Envoy-based service mesh. And look check out [part 2 of the series here](https://konghq.com/blog/envoy-service-mesh-configuration/)!

- [Getting Started With Kuma Service Mesh](https://konghq.com/blog/getting-started-kuma-service-mesh/)
- [Implementing Traffic Policies in Kubernetes](https://konghq.com/blog/traffic-policies-kubernetes/)
- [Automate Service Mesh Observability With Kuma](https://konghq.com/blog/service-mesh-observability)
- [Zero-Trust Service Mesh Security](https://konghq.com/blog/zero-trust-service-mesh-security/)
- [Zero Load Balancers with Kuma/Kong Mesh](https://konghq.com/zerolb/)



------

服务网格，101，开始深入学习中...