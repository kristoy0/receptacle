Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"
  config.vm.provision :shell, path: "bootstrap.sh", privileged: true, binary: false
  config.vm.synced_folder "./", "/home/vagrant/go/src/github.com/kristoy0/receptacle/"
  config.vm.network 'forwarded_port', guest: 8500, host: 8500
end