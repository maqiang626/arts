# Kubernetes main attack vectors tree: an explainer guide		

https://www.cncf.io/blog/2021/11/08/kubernetes-main-attack-vectors-tree-an-explainer-guide/



[CNCF 				Member				Blog Post](https://www.cncf.io/lf-author-category/member/)

Posted on 			November 8, 2021			 		By Andrew Zola			 					

*Guest post originally published on [Magalix’](https://www.magalix.com/blog/kubernetes-main-attack-vectors-tree-an-explainer-guide)*[*s*](https://www.magalix.com/blog/kubernetes-main-attack-vectors-tree-an-explainer-guide)*[ blog](https://www.magalix.com/blog/kubernetes-main-attack-vectors-tree-an-explainer-guide) by Andrew Zola*

Kubernetes is a leader in container orchestration. According to [Statista](https://www.statista.com/statistics/1233945/kubernetes-adoption-level-organization/), as much as 46% of respondents in a recent survey stated that they used  Kubernetes for automating computer application deployments, management,  and scaling.

However, there are some security issues that we need to address  during the container lifecycle phase. This means that you should take  steps to mitigate the risk of known vulnerabilities, including  misconfigurations, during the development and deployment phase. You  should also be well-placed to quickly respond to potential threats  during runtime.

As such, organizations need to take an in-depth look at the K8s  Attack Tree and documentation. It’s important as the attack tree is  essentially a threat model that offers a detailed view of potential  threats and mitigations.  

You can even use this threat model as a checklist to identify  potential vulnerabilities and common attack vectors a bad actor might  leverage to breach the system. You can also use it as a tool to test  Kubernetes security and gain visibility into the logging output that’s  generated in the event of a security incident.

However, it’s important to note that the Attack Tree only focuses on  Kubernetes and not end-to-end container security. This means that it  doesn’t include any other application or components you may have added  during the software development lifecycle.

So, whenever you plan to use Kubernetes in production, you must also  consider the system’s overall security. According to the National[ Security Agency (NSA)](https://media.defense.gov/2021/Aug/03/2002820425/-1/-1/1/CTR_KUBERNETES HARDENING GUIDANCE.PDF), common security threats include insider threats, malicious threat actors, and supply chain risks.

This makes it vital to establish Kubernetes’ security best practices from the beginning to better secure [critical areas of cloud infrastructure](https://www.magalix.com/blog/cloud-vulnerability-management-101-part-1-key-vulnerability-areas) and cloud-native environments.

## Types of Kubernetes Attackers

When it comes to Kubernetes attack vectors, there are three main  types—external attackers, malicious containers, and compromised or  malicious users.

### 1- External Attackers

You can have threat actors who have no access to the cluster but are  able to reach the application running on it. They might also have access to the management ports on the network.

To avoid exposure to untrusted networks, you must also be alert to  controls like management services. In this case, it can take the form of networks without any form of authentication protocol.

### 2- Malicious Containers

If a threat actor manages to breach a single container, they will  attempt to increase their access and take over the entire cluster. If  there are only minimal controls to stop attackers from gaining full  cluster administration rights, they will probably be successful.

To mitigate risk, you must ensure that all management ports are  visible on the cluster network. You must also leverage multi-factor  authentication protocols for all users. You must also avoid mounting  service accounts in containers or have restricted rights.

Through network policies and using [policy-as-code](https://www.magalix.com/blog/what-is-policy-as-code), you can restrict access between pods and namespaces.

### 3- Compromised or Malicious Users

When you’re dealing with compromised accounts or malicious users, an  attacker with stolen yet valid credentials will execute commands against network access and the Kubernetes API. However, you can mitigate this  risk by following a “least privilege” policy. You should also enforce  Role-based Access Control (RABC) policies and best practices for all  users.

## The Main Attack Vectors

Last year, the [CNCF Financial User Group](https://github.com/cncf/financial-user-group) released a threat modeling exercise that targeted a generic Kubernetes cluster.  The primary objective here was to offer a detailed view of potential  threats and mitigations. The accompanying checklist helps teams identify common vulnerabilities and exploits within a Kubernetes cluster.

![img](https://www.magalix.com/hs-fs/hubfs/Google%20Drive%20Integration/Copy%20of%20Magalix%20-%20K8s%20Main%20Attack%20Vectors.png?width=624&height=272&name=Copy%20of%20Magalix%20-%20K8s%20Main%20Attack%20Vectors.png)

Kubernetes Trust Boundaries – Courtesy of CNCF

Using [STRIDE methodology](https://securityintelligence.com/articles/what-is-stride-threat-modeling-anticipate-cyberattacks/), they analyzed each element of the Kubernetes architecture and  formulated a list of potential security issues within the platform.  STRIDE is an acronym that stands for Spoofing identity, Tampering with  data, Repudiation threats, Information disclosure, Denial of Service,  and Elevation of privileges.

Some of the main attack vectors include:

### Service Token

Service tokens are mounted onto each pod by default. In this  scenario, if an attacker manages to compromise a container, they will  have the mechanism to exploit it further using those same credentials.

To mitigate this risk, it’s crucial to establish strict RBAC  policies. It’s also essential to disable automatic service token  mounting protocols.

### Compromised Containers

One of the primary focal points for an attacker is the remote  execution point within the cluster. Besides service token attacks, other attack vectors include the default network exposure of the control pane of all running containers.

### Network Endpoints

You must secure each Kubernetes endpoint from internal threat actors. This approach helps avert an easy attack vector from being exploited.  It’s crucial to note that whenever an attacker manages to compromise a  container, they quickly gain access to the endpoints whenever the pods’  network policies permit it.

### Denial of Service (DOS)

There were relatively few mitigations against DoS attacks before the 1.14 release.

### RBAC Issues

Threat actors depend on RBAC misconfigurations to initiate a data  breach. To mitigate risk, development teams must leverage automated  tooling to verify and validate policies.

## Attack Trees

The CNCF Financial User Group came up with a set of [attack trees](https://github.com/cncf/financial-user-group/tree/master/projects/k8s-threat-model/AttackTrees) that could potentially determine the lineage of the initial attempt to  create a foothold within the cluster. In this case, the working group  came up with two approaches:

1. Bottom-up approach.
2. Scenario approach.

### Bottom-Up Approach

The bottom-up approach satisfies the stated goal throughout the  Kubernetes platform by showing all entry points. This method is helpful  to map out all security controls and standards to understand their  coverage better.

### Scenario Approach

The scenario approach helps identify attack vectors that are open to  threat actors under certain circumstances. In this case, the  scenario-based method leverages much of the detail in the bottom-up  approach but in a highly realistic manner. This approach also provides  more focus on more prevalent attack vectors.

Here’s a summary of the attack trees that are open-sourced and available on GitHub:

### Bottom-Up Approach: [Malicious Code Execution](https://github.com/cncf/financial-user-group/blob/master/projects/k8s-threat-model/AttackTrees/pdfs/Kubernetes Attack Trees v1.4.malicious.pdf)

The main goal here is to execute malicious code on a cluster.  However, to gain a foothold, you must compromise an application that  provides access to the container.

Once a threat actor can access the container, they will load more  malicious code into the environment. If they can obtain the image pull  secret, the threat actor may poison the repository to initiate a wider  distribution of malicious code.

### Bottom-Up Approach: [Establish Persistence](https://github.com/cncf/financial-user-group/blob/master/projects/k8s-threat-model/AttackTrees/pdfs/Kubernetes Attack Trees v1.4-persistence.pdf)

The primary objective of this attack tree is to discover the  different ways a hacker might try to gain access to the cluster. This  method also investigates different periods of longevity.

In this case, one branch will focus on reading secrets held within  the cluster. This approach helps attackers exploit other vulnerable  areas. The second branch focuses on threats after an attacker gains  access to a container. They will leverage misconfigurations and attempt  to establish persistence resilient to containers, nodes, or pod  restarts.

### Bottom-Up Approach: [Access Sensitive Data](https://github.com/cncf/financial-user-group/blob/master/projects/k8s-threat-model/AttackTrees/pdfs/Kubernetes Attack Trees v1.4.sensitive.data.pdf)

Most leading approaches concentrate on exploiting misconfigured RBAC  permissions to read secret data directly from the cluster. Some other  methods include viewing all the data stored within logs. It’s almost  like eavesdropping on network traffic and communications.

### Bottom-Up Approach: [Denial Of Service](https://github.com/cncf/financial-user-group/blob/master/projects/k8s-threat-model/AttackTrees/pdfs/Kubernetes Attack Trees v1.4.dos.pdf)

This attack tree investigates different approaches an attacker can  take to initiate a DoS attack on the cluster. The first method follows a compromised container scenario where an attacker attempts to DoS the  cluster from within. The aim here is to exhaust all its resources.

The second method concentrates on a threat actor who has network  access to the cluster control pane. In this scenario, they might try to  flood the network at the most suitable endpoints. Again, the goal here  is to exhaust all resources.

### Scenario: [Compromised Applications Leads to Foothold in Container](https://github.com/cncf/financial-user-group/blob/master/projects/k8s-threat-model/AttackTrees/pdfs/Kubernetes Attack Trees v1.4.scenario.compromised.pdf)

This scenario focuses on potential vectors that are probably open to  an attacker after exploiting an application running within a container.  This can lead to remote code execution within the container through  programmatic or shell access mechanisms.

### Scenario: [Attacker on the Network](https://github.com/cncf/financial-user-group/blob/master/projects/k8s-threat-model/AttackTrees/pdfs/Kubernetes Attack Trees v1.4.scenario.network.pdf)

This scenario concentrates on insider threats. An internal attacker  with network access to the Kubernetes cluster will probably have many  user privileges without direct access to the cluster. However, you can  quickly mitigate most of these threats with appropriate Kubernetes  configurations and firewalls.

[You can find out more from the ](https://github.com/kubernetes/community/tree/master/wg-security-audit)[Kubernetes Security Audit Working Group](https://github.com/kubernetes/community/tree/master/wg-security-audit) and their findings in performing a security audit and producing artifacts as a [threat model](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/Kubernetes Threat Model.pdf) and [whitepaper](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/Kubernetes White Paper.pdf). They focus on different parts of the Kubernetes cluster, and it’s worth the time and effort to read through it.

Other Kubernetes best practices and recommendations include the following:

- Always run pods and containers following the least privileges policy (whenever possible).
- Demand strong authentication and authorization protocols to limit administrator and user access (and limit the attack surface).
- Encrypt data to ensure confidentiality.
- Leverage firewalls to limit unnecessary connectivity.
- Limit potential damage by using network separation controls.
- Monitor activity via log auditing.
- Regularly review all settings and use vulnerability scans to identify potential risks and apply security patches.
- Scan pods and containers for misconfigurations and vulnerabilities.

## Secure your K8s Environment with Policy As Code

The exuberant number of best practices, configurations, and  maintainability necessary to prevent hackers from exploiting an attack  vector makes managing Kubernetes a challenging task for even seasoned  experts. Magalix provides one of the largest scalable Policy-as-Code  libraries to simplify and reduce the complexity of managing secure K8s.

Policies include:

- Containers should not be running as root.
- Containers are missing securityContext.
- RBAC Protect cluster-admin ClusterRoleBindings.
- Prohibit RBAC Wildcards for Verbs.
- Services should not be using NodePort.
- Containers are missing Liveness Probe.
- Containers should mount the root filesystem as read-only.
- Containers should not share hostIPC.
- Containers should not be using hostPort.
- Containers should not be mounting the Docker sockets.

With Magalix, you can also programmatically enforce[ security](https://www.magalix.com/airtight-security-as-code) and [compliance](https://www.magalix.com/usecases/continuous-compliance-assessment) standards with policy-as-code. This approach helps provide developer-centric  experiences that complement continuous deployment protocols for  cloud-native applications. In this scenario, automated operators within  the cluster will continuously monitor the repositories for changes.

When using policy-as-code, you can also develop a centralized  playbook enacted and enforced across the whole software development  lifecycle. This approach helps accelerate innovation without  compromising on security.

When you build cloud-native applications following the  recommendations listed above, development teams essentially weave  security into the fabric of the whole application. As such, it’ll  provide the foundation to scale securely while delivering enhanced user  experiences.

To learn more about Kubernetes security and how you can take advantage of policy-as-code.



------

安全永远是重中之重！

学习到了很多最佳实践！

