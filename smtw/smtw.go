package smtw

import (
	"crypto/sha256"

	"github.com/celestiaorg/smt"

	"smttestgen/framework"
)

func digest(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	sum := h.Sum(nil)
	return sum
}

type SparseMerkleTreeWrapper struct {
	tree smt.SparseMerkleTree
	test framework.Test
}

func NewSparseMerkleTreeWrapper(testName string) SparseMerkleTreeWrapper {
	nodes := smt.NewSimpleMap()
	values := smt.NewSimpleMap()
	hasher := sha256.New()
	tree := *smt.NewSparseMerkleTree(nodes, values, hasher)

	root := framework.EncodedValue{Value: "", Encoding: ""}
	test := framework.Test{
		Name:  testName,
		Root:  root,
		Steps: make([]framework.Step, 0),
	}

	wrapper := SparseMerkleTreeWrapper{
		tree: tree,
		test: test,
	}

	return wrapper
}

func (smtw *SparseMerkleTreeWrapper) Root() []byte {
	return smtw.tree.Root()
}

func (smtw *SparseMerkleTreeWrapper) Update(key []byte, value []byte) ([]byte, error) {
	stepKey := framework.HexValue(digest(key))
	stepData := framework.Utf8Value(value)
	step := framework.Step{
		Action: "update",
		Key:    &stepKey,
		Data:   &stepData,
	}
	smtw.test.Steps = append(smtw.test.Steps, step)
	return smtw.tree.Update(key, value)
}

func (smtw *SparseMerkleTreeWrapper) Delete(key []byte) ([]byte, error) {
	stepKey := framework.HexValue(digest(key))
	step := framework.Step{
		Action: "delete",
		Key:    &stepKey,
	}
	smtw.test.Steps = append(smtw.test.Steps, step)
	return smtw.tree.Delete(key)
}

func (smtw *SparseMerkleTreeWrapper) GetTest() (test framework.Test) {
	rootBytes := smtw.Root()
	smtw.test.Root = framework.HexValue(rootBytes)
	return smtw.test
}
