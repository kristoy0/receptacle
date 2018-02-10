GO_VERSION=1.9.3

apt-get update
apt-get install git -y
cd /tmp

curl -O https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz
tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz

echo "PATH=$PATH:/usr/local/go/bin:/home/vagrant/go/bin
GOPATH=/home/vagrant/go" >> /etc/profile
chown -R vagrant /home/vagrant/go

add-apt-repository ppa:masterminds/glide && sudo apt-get update
apt-get install -y glide

apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

apt-get update

apt-get install -y docker-ce
usermod -aG docker vagrant

cd /home/vagrant/go/src/github.com/kristoy0/receptacle

go get -v github.com/micro/micro

docker run -p 8500:8500 -p 8300:8300 -d \
  consul agent -server \
  -client='{{ GetInterfaceIP "eth0" }}' \
  -bind='{{ GetInterfaceIP "eth0" }}' \
  -bootstrap-expect=1 \
  -node=dev \
  -ui
