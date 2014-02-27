package h264

/*
#include <wels/codec_api.h>
#include <stdlib.h>
#include <stdio.h>
#cgo LDFLAGS: -lwels

int encoder_initialize(ISVCEncoder *encoder, const SEncParamBase* pParam) {
  return ((*encoder)->Initialize)(encoder, pParam);
}

int encoder_initialize_ext(ISVCEncoder *encoder, const SEncParamExt* pParam) {
  return ((*encoder)->InitializeExt)(encoder, pParam);
}

int encoder_uninitialize(ISVCEncoder *encoder) {
  return ((*encoder)->Uninitialize)(encoder);
}

int encoder_encode_frame(ISVCEncoder *encoder, SSourcePicture *kpSrcPic, SFrameBSInfo *pBsInfo) {
  return ((*encoder)->EncodeFrame)(encoder, kpSrcPic, pBsInfo);
}

int encoder_encode_parameter_sets(ISVCEncoder *encoder, SFrameBSInfo *pBsInfo) {
  return ((*encoder)->EncodeParameterSets)(encoder, pBsInfo);
}

int encoder_pause_frame(ISVCEncoder *encoder, SSourcePicture *kpSrcPic, SFrameBSInfo *pBsInfo) {
  return ((*encoder)->PauseFrame)(encoder, kpSrcPic, pBsInfo);
}

int encoder_force_intra_frame(ISVCEncoder *encoder, bool bIDR) {
  return ((*encoder)->ForceIntraFrame)(encoder, bIDR);
}

int encoder_set_option(ISVCEncoder *encoder, ENCODER_OPTION eOptionId, void *pOption) {
  return ((*encoder)->SetOption)(encoder, eOptionId, pOption);
}

int encoder_get_option(ISVCEncoder *encoder, ENCODER_OPTION eOptionId, void *pOption) {
  return ((*encoder)->GetOption)(encoder, eOptionId, pOption);
}

long decoder_initialize(ISVCDecoder *decoder, const SDecodingParam *pParam) {
  return ((*decoder)->Initialize)(decoder, pParam);
}

long decoder_uninitialize(ISVCDecoder *decoder) {
  return ((*decoder)->Uninitialize)(decoder);
}

DECODING_STATE decoder_decode_frame(ISVCDecoder *decoder, const unsigned char *pSrc, const int iSrcLen, unsigned char **ppDst, int *pStride, int *iWidth, int *iHeight) {
  return ((*decoder)->DecodeFrame)(decoder, pSrc, iSrcLen, ppDst, pStride, iWidth, iHeight);
}

DECODING_STATE decoder_decode_frame2(ISVCDecoder *decoder, const unsigned char *pSrc, const int iSrcLen, void **ppDst, SBufferInfo *pDstInfo) {
  return ((*decoder)->DecodeFrame2)(decoder, pSrc, iSrcLen, ppDst, pDstInfo);
}

DECODING_STATE decoder_decode_frame_ex(ISVCDecoder *decoder, const unsigned char *pSrc, const int iSrcLen, unsigned char *pDst, int iDstStride, int *iDstLen, int *iWidth, int *iHeight, int* iColorFormat) {
  return ((*decoder)->DecodeFrameEx)(decoder, pSrc, iSrcLen, pDst, iDstStride, iDstLen, iWidth, iHeight, iColorFormat);
}

long decoder_set_option(ISVCDecoder *decoder, DECODER_OPTION eOptionId, void *pOption) {
  return ((*decoder)->SetOption)(decoder, eOptionId, pOption);
}

long decoder_get_option(ISVCDecoder *decoder, DECODER_OPTION eOptionId, void *pOption) {
  return ((*decoder)->GetOption)(decoder, eOptionId, pOption);
}

SSysMEMBuffer * get_buffer_info_union_pointer(SBufferInfo *bInfo) {
  return &bInfo->UsrData.sSystemBuffer;
}
*/
import "C"

import (
	"unsafe"
)

////////////////////
// codec_api.h
////////////////////
// change return values to errors instead of int or int64

///ENCODER

type ISVCEncoder struct {
	base *C.ISVCEncoder
}

// CreateSVCEncoder
func CreateSVCEncoder() (encoder *ISVCEncoder, ret int) {
	var ppEncoder *C.ISVCEncoder
	_ppEncoder := (**C.ISVCEncoder)(unsafe.Pointer(&ppEncoder))
	_ret := C.CreateSVCEncoder(_ppEncoder)
	encoder = &ISVCEncoder{ppEncoder}
	ret = int(_ret)
	return
}

// DestroySVCEncoder
func DestroySVCEncoder(encoder *ISVCEncoder) {
	_pEncoder := encoder.base
	C.DestroySVCEncoder(_pEncoder)
}

func (e *ISVCEncoder) Initialize(pParam *EncParamBase) (ret int) {
  pParam.updateBase()
	_ret := C.encoder_initialize(e.base, pParam.base)
	ret = int(_ret)
	return
}

func (e *ISVCEncoder) InitializeExt(pParam *EncParamExt) (ret int) {
	_ret := C.encoder_initialize_ext(e.base, pParam.base)
	ret = int(_ret)
	return
}

func (e *ISVCEncoder) Uninitialize() (ret int) {
	_ret := C.encoder_uninitialize(e.base)
	ret = int(_ret)
	return
}

func (e *ISVCEncoder) EncodeFrame(kpSrcPic *SourcePicture, pBsInfo *FrameBSInfo) (ret int) {
	kpSrcPic.updateBase()
	pBsInfo.updateBase()
	_ret := C.encoder_encode_frame(e.base, kpSrcPic.base, pBsInfo.base)
	pBsInfo.updateStruct()
	kpSrcPic.updateStruct()
	ret = int(_ret)
	return
}

func (e *ISVCEncoder) EncodeParameterSets(pBsInfo *FrameBSInfo) (ret int) {
	_ret := C.encoder_encode_parameter_sets(e.base, pBsInfo.base)
	ret = int(_ret)
	return
}

func (e *ISVCEncoder) PauseFrame(kpSrcPic *SourcePicture, pBsInfo *FrameBSInfo) (ret int) {
	_ret := C.encoder_pause_frame(e.base, kpSrcPic.base, pBsInfo.base)
	ret = int(_ret)
	return
}

func (e *ISVCEncoder) ForceIntraFrame(bIDR bool) (ret int) {
	_bIDR := C.bool(bIDR)
	_ret := C.encoder_force_intra_frame(e.base, _bIDR)
	ret = int(_ret)
	return
}

//think the uintptr goes to interface{}?

func (e *ISVCEncoder) SetOption(eOptionId EncoderOption, pOption uintptr) (ret int) {
	_pOption := unsafe.Pointer(pOption)
	_eOptionId := C.ENCODER_OPTION(eOptionId)
	_ret := C.encoder_set_option(e.base, _eOptionId, _pOption)
	ret = int(_ret)
	return
}

func (e *ISVCEncoder) GetOption(eOptionId EncoderOption, pOption uintptr) (ret int) {
	_pOption := unsafe.Pointer(pOption)
	_eOptionId := C.ENCODER_OPTION(eOptionId)
	_ret := C.encoder_get_option(e.base, _eOptionId, _pOption)
	ret = int(_ret)
	return
}

////DECODER

type ISVCDecoder struct {
	base *C.ISVCDecoder
}

// CreateDecoder
func CreateDecoder() (decoder *ISVCDecoder, ret int) {
	var ppDecoder *C.ISVCDecoder
	_ppDecoder := (**C.ISVCDecoder)(unsafe.Pointer(&ppDecoder))
	_ret := C.CreateDecoder(_ppDecoder)
	decoder = &ISVCDecoder{ppDecoder}
	ret = int(_ret)
	return
}

// DestroyDecoder
func DestroyDecoder(decoder *ISVCDecoder) {
	_pDecoder := decoder.base
	C.DestroyDecoder(_pDecoder)
}

func (d *ISVCDecoder) Initialize(pParam *DecodingParam) (ret int64) {
	_ret := C.decoder_initialize(d.base, pParam.base)
	ret = int64(_ret)
	return
}

func (d *ISVCDecoder) Uninitialize() (ret int64) {
	_ret := C.decoder_uninitialize(d.base)
	ret = int64(_ret)
	return
}


/*
func (d *ISVCDecoder) DecodeFrame (pSrc string, iSrcLen int, ppDst *string, pStride *int, iWidth *int, iHeight *int) (ret DecodingState) {

}
*/

func (d *ISVCDecoder) DecodeFrame2 (source []byte, pDstInfo *BufferInfo) (data [][]byte, ret DecodingState) {
  var _data [3]*byte
  var _pSrc *C.uchar
  _srcLen := C.int(len(source))
  if _srcLen > 0 {
    _pSrc = (*C.uchar)(&source[0])
  } else {
    _pSrc = nil
  }
  pDstInfo.updateBase()
  _ret := C.decoder_decode_frame2(d.base, _pSrc, _srcLen, (*unsafe.Pointer)(unsafe.Pointer(&_data[0])), pDstInfo.base)
  pDstInfo.updateStruct()
  q := make([][]byte, 3)
  var width, height, stride int
  ret = DecodingState(_ret)
  if pDstInfo.BufferStatus == 1 {
    for i, _d := range _data {
      if i == 0 {
        width = pDstInfo.SystemBuffer.Width
        height = pDstInfo.SystemBuffer.Height
        stride = pDstInfo.SystemBuffer.Stride[0]        
      } else {
        width = pDstInfo.SystemBuffer.Width/2
        height = pDstInfo.SystemBuffer.Height/2
        stride = pDstInfo.SystemBuffer.Stride[1]
      }
      q[i] = (*[1<<30]byte)(unsafe.Pointer(_d))[:width+stride*(height-1)]
    }
    data = q
  }
  return
}


/*
{ // y plane
    static_cast<uint8_t*>(data[0]),
    bufInfo.UsrData.sSystemBuffer.iWidth,
    bufInfo.UsrData.sSystemBuffer.iHeight,
    bufInfo.UsrData.sSystemBuffer.iStride[0]
},
{ // u plane
    static_cast<uint8_t*>(data[1]),
    bufInfo.UsrData.sSystemBuffer.iWidth / 2,
    bufInfo.UsrData.sSystemBuffer.iHeight / 2,
    bufInfo.UsrData.sSystemBuffer.iStride[1]
},
{ // v plane
    static_cast<uint8_t*>(data[2]),
    bufInfo.UsrData.sSystemBuffer.iWidth / 2,
    bufInfo.UsrData.sSystemBuffer.iHeight / 2,
    bufInfo.UsrData.sSystemBuffer.iStride[1]
},
*/

/*
func (d *ISVCDecoder) DecodeFrameEx (pSrc string, iSrcLen int, pDst string, iDstStride int, iDstLen *int, iWidth *int, iHeight *ing) (ret DecodingState) {

}
*/

func (d *ISVCDecoder) SetOption(eOptionId EncoderOption, pOption uintptr) (ret int64) {
	_pOption := unsafe.Pointer(pOption)
	_eOptionId := C.DECODER_OPTION(eOptionId)
	_ret := C.decoder_set_option(d.base, _eOptionId, _pOption)
	ret = int64(_ret)
	return
}

func (d *ISVCDecoder) GetOption(eOptionId EncoderOption, pOption uintptr) (ret int64) {
	_pOption := unsafe.Pointer(pOption)
	_eOptionId := C.DECODER_OPTION(eOptionId)
	_ret := C.decoder_get_option(d.base, _eOptionId, _pOption)
	ret = int64(_ret)
	return
}

////////////////////
// codec_app_def.h
////////////////////

const (
	MAX_TEMPORAL_LAYER_NUM = 4
	MAX_SPATIAL_LAYER_NUM  = 4
	MAX_QUALITY_LAYER_NUM  = 4

	MAX_LAYER_NUM_OF_FRAME = 128
	MAX_NAL_UNITS_IN_LAYER = 128

	MAX_RTP_PAYLOAD_LEN     = 1000
	AVERAGE_RTP_PAYLOAD_LEN = 800

	SAVED_NALUNIT_NUM_TMP = (MAX_SPATIAL_LAYER_NUM * MAX_QUALITY_LAYER_NUM) + 1 + MAX_SPATIAL_LAYER_NUM
	MAX_SLICES_NUM_TMP    = (MAX_NAL_UNITS_IN_LAYER - SAVED_NALUNIT_NUM_TMP) / 3
)

// DECODING_STATE
type DecodingState int32

const (
	DsErrorFree          = 0x0000 //0
	DsFramePending       = 0x0001 //1
	DsRefLost            = 0x0002 //2
	DsBitstreamError     = 0x0004 //4
	DsDepLayerLost       = 0x0008 //8
	DsNoParamSets        = 0x0010 //16
	DsInvalidArgument    = 0x1000 //4096
	DsInitialOptExpected = 0x2000 //8192
	DsOutOfMemory        = 0x4000 //16384
	DsDstBufNeedExpand   = 0x8000 //32768
)

// ENCODER_OPTION
type EncoderOption int32

const (
	EncoderOptionDataformat             = 0x00
	EncoderOptionIdrInterval            = 0x01
	EncoderOptionSvcEncodeParamBase     = 0x02
	EncoderOptionSvcEncodeParamExt      = 0x03
	EncoderOptionFrameRate              = 0x04
	EncoderOptionBitrate                = 0x05
	EncoderOptionInterSpatialPred       = 0x06
	EncoderOptionRcMode                 = 0x07
	EncoderPaddingPadding               = 0x08
	EncoderLtrRecoveryRequest           = 0x09
	EncoderLtrMarkingFeedback           = 0x0A
	EncocerLtrMarkingPeriod             = 0x0B
	EncoderOptionLtr                    = 0x0C
	EncoderOptionEnableSsei             = 0x0D
	EncoderOptionEnablePrefixNalAdding  = 0x0E
	EncoderOptionEnableSpsPpsIdAddition = 0x0F
	EncoderOptionCurrentPath            = 0x10
)

// DECODER_OPTION
type DecoderOption int32

const (
	DecoderOptionDataformat        = 0x0
	DecoderOptionEndOfStream       = 0x1
	DecoderOptionVclNal            = 0x2
	DecoderOptionTemporalId        = 0x3
	DecoderOptionFrameNum          = 0x4
	DecoderOptionIdrPicId          = 0x5
	DecoderOptionLtrMarkingFlag    = 0x6
	DecoderOptionLtrMarkedFrameNum = 0x7
)

// FEEDBACK_VCL_NAL_IN_AU
type FeedbackVclNalInAu int32

const (
	FeedbackNonVclNal  = 0x0
	FeedbackVclNal     = 0x1
	FeedbackUnknownNal = 0x2
)

// LAYER_TYPE
type LayerType int32

const (
	NonVideoCodingLayer = 0x0
	VideoCodingLayer    = 0x1
)

// VIDEO_BITSTREAM_TYPE
type VideoBitstreamType int32

const (
	VideoBitstreamAvc     = 0x0
	VideoBitstreamSvc     = 0x1
	VideoBitstreamDefault = 0x1
)

// KEY_FRAME_REQUEST_TYPE
type KeyFrameRequestType int32

const (
	NoRecoveryRequset    = 0x0
	LtrRecoveryRequest   = 0x1
	IdrRecoveryRequest   = 0x2
	NoLtrMarkingFeedback = 0x3
	LtrMarkingSuccess    = 0x4
	LtrMarkingFailed     = 0x5
)

func NewLTRRecoverRequest() *LTRRecoverRequest {
	var fType, id uint
	var lFrameNum, frameNum int
	var _ltrRecoverRequest C.SLTRRecoverRequest

	return &LTRRecoverRequest{fType, id, lFrameNum, frameNum, &_ltrRecoverRequest}
}

func (p *LTRRecoverRequest) updateBase() {
	p.base.uiFeedbackType = C.uint(p.FeedbackType)
	p.base.uiIDRPicId = C.uint(p.IDRPicId)
	p.base.iLastCorrectFrameNum = C.int(p.LastCorrectFrameNum)
	p.base.iCurrentFrameNum = C.int(p.CurrentFrameNum)
}

func (p *LTRRecoverRequest) updateStruct() {
	p.FeedbackType = uint(p.base.uiFeedbackType)
	p.IDRPicId = uint(p.base.uiIDRPicId)
	p.LastCorrectFrameNum = int(p.base.iLastCorrectFrameNum)
	p.CurrentFrameNum = int(p.base.iCurrentFrameNum)
}

type LTRRecoverRequest struct {
	FeedbackType        uint `h264:"uiFeedbackType"`
	IDRPicId            uint `h264:"uiIDRPicId"`
	LastCorrectFrameNum int  `h264:"iLastCorrectFrameNum"`
	CurrentFrameNum     int  `h264:"iCurrentFrameNum"`

	base *C.SLTRRecoverRequest
}

/*
typedef struct {
  unsigned int uiFeedbackType; //IDR request or LTR recovery request
  unsigned int uiIDRPicId; // distinguish request from different IDR
  int     iLastCorrectFrameNum;
  int     iCurrentFrameNum; //specify current decoder frame_num.
} SLTRRecoverRequest;
*/

func NewLTRMarkingFeedback() *LTRMarkingFeedback {
	var fType, id uint
	var frame int
	var _ltrMarkingFeedback C.SLTRMarkingFeedback

	return &LTRMarkingFeedback{fType, id, frame, &_ltrMarkingFeedback}
}

func (p *LTRMarkingFeedback) updateBase() {
	p.base.uiFeedbackType = C.uint(p.FeedbackType)
	p.base.uiIDRPicId = C.uint(p.IDRPicId)
	p.base.iLTRFrameNum = C.int(p.LTRFrameNum)
}

func (p *LTRMarkingFeedback) updateStruct() {
	p.FeedbackType = uint(p.base.uiFeedbackType)
	p.IDRPicId = uint(p.base.uiIDRPicId)
	p.LTRFrameNum = int(p.base.iLTRFrameNum)
}

type LTRMarkingFeedback struct {
	FeedbackType uint `h264:"uiFeedbackType"`
	IDRPicId     uint `h264:"uiIDRPicId"`
	LTRFrameNum  int  `h264:"iLTRFrameNum"`

	base *C.SLTRMarkingFeedback
}

/*
typedef struct {
  unsigned int  uiFeedbackType; //mark failed or successful
  unsigned int  uiIDRPicId; // distinguish request from different IDR
  int       iLTRFrameNum; //specify current decoder frame_num
} SLTRMarkingFeedback;
*/

func NewSliceArgument() *SliceArgument {
	var mbNum [MAX_SLICES_NUM_TMP]uint
	var num, size uint
	var _sliceArgument C.SSliceArgument

	return &SliceArgument{mbNum, num, size, &_sliceArgument}
}

func (p *SliceArgument) updateBase() {
	for index, val := range p.SliceMbNum {
		p.base.uiSliceMbNum[index] = C.uint(val)
	}
	p.base.uiSliceNum = C.uint(p.SliceNum)
	p.base.uiSliceSizeConstraint = C.uint(p.SliceSizeConstraint)
}

func (p *SliceArgument) updateStruct() {
	for index, val := range p.base.uiSliceMbNum {
		p.SliceMbNum[index] = uint(val)
	}
	p.SliceNum = uint(p.base.uiSliceNum)
	p.SliceSizeConstraint = uint(p.base.uiSliceSizeConstraint)
}

type SliceArgument struct {
	SliceMbNum          [MAX_SLICES_NUM_TMP]uint `h264:"uiSliceMbNum"`
	SliceNum            uint                     `h264:"uiSliceNum"`
	SliceSizeConstraint uint                     `h264:"uiSliceSizeConstraint"`

	base *C.SSliceArgument
}

/*
typedef struct {
    unsigned int
    uiSliceMbNum[MAX_SLICES_NUM_TMP];  //here we use a tmp fixed value since MAX_SLICES_NUM is not defined here and its definition may be changed;
    unsigned int    uiSliceNum;
    unsigned int    uiSliceSizeConstraint;
  } SSliceArgument;//not all the elements in this argument will be used, how it will be used depends on uiSliceMode; see below
*/

func NewSliceConfig() *SliceConfig {
	var mode uint
	var _sliceArgument *SliceArgument
	var _sliceConfig C.SSliceConfig

	_sliceArgument = NewSliceArgument()
  _sliceArgument.base = &_sliceConfig.sSliceArgument 

	return &SliceConfig{mode, _sliceArgument, &_sliceConfig}
}

func (p *SliceConfig) updateBase() {
	p.base.uiSliceMode = C.uint(p.SliceMode)
	p.SliceArgument.updateBase()
}

func (p *SliceConfig) updateStruct() {
	p.SliceMode = uint(p.base.uiSliceMode)
	p.SliceArgument.updateStruct()
}

type SliceConfig struct {
	SliceMode     uint           `h264:"uiSliceMode"`
	SliceArgument *SliceArgument `h264:"sSliceArgument"`

	base *C.SSliceConfig
}

/*
typedef struct {

  //# 0 SM_SINGLE_SLICE     | SliceNum==1
  //# 1 SM_FIXEDSLCNUM_SLICE  | according to SliceNum     | Enabled dynamic slicing for multi-thread
  //# 2 SM_RASTER_SLICE     | according to SlicesAssign   | Need input of MB numbers each slice. In addition, if other constraint in SSliceArgument is presented, need to follow the constraints. Typically if MB num and slice size are both constrained, re-encoding may be involved.
  //# 3 SM_ROWMB_SLICE      | according to PictureMBHeight  |  Typical of single row of mbs each slice?+ slice size constraint which including re-encoding
  //# 4 SM_DYN_SLICE      | according to SliceSize    | Dynamic slicing (have no idea about slice_nums until encoding current frame)
  unsigned int uiSliceMode; //by default, uiSliceMode will be 0
  SSliceArgument sSliceArgument;
} SSliceConfig;
*/

func NewSpatialLayerConfig() *SpatialLayerConfig {
	var width, height, bitrate, initialQP int
	var rateOut float64
	var idc uint
	var _sliceConfig *SliceConfig
	var _spatialLayerConfig C.SSpatialLayerConfig

	_sliceConfig = NewSliceConfig()

  _sliceConfig.base = &_spatialLayerConfig.sSliceCfg

	return &SpatialLayerConfig{width, height, rateOut, bitrate, idc, initialQP, _sliceConfig, &_spatialLayerConfig}
}

func (p *SpatialLayerConfig) updateBase() {
	p.base.iVideoWidth = C.int(p.FrameWidth)
	p.base.iVideoHeight = C.int(p.FrameHeight)
	p.base.fFrameRate = C.float(p.FrameRateOut)
	p.base.iSpatialBitrate = C.int(p.SpatialBitrate)
	p.base.uiProfileIdc = C.uint(p.ProfileIdc)
	p.base.iDLayerQp = C.int(p.InitialQP)

	p.SliceConfig.updateBase()
}

func (p *SpatialLayerConfig) updateStruct() {
	p.FrameWidth = int(p.base.iVideoWidth)
	p.FrameHeight = int(p.base.iVideoHeight)
	p.FrameRateOut = float64(p.base.fFrameRate)
	p.SpatialBitrate = int(p.base.iSpatialBitrate)
	p.ProfileIdc = uint(p.base.uiProfileIdc)
	p.InitialQP = int(p.base.iDLayerQp)

	p.SliceConfig.updateStruct()
}

type SpatialLayerConfig struct {
	FrameWidth     int     `h264:"iVideoWidth"`
	FrameHeight    int     `h264:"iVideoHeight"`
	FrameRateOut   float64 `h264:"fFrameRate"`
	SpatialBitrate int     `h264:"iSpatialBitrate"`
	ProfileIdc     uint    `h264:"uiProfileIdc"`
	InitialQP      int     `h264:"iDLayerQp"`

	SliceConfig *SliceConfig `h264:"sSliceCfg"`

	base *C.SSpatialLayerConfig
}

/*
typedef struct {
  int iVideoWidth;    // video size in cx specified for a layer
  int iVideoHeight;   // video size in cy specified for a layer
  float fFrameRate;   // frame rate specified for a layer
  int iSpatialBitrate;  // target bitrate for a spatial layer
  unsigned int  uiProfileIdc; // value of profile IDC (0 for auto-detection)
  int    iDLayerQp;

  SSliceConfig sSliceCfg;
} SSpatialLayerConfig;
*/

func NewEncParamBase() *EncParamBase {
	var usage, input, width, height, bitrate, rcMode int
	var frameRate float64
	var _encParamBase C.SEncParamBase

	return &EncParamBase{usage, input, width, height, bitrate, rcMode, frameRate, &_encParamBase}
}

func (p *EncParamBase) updateBase() {
	p.base.iUsageType = C.int(p.UsageType)
	p.base.iInputCsp = C.int(p.InputCsp)
	p.base.iPicWidth = C.int(p.SourceWidth)
	p.base.iPicHeight = C.int(p.SourceHeight)
	p.base.iTargetBitrate = C.int(p.TargetBitrate)
	p.base.iRCMode = C.int(p.RCMode)
	p.base.fMaxFrameRate = C.float(p.MaxFrameRate)
}

func (p *EncParamBase) updateStruct() {
	p.UsageType = int(p.base.iUsageType)
	p.InputCsp = int(p.base.iInputCsp)
	p.SourceWidth = int(p.base.iPicWidth)
	p.SourceHeight = int(p.base.iPicHeight)
	p.TargetBitrate = int(p.base.iTargetBitrate)
	p.RCMode = int(p.base.iRCMode)
	p.MaxFrameRate = float64(p.base.fMaxFrameRate)
}

type EncParamBase struct {
	UsageType int `h264:"iUsageType"`
	InputCsp  int `h264:"iInputCsp"`

	SourceWidth   int     `h264:"iPicWidth"`
	SourceHeight  int     `h264:"iPicHeight"`
	TargetBitrate int     `h264:"iTargetBitrate"`
	RCMode        int     `h264:"iRCMode"`
	MaxFrameRate  float64 `h264:"fMaxFrameRate"`

	base *C.SEncParamBase
}

/*
typedef struct TagEncParamBase{

  int       iUsageType; //enable_screen_content_signal;// 0: //camera video signal; 1: screen content signal;
  int   iInputCsp;  // color space of input sequence

  int   iPicWidth;      // width of picture in samples
  int   iPicHeight;     // height of picture in samples
  int   iTargetBitrate;   // target bitrate desired
  int       iRCMode;                 // RC mode
  float     fMaxFrameRate;      // input maximal frame rate

} SEncParamBase, *PEncParamBase;
*/

func NewEncParamExt() *EncParamExt {
	var usage, input, width, height, bitrate, rcMode, temporal, numLayers, numRef, pFlag, eFlag, minQ, maxQ, ltrRef, ltrMark, lFDI, lFAO, lFBO, iLLFD, iLLFAO, iLLFBO int
	var maxFrameRate float64
	var _sLayers [MAX_SPATIAL_LAYER_NUM]*SpatialLayerConfig
	var intra, frameEnc uint
	var eSps, pNal, eScal, eRc, eFs, eLTR, eD, eBD, eAQ, eFCF, eSSD bool
	var mulT, couT int16
	var _encParamExt C.SEncParamExt

	for index, _ := range _sLayers {
		_sLayers[index] = NewSpatialLayerConfig()
		_sLayers[index].base = &_encParamExt.sSpatialLayers[index]
	}

	return &EncParamExt{usage, input, width, height, bitrate, rcMode, maxFrameRate, temporal, numLayers, _sLayers, intra, numRef, frameEnc, eSps, pNal, eScal, pFlag, eFlag, eRc, eFs, maxQ, minQ, eLTR, ltrRef, ltrMark, mulT, couT, lFDI, lFAO, lFBO, iLLFD, iLLFAO, iLLFBO, eD, eBD, eAQ, eFCF, eSSD, &_encParamExt}
}

func (p *EncParamExt) updateBase() {
  p.base.iUsageType = C.int(p.UsageType)
  p.base.iInputCsp = C.int(p.InputCsp)
  p.base.iPicWidth = C.int(p.SourceWidth)
  p.base.iTargetBitrate = C.int(p.TargetBitrate)
  p.base.iRCMode = C.int(p.RCMode)
  p.base.fMaxFrameRate = C.float(p.MaxFrameRate)
  p.base.iTemporalLayerNum = C.int(p.TemporalLayerNumber)
  p.base.iSpatialLayerNum = C.int(p.NumLayers)
  
  for _, layer := range p.SpatialLayers {
    layer.updateBase()
  }
  
  p.base.uiIntraPeriod = C.uint(p.IntraPeriod)
  p.base.iNumRefFrame = C.int(p.NumRefFrame)
  p.base.uiFrameToBeCoded = C.uint(p.FramesToBeEncoded)
  p.base.bEnableSpsPpsIdAddition = C.bool(p.EnableSpsPpsIDAddition)
  p.base.bPrefixNalAddingCtrl = C.bool(p.PrefixNALAddingCtrl)
  p.base.bEnableSSEI = C.bool(p.EnableScalableSEI)
  p.base.iPaddingFlag = C.int(p.PaddingFlag)
  p.base.iEtropyCodingModeFlag = C.int(p.EntropyCodingModeFlag)
  p.base.bEnableRc = C.bool(p.EnableRC)
  p.base.bEnableFrameSkip = C.bool(p.EnableFrameSkip)
  p.base.iMaxQp = C.int(p.MaxQp)
  p.base.iMinQp = C.int(p.MinQp)
  p.base.bEnableLongTermReference = C.bool(p.EnableLongTermReference)
  p.base.iLTRRefNum = C.int(p.LTRRefNum)
  p.base.iLtrMarkPeriod = C.int(p.LtrMarkPeriod)
  p.base.iMultipleThreadIdc = C.short(p.MultipleThreadIdc)
  p.base.iCountThreadsNum = C.short(p.CountThreadsNum)
  p.base.iLoopFilterDisableIdc = C.int(p.LoopFilterDisableIDC)
  p.base.iLoopFilterAlphaC0Offset = C.int(p.LoopFilterAlphaC0Offset)
  p.base.iLoopFilterBetaOffset = C.int(p.LoopFilterBetaOffset)
  p.base.iInterLayerLoopFilterDisableIdc = C.int(p.InterLayerLoopFilterDisableIDC)
  p.base.iInterLayerLoopFilterAlphaC0Offset = C.int(p.InterLayerLoopFilterAlphaC0Offset)
  p.base.iInterLayerLoopFilterBetaOffset = C.int(p.InterLayerLoopFilterBetaOffset)
  p.base.bEnableDenoise = C.bool(p.EnableDenoise)
  p.base.bEnableBackgroundDetection = C.bool(p.EnableBackgroundDetection)
  p.base.bEnableAdaptiveQuant = C.bool(p.EnableAdaptiveQuantization)
  p.base.bEnableFrameCroppingFlag = C.bool(p.EnableFrameCroppingFlag)
  p.base.bEnableSceneChangeDetect = C.bool(p.EnableSceneChangeDetection)
}

func (p *EncParamExt) updateStruct() {
  p.UsageType = int(p.base.iUsageType)
  p.InputCsp = int(p.base.iInputCsp)
  p.SourceWidth = int(p.base.iPicWidth)
  p.SourceHeight = int(p.base.iPicHeight)
  p.TargetBitrate = int(p.base.iTargetBitrate)
  p.RCMode = int(p.base.iRCMode)
  p.MaxFrameRate = float64(p.base.fMaxFrameRate)
  p.TemporalLayerNumber = int(p.base.iTemporalLayerNum)
  p.NumLayers = int(p.base.iSpatialLayerNum)
  
  for _, layer := range p.SpatialLayers {
    layer.updateStruct()
  }
  
  p.IntraPeriod = uint(p.base.uiIntraPeriod)
  p.NumRefFrame = int(p.base.iNumRefFrame)
  p.FramesToBeEncoded = uint(p.base.uiFrameToBeCoded)
  p.EnableSpsPpsIDAddition = bool(p.base.bEnableSpsPpsIdAddition)
  p.PrefixNALAddingCtrl = bool(p.base.bPrefixNalAddingCtrl)
  p.EnableScalableSEI = bool(p.base.bEnableSSEI)
  p.PaddingFlag = int(p.base.iPaddingFlag)
  p.EntropyCodingModeFlag = int(p.base.iEtropyCodingModeFlag)
  p.EnableRC = bool(p.base.bEnableRc)
  p.EnableFrameSkip = bool(p.base.bEnableFrameSkip)
  p.MaxQp = int(p.base.iMaxQp)
  p.MinQp = int(p.base.iMinQp)
  p.EnableLongTermReference = bool(p.base.bEnableLongTermReference)
  p.LTRRefNum = int(p.base.iLTRRefNum)
  p.LtrMarkPeriod = int(p.base.iLtrMarkPeriod)
  p.MultipleThreadIdc = int16(p.base.iMultipleThreadIdc)
  p.CountThreadsNum = int16(p.base.iCountThreadsNum)
  p.LoopFilterDisableIDC = int(p.base.iLoopFilterDisableIdc)
  p.LoopFilterAlphaC0Offset = int(p.base.iLoopFilterAlphaC0Offset)
  p.LoopFilterBetaOffset = int(p.base.iLoopFilterBetaOffset)
  p.InterLayerLoopFilterDisableIDC = int(p.base.iInterLayerLoopFilterDisableIdc)
  p.InterLayerLoopFilterAlphaC0Offset = int(p.base.iInterLayerLoopFilterAlphaC0Offset)
  p.InterLayerLoopFilterBetaOffset = int(p.base.iInterLayerLoopFilterBetaOffset)
  p.EnableDenoise = bool(p.base.bEnableDenoise)
  p.EnableBackgroundDetection = bool(p.base.bEnableBackgroundDetection)
  p.EnableAdaptiveQuantization = bool(p.base.bEnableAdaptiveQuant)
  p.EnableFrameCroppingFlag = bool(p.base.bEnableFrameCroppingFlag)
  p.EnableSceneChangeDetection = bool(p.base.bEnableSceneChangeDetect)
}

type EncParamExt struct {
	UsageType int //iUsageType
	InputCsp  int //iInputCsp

	SourceWidth   int     //iPicWidth
	SourceHeight  int     //iPicHeight
	TargetBitrate int     //iTargetBitrate
	RCMode        int     //iRCMode
	MaxFrameRate  float64 //fMaxFrameRate

	TemporalLayerNumber int                                        //iTemporalLayerNum
	NumLayers           int                                        // iSpatialLayerNum
	SpatialLayers       [MAX_SPATIAL_LAYER_NUM]*SpatialLayerConfig // sSpatialLayers[MAX_SPATIAL_LAYER_NUM]

	IntraPeriod            uint //uiIntraPeriod
	NumRefFrame            int  // iNumRefFrame
	FramesToBeEncoded      uint //uiFrameToBeCoded
	EnableSpsPpsIDAddition bool //bEnableSpsPpsIDAddition
	PrefixNALAddingCtrl    bool //bPrefixNalAddingCtrl
	EnableScalableSEI     bool //bEnableSSEI
	PaddingFlag            int  //iPaddingFlag
	EntropyCodingModeFlag  int  //iEntropyCodingModeFlag

	EnableRC        bool //bEnableRc
	EnableFrameSkip bool //bEnableFrameSkip
	MaxQp           int  //iMaxQp
	MinQp           int  //iMinQp

	EnableLongTermReference bool //bEnableLongTermReference
	LTRRefNum               int  //iLTRRefNum
	LtrMarkPeriod           int  //iLtrMarkPeriod

	MultipleThreadIdc int16 //iMultipleThreadIdc
	CountThreadsNum   int16 //iCountThreadsNum

	LoopFilterDisableIDC             int //iLoopFilterDisableIdc
	LoopFilterAlphaC0Offset          int //iLoopFilterAlphaC0Offset
	LoopFilterBetaOffset             int //iLoopFilterBetaOffset
	InterLayerLoopFilterDisableIDC   int //iInterLayerLoopFilterDisableIdc
	InterLayerLoopFilterAlphaC0Offset int //iInterLayerLoopFilterAlphaC0Offset
	InterLayerLoopFilterBetaOffset   int //iInterLayerLoopFilterBetaOffset

	EnableDenoise              bool //bEnableDenoise
	EnableBackgroundDetection  bool //bEnableBackgroundDetection
	EnableAdaptiveQuantization bool //bEnableAdaptiveQuant
	EnableFrameCroppingFlag    bool //bEnableFrameCroppingFlag
	EnableSceneChangeDetection bool //bEnableSceneChangeDetection

	base *C.SEncParamExt
}

/*
typedef struct TagEncParamExt
{
  int       iUsageType; //application type;// 0: //camera video signal; 1: screen content signal;
  int   iInputCsp;  // color space of input sequence

  int   iPicWidth;      // width of picture in samples
  int   iPicHeight;     // height of picture in samples
  int   iTargetBitrate;   // target bitrate desired
  int       iRCMode;                 // RC mode
  float     fMaxFrameRate;      // input maximal frame rate

  int   iTemporalLayerNum;  // layer number at temporal level
  int   iSpatialLayerNum; // layer number at spatial level
  SSpatialLayerConfig sSpatialLayers[MAX_SPATIAL_LAYER_NUM];

  unsigned int    uiIntraPeriod;    // period of Intra frame
  int           iNumRefFrame;   // number of reference frame used
  unsigned int      uiFrameToBeCoded; // frame to be encoded (at input frame rate)
  bool    bEnableSpsPpsIdAddition;
  bool    bPrefixNalAddingCtrl;
  bool    bEnableSSEI;
  int      iPaddingFlag;            // 0:disable padding;1:padding
  int      iEtropyCodingModeFlag;

  bool    bEnableRc;
  bool    bEnableFrameSkip; // allow skipping frames to keep the bitrate within limits
  int     iMaxQp;
  int     iMinQp;

  bool     bEnableLongTermReference; // 0: on, 1: off
  int    iLTRRefNum;
  int      iLtrMarkPeriod;

  short   iMultipleThreadIdc;   // 1  # 0: auto(dynamic imp. internal encoder); 1: multiple threads imp. disabled; > 1: count number of threads;
  short   iCountThreadsNum;     //    # derived from disable_multiple_slice_idc (=0 or >1) means;

  int   iLoopFilterDisableIdc;  // 0: on, 1: off, 2: on except for slice boundaries
  int   iLoopFilterAlphaC0Offset;// AlphaOffset: valid range [-6, 6], default 0
  int   iLoopFilterBetaOffset;  // BetaOffset:  valid range [-6, 6], default 0
  int   iInterLayerLoopFilterDisableIdc; // Employed based upon inter-layer, same comment as above
  int   iInterLayerLoopFilterAlphaC0Offset; // InterLayerLoopFilterAlphaC0Offset
  int   iInterLayerLoopFilterBetaOffset;  // InterLayerLoopFilterBetaOffset

  bool    bEnableDenoise;     // denoise control
  bool    bEnableBackgroundDetection;// background detection control //VAA_BACKGROUND_DETECTION //BGD cmd
  bool    bEnableAdaptiveQuant; // adaptive quantization control
  bool    bEnableFrameCroppingFlag;// enable frame cropping flag: TRUE always in application
  bool    bEnableSceneChangeDetect;
}SEncParamExt;
*/

func NewVideoProperty() *VideoProperty {
	var size uint
	var vType VideoBitstreamType //enum
	var _vidProperty C.SVideoProperty

	return &VideoProperty{size, vType, &_vidProperty}
}

func (p *VideoProperty) updateBase() {
	p.base.size = C.uint(p.Size)
	p.base.eVideoBsType = C.VIDEO_BITSTREAM_TYPE(p.VideoBsType)
}

func (p *VideoProperty) updateStruct() {
	p.Size = uint(p.base.size)
	p.VideoBsType = VideoBitstreamType(p.base.eVideoBsType)
}

type VideoProperty struct {
	Size        uint               //size
	VideoBsType VideoBitstreamType //eVideoBsType

	base *C.SVideoProperty
}

/*
//Define a new struct to show the property of video bitstream.
typedef struct {
  unsigned int          size; //size of the struct
  VIDEO_BITSTREAM_TYPE  eVideoBsType;
} SVideoProperty;
*/

func NewDecodingParam() *DecodingParam {
	var fName string
	var format int
	var load uint
	var target, active uint8
	var _decodingParam C.SDecodingParam

	var _videoProperty = NewVideoProperty()

  _videoProperty.base = &_decodingParam.sVideoProperty

	return &DecodingParam{fName, format, load, target, active, _videoProperty, &_decodingParam}
}

func (p *DecodingParam) updateBase() {
	if p.base.pFileNameRestructed != nil {
		C.free(unsafe.Pointer(p.base.pFileNameRestructed))
	}
	p.base.pFileNameRestructed = C.CString(p.FileNameRestructed)
	p.base.iOutputColorFormat = C.int(p.OutputColorFormat)
	p.base.uiCpuLoad = C.uint(p.CPULoad)
	p.base.uiTargetDqLayer = C.uchar(p.TargetDQLayer)
	p.base.uiEcActiveFlag = C.uchar(p.ECActiveFlag)
	p.VideoProperty.updateBase()
}

func (p *DecodingParam) updateStruct() {
	p.FileNameRestructed = C.GoString(p.base.pFileNameRestructed)
	p.OutputColorFormat = int(p.base.iOutputColorFormat)
	p.CPULoad = uint(p.base.uiCpuLoad)
	p.TargetDQLayer = uint8(p.base.uiTargetDqLayer)
	p.ECActiveFlag = uint8(p.base.uiEcActiveFlag)
	p.VideoProperty.updateStruct()
}

type DecodingParam struct {
	FileNameRestructed string //pFileNameRestructured

	OutputColorFormat int   //iOutputColorFormat
	CPULoad           uint  //uiCpuLoad
	TargetDQLayer     uint8 //uiTargetDqLayer

	ECActiveFlag uint8 //uiEcActiveFlag

	VideoProperty *VideoProperty

	base *C.SDecodingParam
}

/*
typedef struct TagSVCDecodingParam {
  char*   pFileNameRestructed;  // File name of restructed frame used for PSNR calculation based debug

  int       iOutputColorFormat; // color space format to be outputed, EVideoFormatType specified in codec_def.h
  unsigned int  uiCpuLoad;    // CPU load
  unsigned char uiTargetDqLayer;  // Setting target dq layer id

  unsigned char uiEcActiveFlag;   // Whether active error concealment feature in decoder

  SVideoProperty   sVideoProperty;
} SDecodingParam, *PDecodingParam;
*/

func NewLayerBSInfo() *LayerBSInfo {
	var temporal, spatial, quality, priority, layerType uint8
	var nCount int
	var nLengthInByte [MAX_NAL_UNITS_IN_LAYER]int
	var buf *uint8
	var _layerBSInfo C.SLayerBSInfo
  
	return &LayerBSInfo{temporal, spatial, quality, priority, layerType, nCount, nLengthInByte, buf, &_layerBSInfo}
}

func (p *LayerBSInfo) updateBase() {
	p.base.uiTemporalId = C.uchar(p.TemporalId)
	p.base.uiSpatialId = C.uchar(p.SpatialId)
	p.base.uiQualityId = C.uchar(p.QualityId)
	p.base.uiPriorityId = C.uchar(p.PriorityId)
	p.base.uiLayerType = C.uchar(p.LayerType)
	p.base.iNalCount = C.int(p.NALCount)

	for index, val := range p.NALLengthInByte {
		p.base.iNalLengthInByte[index] = C.int(val)
	}
	
	p.base.pBsBuf = (*C.uchar)(unsafe.Pointer(p.BitstreamBuffer))
}

func (p *LayerBSInfo) updateStruct() {
  p.TemporalId = uint8(p.base.uiTemporalId)
  p.SpatialId = uint8(p.base.uiSpatialId)
  p.QualityId = uint8(p.base.uiQualityId)
  p.PriorityId = uint8(p.base.uiPriorityId)
  p.LayerType = uint8(p.base.uiLayerType)
  p.NALCount = int(p.base.iNalCount)
  
  for index, val := range p.base.iNalLengthInByte {
		p.NALLengthInByte[index] = int(val)
	}
	
	p.BitstreamBuffer = (*uint8)(unsafe.Pointer(p.base.pBsBuf))
}

type LayerBSInfo struct {
	TemporalId uint8 //uiTemporalId
	SpatialId  uint8 //uiSpatialId
	QualityId  uint8 //uiQualityId

	PriorityId uint8 //uiPriorityId

	LayerType uint8 //uiLayerType

	NALCount        int                         //iNalCount
	NALLengthInByte [MAX_NAL_UNITS_IN_LAYER]int //iNalLengthInByte
	BitstreamBuffer *uint8                      //pBsBuf

	base *C.SLayerBSInfo
}

/*
typedef struct {
  unsigned char uiTemporalId;
  unsigned char uiSpatialId;
  unsigned char uiQualityId;

  unsigned char uiPriorityId; //ignore it currently

  unsigned char uiLayerType;

  int iNalCount;          // Count number of NAL coded already
  int iNalLengthInByte[MAX_NAL_UNITS_IN_LAYER]; // Length of NAL size in byte from 0 to iNalCount-1
  unsigned char*  pBsBuf;   // Buffer of bitstream contained
} SLayerBSInfo, *PLayerBSInfo;
*/

func NewFrameBSInfo() *FrameBSInfo {
	var temporal, layerNum, oFType int
	var fType uint8
	var _frameBSInfo C.SFrameBSInfo
	var _lInfoArray [MAX_LAYER_NUM_OF_FRAME]*LayerBSInfo

	for index, _ := range _lInfoArray {
		_lInfoArray[index] = NewLayerBSInfo()
		_lInfoArray[index].base = &_frameBSInfo.sLayerInfo[index]		
	}

	return &FrameBSInfo{temporal, fType, layerNum, _lInfoArray, oFType, &_frameBSInfo}
}

func (p *FrameBSInfo) updateBase() {
	p.base.iTemporalId = C.int(p.TemporalId)
	p.base.uiFrameType = C.uchar(p.FrameType)
	p.base.iLayerNum = C.int(p.LayerNum)

	for _, layer := range p.LayerInfo {
		layer.updateBase()
	}

	p.base.eOutputFrameType = C.int(p.OutputFrameType)
}

func (p *FrameBSInfo) updateStruct() {
	p.TemporalId = int(p.base.iTemporalId)
	p.FrameType = uint8(p.base.uiFrameType)
	p.LayerNum = int(p.base.iLayerNum)

	for _, layer := range p.LayerInfo {
		layer.updateStruct()
	}

	p.OutputFrameType = int(p.base.eOutputFrameType)
}

type FrameBSInfo struct {
	TemporalId int   //iTemporalId
	FrameType  uint8 //uiFrameType

	LayerNum  int                                  //iLayerNum
	LayerInfo [MAX_LAYER_NUM_OF_FRAME]*LayerBSInfo //sLayerinfo

	OutputFrameType int //eOutputFrameType

	base *C.SFrameBSInfo
}

/*
typedef struct {
  int   iTemporalId;  // Temporal ID
  unsigned char uiFrameType;

  int   iLayerNum;
  SLayerBSInfo  sLayerInfo[MAX_LAYER_NUM_OF_FRAME];

  int eOutputFrameType;
} SFrameBSInfo, *PFrameBSInfo;
*/

func NewSourcePicture() *SourcePicture {
	var color, width, height int
	var stride [4]int
	var data [4]*uint8
	var _sourcePicture C.SSourcePicture

	return &SourcePicture{color, stride, data, width, height, &_sourcePicture}
}

func (p *SourcePicture) updateBase() {
  p.base.iColorFormat = C.int(p.ColorFormat)
  
  for index, val := range p.Stride {
    p.base.iStride[index] = C.int(val)
  }
  
  for index, val := range p.Data {
    p.base.pData[index] = (*C.uchar)(unsafe.Pointer(val))
  }
  
  p.base.iPicWidth = C.int(p.SourceWidth)
  p.base.iPicHeight = C.int(p.SourceHeight)
}

func (p *SourcePicture) updateStruct() {
  p.ColorFormat = int(p.base.iColorFormat)
  
  for index, val := range p.base.iStride {
    p.Stride[index] = int(val)
  }
  
  for index, val := range p.base.pData {
    p.Data[index] = (*uint8)(unsafe.Pointer(val))
  }
  
  p.SourceWidth = int(p.base.iPicWidth)
  p.SourceHeight = int(p.base.iPicHeight)
}

type SourcePicture struct {
	ColorFormat  int      //iColorFormat
	Stride       [4]int   //iStride
	Data         [4]*uint8 //pData
	SourceWidth  int      //iPicWidth
	SourceHeight int      //iPicHeight

	base *C.SSourcePicture
}

/*
typedef struct Source_Picture_s {
  int       iColorFormat; // color space type
  int     iStride[4];   // stride for each plane pData
  unsigned char*  pData[4];   // plane pData
  int     iPicWidth;        // luma picture width in x coordinate
  int     iPicHeight;       // luma picture height in y coordinate
} SSourcePicture;
*/

////////////////////
// codec_def.h
////////////////////

// EVideoFormatType
type EVideoFormatType int32

const (
	VideoFormatRGB      = 0x000000001
	VideoFormatRGBA     = 0x000000002
	VideoFormatRGB555   = 0x000000003
	VideoFormatRGB565   = 0x000000004
	VideoFormatBGR      = 0x000000005
	VideoFormatBGRA     = 0x000000006
	VideoFormatABGR     = 0x000000007
	VideoFormatARGB     = 0x000000008
	VideoFormatYUY2     = 0x000000014
	VideoFormatYVYU     = 0x000000015
	VideoFormatUYVY     = 0x000000016
	VideoFormatI420     = 0x000000017
	VideoFormatYV12     = 0x000000018
	VideoFormatInternal = 0x000000019
	VideoFormatNV12     = 0x00000001A
	VideoFormatVFlip    = -2147483648
)

// EVideoFrameType
type EVideoFrameType int32

const (
	VideoFrameTypeInvalid = 0x0
	VideoFrameTypeIDR     = 0x1
	VideoFrameTypeI       = 0x2
	VideoFrameTypeP       = 0x3
	VideoFrameTypeSkip    = 0x4
	VideoFrameTypeIPMixed = 0x5
)

// CM_RETURN
type CmReturn int32

const (
	CmResultSuccess   = 0x0
	CmInitParaError   = 0x1
	CmMachPerfIsBad   = 0x2
	CmUnkonwReason    = 0x3
	CmMallocMemeError = 0x4
	CmInitExpected    = 0x5
)

// ENalUnitType
type ENalUnitType int32

const (
	NalUnknown  = 0x0
	NalSlice    = 0x1
	NalSliceDpa = 0x2
	NalSliceDpb = 0x3
	NalSliceDpc = 0x4
	NalSliceIdr = 0x5
	NalSei      = 0x6
	NalSps      = 0x7
	NalPps      = 0x8
)

// ENalPriority
type ENalPriority int32

const (
	NalPriorityDisposable = 0x0
	NalPriorityLow        = 0x1
	NalPriorityHigh       = 0x2
	NalPriorityHighest    = 0x3
)

// ERR_TOOL
type ErrTool uint8

const (
	EtNone    = 0x00
	EtIpScale = 0x01
	EtFmo     = 0x02
	EtIrR1    = 0x04
	EtIrR2    = 0x08
	EtIrR3    = 0x10
	EtFecHalf = 0x20
	EtFecFull = 0x40
	EtRfs     = 0x80
)

func isParameterSetNal(eNalRefIdc ENalPriority, eNalType ENalUnitType) bool {
	return (eNalRefIdc == NalPriorityHighest) && (eNalType == (NalSps|NalPps) || eNalType == NalSps)
}

func isIdrNal(eNalRefIdc ENalPriority, eNalType ENalUnitType) bool {
	return (eNalRefIdc == NalPriorityHighest) && (eNalType == NalSliceIdr)
}

const (
	FRAME_NUM_PARAM = -1
	FRAME_NUM_IDR   = 0
)

func NewSliceInfo() *SliceInfo {
	var bufOfSlices *uint8
	var sliceCount, fec int
	var idx, count, idc, nType, cFN uint8
	var fIndex int8
	var length *uint
	var _sliceInfo C.SliceInfo

	return &SliceInfo{bufOfSlices, sliceCount, length, fec, idx, count, fIndex, idc, nType, cFN, &_sliceInfo}
}

func (p *SliceInfo) updateBase() {
  p.base.pBufferOfSlices = (*C.uchar)(p.BufferOfSlices)
  p.base.iCodedSliceCount = C.int(p.CodedSliceCount)
  p.base.pLengthOfSlices = (*C.uint)(unsafe.Pointer(p.LengthOfSlices))
  p.base.iFecType = C.int(p.FecType)
  p.base.uiSliceIdx = C.uchar(p.SliceIdx)
  p.base.uiSliceCount = C.uchar(p.SliceCount)
  p.base.iFrameIndex = C.char(p.FrameIndex)
  p.base.uiNalRefIdc = C.uchar(p.NalRefIdc)
  p.base.uiNalType = C.uchar(p.NalType)
  p.base.uiContainingFinalNal = C.uchar(p.ContainingFinalNal)
}

func (p *SliceInfo) updateStruct() {
  p.BufferOfSlices = (*uint8)(p.base.pBufferOfSlices)
  p.CodedSliceCount = int(p.base.iCodedSliceCount)
  p.LengthOfSlices = (*uint)(unsafe.Pointer(p.base.pLengthOfSlices))
  p.FecType = int(p.base.iFecType)
  p.SliceIdx = uint8(p.base.uiSliceIdx)
  p.SliceCount = uint8(p.base.uiSliceCount)
  p.FrameIndex = int8(p.base.iFrameIndex)
  p.NalRefIdc = uint8(p.base.uiNalRefIdc)
  p.NalType = uint8(p.base.uiNalType)
  p.ContainingFinalNal = uint8(p.base.uiContainingFinalNal)
}

type SliceInfo struct {
	BufferOfSlices      *uint8 //pBufferOfSlices
	CodedSliceCount     int    //iCodedSliceCount
	LengthOfSlices      *uint //pLengthOfSlices
	FecType             int    //iFecType
	SliceIdx            uint8  //uiSliceIdx
	SliceCount          uint8  //uiSliceCount
	FrameIndex          int8   //iFrameIndex
	NalRefIdc           uint8  //uiNalRefIdc
	NalType             uint8  //uiNalType
	ContainingFinalNal uint8  //uiContainingFinalNal

	base *C.SliceInfo
}

/*
typedef struct SliceInformation {
  unsigned char*  pBufferOfSlices;    // base buffer of coded slice(s)
  int       iCodedSliceCount; // number of coded slices
  unsigned int* pLengthOfSlices;    // array of slices length accordingly by number of slice
  int       iFecType;     // FEC type[0, 50%FEC, 100%FEC]
  unsigned char uiSliceIdx;   // index of slice in frame [FMO: 0,..,uiSliceCount-1; No FMO: 0]
  unsigned char uiSliceCount;   // count number of slice in frame [FMO: 2-8; No FMO: 1]
  char      iFrameIndex;    // index of frame[-1, .., idr_interval-1]
  unsigned char uiNalRefIdc;    // NRI, priority level of slice(NAL)
  unsigned char uiNalType;      // NAL type
  unsigned char
  uiContainingFinalNal; // whether final NAL is involved in buffer of coded slices, flag used in Pause feature in T27
} SliceInfo, *PSliceInfo;
*/

const (
	CIF_WIDTH    = 352
	CIF_HEIGHT   = 288
	QVGA_WIDTH   = 320
	QVGA_HEIGHT  = 240
	QCIF_WIDTH   = 176
	QCIF_HEIGHT  = 144
	SQCIF_WIDTH  = 128
	SQCIF_HEIGHT = 96
)

func NewRateThresholds() *RateThresholds {
	var width, height, initRate, maxRate, minRate, frameRate, skipFrameRate, skipFrameStep int
	var _rateThresholds C.SRateThresholds
	return &RateThresholds{width, height, initRate, maxRate, minRate, frameRate, skipFrameRate, skipFrameStep, &_rateThresholds}
}

func newRateThresholds(width, height, initRate, maxRate, minRate, frameRate, skipFrameRate, skipFrameStep int) *RateThresholds {
	var _rateThresholds C.SRateThresholds
	p := &RateThresholds{width, height, initRate, maxRate, minRate, frameRate, skipFrameRate, skipFrameStep, &_rateThresholds}
	p.updateBase()
	return p
}

func (p *RateThresholds) updateBase() {
	p.base.iWidth = C.int(p.Width)
	p.base.iHeight = C.int(p.Height)
	p.base.iThresholdOfInitRate = C.int(p.ThresholdOfInitRate)
	p.base.iThresholdOfMaxRate = C.int(p.ThresholdOfMaxRate)
	p.base.iThresholdOfMinRate = C.int(p.ThresholdOfMinRate)
	p.base.iMinThresholdFrameRate = C.int(p.MinThresholdFrameRate)
	p.base.iSkipFrameRate = C.int(p.SkipFrameRate)
	p.base.iSkipFrameStep = C.int(p.SkipFrameStep)
}

func (p *RateThresholds) updateStruct() {
	p.Width = int(p.base.iWidth)
	p.Height = int(p.base.iHeight)
	p.ThresholdOfInitRate = int(p.base.iThresholdOfInitRate)
	p.ThresholdOfMaxRate = int(p.base.iThresholdOfMaxRate)
	p.ThresholdOfMinRate = int(p.base.iThresholdOfMinRate)
	p.MinThresholdFrameRate = int(p.base.iMinThresholdFrameRate)
	p.SkipFrameRate = int(p.base.iSkipFrameRate)
	p.SkipFrameStep = int(p.base.iSkipFrameStep)
}

type RateThresholds struct {
	Width                 int
	Height                int
	ThresholdOfInitRate   int
	ThresholdOfMaxRate    int
	ThresholdOfMinRate    int
	MinThresholdFrameRate int
	SkipFrameRate         int
	SkipFrameStep         int

	base *C.SRateThresholds
}

/*
typedef struct {
  int iWidth;     // frame width
  int iHeight;      // frame height
  int iThresholdOfInitRate; // threshold of initial rate
  int iThresholdOfMaxRate;  // threshold of maximal rate
  int iThresholdOfMinRate;  // threshold of minimal rate
  int iMinThresholdFrameRate;   //min frame rate min
  int iSkipFrameRate; //skip to frame rate min
  int iSkipFrameStep; //how many frames to skip
} SRateThresholds, *PRateThresholds;
*/

func NewSysMemBuffer() *SysMemBuffer {
	var width, height, format int
	var stride [2]int
	var _sysMemBuffer C.SSysMEMBuffer
	return &SysMemBuffer{width, height, format, stride, &_sysMemBuffer}
}

func (p *SysMemBuffer) updateBase() {
	p.base.iWidth = C.int(p.Width)
	p.base.iHeight = C.int(p.Height)
	p.base.iFormat = C.int(p.Format)
	p.base.iStride[0] = C.int(p.Stride[0])
	p.base.iStride[1] = C.int(p.Stride[1])
}

func (p *SysMemBuffer) updateStruct() {
	p.Width = int(p.base.iWidth)
	p.Height = int(p.base.iHeight)
	p.Format = int(p.base.iFormat)
	p.Stride[0] = int(p.base.iStride[0])
	p.Stride[1] = int(p.base.iStride[1])
}

type SysMemBuffer struct {
	Width  int    //iWidth
	Height int    //iHeight
	Format int    //iFormat
	Stride [2]int //iStride

	base *C.SSysMEMBuffer
}

/*
typedef struct TagSysMemBuffer {
  int iWidth;     //width of decoded pic for display
  int iHeight;      //height of decoded pic for display
  int iFormat;    // type is "EVideoFormatType"
  int iStride[2];   //stride of 2 component
} SSysMEMBuffer;
*/

func NewBufferInfo() *BufferInfo {
	var bStatus int
	var _bufferInfo C.SBufferInfo
	var _sysMemBuffer = NewSysMemBuffer()

  _sysMemBuffer.base = C.get_buffer_info_union_pointer(&_bufferInfo)

	return &BufferInfo{bStatus, _sysMemBuffer, &_bufferInfo}
}

func (p *BufferInfo) updateBase() {
	p.base.iBufferStatus = C.int(p.BufferStatus)
	p.SystemBuffer.updateBase()
}

func (p *BufferInfo) updateStruct() {
	p.BufferStatus = int(p.base.iBufferStatus)
	
/*  C.get_buffer_info_union_pointer(p.base)*/
  
	p.SystemBuffer.updateStruct()
}

type BufferInfo struct {
	BufferStatus int           //iBufferStatus
	SystemBuffer *SysMemBuffer //sSystemBuffer

	base *C.SBufferInfo
}

/*
typedef struct TagBufferInfo {
  int iBufferStatus;  // 0: one frame data is not ready; 1: one frame data is ready
  union {
    SSysMEMBuffer sSystemBuffer;
  } UsrData;
} SBufferInfo;
*/

// ksRateThrMap
var KsRateThrMap = [4]*RateThresholds{
	newRateThresholds(CIF_WIDTH, CIF_HEIGHT, 225000, 384000, 96000, 3, 1, 1),      //CIF
	newRateThresholds(QVGA_WIDTH, QVGA_HEIGHT, 192000, 320000, 80000, -1, -1, -1), // QVGA
	newRateThresholds(QCIF_WIDTH, QCIF_HEIGHT, 150000, 256000, 64000, 8, 4, 2),    // QCIF
	newRateThresholds(SQCIF_WIDTH, SQCIF_HEIGHT, 120000, 192000, 48000, 5, 3, 1),  // SQCIF
}

// kiKeyNumMultiple
var KiKeyNumMultiple = []byte{1, 1, 2, 4, 8, 16}
