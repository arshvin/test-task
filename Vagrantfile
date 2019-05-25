# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "bento/centos-7.4"

  # config.vm.box_check_update = false
  config.vm.network "forwarded_port", guest: 8080, host: 8080
  config.vm.network "forwarded_port", guest: 8443, host: 8443

  config.vm.hostname = "testhost.local"
  # config.vm.network "public_network"
  
  config.vm.provision "shell", inline: <<-SHELL
     yum -y install epel-release
     yum -y install ansible
     
     ansible-playbook /vagrant/provision/vagrant.yml
  SHELL
end
