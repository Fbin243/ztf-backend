# Zalopay Coupon System (ZCS)

Microservices-based coupon management system for Zalopay Tech Fresher 2025.

## Prerequisites

Install these tools (see official docs for installation instructions):

- **[Docker](https://docs.docker.com/get-docker/)** & **[Docker Compose](https://docs.docker.com/compose/install/)** (Optional)
- **Kubernetes** cluster: Use [Orbstack](https://orbstack.dev/docs/kubernetes) as a control plane node and a worker node.
- **[kubectl](https://kubernetes.io/docs/tasks/tools/)**
- **[Helm](https://helm.sh/docs/intro/install/)** (for K8s deployment)
- **[Go](https://go.dev/doc/install)** (for local development)
- **[Make](https://www.gnu.org/software/make/)**

Quick verification:

```bash
docker --version && kubectl version --client && helm version
```

## Deployment

### Kubernetes

```bash
# 1. Clone and setup
git clone <repository-url>
cd ztf-backend

# 2. Deploy TiDB cluster + ZCS 
make k8s-up

# 3. Wait some minutes for deploying ... 

# 4. Forward port for accessing tidb by DBeaver (Optional)
kubectl port-forward svc/basic-tidb 4000:4000 -n tidb-cluster

# 5. Delete TiDB cluster + ZCS
make k8s-down

# 6. Wait a few minutes for the deletion process to complete (especially namespace removal may take longer).
```

### Docker Compose

```bash
# 1. Clone and setup
git clone <repository-url>
cd ztf-backend

# 2. Create environment file .env.dev by copying .env.
cp env.example env.dev

# 3. Deploy a tidb local by tidb playgroud
tiup playground --tag ztf_db

# 4. Start ZCS
make up

# 5. Stop ZCS
make down
```

## Testing

```bash
# Load testing with K6
cd k6
npm install
k6 run src/order-test.js

# Unit tests
make test
```
