## stand-alone
### start daemon
```bash
docker run -d ttbb/kafka:mate
```
### start it port expose
```bash
docker run -p 9092:9092 -it ttbb/kafka:mate
```
### start daemon
```bash
docker run -p 9092:9092 -d ttbb/kafka:mate
```
### start daemon mac port expose
```bash
docker run  -p 9092:9092 -e KAFKA_ADVERTISE_ADDRESS=localhost -d ttbb/kafka:mate
```
## raft
### start it port expose
```bash
docker run -p 9092:9092 -e REMOTE_MODE=false -it ttbb/kafka:mate bash
```
### start daemon port expose
```bash
docker run -p 9092:9092 -e REMOTE_MODE=false -d ttbb/kafka:mate
```
### start daemon mac port expose
```bash
docker run -p 9092:9092 -e KAFKA_ADVERTISE_ADDRESS=localhost -e REMOTE_MODE=false -d ttbb/kafka:mate
```