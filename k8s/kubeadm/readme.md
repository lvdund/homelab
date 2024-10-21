## Setup cluster with kubeadm

!! Need to check version

- Disable apparmor (Optional)
```bash
systemctl stop apparmor
systemctl disable apparmor
apt remove apparmor -y
apt autoremove
```

- Config network:
```bash
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

modprobe overlay
modprobe br_netfilter

# sysctl params required by setup, params persist across reboots
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF

# Apply sysctl params without reboot
sudo sysctl --system

lsmod | grep br_netfilter
lsmod | grep overlay

sysctl net.bridge.bridge-nf-call-iptables net.bridge.bridge-nf-call-ip6tables net.ipv4.ip_forward
```

- Install containerd:
```bash
apt-get install containerd
mkdir /etc/containerd
containerd config default > /etc/containerd/config.toml
```
```bash
# Config containerd - enable systemd:
nano /etc/containerd/config.toml
# find ``SystemdCgroup`` --> ``SystemdCgroup = true``
```
```bash
systemctl restart containerd.service
ps -ef|grep containerd
```

- Install Kubernetes:
```bash
apt-get install -y apt-transport-https ca-certificates curl gpg
mkdir /etc/apt/keyrings/
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.29/deb/Release.key | gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.29/deb/Release.key | gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.29/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list

apt-get install -y kubelet kubeadm kubectl
apt-get update
apt-get install -y kubelet kubeadm kubectl
apt-mark hold kubelet kubeadm kubectl
```

- Build master node:
```bash
kubeadm init --apiserver-advertise-address=192.168.56.101 --pod-network-cidr=10.244.0.0/16
# Installing a Pod network add-on
mkdir /opt/sofware
cd /opt/sofware
wget https://github.com/flannel-io/flannel/releases/latest/download/kube-flannel.yml
kubectl apply -f kube-flannel.yml 
```