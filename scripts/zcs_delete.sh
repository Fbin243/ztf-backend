#!/bin/bash

set -e

echo "[1/7] Stop port forwarding..."
pgrep -lfa kubectl || true

echo "[2/7] Deleting TiDB Cluster resources..."
# kubectl -n tidb-cluster delete -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.6.1/examples/basic/tidb-dashboard.yaml || true
# kubectl -n tidb-cluster delete -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.6.1/examples/basic/tidb-monitor.yaml || true
kubectl -n tidb-cluster delete -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.6.1/examples/basic/tidb-cluster.yaml || true

echo "[3/7] Delete PV data ..."
kubectl delete pvc -n tidb-cluster -l app.kubernetes.io/instance=basic,app.kubernetes.io/managed-by=tidb-operator || true
kubectl get pv -l app.kubernetes.io/namespace=tidb-cluster,app.kubernetes.io/managed-by=tidb-operator,app.kubernetes.io/instance=basic -o name | xargs -I {} kubectl patch {} -p '{"spec":{"persistentVolumeReclaimPolicy":"Delete"}}' || true

echo "[4/7] Deleting all TiDB-related CRDs..."
kubectl get crd | grep pingcap.com | awk '{print $1}' | xargs -r kubectl delete crd || true

echo "[5/7] Uninstalling TiDB Operator (Helm)..."
helm uninstall tidb-operator -n tidb-admin || true

echo "[6/7] Deleting namespaces..."
kubectl delete ns tidb-cluster || true
kubectl delete ns tidb-admin || true

echo "[7/7] Removing ZCS from Kubernetes..."
kubectl delete -f ./k8s || true

echo "TiDB cluster and ZCS have been deleted completely."
