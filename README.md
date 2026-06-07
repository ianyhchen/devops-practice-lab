# devops-practice-api

A local cloud-native DevOps practice project built on macOS using Go, Docker, Kubernetes, GitHub Actions, Prometheus/Grafana, and ArgoCD.

This project was created to practice and understand modern DevOps and cloud-native workflows, including containerization, CI/CD, Kubernetes deployment, GitOps, observability, and troubleshooting.

---

# Architecture

```text
GitHub Repository
        ↓
GitHub Actions CI
        ↓
Docker Image (OrbStack)
        ↓
minikube Kubernetes Cluster
        ↓
Deployment / Service / ConfigMap
        ↓
Prometheus + Grafana Monitoring
        ↓
ArgoCD GitOps Sync
```

---

# Features

* Go backend API service
* Unit testing with Go test
* Multi-stage Docker build
* GitHub Actions CI pipeline
* Kubernetes Deployment / Service / ConfigMap
* Readiness and liveness probes
* Resource requests and limits
* Local Kubernetes cluster with minikube
* Prometheus/Grafana monitoring stack
* ArgoCD GitOps deployment
* Kubernetes troubleshooting practice

---

# Tech Stack

| Category           | Tools                |
| ------------------ | -------------------- |
| Backend            | Go                   |
| Container Runtime  | OrbStack             |
| Containerization   | Docker               |
| Kubernetes         | minikube             |
| CI                 | GitHub Actions       |
| GitOps             | ArgoCD               |
| Monitoring         | Prometheus + Grafana |
| Package Management | Helm                 |
| Version Control    | Git + GitHub         |

---

# Project Structure

```text
devops-practice-api/
├── app/
│   ├── main.go
│   ├── main_test.go
│   └── go.mod
│
├── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── configmap.yaml
│
├── .github/
│   └── workflows/
│       └── ci.yml
│
├── Dockerfile
├── .gitignore
└── README.md
```

---

# Local Setup

## Prerequisites

* OrbStack
* Docker
* kubectl
* minikube
* Helm
* ArgoCD CLI
* Go

---

# Run Locally

## Start the API

```bash
cd app

go test ./...
go run main.go
```

Test API:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
curl http://localhost:8080/api/v1/items
```

---

# Docker

## Build Image

```bash
docker build -t devops-practice-api:local .
```

## Run Container

```bash
docker run --rm -p 8080:8080 devops-practice-api:local
```

---

# GitHub Actions CI

The CI pipeline includes:

* Checkout source code
* Setup Go
* Run unit tests
* Build Docker image

Workflow file:

```text
.github/workflows/ci.yml
```

---

# Kubernetes Deployment

## Start minikube

```bash
minikube start --driver=docker
```

## Create namespace

```bash
kubectl create namespace devops-practice-lab
```

## Load local Docker image into minikube

```bash
minikube image load devops-practice-api:local
```

## Deploy resources

```bash
kubectl apply -n devops-practice-lab -f k8s/
```

## Verify resources

```bash
kubectl get pods -n devops-practice-lab
kubectl get svc -n devops-practice-lab
```

---

# Port Forward

```bash
kubectl port-forward svc/devops-practice-api 8080:80 -n devops-practice-lab
```

Test API:

```bash
curl http://localhost:8080/healthz
```

---

# Kubernetes Concepts Practiced

## Deployment

* Replica management
* Rolling updates
* Self-healing
* Resource limits
* Health checks

## Service

* Stable networking endpoint
* Service discovery
* Internal load balancing

## ConfigMap

* Environment variable injection
* Configuration management

---

# Health Checks

## Readiness Probe

Used to determine whether the pod is ready to receive traffic.

## Liveness Probe

Used to determine whether the container should be restarted.

---

# Troubleshooting Practice

The following Kubernetes troubleshooting scenarios were practiced:

## ImagePullBackOff

* Wrong image tag
* Invalid image configuration

## CrashLoopBackOff

* Invalid startup command
* Application startup failure

## Readiness Probe Failure

* Wrong health check path
* Pod running but not receiving traffic

## Common Debug Commands

```bash
kubectl get pods
kubectl describe pod <pod-name>
kubectl logs <pod-name>
kubectl logs <pod-name> --previous
kubectl get endpoints
```

---

# Monitoring Stack

Installed using Helm:

```bash
helm install monitoring prometheus-community/kube-prometheus-stack \
  -n monitoring \
  --create-namespace
```

Practiced monitoring:

* Pod CPU usage
* Memory usage
* Pod restart count
* Namespace resource usage

---

# ArgoCD GitOps

ArgoCD was used to practice GitOps deployment workflows.

## GitOps Flow

```text
Git Push
↓
ArgoCD detects repository changes
↓
Automatic Kubernetes sync
↓
Cluster state updated
```

## Features Practiced

* Automated sync
* Self-healing
* Declarative deployment
* Git as source of truth

---

# ArgoCD Setup

## Install ArgoCD

```bash
kubectl create namespace argocd

kubectl create -n argocd \
-f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

## Access UI

```bash
kubectl port-forward svc/argocd-server -n argocd 8081:443
```

Open:

```text
https://localhost:8081
```

---

# Lessons Learned

This project helped practice:

* Cloud-native deployment workflows
* CI/CD pipeline concepts
* Kubernetes networking and health checks
* GitOps deployment models
* Kubernetes operational troubleshooting
* Monitoring and observability concepts
* Production-style deployment thinking

---

# Future Improvements

* Add Ingress controller
* Add OpenTelemetry tracing
* Add Helm chart support
* Add canary deployment strategy
* Push Docker images to container registry
* Add integration tests
* Add Terraform for infrastructure management
* Add Kafka / Redis integration

---

# References

* Kubernetes
* ArgoCD
* Prometheus
* Grafana
* GitHub Actions
* minikube
* OrbStack
