# volume uss nfs

yudady/jenkins:lts-jdk11

docker run -d --privileged --restart=always \
-v /tmp:/nfs \
-e NFS_EXPORT_DIR_1=/nfs \
-e NFS_EXPORT_DOMAIN_1=\* \
-e NFS_EXPORT_OPTIONS_1=ro,insecure,no_subtree_check,no_root_squash,fsid=1 \
-p 111:111 -p 111:111/udp \
-p 2049:2049 -p 2049:2049/udp \
-p 32765:32765 -p 32765:32765/udp \
-p 32766:32766 -p 32766:32766/udp \
-p 32767:32767 -p 32767:32767/udp \
fuzzle/docker-nfs-server:latest



docker pull atlassian/nfs-server-test



https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner








