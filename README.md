# hello-world
Training purpose. Will not contains any valuable things. ^^


Hi,

This is my first Git repository. It will be used for testing only.
If you are looking for test repository too? then let me know, or request to pull.

학습 목적으로 만든 repository랍니다.  GIT 사용법을 익히고자 생성했는데, 혹시 테스트를 원하는 분은 테스트 목적으로 사용하셔도...


약간의 코드도 추가할 생각인데, Golang 을 사용할 생각입니다.

내용 추가 : 2016.08.16

* http proxy 추가.

firewall 내에서 http/https 만을 허용하는 환경은 git 를 사용할 수 없다. 이 경우 아래 proxy를 추가하고 git를 대신해 https protocol을 사용해야한다.

git config --global http.proxy http://proxyname:port
git config --global https.proxy http://proxyname:port

이 후 clone한다면

git clone https://gitbug.com/jnshin/hello-world.git
