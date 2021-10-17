# 			GitOps 101: What’s it all about?		

[CNCF 				Member				Blog Post](https://www.cncf.io/lf-author-category/member/)

Posted on 			September 28, 2021			 		By Andrew Zola			 					

https://www.cncf.io/blog/2021/09/28/gitops-101-whats-it-all-about/



*Guest post originally published on* [Magalix*‘ blog*](https://www.magalix.com/blog/what-is-gitops) *by Andrew Zola*

If you think [GitOps](https://about.gitlab.com/topics/gitops/) sounds a bit like [DevOps](https://www.magalix.com/blog/what-is-devsecops-and-why-is-it-important), you would be right. GitOps is essentially an operational framework that uses DevOps best practices. In this scenario, we basically move all  cloud operations to Git.

[Git](https://searchitoperations.techtarget.com/definition/Git) is essentially a source code management (SCM) tool developed in 2005 to  support software development projects. As Git is already a core part of  application development, leveraging Git enforces DevOps best practices  and provides access to a robust version control system.

Other benefits include infrastructure automation, continuous integration/continuous delivery (CI/CD), and collaboration.

## What is GitOps?

[Weaveworks](https://www.weave.works/blog/what-is-gitops-really) coined the term “GitOps” in 2017 to share the idea that all deployments should be as easy as enacting a code change.

GitOps is a standardized workflow for configuring, deploying,  monitoring, managing and updating infrastructure-as-code Kubernetes and  all of its components as code. This includes all the applications that  run on it.

The core idea here is to have declarative descriptions of the  infrastructure and all related elements in its currently desired state  in the production environment. In this scenario, an automated process  ensures that the described state in the repository and the production  environment always match.

In this case, it considers everything related to software deployment:

- Codebase changes.
- Build.
- Packaging.
- Application validation.
- Infrastructure changes.

Because of everything that was required, Kubernetes was the ecosystem of choice for GitOps projects. The Kuberentes platform is the perfect  solution because of improved developer productivity, higher reliability, increased flexibility, enhanced operational flexibility, improved  auditability, compliance, and security.

As GitOps evolved, the definition of “microservices” also changed to accommodate more movement.

As such, we can summarize GitOps as:

- An operating model for cloud-native technologies and Kubernetes  that comes with a set of best practices to achieve unified Git  deployment, management, and monitoring for containerized applications  and clusters.
- A pathway towards efficiently managing  applications by applying Git workflows and end-to-end CI/CD pipelines to operations and development.

At its most basic, GitOps is about merging intelligent source control with automated CI/CD tooling. For example, whenever you add some code  into a Git repository, a lot happens to ensure that your code gets to a  relevant target automatically.

For example, if you have code for a new application feature, it  automatically ends up in the existing application. Whenever your code  declares a network policy update, it’s automatically propagated into the network infrastructure.

![GitOps Operating Model](https://www.magalix.com/hs-fs/hubfs/Google%20Drive%20Integration/Copy%20of%20Magalix%20-%20The%20Ultimate%20Guide%20to%20GitOps.png?width=624&name=Copy%20of%20Magalix%20-%20The%20Ultimate%20Guide%20to%20GitOps.png)

Figure 1: GitOps Operating Model. *Source Weaveworks*

## What are the GitOps Principles?

In a cloud-native environment with Git as a single source of truth of the system’s current desired state, you can commit all intended  operations with a pull request. All changes are observable and  auditable, and automatic convergence highlights all differences between  the intended and observed states.

GitOps encompasses several guiding principles. **These include the following:**

### The Entire System is in a Declarative State

GitOps demands an infrastructure that’s always declarative. It should also concentrate on the target configuration. In other words, it  focuses on the desired state and enables the system to execute whatever  it needs to realize a desired state.

In contrast, an imperative approach concentrates on a set of explicit commands to change the desired state. This makes reconciling a  challenge as imperative infrastructure is unaware of the state. You must store the declarative state of the entire system in Git. In this case,  Kubernetes is the most prolific declarative infrastructure that allows  you to keep its state in Git.

### The Canonical Desired State Versioned in Git

When it comes to GitOps, the canonical state is essentially the  “source of truth” state. For example, when the state is stored and  versioned in source control, it must be viewed as a source of truth.

In this case, you can test objects on how equal they are when  compared to the canonical form. Whenever there’s a deviation in the  state, it can quickly recognize and reconcile it back to the canonical  state in source control.

### Automatically Apply Approved Changes in the System

Once you store the declared state in Git, you must allow all changes  to that state to be applied automatically to your system through pull  requests (PR) or merge requests (MR). You won’t need cluster credentials to make changes to the system.

There’s also a segregated environment in GitOps where the state  definition lives outside. As a result, you can separate what you do and  how you do it. What’s excellent about GitOps here is that it favors a  low barrier of entry. In this case, you won’t achieve immediate  deployment or reconciliation until you achieve a new canonical state.

Once you declare the state of your system and keep it under version  control, you can use software agents to alert you whenever reality and  your expectations don’t match.

Software agents also help ensure that the whole system is  self-healing to mitigate the risk of human error and more. In this  scenario, software agents act like an operational control and feedback  loop.

## What is the Role of Infrastructure-as-Code (IAC)?

Many GitOps concepts and benefits are rooted in IaC. This also  reflects a core foundational concept of GitOps. However, to put it  simply, GitOps concentrates on automating tasks around Kubernetes  clusters and deployments. IaC essentially focuses on the underly  infrastructure like everything Kubernetes clusters need to run smoothly.

But this proves to be a challenge in the complete GitOps model as  GitOps processes aren’t aware of the infrastructure. However, as most  GitOps adopters already have IaC skills, it hasn’t become a significant  issue.

You can also take it a step further to ensure security by enforcing security standards as by using [policy-as-code](https://www.magalix.com/blog/policy-as-code-for-kubernetes) and [security-as-code](https://www.magalix.com/airtight-security-as-code). This approach helps create developer-centric experiences with CD for cloud-native applications.

In this scenario, “automated operators” within a Kubernetes cluster  of cloud infrastructure will be continuously monitoring the repository  for changes. Whenever it discovers a change, the operators will  automatically trigger an update.

You can also create a centralized playbook and enforce the proper  workflows across the software development lifecycle. This approach helps development teams innovate faster without compromising security. You  can also use this method to help ensure compliance.

## What are the Key Benefits of GitOps Adoption?

Software developers often find a massive gap between the code they  write and the idea that finally goes into production. We rarely develop  applications in a vacuum, and there’s a lot of source code management  that takes place from day one.

With GitOps, you can take advantage of the following:

### 1- Better Ops

With GitOps, you have a complete end-to-end pipeline. In this  scenario, PR can drive your CI/CD pipelines and reproduce operations  tasks through Git.

### 2- Rapid Development

When you adopt and follow GitOps best practices, you must leverage  tools like Git to quickly manage Kubernetes updates and features. When  businesses continuously push feature updates, they also become more  agile. This makes it easier to respond to customer requests more rapidly and achieve a competitive advantage.

### 3- Stronger Security Protocols Guaranteed

As Git is backed by robust cryptography that helps track and manage  changes securely, your security is more or less guaranteed. In this  case, the ability to sign changes and prove authorship origin is crucial to secure and correct the definition of the desired state of a cluster.

In the unfortunate event of a security incident, the immutable and  auditable source of truth helps recreate a new system independently.  This approach helps minimize downtime and sets the stage for  much-improved incident response.

The separation of responsibility between packing and releasing  software into production also helps ensure security. This approach  follows the least privilege principle and minimizes the impact of a  breach because of a smaller attack surface.

### 4- Simple Auditing and Compliance

As we track and log changes securely, auditing and ensuring  compliance is a breeze. In this case, you can use a wide range of  comparison tools like [ansiblediff](https://docs.ansible.com/ansible/latest/collections/ansible/utils/fact_diff_module.html) and [kubediff](https://github.com/weaveworks/kubediff) to compare the trusted definition of the state of a cluster with the  cluster running in the real world. This approach helps ensure that  auditable and tracked changes match reality.

To tie this all together, GitOps helps businesses with the internet  at their foundation innovate faster and maintain a competitive  advantage. In the digital age, only enterprises that can keep up with  this relentless cycle will survive.

At the same time, companies also need to leverage policy-as-code to ensure that their digital products remain secure.





------

GitOps

IaC



非常强大，Git 本身很给力，DevOps，GitOps，期待...





