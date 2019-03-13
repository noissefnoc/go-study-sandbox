package mock

import (
	"github.com/golang/mock/gomock"
	"github.com/noissefnoc/go-study-sandbox/test/mock/mock_base"
	"testing"
)

func TestSample1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSample := mock_mock.NewMockSample(ctrl)
	mockSample.EXPECT().Method("hoge").Return(1)

	t.Log("result:", mockSample.Method("hoge"))
}
