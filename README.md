# build and run
docker-compose up -d


# jwt key generate
ssh-keygen -t rsa -f revel
ssh-keygen -f revel.pub -e -m pkcs8 > revel.pub
https://qiita.com/AkiTakeU/items/e2133eeb94f57629b5e7