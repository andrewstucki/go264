package h264

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"unsafe"
)

func TestSmoke(*testing.T) {
	a := NewLTRRecoverRequest()
	b := NewLTRMarkingFeedback()
	c := NewSliceArgument()
	d := NewSliceConfig()
	e := NewSpatialLayerConfig()
	f := NewEncParamBase()
	g := NewEncParamExt()
	h := NewVideoProperty()
	i := NewDecodingParam()
	j := NewLayerBSInfo()
	k := NewFrameBSInfo()
	l := NewSourcePicture()
	m := NewSliceInfo()
	n := NewRateThresholds()
	o := NewSysMemBuffer()
	q := NewBufferInfo()

	a.updateBase()
	b.updateBase()
	c.updateBase()
	d.updateBase()
	e.updateBase()
	f.updateBase()
	g.updateBase()
	h.updateBase()
	i.updateBase()
	j.updateBase()
	k.updateBase()
	l.updateBase()
	m.updateBase()
	n.updateBase()
	o.updateBase()
	q.updateBase()

	a.updateStruct()
	b.updateStruct()
	c.updateStruct()
	d.updateStruct()
	e.updateStruct()
	f.updateStruct()
	g.updateStruct()
	h.updateStruct()
	i.updateStruct()
	j.updateStruct()
	k.updateStruct()
	l.updateStruct()
	m.updateStruct()
	n.updateStruct()
	o.updateStruct()
	q.updateStruct()
}

func TestStreamEncode(t *testing.T) {
	file, _ := os.Open("res/CiscoVT2people_320x192_12fps.yuv")
	fileHash := "06441376891cbc237a36e59b62131cd94ff9cb19"
	hash := streamEncode(file, 320, 192, 12.0)
	if hash != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/CiscoVT2people_160x96_6fps.yuv")
	fileHash = "4f3759fc44125b27a179ebff158dbba9e431bd0b"
	hash = streamEncode(file, 160, 96, 6.0)
	if fileHash != hash {
		t.FailNow()
	}

	file, _ = os.Open("res/Static_152_100.yuv")
	fileHash = "a004c7410a78bfe00af65ae7071ce1c485cf036e"
	hash = streamEncode(file, 152, 100, 6.0)
	if hash != fileHash {
		t.FailNow()
	}
}
func TestStreamDecode(t *testing.T) {
	file, _ := os.Open("res/test_vd_1d.264")
	if streamDecode(file) != "5827d2338b79ff82cd091c707823e466197281d3" {
		t.FailNow()
	}

	file, _ = os.Open("res/test_vd_rc.264")
	if streamDecode(file) != "eea02e97bfec89d0418593a8abaaf55d02eaa1ca" {
		t.FailNow()
	}

	file, _ = os.Open("res/Static.264")
	if streamDecode(file) != "91dd4a7a796805b2cd015cae8fd630d96c663f42" {
		t.FailNow()
	}
}

func streamDecode(stream io.Reader) string {
	decParam := NewDecodingParam()
	decParam.OutputColorFormat = VideoFormatI420
	decParam.TargetDQLayer = 0xff
	decParam.ECActiveFlag = 1
	decParam.VideoProperty.VideoBsType = VideoBitstreamDefault

	reader := bufio.NewReader(stream)

	startFrame := []byte{0x0, 0x0, 0x0, 0x1}
	file, _ := ioutil.ReadAll(reader)
	frames := bytes.Split(file, startFrame)

	decoder, _ := CreateDecoder()
	decoder.Initialize(decParam)
	defer DestroyDecoder(decoder)
	hash := sha1.New()
	frameCount := 0
	var bufInfo *BufferInfo
	for _, frame := range frames[1:] {
		bufInfo = NewBufferInfo()
		f := append(startFrame, frame...)
		data, _ := decoder.DecodeFrame2(f, bufInfo)
		if bufInfo.BufferStatus == 1 {
			frameCount++
			j := 0
			//y
			for i := 0; i < bufInfo.SystemBuffer.Height; i++ {
				hash.Write(data[0][j : j+bufInfo.SystemBuffer.Width])
				j += bufInfo.SystemBuffer.Stride[0]
			}

			j = 0
			//u
			for i := 0; i < bufInfo.SystemBuffer.Height/2; i++ {
				hash.Write(data[1][j : j+bufInfo.SystemBuffer.Width/2])
				j += bufInfo.SystemBuffer.Stride[1]
			}

			j = 0
			//v
			for i := 0; i < bufInfo.SystemBuffer.Height/2; i++ {
				hash.Write(data[2][j : j+bufInfo.SystemBuffer.Width/2])
				j += bufInfo.SystemBuffer.Stride[1]
			}
		}
	}
	var a []byte
	bufInfo = NewBufferInfo()
	data, _ := decoder.DecodeFrame2(a, bufInfo)
	if bufInfo.BufferStatus == 1 {
		frameCount++
		j := 0
		//y
		for i := 0; i < bufInfo.SystemBuffer.Height; i++ {
			hash.Write(data[0][j : j+bufInfo.SystemBuffer.Width])
			j += bufInfo.SystemBuffer.Stride[0]
		}

		j = 0
		//u
		for i := 0; i < bufInfo.SystemBuffer.Height/2; i++ {
			hash.Write(data[1][j : j+bufInfo.SystemBuffer.Width/2])
			j += bufInfo.SystemBuffer.Stride[1]
		}

		j = 0
		//v
		for i := 0; i < bufInfo.SystemBuffer.Height/2; i++ {
			hash.Write(data[2][j : j+bufInfo.SystemBuffer.Width/2])
			j += bufInfo.SystemBuffer.Stride[1]
		}
	}

	shaVal := hash.Sum(nil)
	return hex.EncodeToString(shaVal)
}

func streamEncode(stream io.Reader, width, height int, frameRate float64) string {
	reader := bufio.NewReader(stream)
	encoder, _ := CreateSVCEncoder()
	defer DestroySVCEncoder(encoder)

	encParam := NewEncParamBase()
	encParam.MaxFrameRate = frameRate
	encParam.SourceWidth = width
	encParam.SourceHeight = height
	encParam.TargetBitrate = 5000000
	encParam.InputCsp = VideoFormatI420

	encoder.Initialize(encParam)

	frameSize := width * height * 3 / 2

	frameInfo := NewFrameBSInfo()
	sourcePicture := NewSourcePicture()
	sourcePicture.SourceWidth = width
	sourcePicture.SourceHeight = height
	sourcePicture.ColorFormat = VideoFormatI420
	sourcePicture.Stride[0] = sourcePicture.SourceWidth
	sourcePicture.Stride[1] = sourcePicture.SourceWidth >> 1
	sourcePicture.Stride[2] = sourcePicture.SourceWidth >> 1

	buf := make([]byte, frameSize)
	hash := sha1.New()
	for numBytes, _ := reader.Read(buf); numBytes == frameSize; numBytes, _ = reader.Read(buf) {
		sourcePicture.Data[0] = (*uint8)(&buf[0])
		sourcePicture.Data[1] = (*uint8)(&buf[width*height])
		sourcePicture.Data[2] = (*uint8)(&buf[width*height+width*height>>2])
		ret := encoder.EncodeFrame(sourcePicture, frameInfo)
		if ret == VideoFrameTypeInvalid {
			log.Printf("Invalid Frame!")
		} else if ret != VideoFrameTypeSkip {
			for i := 0; i < frameInfo.LayerNum; i++ {
				layerInfo := frameInfo.LayerInfo[i]
				layerSize := 0
				for j := 0; j < layerInfo.NALCount; j++ {
					layerSize += layerInfo.NALLengthInByte[j]
				}
				bSBuf := (*[1 << 30]byte)(unsafe.Pointer(layerInfo.BitstreamBuffer))[:layerSize]
				hash.Write(bSBuf)
			}
		}
	}
	shaVal := hash.Sum(nil)
	return hex.EncodeToString(shaVal)
}
