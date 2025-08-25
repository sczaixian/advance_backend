package test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	URL    = "https://eth-sepolia.g.alchemy.com/v2/RmsPYhly5O6-XH8UdmqCQ"
	WS_URL = "wss://eth-mainnet.g.alchemy.com/v2/RmsPYhly5O6-XH8UdmqCQ"

	PRIVATE_KEY_HOM1 = "e039af6407e7622a8354bd45ea44de86ca663c81d6176ab698fe788e603b2682"
	PRIVATE_KEY_HOM2 = "d36fe4707c15f011256ae8813d298b05511a889a09c30e3813adc100117af9be"
	ADDRESS_HOM_1    = "0xd96AF416b2060500f828A3f31e617F15CBEA1e4b"
	ADDRESS_HOM_2    = "0xD930e6b5C1C1112d4c0Db00c4888557be1b58d0D"

	ADDRESS_CMP_1     = "0x3E0bDb54f94D735dDCf8D2074c852a8C22914aA7"
	PRIVATE_KEY_CMP   = "decf550fcc469204df8d024977ad887c888a3164b0977f588ae645d3786b4511"
	ADDRESS_CMP_2     = "0x1984EE3bCE0c863d24905dF9D54322856e54CCea"
	PRIVATE_KEY_CMP_2 = "6886c0e89cf2eca81e01cf1cb49462b6c18ac9d74992f77ce6192a9f87578835"

	BLOCK_HEIGHT int64 = 9058703
	BLOCK_HASH         = "0x632ff811cfa17f99cba09a5cc81a42ebc0f7557c532c2cb1c383901dda957619"
	TX_HASH            = "0x5caebec23f825ec46425db409573fffcc0909a105316f28735908962ee972ba3"

	CONTRACT_ADDRESS = "0xbaaa34743675c22be153484e0b452ff08472c345"
)

type Bloom [256]byte
type BlockNonce [8]byte

// Header 代表以太坊区块链中的一个区块头，包含了区块的元数据和状态信息
type HeaderLocal struct {
	// ParentHash 父区块的哈希值，用于构建区块链的链式结构
	// 每个区块通过指向其父区块来形成链，这是区块链不可篡改特性的基础
	ParentHash common.Hash `json:"parentHash"       gencodec:"required"`

	// UncleHash 叔区块（ommer）的默克尔树根哈希
	// 叔区块是指与当前区块同一层级但未被主链采纳的区块，用于以太坊的GHOST协议
	UncleHash common.Hash `json:"sha3Uncles"       gencodec:"required"`

	// Coinbase 挖出此区块的矿工地址，将获得此区块的挖矿奖励和交易费用
	// 在以太坊转向权益证明(PoS)后，此字段可能代表验证者地址
	Coinbase common.Address `json:"miner"`

	// Root 世界状态的默克尔树(Merkle Trie)根哈希
	// 代表了执行完此区块中所有交易后的全局状态
	Root common.Hash `json:"stateRoot"        gencodec:"required"`

	// TxHash 此区块中所有交易组成的默克尔树根哈希
	// 用于快速验证某笔交易是否包含在此区块中
	TxHash common.Hash `json:"transactionsRoot" gencodec:"required"`

	// ReceiptHash 此区块中所有交易收据组成的默克尔树根哈希
	// 交易收据包含了交易执行的结果信息，如日志和消耗的Gas
	ReceiptHash common.Hash `json:"receiptsRoot"     gencodec:"required"`

	// Bloom 布隆过滤器，用于快速查询此区块中的日志事件
	// 是一种空间效率高的概率数据结构，用于测试元素是否在集合中
	Bloom Bloom `json:"logsBloom"        gencodec:"required"`

	// Difficulty 此区块的挖矿难度，用于工作量证明(PoW)共识机制
	// 难度值越高，找到有效哈希所需的计算量就越大
	Difficulty *big.Int `json:"difficulty"       gencodec:"required"`

	// Number 区块高度，从创世区块(0)开始递增的区块序号
	// 代表了此区块在区块链中的位置
	Number *big.Int `json:"number"           gencodec:"required"`

	// GasLimit 此区块允许的最大Gas总量，限制了区块中可以包含的交易数量
	// 矿工可以投票调整后续区块的Gas限制，以适应网络需求
	GasLimit uint64 `json:"gasLimit"         gencodec:"required"`

	// GasUsed 此区块中所有交易实际消耗的Gas总量
	// 必须小于或等于GasLimit
	GasUsed uint64 `json:"gasUsed"          gencodec:"required"`

	// Time 区块时间戳，Unix时间格式（自1970年1月1日以来的秒数）
	// 必须大于其父区块的时间戳
	Time uint64 `json:"timestamp"        gencodec:"required"`

	// Extra 额外数据字段，可由矿工自由添加任意数据
	// 通常用于存储签名或其他元数据，最大长度为32字节
	Extra []byte `json:"extraData"        gencodec:"required"`

	// MixDigest 混合哈希，用于工作量证明算法的随机数
	// 与Nonce一起用于证明矿工完成了足够的计算工作
	MixDigest common.Hash `json:"mixHash"`

	// Nonce 随机数，用于工作量证明算法
	// 矿工通过不断更改此值来寻找满足难度目标的区块哈希
	Nonce BlockNonce `json:"nonce"`

	// BaseFee 基础费用，由EIP-1559引入，用于交易费用市场改革
	// 表示每单位Gas的基础费用，会被燃烧而不是支付给矿工
	// 在传统的工作量证明区块中此字段被忽略
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash 提款默克尔树根哈希，由EIP-4895引入
	// 代表了信标链提款到执行层的默克尔根，仅在合并后的区块中存在
	// 在传统的工作量证明区块中此字段被忽略
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed 使用的Blob Gas数量，由EIP-4844引入
	// 用于Proto-Danksharding，表示此区块中blob交易使用的Gas总量
	// 在传统的工作量证明区块中此字段被忽略
	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas 超额Blob Gas，由EIP-4844引入
	// 用于调整blob基础费用，类似于EIP-1559的基础费用机制
	// 在传统的工作量证明区块中此字段被忽略
	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot 父信标链根，由EIP-4788引入
	// 将信标链状态根包含在执行层区块头中，使智能合约能够访问共识层状态
	// 在传统的工作量证明区块中此字段被忽略
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`

	// RequestsHash 请求默克尔树根哈希，由EIP-7685引入
	// 用于表示跨域请求的默克尔根，支持更复杂的跨链交互
	// 在传统的工作量证明区块中此字段被忽略
	RequestsHash *common.Hash `json:"requestsHash" rlp:"optional"`
}

// Log 结构体表示以太坊智能合约事件日志
// 事件日志是智能合约通过LOG操作码发出的结构化数据，用于记录合约执行中的重要信息
type Log struct {
	// 共识字段：这些字段是区块链共识的一部分，由智能合约生成并在全网节点间保持一致

	// Address 生成事件的合约地址
	// 表示哪个合约发出了这个事件日志
	Address common.Address `json:"address" gencodec:"required"`

	// Topics 事件主题列表
	// 第一个主题通常是事件签名的Keccak-256哈希（如Transfer(address,address,uint256)）
	// 后续主题通常包含索引参数（indexed parameters），用于高效过滤和检索
	Topics []common.Hash `json:"topics" gencodec:"required"`

	// Data 事件数据
	// 包含非索引参数的ABI编码数据，通常需要解码才能读取具体内容
	// 与Topics不同，Data内容不参与日志过滤，但包含更详细的事件信息
	Data []byte `json:"data" gencodec:"required"`

	// 派生字段：这些字段由节点填充，但不属于共识部分
	// 它们提供关于日志来源的上下文信息，但不影响日志内容的有效性

	// BlockNumber 包含此日志的交易的区块号
	BlockNumber uint64 `json:"blockNumber" rlp:"-"` // rlp:"-" 表示此字段不参与RLP编码

	// TxHash 包含此日志的交易哈希
	TxHash common.Hash `json:"transactionHash" gencodec:"required" rlp:"-"`

	// TxIndex 交易在区块中的索引位置
	TxIndex uint `json:"transactionIndex" rlp:"-"`

	// BlockHash 包含此日志的区块哈希
	BlockHash common.Hash `json:"blockHash" rlp:"-"`

	// BlockTimestamp 区块时间戳
	// 区块被挖出时的Unix时间戳
	BlockTimestamp uint64 `json:"blockTimestamp" rlp:"-"`

	// Index 日志在区块中的索引位置（日志索引）
	// 同一交易中可能包含多个日志，每个日志有唯一的索引
	Index uint `json:"logIndex" rlp:"-"`

	// Removed 标识此日志是否因链重组而被移除
	// 如果因为链重组（reorg）导致包含此日志的交易被回滚，此字段为true
	// 在使用过滤器查询日志时，必须注意此字段的值
	Removed bool `json:"removed" rlp:"-"`
}

// Receipt 结构体表示以太坊交易的执行回执
// 交易回执包含了交易执行后的状态信息，包括Gas消耗、日志、状态等
type ReceiptLocal struct {
	// 共识字段：这些字段在以太坊黄皮书中定义，是区块链共识的一部分

	// Type 表示交易类型（EIP-2718引入的交易类型）
	// 0: 传统交易，1: EIP-2930访问列表交易，2: EIP-1559动态费用交易
	Type uint8 `json:"type,omitempty"`

	// PostState 交易执行后的状态根（仅用于前拜占庭分叉的交易）
	// 在拜占庭分叉后，被Status字段取代
	PostState []byte `json:"root"`

	// Status 交易执行状态（拜占庭分叉后引入）
	// 0: 交易执行失败，1: 交易执行成功
	Status uint64 `json:"status"`

	// CumulativeGasUsed 累计Gas使用量
	// 从区块开始到当前交易执行完毕累计使用的Gas总量
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed" gencodec:"required"`

	// Bloom 日志布隆过滤器
	// 一个高效的数据结构，用于快速检查日志中是否可能包含某些主题
	Bloom Bloom `json:"logsBloom"         gencodec:"required"`

	// Logs 交易执行过程中产生的日志数组
	// 包含智能合约执行过程中发出的事件日志
	Logs []*Log `json:"logs"              gencodec:"required"`

	// 实现字段：这些字段由Geth客户端在处理交易时添加

	// TxHash 交易哈希
	// 对应交易的Keccak-256哈希值，用于唯一标识交易
	TxHash common.Hash `json:"transactionHash" gencodec:"required"`

	// ContractAddress 合约地址（仅适用于合约创建交易）
	// 如果是合约创建交易，这里存储新创建的合约地址
	// 对于普通转账交易，此字段为零地址
	ContractAddress common.Address `json:"contractAddress"`

	// GasUsed 本交易实际使用的Gas数量
	GasUsed uint64 `json:"gasUsed" gencodec:"required"`

	// EffectiveGasPrice 实际生效的Gas价格
	// 对于EIP-1559交易，这是基础费用加上优先费用
	// 对于传统交易，这与交易的GasPrice相同
	EffectiveGasPrice *big.Int `json:"effectiveGasPrice"` // 必需字段，但为了向后兼容省略了标签

	// BlobGasUsed Blob Gas使用量（EIP-4844引入）
	// 用于存储Blob数据的Gas消耗量
	BlobGasUsed uint64 `json:"blobGasUsed,omitempty"`

	// BlobGasPrice Blob Gas价格（EIP-4844引入）
	// 用于Blob数据的Gas价格
	BlobGasPrice *big.Int `json:"blobGasPrice,omitempty"`

	// 包含信息：这些字段提供关于交易所在区块的信息

	// BlockHash 区块哈希
	// 包含此交易的区块的哈希值
	BlockHash common.Hash `json:"blockHash,omitempty"`

	// BlockNumber 区块号
	// 包含此交易的区块的编号（高度）
	BlockNumber *big.Int `json:"blockNumber,omitempty"`

	// TransactionIndex 交易在区块中的索引位置
	// 表示此交易在区块交易列表中的位置（从0开始）
	TransactionIndex uint `json:"transactionIndex"`
}
