# Backup and Restore Etcd Deployments on Kubernetes

更新时间2021年10月27日

https://docs.vmware.com/en/VMware-Tanzu-Application-Catalog/services/tac-docs-tutorials/GUID-backup-restore-data-etcd-kubernetes-index.html



## Introduction

[etcd](https://etcd.io/) is a reliable and efficient  key-value store, most commonly used for data storage in distributed  systems. It offers a simple interface for reading and writing data and  is available under an open source license.

VMware Application Catalog offers an [etcd Helm chart](https://docs.vmware.com/en/VMware-Application-Catalog/services/apps/GUID-apps-etcd-index.html) that enables quick and easy deployment of an etcd cluster on  Kubernetes. This Helm chart is compliant with current best practices and is suitable for use in production environments, with built-in features  for role-based access control (RBAC), horizontal scaling, disaster  recovery and TLS.

Of course, the true business value of an etcd cluster comes not from  the cluster itself, but from the data that resides within it. It is  critical to protect this data, by backing it up regularly and ensuring  that it can be easily restored as needed. Data backup and restore  procedures are also important for other scenarios, such as off-site data migration/data analysis or application load testing.

This guide walks you through two different approaches you can follow  when backing up and restoring VMware Application Catalog etcd Helm chart deployments on Kubernetes:

- Back up the data from the source deployment and restore it in a new deployment using etcd's built-in backup/restore tools.
- Back up the persistent volumes from the source deployment and attach them to a new deployment using [Velero](https://velero.io/), a Kubernetes backup/restore tool.

## Assumptions and prerequisites

This guide makes the following assumptions:

- You have two separate Kubernetes clusters - a source cluster and a destination cluster - with `kubectl` and Helm v3 installed. This guide uses Google Kubernetes Engine (GKE)  clusters but you can also use any other Kubernetes provider. Learn how  to install [`kubectl`](https://docs.vmware.com/en/VMware-Application-Catalog/services/main/GUID-faq-kubernetes.html#how-do-i-install-kubectl-1) and [Helm v3.x](https://docs.vmware.com/en/VMware-Application-Catalog/services/main/GUID-faq-kubernetes.html#how-do-i-install-and-configure-helm-2).

- You have configured Helm to use the VMware Application Catalog chart repository following the [instructions for VMware Application Catalog](https://docs.vmware.com/en/VMware-Application-Catalog/services/main/GUID-get-started-get-started-vmware-application-catalog.html#consume-the-tanzu-application-catalog-artifacts-in-your-local-machine-7) or the [instructions for VMware Application Catalog for Tanzu Advanced](https://docs.vmware.com/en/VMware-Application-Catalog/services/main/GUID-get-started-get-started-vmware-application-catalog-tanzu-advanced.html#consume-helm-charts-8).

- You have previously deployed the VMware Application Catalog etcd Helm chart on the source cluster and added some data to it. Example  command sequences to perform these tasks are shown below, where the  PASSWORD placeholder refers to the etcd administrator password. Replace  the REPOSITORY placeholder with a reference to your VMware Application  Catalog chart repository.

  ```bash
  helm install etcd REPOSITORY/etcd \
    --set auth.rbac.rootPassword=PASSWORD \
    --set statefulset.replicaCount=3
  export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=etcd,app.kubernetes.io/instance=etcd" -o jsonpath="{.items[0].metadata.name}")
  kubectl exec -it $POD_NAME -- etcdctl --user root:PASSWORD put /message1 foo
  kubectl exec -it $POD_NAME -- etcdctl --user root:PASSWORD put /message2 bar
  ```

## Method 1: Backup and restore data using etcd's built-in tools

This method involves using etcd's `etcdctl` tool to create a snapshot of the data in the source cluster and the VMware Application Catalog's Helm chart's recovery features to create a new cluster using  the data from the snapshot.

### Step 1: Create a data snapshot

The first step is to back up the data in the etcd deployment on the source cluster. Follow these steps:

1. Forward the etcd service port and place the process in the background:

   ```bash
   kubectl port-forward --namespace default svc/etcd 2379:2379 &
   ```

2. Create a directory for the backup files and make it the current working directory:

   ```bash
   mkdir etcd-backup
   chmod o+w etcd-backup
   cd etcd-backup
   ```

3. Use the `etcdctl` tool to create a snapshot of the  etcd cluster and save it to the current directory. If this tool is not  installed on your system, use VMware Application Catalog's etcd Docker  image to perform the backup, as shown below (replace the example  container registry shown with your VMware Application Catalog container  registry). Replace the PASSWORD placeholder with the administrator  password set at deployment-time and the REGISTRY placeholder with a  reference to your VMware Application Catalog container registry.

   ```bash
   docker run -it --env ALLOW_NONE_AUTHENTICATION=yes --rm --network host  -v $(pwd):/backup REGISTRY/etcd etcdctl --user root:PASSWORD --endpoints http://127.0.0.1:2379 snapshot save /backup/mybackup
   ```

   Here, the `--net` parameter lets the Docker container use the host's network stack and thereby gain access to the forwarded port. The `etcdctl` command connects to the etcd service and creates a snapshot in the `/backup` directory, which is mapped to the current directory (`etcd-backup/`) on the Docker host with the `-v` parameter. Finally, the `--rm` parameter deletes the container after the `etcdctl` command completes execution.

4. Stop the service port forwarding by terminating the corresponding background process.

5. Adjust the permissions of the backup file to make it world-readable:

   ```bash
   sudo chmod -R 644 mybackup
   ```

At the end of this step, the backup directory should contain a file named `mybackup`, which is a snapshot of the data from the etcd deployment.

### Step 2: Copy the snapshot to a PVC

The etcd cluster can be restored from the snapshot created in the  previous step. There are different ways to do this; one simple approach  is to make the snapshot available to the pods using a Kubernetes  PersistentVolumeClaim (PVC). Therefore, the next step is to create a PVC and copy the snapshot file into it. Further, since each node of the  restored cluster will access the PVC, it is important to create the PVC  using a storage class that supports ReadWriteMany access, such as NFS.

1. Begin by installing the NFS Server Provisioner. The easiest way  to get this running on any platform is with the stable Helm chart. Use  the command below, remembering to adjust the storage size to reflect  your cluster's settings:

   ```bash
   helm repo add stable https://charts.helm.sh/stable
   helm install nfs stable/nfs-server-provisioner \
     --set persistence.enabled=true,persistence.size=5Gi
   ```

2. Create a Kubernetes manifest file named `etcd.yaml`  to configure an NFS-backed PVC and a pod that uses it, as below. Replace the REGISTRY placeholder with a reference to your VMware Application  Catalog container registry.

   ```yaml
   ## etcd.yaml
   kind: PersistentVolumeClaim
   apiVersion: v1
   metadata:
     name: etcd-backup-pvc
   spec:
     accessModes:
       - ReadWriteMany
     resources:
       requests:
         storage: 2Gi
     storageClassName: nfs
   ---
   apiVersion: v1
   kind: Pod
   metadata:
     name: etcd-backup-pod
   spec:
     volumes:
       - name: etcd-backup
         persistentVolumeClaim:
           claimName: etcd-backup-pvc
     containers:
       - name: inspector
         image: REGISTRY/tac-shell:latest
         command:
           - sleep
           - infinity
         volumeMounts:
           - mountPath: "/backup"
             name: etcd-backup
   ```

3. Apply the manifest to the Kubernetes cluster:

   ```bash
   kubectl apply -f etcd.yaml
   ```

   This will create a pod named `etcd-backup-pod` with an attached PVC named `etcd-backup-pvc`. The PVC will be mounted at the `/backup` mount point of the pod.

4. Copy the snapshot to the PVC using the mount point:

   ```bash
   kubectl cp mybackup etcd-backup-pod:/backup/mybackup
   ```

5. Verify that the snapshot exists in the PVC, by connecting to the pod command-line shell and inspecting the `/backup` directory:

   ```bash
   kubectl exec -it etcd-backup-pod -- ls -al /backup
   ```

   The command output should display a directory listing containing the snapshot file, as shown below:

   [![Directory listing](https://docs.vmware.com/en/VMware-Tanzu-Application-Catalog/services/tac-docs-tutorials/Images/pvc-contents.png)](https://docs.vmware.com/en/VMware-Tanzu-Application-Catalog/services/tac-docs-tutorials/Images/pvc-contents.png)

6. Delete the pod, as it is not longer required:

   ```bash
   kubectl delete pod etcd-backup-pod
   ```

### Step 3: Restore the snapshot in a new cluster

The next step is to create an empty etcd deployment on the  destination cluster and restore the data snapshot into it. The VMware  Application Catalog etcd Helm chart provides built-in capabilities to do this, via its `startFromSnapshot.*` parameters.

1. Create a new etcd deployment. Replace the PASSWORD placeholder  with the same password used in the original deployment and replace the  REPOSITORY placeholder with a reference to your VMware Application  Catalog chart repository.

   ```bash
   helm install etcd-new REPOSITORY/etcd \
     --set startFromSnapshot.enabled=true \
     --set startFromSnapshot.existingClaim=etcd-backup-pvc \
     --set startFromSnapshot.snapshotFilename=mybackup \
     --set auth.rbac.rootPassword=PASSWORD \
     --set statefulset.replicaCount=3
   ```

   This command creates a new etcd cluster and initializes it using an existing data snapshot. The `startFromSnapshot.existingClaim` and `startFromSnapshot.snapshotFilename` define the source PVC and source filename for the snapshot respectively.

   **NOTE**: It is important to create the new deployment on the destination cluster using the same credentials as the original deployment on the source  cluster.

2. Connect to the new deployment and confirm that your data has  been successfully restored. Replace the PASSWORD placeholder with the  administrator password set at deployment-time.

   ```bash
   export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=etcd,app.kubernetes.io/instance=etcd-new" -o jsonpath="{.items[0].metadata.name}")
   kubectl exec -it $POD_NAME -- etcdctl --user root:PASSWORD get /message1
   kubectl exec -it $POD_NAME -- etcdctl --user root:PASSWORD get /message2
   ```

   Here is an example of what you should see:

   [![Query results](https://docs.vmware.com/en/VMware-Tanzu-Application-Catalog/services/tac-docs-tutorials/Images/tac-documentation-backup-restore-data-etcd-kubernetes-query-results.png)](https://docs.vmware.com/en/VMware-Tanzu-Application-Catalog/services/tac-docs-tutorials/Images/tac-documentation-backup-restore-data-etcd-kubernetes-query-results.png)

**NOTE**: VMware Application Catalog's etcd Helm chart  also supports auto disaster recovery by periodically snapshotting the  keyspace. If the cluster permanently loses more than (N-1)/2 members, it tries to recover the cluster from a previous snapshot. [Learn more about this feature](https://docs.vmware.com/en/VMware-Application-Catalog/services/apps/GUID-apps-etcd-administration-enable-disaster-recovery.html).

## Method 2: Back up and restore persistent data volumes

This method involves copying the persistent data volumes for the etcd nodes and reusing them in a new deployment with [Velero](https://velero.io/), an open source Kubernetes backup/restore tool. This method is only suitable when:

- The Kubernetes provider is [supported by Velero](https://velero.io/docs/latest/supported-providers/).
- Both clusters are on the same Kubernetes provider, as this is a requirement of [Velero's native support for migrating persistent volumes](https://velero.io/docs/latest/migration-case/).
- The restored deployment on the destination cluster will have the  same name, namespace, topology and credentials as the original  deployment on the source cluster.

**NOTE**: For persistent volume migration across cloud providers with Velero, you have the option of using Velero's [Restic integration](https://velero.io/docs/latest/restic/). This integration is not covered in this guide.

### Step 1: Install Velero on the source cluster

[Velero](https://velero.io/) is an open source tool that  makes it easy to backup and restore Kubernetes resources. It can be used to back up an entire cluster or specific resources such as persistent  volumes.

1. Modify your context to reflect the source cluster (if not already done).

2. Follow the [Velero plugin setup instructions for your cloud provider](https://velero.io/docs/latest/supported-providers/). For example, if you are using Google Cloud Platform (as this guide does), follow [the GCP plugin setup instructions](https://github.com/vmware-tanzu/velero-plugin-for-gcp##setup) to create a service account and storage bucket and obtain a credentials file.

3. Then, install Velero on the source cluster by executing the  command below, remembering to replace the BUCKET-NAME placeholder with  the name of your storage bucket and the SECRET-FILENAME placeholder with the path to your credentials file:

   ```bash
   velero install --provider gcp --plugins velero/velero-plugin-for-gcp:v1.0.0 --bucket BUCKET-NAME --secret-file SECRET-FILENAME
   ```

   You should see output similar to the screenshot below as Velero is installed:

   [![Velero installation](https://docs.vmware.com/en/VMware-Tanzu-Application-Catalog/services/tac-docs-tutorials/Images/tac-documentation-backup-restore-data-etcd-kubernetes-velero-installation.png)](https://docs.vmware.com/en/VMware-Tanzu-Application-Catalog/services/tac-docs-tutorials/Images/tac-documentation-backup-restore-data-etcd-kubernetes-velero-installation.png)

4. Confirm that the Velero deployment is successful by checking for a running pod using the command below:

   ```bash
   kubectl get pods -n velero
   ```

### Step 2: Back up the etcd deployment on the source cluster

Next, back up the persistent volumes using Velero.

1. Create a backup of the volumes in the running etcd deployment on the source cluster. This backup will contain all the node volumes.

   ```bash
   velero backup create etcd-backup --include-resources pvc,pv --selector app.kubernetes.io/instance=etcd
   ```

2. Execute the command below to view the contents of the backup and confirm that it contains all the required resources:

   ```bash
   velero backup describe etcd-backup  --details
   ```

3. To avoid the backup data being overwritten, switch the bucket to read-only access:

   ```bash
   kubectl patch backupstoragelocation default -n velero --type merge --patch '{"spec":{"accessMode":"ReadOnly"}}'
   ```

### Step 3: Restore the etcd deployment on the destination cluster

You can now restore the persistent volumes and integrate them with a new etcd deployment on the destination cluster.

1. Modify your context to reflect the destination cluster.

2. Install Velero on the destination cluster as described in Step  1. Remember to use the same values for the BUCKET-NAME and  SECRET-FILENAME placeholders as you did originally, so that Velero is  able to access the previously-saved backups.

   ```bash
   velero install --provider gcp --plugins velero/velero-plugin-for-gcp:v1.0.0 --bucket BUCKET-NAME --secret-file SECRET-FILENAME
   ```

3. Confirm that the Velero deployment is successful by checking for a running pod using the command below:

   ```bash
   kubectl get pods -n velero
   ```

4. Restore the persistent volumes in the same namespace as the source cluster using Velero.

   ```bash
   velero restore create --from-backup etcd-backup
   ```

5. Confirm that the persistent volumes have been restored:

   ```bash
   kubectl get pvc --namespace default
   ```

6. Create a new etcd deployment. Use the same deployment name,  credentials and other parameters as the original deployment. Replace the PASSWORD placeholder with the same database password used in the  original release and the REPOSITORY placeholder with a reference to your VMware Application Catalog chart repository.

   ```bash
   helm install etcd REPOSITORY/etcd \
       --set auth.rbac.rootPassword=PASSWORD \
       --set statefulset.replicaCount=3
   ```

   **NOTE**: It is important to create the  new deployment on the destination cluster using the same namespace,  deployment name, credentials and number of replicas as the original  deployment on the source cluster.

   This will create a new deployment that uses the original volumes (and hence the original data).

7. Connect to the new deployment and confirm that your original data is intact:

   ```bash
   export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=etcd,app.kubernetes.io/instance=etcd" -o jsonpath="{.items[0].metadata.name}")
   kubectl exec -it $POD_NAME -- etcdctl --user root:PASSWORD get /message1
   kubectl exec -it $POD_NAME -- etcdctl --user root:PASSWORD get /message2
   ```

   Here is an example of what you should see:

   [![Query results](https://docs.vmware.com/en/VMware-Tanzu-Application-Catalog/services/tac-docs-tutorials/Images/tac-documentation-backup-restore-data-etcd-kubernetes-query-results.png)](https://docs.vmware.com/en/VMware-Tanzu-Application-Catalog/services/tac-docs-tutorials/Images/tac-documentation-backup-restore-data-etcd-kubernetes-query-results.png)

## Useful links

- [VMware Application Catalog etcd Helm chart](https://docs.vmware.com/en/VMware-Application-Catalog/services/apps/GUID-apps-etcd-index.html)
- etcd client application [`etcdctl`](https://github.com/etcd-io/etcd/tree/master/etcdctl)
- [Velero documentation](https://velero.io/docs/latest/)