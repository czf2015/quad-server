# 先安装nvm，利用nvm安装指定版本node，再安装node本地包
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.0/install.sh | bash

export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

nvm install 12.13.0 && npm install

#################################################################################################

apt install -y mysql-server

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
if [ -z `which go` ]; then
    wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    rm go1.13.linux-amd64.tar.gz
fi

# go get -u github.com/go-sql-driver/mysql
# go get -u github.com/gin-gonic/gin
# go get -u github.com/gin-gonic/contrib/sessions
# go get -u golang.org/x/oauth2/google
# go get -u google.golang.org/api/gmail/v1