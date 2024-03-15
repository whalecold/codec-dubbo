#!/bin/bash

export KITEX_XDS_METAS={"CLUSTER_ID":"fakeclusterid","DNS_AUTO_ALLOCATE":"true","DNS_CAPTURE":"true","INSTANCE_IPS":"","ISTIO_VERSION":"1.16.3","MESH_ID":"mse-e2e-mesh-id","NAMESPACE":"e2e-public-default-group"}
export POD_NAMESPACE=e2e-public-default-group
export POD_NAME=helloworld
export INSTANCE_IP=127.0.0.1
export KITEX_XDS_ISTIO_ADDR=127.0.0.1:15010
