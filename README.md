# sysKubernetes

I created an account in AWS and I created a cluster and I added a nodegroup with one node.

![aws](https://github.com/carlos2380/webCarlos2380/blob/master/aws.png)

```
eksctl create cluster 
  -name test-cluster  
  --without-nodegroup 
```

```
eksctl create nodegroup   
  --cluster test-cluster   
   --region us-east-2   
   --name test-workers   
   --node-type m5.large   
   --nodes 1   
   -nodes-min 1   
   --nodes-max 3   
   --asg-access
```

I programmed two apps in go. One to get the Sys info and write it in a txt as JSON and other to read the txt and print the info as logs.


- Go get Sys info: https://github.com/carlos2380/sysKubernetes/blob/main/systeminfo/main.go
- Go print logs: https://github.com/carlos2380/sysKubernetes/blob/main/printsysinfo/main.go


Then I created the Dockerfiles and I update the both images to DockerHub


- image get sys info: https://hub.docker.com/repository/docker/carlos2380/sysinfo
- image print logs: https://hub.docker.com/repository/docker/carlos2380/printinfo


Finally, I created the Kubernetes yaml and apply it in the cluster. I set Kind as a Job because only runs one time.

- yaml: https://github.com/carlos2380/sysKubernetes/blob/main/sysprintinfo.yaml

```
kubectl apply -f sysprintinfo.yaml 
```

![k8s](https://github.com/carlos2380/webCarlos2380/blob/master/k8s.png)

- describe pod: https://github.com/carlos2380/sysKubernetes/blob/main/describe_job_sysinfo.txt



