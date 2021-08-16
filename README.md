# sysKubernetes

eksctl create cluster --name test-cluster --without-nodegroup
eksctl create nodegroup   --cluster test-cluster   --region us-east-2   --name test-workers   --node-type m5.large   --nodes 1   --nodes-min 1   --nodes-max 3   --asg-access

kubectl apply -f sysprintinfo.yaml 