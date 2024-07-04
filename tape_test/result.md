```shell
./network.sh down
./network.sh up createChannel 
./network.sh deployCC -ccn DTModeling -ccp DTModeling -ccl go
```

```shell
rm -rf organizations
cp -R ../fabric-samples/test-network/organizations .
```

```shell
./tape -c config_modeling.yaml -n 200
./tape -c config_save_ipfs.yaml -n 200
./tape -c config_save_ct.yaml -n 200
```




| orderer num <br/> tps  | 1       | 2       | 3       | 4 | 5      | 6 | 7       | 8 | 9       |
|------------------------|---------|---------|---------|---|--------|---|---------|---|---------|
| Modeling               | 401.67  | 395.41  | 383.7   |   | 339.45 |   | 311.21  |   | 285.8   |
| SaveIPFSAddr           | 404.86  | 385.5   | 371.65  |   | 355.03 |   | 310.63  |   | 303.72  |
| SaveCT                 | 308.71  | 281.82  | 241.62  |   | 233.07 |   | 223.99  |   | 202.73  |

| 3order<br/> block size | 16KB   | 32KB     | 48KB   | 64KB    | 128KB  | 256KB  | 512KB   |
|------------------------|--------|----------|--------|---------|--------|--------|---------|
| Modeling               | 166.71 | 231.97   | 281.15 | 336.32  | 486.50 | 479.57 | 484.91  |
| SaveIPFSAddr           | 170.51 | 245.439  | 298.92 | 355.21  | 539.77 | 561.22 | 567.00  |
| SaveCT                 | 102.02 | 160.174  | 222.06 | 260.23  | 386.28 | 417.15 | 447.68  |

