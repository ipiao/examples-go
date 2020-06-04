# For each machine
TOKEN=my-etcd-token-1
CLUSTER_STATE=existing

NAME_1=etcd-node-1
NAME_2=etcd-node-2
NAME_3=etcd-node-3
NAME_4=etcd-node-4

HOST_1=127.0.0.1
HOST_2=127.0.0.1
HOST_3=127.0.0.1
HOST_4=127.0.0.1

CLUSTER=${NAME_1}=http://${HOST_1}:2380,${NAME_2}=http://${HOST_2}:2381,${NAME_4}=http://${HOST_4}:2383

# curl https://discovery.etcd.io/new?size=3
DISCOVERY=https://discovery.etcd.io/2e7e07b94bdb96107edb6b20baa5cd3c

# For node 3
THIS_NAME=${NAME_4}
THIS_IP=${HOST_4}
DATADIR=${THIS_NAME}.etcd

rm -rf ${DATADIR}

etcd --data-dir=${DATADIR} --name ${THIS_NAME} \
	--initial-advertise-peer-urls http://${THIS_IP}:2383 \
	--listen-peer-urls http://${THIS_IP}:2383 \
	--advertise-client-urls http://${THIS_IP}:2376 \
	--listen-client-urls http://${THIS_IP}:2376 \
  --initial-cluster ${CLUSTER} \
	--initial-cluster-state ${CLUSTER_STATE} \
	--initial-cluster-token ${TOKEN} \
	--log-outputs 'stdout' \
	--log-level 'debug'
#	--logger 'zap' \
#	--discovery ${DISCOVERY} \