package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/thrift-iterator/go/test"
)

func Test_skip_map_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMapBegin(thrift.I64, thrift.LIST, 1)
		proto.WriteI64(1)
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(1)
		proto.WriteListEnd()
		proto.WriteMapEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipMap(nil))
	}
}

func Test_decode_map_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMapBegin(thrift.I64, thrift.LIST, 1)
		proto.WriteI64(1)
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(1)
		proto.WriteListEnd()
		proto.WriteMapEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal([]interface{}{
			int64(1),
		}, iter.ReadMap()[int64(1)])
	}
}

func Test_unmarshal_map_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMapBegin(thrift.I64, thrift.LIST, 1)
		proto.WriteI64(1)
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(1)
		proto.WriteListEnd()
		proto.WriteMapEnd()
		var val map[int64][]int64
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(map[int64][]int64{
			1: {1},
		}, val)
	}
}

func Test_encode_map_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteMap(map[interface{}]interface{}{
			int64(1): []interface{}{int64(1)},
		})
		iter := c.CreateIterator(stream.Buffer())
		should.Equal([]interface{}{
			int64(1),
		}, iter.ReadMap()[int64(1)])
	}
}

func Test_marshal_map_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(map[int64][]int64{
			1: {1},
		})
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal([]interface{}{
			int64(1),
		}, iter.ReadMap()[int64(1)])
	}
}
