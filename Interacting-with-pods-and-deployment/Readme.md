## Steps Followed

1. First created a golang application which interacts with the k8s api server to list pods and deployment resources.
2. Then build this golang app and dockerize the application and pushed to dockerhub.

3. Create a deployment using the docker image .
```
kubectl create deployment lister --image lister:1.0.1 
```

4. Then checked the logs . Here i found out that there is some authorization issue. To solve this we have to create a role and role binding for the default service account .

#### Role
```
kubectl create role lister -n default --verb list --resource pod,deployment
```

#### Role Binding

```
kubectl create rolebinding listerrb -n default --role lister --serviceaccount default:default
```

5. Again created deployment and checked the logs . Finally i observe it successfully fetches all the resources mentioned the application . 