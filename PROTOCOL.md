# Protocol
### - authentication
since we are sending packets to the receiver via connector, we need to authenticate the connector from both side first\
just incase. as we don't want to send our packets to a random person

```
sender <----(authenticate)------ connector ------(authenticate)----> receiver
sender ---(authenticate-ack)---> connector <---(authenticate-ack)--- receiver
```

### - encryption key exchange:
after authenticating with the connector, the receiver will send us a public key, to ensure that the connector doesn't\
go rogue we need to check if the public key is a valid one
```
sender <----(publickey)----- connector <----(publickey)----- receiver
  after receiving the public key, we then verify it, if it's
  valid we then send our aes key to the receiver, this is to
  ensure that the connector doesn't send us fake public key

sender ------(aeskey)------> connector ------(aeskey)------> receiver
  now that we have exchanged the aes key, we can still send
  our packet trough the connector, without the fear of the
  connector reading our packet as it's encrypted end to end
```
### - communication:
now that we have exchanged the aes key, we can now send our packet to the receiver, the receiver will then decrypt\
the packet and forward it to the destination (i.e. internet)
```
sender -------(data)------> receiver -------(data)------> internet
sender <------(data)------- receiver <------(data)------- internet
```