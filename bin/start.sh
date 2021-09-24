npm run build
cd goserver/
go build -o revstream-server
cd ..

PROCESSES=`ps aux -P | grep nuxt | sed -e 's/ubuntu *//' | cut -d" " -f 1`
for i in $PROCESSES;do
    kill $i
done
npm run start &

PROCESSES=`ps aux -P | grep "revstream-server" | sed -e 's/ubuntu *//' | cut -d" " -f 1`
for i in $PROCESSES;do
    kill $i
done
cd ./goserver
./revstream-server &