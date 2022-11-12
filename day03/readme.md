## 2.实现POW共识算法
1.在区块链结构中完善了Nonce属性，方便其进行计算
2.产生了ProofOfWork结构，并完善了其方法：
    NewProofOfWork：根据目标区块，生成POW结构
    run():执行POW算法，比较哈希
    prepareData():生成准备数据，对ProofOfWork数据拼接形成哈希值并返回