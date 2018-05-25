package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodrigodiez/zorro/pkg/protobuf"

	serviceMocks "github.com/rodrigodiez/zorro/lib/mocks/service"
)

func TestMask(t *testing.T) {
	t.Parallel()

	var (
		resp *protobuf.MaskResponse
		err  error
	)

	zorro := serviceMocks.Zorro{}
	server := &server{
		zorro: &zorro,
	}

	zorro.On("Mask", "foo").Return("bar").Once()

	resp, err = server.Mask(context.TODO(), &protobuf.MaskRequest{Key: "foo"})

	assert.Equal(t, "bar", resp.GetValue())
	assert.Nil(t, err)
	zorro.AssertExpectations(t)
}

func TestUnMaskKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	var (
		resp *protobuf.UnmaskResponse
		err  error
	)

	zorro := serviceMocks.Zorro{}
	server := &server{
		zorro: &zorro,
	}

	zorro.On("Unmask", "bar").Return("", false).Once()

	resp, err = server.Unmask(context.TODO(), &protobuf.UnmaskRequest{Value: "bar"})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	zorro.AssertExpectations(t)
}
func TestUnMaskKeyExists(t *testing.T) {
	t.Parallel()

	var (
		resp *protobuf.UnmaskResponse
		err  error
	)

	zorro := serviceMocks.Zorro{}
	server := &server{
		zorro: &zorro,
	}

	zorro.On("Unmask", "bar").Return("foo", true).Once()

	resp, err = server.Unmask(context.TODO(), &protobuf.UnmaskRequest{Value: "bar"})

	assert.Nil(t, err)
	assert.Equal(t, "foo", resp.GetKey())
	zorro.AssertExpectations(t)
}
