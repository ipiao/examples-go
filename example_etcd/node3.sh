# For each machine
TOKEN=my-etcd-token-1
CLUSTER_STATE=new

NAME_1=etcd-node-1
NAME_2=etcd-node-2
NAME_3=etcd-node-3
NAME_4=etcd-node-4

HOST_1=127.0.0.1
HOST_2=127.0.0.1
HOST_3=127.0.0.1
HOST_4=127.0.0.1

CLUSTER=${NAME_1}=http://${HOST_1}:2380,${NAME_2}=http://${HOST_2}:2381,${NAME_3}=http://${HOST_3}:2382

# curl https://discovery.etcd.io/new?size=3
DISCOVERY=https://discovery.etcd.io/aaf7408e6e8034c4c5d06416c737aa94

# For node 3
THIS_NAME=${NAME_3}
THIS_IP=${HOST_3}
DATADIR=${THIS_NAME}.etcd

rm -rf ${DATADIR}

etcd --data-dir=${DATADIR} --name ${THIS_NAME} \
	--initial-advertise-peer-urls http://${THIS_IP}:2382 \
	--listen-peer-urls http://${THIS_IP}:2382 \
	--advertise-client-urls http://${THIS_IP}:2377 \
	--listen-client-urls http://${THIS_IP}:2377 \
  --initial-cluster ${CLUSTER} \
	--initial-cluster-state ${CLUSTER_STATE} \
	--initial-cluster-token ${TOKEN} \
	--log-outputs 'stdout' \
	--log-level 'debug'
#	--discovery ${DISCOVERY} \
#	--logger 'zap' \
