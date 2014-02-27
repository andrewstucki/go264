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
	fileHash := "5827d2338b79ff82cd091c707823e466197281d3"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/test_vd_rc.264")
	fileHash = "eea02e97bfec89d0418593a8abaaf55d02eaa1ca"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/Static.264")
	fileHash = "91dd4a7a796805b2cd015cae8fd630d96c663f42"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/BA1_FT_C.264")
	fileHash = "418d152fb85709b6f172799dcb239038df437cfa"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/BA1_Sony_D.jsv")
	fileHash = "d94b5ceed5686a03ea682b53d415dee999d27eb6"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/BAMQ1_JVC_C.264")
	fileHash = "613cf662c23e5d9e1d7da7fe880a3c427411d171"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/BAMQ2_JVC_C.264")
	fileHash = "11bcf3713f520e606a8326d37e00e5fd6c9fd4a0"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/BA_MW_D.264")
	fileHash = "afd7a9765961ca241bb4bdf344b31397bec7465a"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/BANM_MW_D.264")
	fileHash = "92d924a857a1a7d7d9b224eaa3887830f15dee7f"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/BASQP1_Sony_C.jsv")
	fileHash = "3986c8c9d2876d2f0748b925101b152c6ec8b811"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/CI1_FT_B.264")
	fileHash = "cbfec15e17a504678b19a1191992131c92a1ac26"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/CI_MW_D.264")
	fileHash = "289f29a103c8d95adf2909c646466904be8b06d7"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/CVFC1_Sony_C.jsv")
	fileHash = "4641abd7419a5580b97f16e83fd1d566339229d0"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/CVPCMNL1_SVA_C.264")
	fileHash = "c2b0d964de727c64b9fccb58f63b567c82bda95a"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/LS_SVA_D.264")
	fileHash = "72118f4d1674cf14e58bed7e67cb3aeed3df62b9"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/MIDR_MW_D.264")
	fileHash = "9467030f4786f75644bf06a7fc809c36d1959827"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/MPS_MW_A.264")
	fileHash = "67f1cfbef0e8025ed60dedccf8d9558d0636be5f"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/MR1_BT_A.h264")
	fileHash = "6e585f8359667a16b03e5f49a06f5ceae8d991e0"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/MR1_MW_A.264")
	fileHash = "d9e2bf34e9314dcc171ddaea2c5015d0421479f2"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/MR2_MW_A.264")
	fileHash = "628b1d4eff04c2d277f7144e23484957dad63cbe"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/MR2_TANDBERG_E.264")
	fileHash = "74d618bc7d9d41998edf4c85d51aa06111db6609"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/NL1_Sony_D.jsv")
	fileHash = "e401e30669938443c2f02522fd4d5aa1382931a0"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/NLMQ1_JVC_C.264")
	fileHash = "f3265c6ddf8db1b2bf604d8a2954f75532e28cda"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/NLMQ2_JVC_C.264")
	fileHash = "350ae86ef9ba09390d63a09b7f9ff54184109ca8"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/NRF_MW_E.264")
	fileHash = "20732198c04cd2591350a361e4510892f6eed3f0"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/SVA_BA1_B.264")
	fileHash = "c4543b24823b16c424c673616c36c7f537089b2d"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/SVA_BA2_D.264")
	fileHash = "98ff2d67860462d8d8bcc9352097c06cc401d97e"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/SVA_Base_B.264")
	fileHash = "91f514d81cd33de9f6fbf5dbefdb189cc2e7ecf4"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/SVA_CL1_E.264")
	fileHash = "4fe09ab6cdc965ea10a20f1d6dd38aca954412bb"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/SVA_FM1_E.264")
	fileHash = "fad08c4ff7cf2307b6579853d0f4652fc26645d3"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/SVA_NL1_B.264")
	fileHash = "6d63f72a0c0d833b1db0ba438afff3b4180fb3e6"
	if streamDecode(file) != fileHash {
		t.FailNow()
	}

	file, _ = os.Open("res/SVA_NL2_E.264")
	fileHash = "70453ef8097c94dd190d6d2d1d5cb83c67e66238"
	if streamDecode(file) != fileHash {
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
