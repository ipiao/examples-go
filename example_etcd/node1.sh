# For each machine
TOKEN=my-etcd-token-1
CLUSTER_STATE=new

NAME_1=etcd-node-1
NAME_2=etcd-node-2
NAME_3=etcd-node-3

HOST_1=127.0.0.1
HOST_2=127.0.0.1
HOST_3=127.0.0.1

CLUSTER=${NAME_1}=http://${HOST_1}:2380,${NAME_2}=http://${HOST_2}:2381,${NAME_3}=http://${HOST_3}:2382

# curl https://discovery.etcd.io/new?size=3
DISCOVERY=https://discovery.etcd.io/2e7e07b94bdb96107edb6b20baa5cd3c

# For node 1
THIS_NAME=${NAME_1}
THIS_IP=${HOST_1}
DATADIR=${THIS_NAME}.etcd

#rm -rf ${DATADIR}

etcd --data-dir=${DATADIR} --name ${THIS_NAME} \
  --initial-advertise-peer-urls http://${THIS_IP}:2380 \
	--listen-peer-urls http://${THIS_IP}:2380 \
	--advertise-client-urls http://${THIS_IP}:2379 \
	--listen-client-urls http://${THIS_IP}:2379 \
  --initial-cluster ${CLUSTER} \
	--initial-cluster-state ${CLUSTER_STATE} \
	--initial-cluster-token ${TOKEN} \
	--log-outputs 'stdout' \
	--log-level 'debug'
#	--logger 'zap' \
#	--discovery ${DISCOVERY} \

