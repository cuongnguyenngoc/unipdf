//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package sampling ;import (_b "github.com/unidoc/unipdf/v3/internal/bitwise";_eb "github.com/unidoc/unipdf/v3/internal/imageutil";_f "io";);func (_efg *Writer )WriteSample (sample uint32 )error {if _ ,_bea :=_efg ._ff .WriteBits (uint64 (sample ),_efg ._db .BitsPerComponent );_bea !=nil {return _bea ;};_efg ._dae --;if _efg ._dae ==0{_efg ._dae =_efg ._db .ColorComponents ;_efg ._fa ++;};if _efg ._fa ==_efg ._db .Width {if _efg ._beg {_efg ._ff .FinishByte ();};_efg ._fa =0;};return nil ;};func (_fe *Reader )ReadSample ()(uint32 ,error ){if _fe ._a ==_fe ._g .Height {return 0,_f .EOF ;};_eg ,_gg :=_fe ._d .ReadBits (byte (_fe ._g .BitsPerComponent ));if _gg !=nil {return 0,_gg ;};_fe ._ea --;if _fe ._ea ==0{_fe ._ea =_fe ._g .ColorComponents ;_fe ._gf ++;};if _fe ._gf ==_fe ._g .Width {if _fe ._ed {_fe ._d .ConsumeRemainingBits ();};_fe ._gf =0;_fe ._a ++;};return uint32 (_eg ),nil ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _dd []uint32 ;_ga :=bitsPerSample ;var _da uint32 ;var _fg byte ;_bc :=0;_fc :=0;_egf :=0;for _egf < len (data ){if _bc > 0{_bb :=_bc ;if _ga < _bb {_bb =_ga ;};_da =(_da <<uint (_bb ))|uint32 (_fg >>uint (8-_bb ));_bc -=_bb ;if _bc > 0{_fg =_fg <<uint (_bb );}else {_fg =0;};_ga -=_bb ;if _ga ==0{_dd =append (_dd ,_da );_ga =bitsPerSample ;_da =0;_fc ++;};}else {_ebb :=data [_egf ];_egf ++;_bf :=8;if _ga < _bf {_bf =_ga ;};_bc =8-_bf ;_da =(_da <<uint (_bf ))|uint32 (_ebb >>uint (_bc ));if _bf < 8{_fg =_ebb <<uint (_bf );};_ga -=_bf ;if _ga ==0{_dd =append (_dd ,_da );_ga =bitsPerSample ;_da =0;_fc ++;};};};for _bc >=bitsPerSample {_fed :=_bc ;if _ga < _fed {_fed =_ga ;};_da =(_da <<uint (_fed ))|uint32 (_fg >>uint (8-_fed ));_bc -=_fed ;if _bc > 0{_fg =_fg <<uint (_fed );}else {_fg =0;};_ga -=_fed ;if _ga ==0{_dd =append (_dd ,_da );_ga =bitsPerSample ;_da =0;_fc ++;};};return _dd ;};type Reader struct{_g _eb .ImageBase ;_d *_b .Reader ;_gf ,_a ,_ea int ;_ed bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _gd []uint32 ;_aa :=bitsPerOutputSample ;var _edd uint32 ;var _ggf uint32 ;_ag :=0;_gdb :=0;_ge :=0;for _ge < len (data ){if _ag > 0{_ggc :=_ag ;if _aa < _ggc {_ggc =_aa ;};_edd =(_edd <<uint (_ggc ))|uint32 (_ggf >>uint (bitsPerInputSample -_ggc ));_ag -=_ggc ;if _ag > 0{_ggf =_ggf <<uint (_ggc );}else {_ggf =0;};_aa -=_ggc ;if _aa ==0{_gd =append (_gd ,_edd );_aa =bitsPerOutputSample ;_edd =0;_gdb ++;};}else {_be :=data [_ge ];_ge ++;_df :=bitsPerInputSample ;if _aa < _df {_df =_aa ;};_ag =bitsPerInputSample -_df ;_edd =(_edd <<uint (_df ))|uint32 (_be >>uint (_ag ));if _df < bitsPerInputSample {_ggf =_be <<uint (_df );};_aa -=_df ;if _aa ==0{_gd =append (_gd ,_edd );_aa =bitsPerOutputSample ;_edd =0;_gdb ++;};};};for _ag >=bitsPerOutputSample {_de :=_ag ;if _aa < _de {_de =_aa ;};_edd =(_edd <<uint (_de ))|uint32 (_ggf >>uint (bitsPerInputSample -_de ));_ag -=_de ;if _ag > 0{_ggf =_ggf <<uint (_de );}else {_ggf =0;};_aa -=_de ;if _aa ==0{_gd =append (_gd ,_edd );_aa =bitsPerOutputSample ;_edd =0;_gdb ++;};};if _aa > 0&&_aa < bitsPerOutputSample {_edd <<=uint (_aa );_gd =append (_gd ,_edd );};return _gd ;};func NewReader (img _eb .ImageBase )*Reader {return &Reader {_d :_b .NewReader (img .Data ),_g :img ,_ea :img .ColorComponents ,_ed :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func (_beaf *Writer )WriteSamples (samples []uint32 )error {for _bbd :=0;_bbd < len (samples );_bbd ++{if _gb :=_beaf .WriteSample (samples [_bbd ]);_gb !=nil {return _gb ;};};return nil ;};func NewWriter (img _eb .ImageBase )*Writer {return &Writer {_ff :_b .NewWriterMSB (img .Data ),_db :img ,_dae :img .ColorComponents ,_beg :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_ec []uint32 )error ;};type SampleWriter interface{WriteSample (_bd uint32 )error ;WriteSamples (_gda []uint32 )error ;};type Writer struct{_db _eb .ImageBase ;_ff *_b .Writer ;_fa ,_dae int ;_beg bool ;};func (_ab *Reader )ReadSamples (samples []uint32 )(_ef error ){for _ebd :=0;_ebd < len (samples );_ebd ++{samples [_ebd ],_ef =_ab .ReadSample ();if _ef !=nil {return _ef ;};};return nil ;};