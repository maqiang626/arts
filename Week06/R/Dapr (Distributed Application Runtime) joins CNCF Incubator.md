# Dapr (Distributed Application Runtime) joins CNCF Incubator	

https://www.cncf.io/blog/2021/11/03/dapr-distributed-application-runtime-joins-cncf-incubator/



[CNCF 				Staff				Blog Post](https://www.cncf.io/lf-author-category/staff/)

Posted on 			November 3, 2021			 					 					

The CNCF [Technical Oversight Committee](https://github.com/cncf/toc) (TOC) has voted to accept Dapr as a CNCF incubating project. 

[Dapr](https://dapr.io/) is a set of APIs that makes it  easy for developers to write distributed applications. Dapr runs as a  sidecar process next to an application, whether on Kubernetes or any  other environment, and gives developers a set of primitives in the form  of pub/sub, state management, secrets management, event triggers, and  service-to-service calls that are secured and reliable. With Dapr,  developers focus on building business logic instead of infrastructure.

“Hearing from developers how Dapr has accelerated their time to build scalable, distributed applications on Kubernetes and other hosting  platforms, and solve their business needs, has been tremendously  impactful for me,” said Mark Fussell, Dapr maintainer, and Steering  Committee member. “Now, with Dapr as part of CNCF, this enables  developers to more easily build, use, and adopt cloud native  technologies.”

The project was created in 2019 at Microsoft. Over time, many  community members have joined and contributed to the project, extending  it and helping it reach a stable 1.0 release in February 2021. Today,  the Dapr Steering and Technical Committee governs the project with  representatives from Alibaba, Intel, and Microsoft.

“I’m most proud of the larger Dapr community contributing new APIs  and building blocks to the project,” said Yaron Schneider, Dapr  maintainer, and Steering Committee member. “We had contributions over  all of the project’s 20+ repositories, from our developer tooling and  SDKs to the runtime itself. Seeing developers come to the project and  suggest new APIs that help address distributed systems challenges is a  significant accomplishment of the Dapr community.”

Dapr integrates with several CNCF projects. For example, it uses gRPC for internal sidecar communications, creates SPIFFIE identities for ACLs, emits telemetry in OpenTelemetry format,  uses Prometheus for metrics gathering, leverages CloudEvents as the  pub/sub message format, and runs natively on Kubernetes using an  operator.

The project is used in production by organizations including Alibaba  Cloud, Bosch, Legentic, Tdcare, Tencent, Swoop Funding, Man Group,  Zeiss, and [many more.](https://github.com/dapr/community/blob/master/ADOPTERS.md) Adopters run Dapr on all major cloud vendors as well as on on-premises environments.

“At Alibaba Cloud, we believe that Dapr will lead the way in  microservice development. By adopting Dapr, our customers can build  portable and robust distributed systems faster,” said Li Xiang, senior  staff engineer at Alibaba Cloud.

“Using Dapr makes it very easy to bolt in new pieces of  infrastructure without changing anything else. It changed our business,” said Russell Stather, chief digital transformation officer at Ignition  Group.

“In our multi-cloud environment, Dapr gave us the flexibility we  needed,” said Kai Walter, lead architect at Zeiss. “It provides a layer  of abstraction that allows the developers to focus on the business case  at hand.” 

**Main Components:**

- **Dapr sidecar** **–** runs next to applications and contains developer-facing APIs. 
- **CLI and SDKs** **–** make up the developer tooling experience for the project. 
- **Components-contrib repository –** developers can extend Dapr to integrate and support a diverse set of cloud services and open-source technologies.

**Notable Milestones:**

- 15.1k GitHub Stars 
- 1,940 pull requests 
- 3,703 issues
- 1.3k contributors
- 14 Releases at stable v1.4
- 26M Docker pulls

“Distributed applications and microservices form the basis for  containers and cloud native, but writing distributed applications that  are scalable and reliable can be incredibly difficult,” said Chris  Aniszczyk, CTO of CNCF. “Dapr integrates well with other CNCF projects  and provides best practices that developers can build on top of using  any language or framework. We’re excited to welcome Dapr to the CNCF and work to cultivate their community.” 

The Dapr [project roadmap](https://github.com/dapr/dapr/issues/3335) includes the addition of a new Configuration API that makes it easier  for developers to manage configuration for their applications and get  notified whenever configurations change, as well as a Query API that  makes it easier for developers to query and filter data in Dapr state  stores. In addition, the project is looking to add support for gRPC and  WASM-based components that’ll allow dynamic discoverability of state  stores, pub/sub brokers, bindings, and other Dapr components. Finally,  new Concurrency APIs will unblock scenarios such as leader election are  also being discussed in the Dapr community.

As a CNCF-hosted project, Dapr is part of a neutral foundation  aligned with its technical interests, as well as the larger Linux  Foundation, which provides governance, marketing support, and community  outreach. Dapr joins incubating technologies Argo, Buildpacks, Cilium,  CloudEvents, CNI, Contour, Cortex, CRI-O, Crossplane, Dragonfly,  emissary-ingress, Falco, Flagger, Flux, gRPC, KEDA, KubeEdge, Longhorn,  NATS, Notary, OpenTelemetry, Operator Framework, SPIFFE, SPIRE, and  Thanos. For more information on maturity requirements for each level,  please visit the [CNCF Graduation Criteria](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc).



------

分布式应用，值得学习！





