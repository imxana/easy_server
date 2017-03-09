静态文件服务器,总有一款适合你...

## NodeJS:

    cd NodeJS
    npm install
    node app.js

### or

    npm install -g node-static
    static -p 8000

### or
    
    npm install -g http-server
    http-server -p 8000

## Python 3.x

    cd Python3
    pip3 install -r requirements
    python3 app.py

### or

    python3 -m http.server 8000

## Python 2.x

    python -m SimpleHTTPServer 8000

### or

    python -c 'from twisted.web.server import Site; from twisted.web.static import File; from twisted.internet import reactor; reactor.listenTCP(8000, Site(File("."))); reactor.run()'

## Go

    cd Golang
    go run app.go

## Ruby on Rails

    ruby -run -e httpd . -p 8000

### or

    ruby -rwebrick -e'WEBrick::HTTPServer.new(:Port => 8000, :DocumentRoot => Dir.pwd).start'

### or

    gem install adsf
    adsf -p 8000

## PHP

    php -S 127.0.0.1:8000
