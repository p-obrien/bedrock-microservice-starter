# Infrastructure

This Terraform code will deploy a new VPC with a Graviton based EKS Cluster in your AWS environment. Graviton was chosen for this starter kit because it's highly performant and keeps the project costs to a minimum.

To use this code you will need to create a file called terraform.tfvars and add the following:
```
cluster_name = "<the name of the cluster you would like to use>"
region = "< aws region to use>"
```

Then ensure your AWS Environment variables are set appropriately.

Then execute:
```
terraform init
terraform apply
```

After the cluster is created you will recieve a command to update your Kubeconf file with the cluster details.