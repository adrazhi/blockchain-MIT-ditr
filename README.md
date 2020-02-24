# Simple Blockchain
Example low-level blockchain implementation for MIT DITR 20 

## Objective
The purpose of this implementation is to provide a lower level explanation 
of how Blockchain works through the implementation of a simple cryptocurrency. 

## Instructions - How can I run this on my computer?

1. Make sure you have [Go](https://golang.org/doc/install) for your operating system installed 
2. Open `terminal` / `commandline`
3. Clone this repository: `git clone https://github.com/albtop/blockchain-MIT-ditr-20`
3. Change directory `cd` inside the local directory:  `cd blockchain-MIT-ditr-20`
2. Run `go build`
3. Run `./blockchain-MIT-ditr-20`

**Output should look something like the following**

```
Prev. hash: 
Information: Initial Block
RootHash: 836783123fcdfcd6a96082ef3bd1a42accc26c12de7288c47357ba62ba2b120f

Prev. hash: 836783123fcdfcd6a96082ef3bd1a42accc26c12de7288c47357ba62ba2b120f
Information: Send 1 BTC to your learning facilitator
RootHash: 65498fda56c820784362cbfbd90fafdd532f1f77bff31568d334bae3fae408b0

Prev. hash: 65498fda56c820784362cbfbd90fafdd532f1f77bff31568d334bae3fae408b0
Information: Send 1 BTC to Professor John
RootHash: cac7a9fc0781b542a72c11027750ad03afeb5c4e32b368862150083eee28f8cf

Prev. hash: cac7a9fc0781b542a72c11027750ad03afeb5c4e32b368862150083eee28f8cf
Information: Send 1 BTC to Professor Sanchez
RootHash: 123a64a90d5acf971b58a88283a75af82fcad187e0c0d28b04f56040bb06b7e0

Prev. hash: 123a64a90d5acf971b58a88283a75af82fcad187e0c0d28b04f56040bb06b7e0
Information: Send 1 BTC to Alice
RootHash: 6ef6d6d5c561d21cc71f3d62295b9fdf359a30a09c375fb1f1ccfcd4a20b7138

Prev. hash: 6ef6d6d5c561d21cc71f3d62295b9fdf359a30a09c375fb1f1ccfcd4a20b7138
Information: Send 1 BTC to Bob
RootHash: 4f5293c6993dd8707cd894b163704ffa7bacffb2322040d0a3bef30d4bf35133
```

## Questions
Reach out to `albano.drazhi@gmail.com`

## Refs
- Block Hashing Algorithm: https://en.bitcoin.it/wiki/Block_hashing_algorithm
- Bitcoin White Paper:  https://bitcoin.org/bitcoin.pdf
