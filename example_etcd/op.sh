
# get member ID
export ETCDCTL_API=3
HOST_1=127.0.0.1
HOST_2=127.0.0.1
HOST_3=127.0.0.1
HOST_4=127.0.0.1
ENDPOINTS=${HOST_1}:2379,${HOST_2}:2378,${HOST_3}:2377,${HOST_3}:2376

etcdctl --endpoints=${ENDPOINTS} member list

# remove the member
MEMBER_ID=278c654c9a6dfd3b
etcdctl --endpoints=${ENDPOINTS} member remove ${MEMBER_ID}

# add a new member (node 4)
TOKEN=my-etcd-token-1
CLUSTER_STATE=new

NAME_1=etcd-node-1
NAME_2=etcd-node-2
NAME_3=etcd-node-3

HOST_1=127.0.0.1
HOST_2=127.0.0.1
HOST_3=127.0.0.1
ENDPOINTS=${HOST_1}:2379,${HOST_2}:2378
etcdctl --endpoints=${ENDPOINTS} member add ${NAME_3} \
	--peer-urls=http://${HOST_3}:2377