package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {

	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("Newからの戻り値がnilです")
	} else {
		tracer.Trace("こんにちは、traceパッケージ")
		if buf.String() != "こんにちは、traceパッケージ\n" {
			t.Errorf("'%s'という謝った文字列が出力されました", buf.String())
		}
	}
}

func Testoff(t *testing.T) {
	var silentTracer Tracer = off()
	silentTracer.Trace("データ")
}
