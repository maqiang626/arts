# Kubernetes RBAC鉴权

https://yinwu.blog.csdn.net/article/details/121311029



使用RBAC在Kubernetes中配置权限

确保控制谁有权访问k8s系统以及哪些用户有权访问是身份和访问管理系统的目标。它是安全管理中的基本过程之一，应予以彻底照顾。

在Kubernetes中，身份和用户管理未集成在平台中，应由外部IAM平台（如Keycloak，Active Directory，Google的IAM等）进行管理。但是，身份验证和授权由Kubernetes处理。

在本文中，将重点介绍Kubernetes中IAM的授权方面，更具体地说，是如何使用基于角色的访问控制模型来确保用户对正确的资源具有正确的权限。
先决条件

RBAC是Kubernetes 1.8赋予的一项稳定功能。在本文中，假定正在使用Kubernetes 1.9+。还将假设通过–authorization-mode=RBACKubernetes API服务器中的选项在集群中启用了RBAC 。可以通过执行命令进行检查kubectl api-versions；如果启用了RBAC，则应该看到API版本.rbac.authorization.k8s.io/v1。
Kubernetes中的RBAC概念概述

Kubernetes中的RBAC模型基于三个要素：角色：定义每种Kubernetes资源类型的权限；主题：用户（人员或机器用户）或用户组；RoleBindings：定义哪些主体具有哪些角色；探索这些元素如何工作。在下面的示例中，可以看到一个角色，该角色允许在名称空间“ mynamespace”中获取，列出和监视Pod：

```yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: mynamespace
  name: example-role
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "watch", "list"]

```



为了给用户上一角色中描述的权限，必须创建一个RoleBinding。在下面的示例中，RoleBinding“ example-rolebinding”将角色“ example-role”绑定到用户“ example-user”：

```yaml
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: example-rolebinding
  namespace: mynamespace
subjects:
- kind: User
  name: example-user
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: example-role
  apiGroup: rbac.authorization.k8s.io

```



应该注意的是，Roles和RoleBindings是命名空间的，这意味着只能为与Role和RoleBinding位于相同命名空间中的Kubernetes资源赋予权限。另外，不用说RoleBinding只能引用其命名空间中存在的Role。
角色，ClusterRoles，RoleBindings和ClusterRoleBindings

在前面的示例中，使用了Roles和RoleBindings。但是，也可以使用在以下情况下有用的ClusterRoles和ClusterRoleBindings：

授予非命名空间资源（如节点）的权限 为群集的所有命名空间中的资源授予权限 为非资源终结点（如/ healthz）授予权限 角色和角色绑定的群集范围版本的定义与非群集范围的版本非常相似。如果我们采用前面的示例并将其改编，将具有以下定义。

ClusterRole：

```yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: example-clusterrole
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "watch", "list"]


```



ClusterRoleBinding：

```yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: example-clusterrolebinding
subjects:
- kind: User
  name: example-user
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: example-clusterrole
  apiGroup: rbac.authorization.k8s.io

```



如何在Roles和ClusterRoles中启用权限

在第一个示例中，我们向用户授予了获取，监视和列出Pod的基本权限。让我们探索不同资源和权限的其他可能性。在下面的示例中，角色允许对部署资源执行任何操作：

```yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: mynamespace
  name: example-role-2
rules:
- apiGroups: ["extensions", "apps"]
  resources: ["deployments"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]

```



请注意，在这种情况下，apiGroups字段已用Deployment的API组填充。根据Kubernetes版本的不同，可以在API apps / v1或extensions / v1beta2中找到Deployment资源；API组是斜线之前的部分。可以在一个Role中定义多个规则，如下面的示例：

```yaml
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["batch", "extensions"]
  resources: ["jobs"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]

```



在这种情况下，我们根据目标资源是Pod还是Job来授予不同的权限。还可以通过其名称来定位资源，如以下示例：

```yaml
rules:
- apiGroups: [""]
  resources: ["configmaps"]
  resourceNames: ["my-config"]
  verbs: ["get"]


```



如何将主题绑定到角色或ClusterRole

在第一个示例中，我们已经看到了如何将人类用户绑定到角色。但是，也有可能绑定服务帐户（非人类用户）或一组人类用户和/或服务帐户。在下面的示例中，RoleBindingexample-rolebinding将ServiceAccount绑定example-s到Role example-role：

```yaml
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: example-rolebinding
  namespace: mynamespace
subjects:
- kind: ServiceAccount
  name: example-sa
  namespace: mynamespace
roleRef:
  kind: Role
  name: example-role
  apiGroup: rbac.authorization.k8s.io

```



可以使用以下命令创建ServiceAccount：

```shell
kubectl create serviceaccount example-sa --namespace mynamespace

```

在前面的RoleBinding定义中，我们还可以替换主题以绑定组。在下面的示例中，我们绑定frontend-admins组：

```yaml
subjects:
- kind: Group
  name: "frontend-admins"
  apiGroup: rbac.authorization.k8s.io

```



另一种可能性是绑定服务帐户组。在这里，我们将所有服务帐户绑定到mynamespace名称空间中：

```yaml
subjects:
- kind: Group
  name: system:serviceaccounts:mynamespace
  apiGroup: rbac.authorization.k8s.io

```



或集群中的所有服务帐户：

```yaml
subjects:
- kind: Group
  name: system:serviceaccounts
  apiGroup: rbac.authorization.k8s.io

```



结语

我们已经看到了如何使用基于角色的访问控制模型向用户或服务帐户授予权限。这是在Kubernetes中实现授权的一种有效方法，它可能是最受欢迎的一种，但它不是唯一可用的模型：额外还可以使用其他模型，例如ABAC（基于属性的访问控制），Node Authorization模型和Webhook模式。可以查阅相关资料进行深入了解。

