# K8s 原理：Kubernetes 架构解析

只要您稍微了解 [Kubernetes 的基础知识](https://www.redhat.com/zh/topics/containers/what-is-kubernetes)，您就知道它是一个用于大规模运行分布式应用和服务的[开源](https://www.redhat.com/zh/topics/open-source/what-is-open-source)容器编排平台。但是，您可能并不了解它的组件以及这些组件的交互方式。

不妨让我们来简要了解一下 Kubernetes 的设计原理，然后探讨 Kubernetes 的不同组件是如何协同工作的。

## K8s 设计原理是什么？

正如 [Kubernetes 实施细节](https://kubernetes.io/docs/reference/setup-tools/kubeadm/implementation-details/)中所述，Kubernetes 集群的设计基于 3 个原则。

Kubernetes 集群应做到：

- **安全。**它应遵循最新的安全最佳实践。
- **易于使用。**它应能通过一些简单的命令进行操作。 
- **可扩展。**不应偏向于某一个提供商，而是能通过配置文件进行自定义。

[利用 Kubernetes 模式创建云原生应用](https://www.redhat.com/zh/topics/cloud-native-apps)

## K8s 部署需要哪些组件？

我们把一个有效的 Kubernetes 部署称为[集群](https://www.redhat.com/zh/topics/containers/what-is-a-kubernetes-cluster)。您可以将 Kubernetes 集群可视化为两个部分：控制平面与计算设备（或称为节点）。每个节点都是其自己的 [Linux®](https://www.redhat.com/zh/topics/linux) 环境，并且可以是物理机或[虚拟机](https://www.redhat.com/zh/topics/virtualization/what-is-a-virtual-machine)。每个节点都运行由[若干容器](https://www.redhat.com/zh/topics/containers)组成的[容器集](https://www.redhat.com/zh/topics/containers/what-is-kubernetes-pod)。

## K8s 集群架构图

以下 K8s 架构图显示了 Kubernetes 集群的各部分之间的联系： 



![k8s 架构图](https://www.redhat.com/cms/managed-files/kubernetes_diagram-v3-770x717_0.svg)



## K8s 架构原理是什么？ 

### K8s 集群的神经中枢：控制平面

让我们从 Kubernetes 集群的神经中枢（即控制平面）开始说起。在这里，我们可以找到用于控制集群的 Kubernetes  组件以及一些有关集群状态和配置的数据。这些核心 Kubernetes 组件负责处理重要的工作，以确保容器以足够的数量和所需的资源运行。 

控制平面会一直与您的计算机保持联系。集群已被配置为以特定的方式运行，而控制平面要做的就是确保万无一失。

### K8s 集群API: kube-apiserver

如果需要与您的 Kubernetes 集群进行交互，就要通过 API。[Kubernetes API](https://www.redhat.com/zh/topics/containers/what-is-the-kubernetes-API) 是 Kubernetes 控制平面的前端，用于处理内部和外部请求。API 服务器会确定请求是否有效，如果有效，则对其进行处理。您可以通过 REST 调用、kubectl 命令行界面或其他命令行工具（例如 kubeadm）来访问 API。

### K8s 调度程序：kube-scheduler

您的集群是否状况良好？如果需要新的容器，要将它们放在哪里？这些是 Kubernetes 调度程序所要关注的问题。

调度程序会考虑容器集的资源需求（例如 CPU 或内存）以及集群的运行状况。随后，它会将容器集安排到适当的计算节点。

### K8s 控制器：kube-controller-manager

控制器负责实际运行集群，而 Kubernetes  控制器管理器则是将多个控制器功能合而为一。控制器用于查询调度程序，并确保有正确数量的容器集在运行。如果有容器集停止运行，另一个控制器会发现并做出响应。控制器会将服务连接至容器集，以便让请求前往正确的端点。还有一些控制器用于创建帐户和 API 访问令牌。

### 键值存储数据库 etcd

配置数据以及有关集群状态的信息位于 [etcd](https://www.redhat.com/zh/topics/containers/what-is-etcd)（一个键值存储数据库）中。etcd 采用分布式、容错设计，被视为集群的最终事实来源。

## Kubernetes 节点中会发生什么？

### K8s 节点 

Kubernetes 集群中至少需要一个计算节点，但通常会有多个计算节点。容器集经过调度和编排后，就会在节点上运行。如果需要扩展集群的容量，那就要添加更多的节点。

### 容器集

容器集是 Kubernetes 对象模型中最小、最简单的单元。它代表了[应用](https://www.redhat.com/zh/topics/cloud-native-apps)的单个实例。每个容器集都由一个容器（或一系列紧密耦合的容器）以及若干控制容器运行方式的选件组成。容器集可以连接至持久存储，以运行有状态应用。

### 容器运行时引擎

为了运行容器，每个计算节点都有一个容器运行时引擎。比如 [Docker](https://www.redhat.com/zh/topics/containers/what-is-docker)，但 Kubernetes 也支持其他符合开源容器运动（OCI）标准的运行时，例如 rkt 和 CRI-O。

### kubelet

每个计算节点中都包含一个 kubelet，这是一个与控制平面通信的微型应用。kublet 可确保容器在容器集内运行。当控制平面需要在节点中执行某个操作时，kubelet 就会执行该操作。

### kube-proxy

每个计算节点中还包含 kube-proxy，这是一个用于优化 Kubernetes 网络服务的网络代理。kube-proxy 负责处理集群内部或外部的网络通信——靠操作系统的数据包过滤层，或者自行转发流量。

## K8s 集群部署还需要什么？

### 持久存储

除了管理运行应用的容器外，Kubernetes 还可以管理附加在集群上的应用数据。Kubernetes 允许用户请求存储资源，而无需了解底层存储基础架构的详细信息。持久卷是集群（而非容器集）所特有的，因此其寿命可以超过容器集。

### 容器镜像仓库

Kubernetes 所依赖的容器镜像存储于容器镜像仓库中。这个镜像仓库可以由您自己配置的，也可以由第三方提供。

### 底层基础架构

您可以自己决定具体在哪里运行 Kubernetes。答案可以是裸机服务器、虚拟机、公共云提供商、私有云和[混合云](https://www.redhat.com/zh/topics/cloud-computing/what-is-hybrid-cloud)环境。Kubernetes 的一大优势就是它可以在许多不同类型的基础架构上运行。

## K8s 集群部署方案：知易行难，直面挑战

这份有关 Kubernetes 架构的概述只不过是走马观花。一旦您开始考虑这些组件之间以及与外部资源和基础架构如何通信时，您就会意识到在配置和[保护 Kubernetes 集群](https://www.redhat.com/zh/topics/security/container-security)方面所面临的挑战。

Kubernetes 提供了编排复杂的大型容器化应用所需的工具，但也为您留出了许多决策空间。您要选择操作系统、容器运行时、[持续集成/持续交付（CI/CD）](https://www.redhat.com/zh/topics/devops/what-is-ci-cd)工具、应用服务、存储以及其他大多数组件。另外，还有管理角色、访问控制、多租户和安全默认设置等工作需要您来完成。不仅如此，您也可以选择自行运行 Kubernetes，或是与能够提供受支持版本的供应商展开合作。

这种选择的自由度正是 Kubernetes 灵活性的一大表现。尽管实施起来可能比较复杂，但 Kubernetes 赋予了您强大的功能，可以按照自己的条件运行容器化应用，并以敏捷的方式对组织的变化做出回应。

### 利用 Kubernetes 创建云原生应用

观看本网络研讨会系列，获取帮助您在需要构建、运行、部署和现代化应用的企业 Kubernetes 上建立数据平台的专家视角。 