#!/bin/bash

set -e

echo "[1/6] Creating namespaces..."
kubectl create namespace tidb-admin || true
kubectl create namespace tidb-cluster || true

echo "[2/6] Installing TiDB CRDs..."
kubectl create -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.6.1/manifests/crd.yaml || true

echo "[3/6] Adding PingCAP Helm repo..."
helm repo add pingcap https://charts.pingcap.org/ || true
helm repo update

echo "[4/6] Installing TiDB Operator..."
helm install tidb-operator pingcap/tidb-operator \
  --namespace tidb-admin \
  --version v1.6.1 || true

echo "[5/6] Applying TiDB Cluster definition..."
kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.6.1/examples/basic/tidb-cluster.yaml

# echo "Deploying Monitor (Prometheus + Grafana)..."
# kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.6.1/examples/basic/tidb-monitor.yaml

# echo "Deploying Dashboard..."
# kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.6.1/examples/basic/tidb-dashboard.yaml

# echo "Waiting for TiDB pods to be ready..."
# kubectl wait --for=condition=Ready pod -l app.kubernetes.io/component=tidb -n tidb-cluster --timeout=300s

# echo "🌐 Port-forward to access TiDB (SQL) and Dashboard..."
# echo "  ➜ SQL:       mysql --host=127.0.0.1 --port=4000 -u root"
# echo "  ➜ Dashboard: http://localhost:12333"
# kubectl port-forward svc/basic-tidb 4000:4000 -n tidb-cluster &
# kubectl port-forward svc/basic-tidb-dashboard-exposed 12333:12333 -n tidb-cluster &

echo "[6/6] Deploy ZCS system to Kubernetes..."
kubectl apply -f ./k8s || true

echo "ZCS deployed successfully!"
