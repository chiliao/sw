# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.require_version ">=2.0.2"

BOX_IMAGE = "turin/centos74"
BOX_VERSION = "0"

VENICE_VER = "v0.1"
VENICE_DOCKER_IMAGE = "pensando/venice:#{VENICE_VER}"
VENICE_IMAGE_FILE = "venice-sim.tar"
VENICE_CPUS_PER_NODE = "4"
VENICE_MEMORY_PER_NODE = "4096"
VENICE_PORT=10001

# Ensure docker creds are set if image is not provided locally
if File.exist?(VENICE_IMAGE_FILE) == false
    unless ENV["DOCKER_USERNAME"] && ENV["DOCKER_PASSWORD"]
      puts "Please set your DOCKER_USERNAME, DOCKER_PASSWORD environment variables"
      exit
    end

    docker_user = ENV["DOCKER_USERNAME"]
    docker_password = ENV["DOCKER_PASSWORD"]
end

Vagrant.configure("2") do |config|
  config.vm.box = BOX_IMAGE

      config.vm.define  "venice" do |subconfig|
          subconfig.vm.box = BOX_IMAGE
          subconfig.vm.box_version = BOX_VERSION
          subconfig.vm.box = "venice/centos74"
          subconfig.vm.box_version = "0.8"
          config.ssh.insert_key = false
          subconfig.vm.provision :shell, inline: 'ulimit -n 4096'
          subconfig.vm.hostname = "venice"

          # control network (eth1)
          subconfig.vm.network :private_network, ip: "0.0.0.0", virtualbox__intnet: "control_net", auto_config: false

          # venice port forwarding (out of the VM)
          subconfig.vm.network  "forwarded_port", guest: VENICE_PORT, host: VENICE_PORT

          subconfig.vm.provider "virtualbox" do |vb|
              vb.memory = VENICE_MEMORY_PER_NODE
              vb.cpus = VENICE_CPUS_PER_NODE
              vb.customize ['modifyvm', :id, '--nictype2', '82545EM']
              vb.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
              vb.linked_clone = true # use base image and clone from it. for multi-VM, saves space
          end

          subconfig.vm.provision "shell", inline: <<-SCRIPT
            systemctl stop firewalld && systemctl disable firewalld
            systemctl stop chronyd && systemctl disable chronyd
            systemctl stop kubelet && systemctl disable kubelet
            setenforce 0
            usermod -a -G root vagrant
            usermod -a -G docker vagrant
            mkdir -p /etc/docker
            echo '{ "insecure-registries" : ["registry.test.pensando.io:5000"] }' > /etc/docker/daemon.json
            systemctl start docker
            systemctl enable docker
            ip link set up dev eth1
            sysctl -w vm.max_map_count=262144
            chmod 777 /var/run/docker.sock
          SCRIPT

          if File.exist?(VENICE_IMAGE_FILE) == false
              subconfig.vm.provision "shell", inline: <<-SCRIPT
                  docker login -u #{docker_user} -p #{docker_password}
                  docker pull #{VENICE_DOCKER_IMAGE}
              SCRIPT
          end

          subconfig.vm.provision "shell", privileged: false, inline: <<-SCRIPT
              curl --silent --location https://rpm.nodesource.com/setup_8.x | sudo bash -
              sudo yum install -y nodejs
              sudo npm install -g newman

              mkdir -p /home/vagrant/venice
              if ! [ -f /vagrant/#{VENICE_IMAGE_FILE} ]; then
                  docker run --user $(id -u):$(id -g) --rm -v /vagrant:/import #{VENICE_DOCKER_IMAGE}
              fi
              cd /home/vagrant/venice
              tar xvf /vagrant/#{VENICE_IMAGE_FILE}
              docker load -i /home/vagrant/venice/pen-dind.tar
              docker load -i /home/vagrant/venice/pen-e2e.tar
              /home/vagrant/venice/venice-bootstrap.sh
          SCRIPT
      end
end
