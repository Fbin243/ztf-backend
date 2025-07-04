# Zalopay Coupon System (ZCS)

Microservices-based coupon management system for Zalopay Tech Fresher 2025.

## Prerequisites

Install these tools (see official docs for installation instructions):

- **[Docker](https://docs.docker.com/get-docker/)**: Platform for building, running, and managing containers.
- **[Docker Compose](https://docs.docker.com/compose/install/)** (optional): Tool for defining and running multi-container Docker applications.
- **[Kubernetes cluster](https://kubernetes.io/)**: Container orchestration platform; you can use **[Orbstack](https://orbstack.dev/docs/kubernetes)** for local clusters.
- **[kubectl](https://kubernetes.io/docs/tasks/tools/)**: Command-line tool to interact with Kubernetes clusters.
- **[Helm](https://helm.sh/docs/intro/install/)**: Package manager for Kubernetes, used to manage charts and deployments.
- **[Go](https://go.dev/doc/install)**: Programming language required for backend development.
- **[Make](https://www.gnu.org/software/make/)**: Build automation tool to run common development tasks.

Quick verification:

```bash
docker --version && kubectl version --client && helm version && go version && make -v
```

## Deployment

### Kubernetes

```bash
# 1. Clone and setup
git clone <repository-url>
cd ztf-backend

# 2. Enable kubernetes in orbstack and check kubectl context and set it to orbstack (if needed)
kubectl config current-context # View current context
kubectl config get-contexts # View all contexts and the currently active one
kubectl config use-context orbstack # Switch context

# 3. Deploy TiDB cluster + ZCS 
make k8s-up

# 4. Wait some minutes for deploying ... 

# 5. Forward port for accessing tidb by DBeaver (Optional)
kubectl port-forward svc/basic-tidb 4000:4000 -n tidb-cluster

# 6. Delete TiDB cluster + ZCS
make k8s-down

# 7. Wait a few minutes for the deletion process to complete (especially namespace removal may take longer).
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
