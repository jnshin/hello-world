# hello-world
Training purpose. Will not contains any valuable things. ^^

Hi,

This is my first Git repository. It will be used for testing only.
If you are looking for test repository too? then let me know, or request to pull.

학습 목적으로 만든 repository 입니다. GIT 사용법을 익히고자 생성했는데, 혹시 테스트를 원하는 분은 테스트 목적으로 사용하셔도...

golang 학습하면서 필요한 func을 만들면 이 곳에 추가할 생각입니다. 기능이 다듬어지면 언제 다른 repository로 옮겨갈지 모르니... 뭐 대단한 코드가 올라오지는 않겠지만요.




git 사용 관련 Tip

* http proxy 추가.

firewall 내에서 http/https 만을 허용하는 환경은 git 를 사용할 수 없다. 이 경우 아래 proxy를 추가하고 git를 대신해 https protocol을 사용해야한다.

git config --global http.proxy http://proxyname:port
git config --global https.proxy http://proxyname:port

이 후 clone한다면

git clone https://gitbug.com/jnshin/hello-world.git

* push 과정에서 에러가 있다면...

force option을 사용하지 말고, remote repository에 올려진 변경 사항을 먼저 가져와 merge해 볼 것.
즉, get -> merge -> push 순서로 재작업 해 보세요.
