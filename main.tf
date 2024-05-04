provider "yandex" {
  token     = "your_yandex_cloud_token"
  cloud_id  = "your_cloud_id"
  folder_id = "your_folder_id"
  zone      = "ru-central1-a"
}

resource "yandex_compute_instance" "servers" {
  count         = 4
  name          = "server-${count.index}"
  zone          = "ru-central1-a"
  platform_id   = "standard-v2"
  resources {
    cores  = 2
    memory = 2
  }

  boot_disk {
    initialize_params {
      image_id = "your_image_id"
    }
  }

  network_interface {
    subnet_id = "your_subnet_id"
    nat       = true
  }

  metadata = {
    user-data = <<-EOF
      #!/bin/bash
      sudo apt-get update
      sudo apt-get install -y docker.io
      sudo systemctl start docker
      sudo systemctl enable docker
      curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
      echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee -a /etc/apt/sources.list.d/kubernetes.list
      sudo apt-get update
      sudo apt-get install -y kubelet kubeadm kubectl
      sudo swapoff -a
      sudo kubeadm init --pod-network-cidr=10.244.0.0/16
      mkdir -p $HOME/.kube
      sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
      sudo chown $(id -u):$(id -g) $HOME/.kube/config
    EOF
  }
}

output "server_ips" {
  value = yandex_compute_instance.servers[*].network_interface[0].nat_ip_address
}
