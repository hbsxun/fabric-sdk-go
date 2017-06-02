## Distribute Test
### org1-ca0-orderer.yaml
1. ca0
2. orderer0
3. peer0 and peer1

### org2-ca1.yaml
Resides on another physical machine.  
1. ca1  
2. peer2 and peer3  
_Note:_  Two items should be set.  
For each peer in the docker-compose file, add  
```
extra_hosts:
  - "orderer0:192.168.50.147"
```
```
environment:
  - CORE_PEER_COMMITTER_LEDGER_ORDERER=192.168.50.147:7050
```
