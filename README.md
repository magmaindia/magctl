# magctl

magctl is CLI Tool for magma, through which we will able to get results of list of tenants,feg,gateways on terminal itself. It will work similar to as kubectl in kubernetes. We are trying to use config in magctl similar to config used in kubernetes under .magmacli folder, so other can also easily access it.

![](/readme_photos/diagm.png)

# How to use
First we need to provide the certificates generates during the deployment of Magma Orchestrator under the folder **magctl/cmd/** <br />
After that we need to follow these steps: <br />
```bash
git clone magctl
install go
cd magctl
go install
go build
sudo mv magctl /usr/local/bin
```

Now we are able to use **magctl** command on our terminal. To verify it, run
```bash
magctl get tenants
```
![](/readme_photos/result2.png)

**Note:** Remember to provide the correct certificates.
