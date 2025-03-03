# Protocol
### - authentication
since we are sending packets to the receiver via connector, 
we need to authenticate the connector from both side first\
just incase. as we don't want to send our packets to a random 
person

```
sender <----(authenticate)------ connector ------(authenticate)----> receiver
sender ---(authenticate-ack)---> connector <---(authenticate-ack)--- receiver
```

### - initiate
after both sender and receiver authenticate the connector, 
the connector would then send init to both receiver and sender
```
sender <----(init)----- connector -----(init)----> receiver
```

in which only the receiver would respond with a "helloclient", 
then the connector would passed on this message to the sender\
in which it would reply back to the receiver with "helloserver"

```
sender <---(helloclient)---- connector <---(helloclient)---- receiver
sender ----(helloserver)---> connector ----(helloserver)---> receiver
```


### - encryption key exchange:
then after the receiver get the (helloserver), the receiver will 
send us a public key.
```
sender <----(publickey)----- connector <----(publickey)----- receiver
  after receiving the public key, we then verify it, if it's
  valid we then send our aes key to the receiver, this is to
  ensure that the connector doesn't send us fake public key

sender ------(aeskey)------> connector ------(aeskey)------> receiver
  now that we have verify the public key we can encrypt the 
  aes key and send it to the receiver, and since the publickey 
  was verified  we can still send our packet trough the connector, 
  without the fear of the connector reading our packet
```
### - communication:
now that we have exchanged the aes key, we can now send our 
packet to the receiver, the receiver will then decrypt\
the packet and forward it to the destination (i.e. internet)
and vice versa
```
sender -------(data)------> receiver -------(data)------> internet
sender <------(data)------- receiver <------(data)------- internet
```