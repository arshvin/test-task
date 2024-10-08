# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "bento/centos-stream-9"

  # config.vm.box_check_update = false
  config.vm.network "forwarded_port", guest: 8080, host: 8080
  config.vm.network "forwarded_port", guest: 8443, host: 8443

  config.vm.hostname = "testhost.local"
  # config.vm.network "public_network"

  config.vm.provider "virtualbox" do |v|
    v.memory = 2048
    v.cpus = 2
  end

  config.vm.provision "shell", inline: <<-SHELL
     dnf -y install epel-release
     dnf -y install ansible
     dnf -y install docker # If it's required to build docker images through ./gradlew docker

     ansible-playbook /vagrant/provision/vagrant.yml --diff
  SHELL
end
