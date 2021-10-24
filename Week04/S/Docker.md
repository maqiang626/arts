# Docker 

https://www.ibm.com/cn-zh/cloud/learn/docker        

Docker 是一个开源平台，用于构建、部署和管理容器化应用程序。  了解容器、它们与虚拟机的区别，以及 Docker 应用如此之广的原因。   

## 什么是 Docker？

Docker 是一个开源的[容器化](https://www.ibm.com/cloud/learn/containerization)平台。 开发人员可使用 Docker 将应用程序打包到容器中；容器是标准化的可执行组件，可以将应用程序源代码与在任何环境中运行该代码所需的操作系统 (OS) 库和依赖项相结合。 容器简化了分布式应用程序的交付，并且随着组织转向云原生开发和混合[多云](https://www.ibm.com/cloud/learn/multicloud)环境而变得越来越流行。

开发人员可以在不使用 Docker 平台的情况下创建容器，但该平台可以使构建、部署和管理容器的过程变得更便捷且更安全。 Docker 本质上是一个工具包，支持开发人员通过单个 API，使用简单的命令和省力的自动化技术来构建、部署、运行、更新和停止容器。

Docker 可以指代 [Docker, Inc.](https://www.docker.com/company)（链接位于 IBM 外部），这是销售商业版 Docker 的公司，也可以指代 [Docker 开源项目](https://github.com/docker)（链接位于 IBM 外部），Docker, Inc. 及许多其他组织和个人都为此项目做出了贡献 。

## 容器如何工作以及容器为何如此受欢迎

[容器](https://www.ibm.com/cn-zh/cloud/learn/containers)是通过 Linux 内核中内置的过程隔离和虚拟化功能来实现。 这些功能 - 例如 *控制组*（即 Cgroup，用于在进程之间分配资源）以及*名称空间*（用于限制进程访问或系统的其他资源或区域的可视性）- 使多个应用程序组件能够共享主机操作系统的单个实例的资源，这与管理程序使多个[虚拟机 (VM)](https://www.ibm.com/cn-zh/cloud/learn/virtual-machines) 共享单个硬件服务器的 CPU、内存和其他资源的方式大致相同。

因此，容器技术不仅提供了虚拟机的所有功能和优势（包括应用程序隔离、经济高效的可扩展性和可处置性）， 而且还提供了其他的重要优势：

- **更轻巧：**与虚拟机不同，容器不会承载整个操作系统实例和管理程序的有效负载； 它们仅包括执行代码所需的操作系统进程和依赖项。 容器大小以兆字节为单位（某些虚拟机则是以千兆字节为单位），因此它们可以更好地利用硬件容量，并具有更快的启动时间。
- **提高资源效率：**在同一硬件上使用容器运行的应用程序副本数是使用虚拟机时的数倍。 这可以减少云支出。
- **提高开发人员工作效率：**与虚拟机相比，容器的部署、配置和重启过程更快且更简单。 这使得容器非常适合在[持续集成](https://www.ibm.com/cn-zh/cloud/learn/continuous-integration)和[持续交付](https://www.ibm.com/cloud/learn/continuous-delivery) (CI/CD) 管道中使用，并且更适合采取敏捷和 [DevOps](https://www.ibm.com/cn-zh/cloud/learn/devops-a-complete-guide) 实践的开发团队。

使用容器的公司还报告了一些其他好处，包括提高应用程序质量、更快地应对市场变化等等。 了解有关此交互式工具的更多信息：

## topic.topicOverviewTitle

topic.topicOverviewSubTitle





















[**下载完整报告：\*企业中的容器\***](https://www.ibm.com/downloads/cas/VG8KRPRM) (PDF, 1.4MB)

## 为何选择使用 Docker？

Docker 目前非常受欢迎，以至于“Docker”和“容器”可以互换使用。 但是在 2013 年 Docker 面世之前，与容器相关的首批技术已经出现了数年[甚至是数十年](https://blog.aquasec.com/a-brief-history-of-containers-from-1970s-chroot-to-docker-2016)（链接位于 IBM 外部）之久。

最值得注意的是，在 2008 年，Linux 内核中实现了 **L**inu**XC**ontainers (LXC)，LXC 完全支持单个 Linux 实例的虚拟化。 虽然目前仍在使用 LXC，但却提供了使用 Linux 内核的新技术。现代的开源 Linux 操作系统 Ubuntu 也提供了此功能。

Docker 通过以下技术增强了本机 Linux 容器化功能：

- **增强的无缝可移植性：**虽然 LXC 容器通常引用特定于机器的配置，但 Docker 容器无需修改即可在任何桌面、数据中心和云环境中运行。
- **更轻巧且更精细的更新：**使用 LXC，可以在单个容器中组合多个进程。 对于 Docker 容器，每个容器中只能运行一个进程。 这样就可以构建如下应用程序：在其中一个部分关闭以进行更新或修复时可以继续运行的应用程序。
- **自动创建容器：**Docker 可以根据应用程序源代码自动构建容器。
- **容器版本控制：**Docker 可以跟踪容器映像的版本，回滚到先前的版本，以及跟踪版本的构建者和构建方式。 它甚至可以只上载现有版本和新版本之间的变更内容。
- **容器复用：**现有容器可用作*基本映像*（本质上类似于用于构建新容器的模板）。
- **共享容器库：**开发人员可以访问包含数千个用户贡献的容器的开源注册表。

如今，Docker 容器化也适用于 Microsoft Windows 服务器。 大多数云提供商都提供了一些专用服务，这些服务可帮助开发人员构建、交付和运行使用 Docker 进行容器化的应用程序。

由于这些原因，Docker 迅速被采用并且其采用量持续激增。 在撰写本文时，Docker Inc. 报告[每月有 1100 万开发人员和 130 亿次容器映像下载](https://www.docker.com/)（链接位于 IBM 外部）。

## Docker 工具和术语

使用 Docker 时，您会遇到以下工具和术语：

### **DockerFile**

每个 Docker 容器都从一个简单的文本文件开始，其中包含有关如何构建 Docker 容器映像的指示信息。*DockerFile* 将自动执行 Docker 映像创建过程。 它本质上是一个命令行界面 (CLI) 指令列表，Docker 引擎将运行这些指令来组装映像。

### **Docker 映像**

*Docker 映像*包含可执行的应用程序源代码以及在将应用程序代码作为容器运行时所需的所有工具、库和依赖项。 当运行 Docker 映像时，它会成为容器的一个（或多个）实例。

虽然可以从头开始构建 Docker 映像，但大多数开发人员都是从公共存储库中提取 Docker 映像。 可以通过一个基本映像来创建多个 Docker 映像，并且它们将共享一个公共堆栈。

Docker 映像由*层*组成，每个层对应于该映像的一个版本。 每当开发人员对映像进行更改时，都会创建一个新的顶层，该顶层将替换之前的顶层作为映像的当前版本。 之前的层会被保存下来，以便执行回滚或在其他项目中进行复用。

每次通过 Docker 映像创建容器时，都会创建另一个称为“容器层”的新层。  对容器所做的更改（例如添加或删除文件）仅保存到容器层中，并且仅在容器运行时才会出现。  这种迭代式映像创建过程可以提高整体效率，因为可以仅从一个基本映像中运行多个实时的容器实例，如果这样做，这些实例将使用一个公共堆栈。

### **Docker 容器**

Docker 容器是实时运行的 Docker 映像实例。 Docker 映像是只读文件，而容器是实时的瞬态可执行内容。 用户可以与它们进行交互，管理员可以使用 docker 命令调整其设置和条件。

### **Docker Hub**

[*Docker Hub*](https://hub.docker.com/)（链接位于 IBM 外部）是 Docker 映像的公共存储库，自称为“世界上最大的容器映像库和社区”。它拥有超过 100,000  个来自商业软件供应商、开源项目和个人开发人员的容器映像。 它包括由 Docker, Inc. 生成的映像、属于 Docker Trusted  Registry 的经过认证的映像以及成千上万的其他映像。

所有 Docker Hub 用户都可以随意共享其映像。 他们还可以从 Docker 文件系统下载预定义的基础映像，用作任何容器化项目的起点。

### **Docker 守护程序**

Docker 守护程序是在操作系统上运行的服务，例如 Microsoft Windows、Apple MacOS 或 iOS。 此服务使用来自客户机的命令为您创建和管理 Docker 映像，充当 Docker 实现的控制中心。

### **Docker 注册表**

Docker 注册表是一个用于 Docker 映像的可扩展开源存储和分发系统。 此注册表使您可以使用标记进行标识来跟踪存储库中的映像版本。 可使用 git 版本控制工具来完成此操作。

## Docker 部署和编排

如果您只运行几个容器，那么在 Docker Engine（行业实际使用的运行时）中管理应用程序将相当简单。 但是，如果您的部署包含数千个容器和数百个服务，在没有这些专用工具的帮助下，您几乎不可能管理该工作流程。

### **Docker Compose**

如果您正在使用位于同一主机上的多个容器中的进程来构建应用程序，那么可以使用 *Docker Compose* 来管理应用程序的架构。Docker Compose 将创建一个 YAML 文件来指定应用程序中要包含的服务，并且可以使用单个命令来部署和运行容器。 使用 Docker Compose，您还可以定义用于存储的持久卷、指定基本节点以及记录和配置服务依赖项。

### **Kubernetes**

要在更复杂的环境中监视和管理容器生命周期，您需要转为使用[容器编排](https://www.ibm.com/cloud/learn/container-orchestration)工具。 虽然 Docker 提供了自己的编排工具（称为 Docker Swarm），但大多数开发人员还是会选择 [Kubernetes](https://www.ibm.com/cn-zh/cloud/learn/kubernetes)。

Kubernetes 是一个开源的容器编排平台，源自为 Google 内部使用而开发的项目。Kubernetes 可调度和自动执行对于基于容器的架构管理而言不可或缺的任务，包括容器部署、更新、服务发现、存储配置、[负载均衡](https://www.ibm.com/cloud/learn/load-balancing)、运行状况监控等。 此外，Kubernetes 工具的开源生态系统（包括 [Istio](https://www.ibm.com/cn-zh/cloud/learn/istio) 和 [Knative](https://www.ibm.com/cloud/learn/knative)）使组织能够为容器化应用程序部署高生产力的 [Platform-as-a-Service (PaaS)](https://www.ibm.com/cn-zh/cloud/learn/paas)，并更快地进入[无服务器计算](https://www.ibm.com/cloud/learn/serverless)时代。

要更深入地了解 Kubernetes，请观看视频“Kubernetes 解释”：



## Docker 和 IBM Cloud

企业容器平台可提供跨多个公共云和私有云的编排，将您的环境统一起来以提高业务绩效和运营绩效。 它是开放式混合云战略的关键组成部分，可让您避免供应商锁定、随时随地以一致的方式构建和运行工作负载，以及[优化所有 IT 并将其现代化](https://www.ibm.com/cn-zh/cloud/hybrid-infrastructure)。

采取下一步行动：

- 使用 [Red Hat OpenShift on IBM Cloud](https://www.ibm.com/cn-zh/cloud/openshift) 部署高度可用、完全托管的 Kubernetes 集群；Red Hat OpenShift on IBM Cloud 是一项托管的 OpenShift 服务，利用 IBM Cloud 的企业规模和安全性来自动执行更新、扩展和配置任务。
- 使用 [IBM Cloud Satellite](https://www.ibm.com/cn-zh/cloud/satellite)（一种托管的分布式云解决方案）跨本地、边缘计算和公共云环境部署和运行来自任何供应商的应用程序
- 通过使用 [IBM 混合云存储解决方案](https://www.ibm.com/cn-zh/it-infrastructure/storage/hybrid-cloud-storage)跨本地和公共云环境无缝部署支持容器的企业存储，以简化和整合您的数据湖
- 使用 [IBM Cloud 托管服务](https://www.ibm.com/cn-zh/services/cloud/managed)来简化复杂的混合 IT 管理。