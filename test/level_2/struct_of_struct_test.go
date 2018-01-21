package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/thrift-iterator/go/protocol"
	"github.com/thrift-iterator/go/test"
	"github.com/thrift-iterator/go/test/level_2/struct_of_struct_test"
)

func Test_skip_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRUCT, 1)

		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRING, 1)
		proto.WriteString("abc")
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()

		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipStruct(nil))
	}
}

func Test_decode_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRUCT, 1)

		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRING, 1)
		proto.WriteString("abc")
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()

		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(map[protocol.FieldId]interface{}{
			protocol.FieldId(1): "abc",
		}, iter.ReadStruct()[protocol.FieldId(1)])
	}
}

func Test_unmarshal_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRUCT, 1)

		proto.WriteStructBegin("hello")
		proto.WriteFieldBegin("field1", thrift.STRING, 1)
		proto.WriteString("abc")
		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()

		proto.WriteFieldEnd()
		proto.WriteFieldStop()
		proto.WriteStructEnd()
		var val struct_of_struct_test.TestObject
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(struct_of_struct_test.TestObject{
			struct_of_struct_test.EmbeddedObject{"abc"},
		}, val)
	}
}

func Test_encode_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteStruct(map[protocol.FieldId]interface{}{
			protocol.FieldId(1): map[protocol.FieldId]interface{}{
				protocol.FieldId(1): "abc",
			},
		})
		iter := c.CreateIterator(stream.Buffer())
		should.Equal(map[protocol.FieldId]interface{}{
			protocol.FieldId(1): "abc",
		}, iter.ReadStruct()[protocol.FieldId(1)])
	}
}

func Test_marshal_struct_of_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(struct_of_struct_test.TestObject{
			struct_of_struct_test.EmbeddedObject{"abc"},
		})
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(map[protocol.FieldId]interface{}{
			protocol.FieldId(1): "abc",
		}, iter.ReadStruct()[protocol.FieldId(1)])
	}
}