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

package huffman ;import (_g "errors";_c "fmt";_aa "github.com/unidoc/unipdf/v3/internal/bitwise";_f "math";_b "strings";);func (_bfg *StandardTable )InitTree (codeTable []*Code )error {_cbga (codeTable );for _ ,_bcb :=range codeTable {if _aaa :=_bfg ._bc .append (_bcb );_aaa !=nil {return _aaa ;};};return nil ;};type Node interface{Decode (_eea _aa .StreamReader )(int64 ,error );String ()string ;};func (_ef *EncodedTable )InitTree (codeTable []*Code )error {_cbga (codeTable );for _ ,_d :=range codeTable {if _cc :=_ef ._e .append (_d );_cc !=nil {return _cc ;};};return nil ;};func (_de *InternalNode )String ()string {_efc :=&_b .Builder {};_efc .WriteString ("\u000a");_de .pad (_efc );_efc .WriteString ("\u0030\u003a\u0020");_efc .WriteString (_de ._cabd .String ()+"\u000a");_de .pad (_efc );_efc .WriteString ("\u0031\u003a\u0020");_efc .WriteString (_de ._ge .String ()+"\u000a");return _efc .String ();};type OutOfBandNode struct{};type BasicTabler interface{HtHigh ()int32 ;HtLow ()int32 ;StreamReader ()_aa .StreamReader ;HtPS ()int32 ;HtRS ()int32 ;HtOOB ()int32 ;};func NewFixedSizeTable (codeTable []*Code )(*FixedSizeTable ,error ){_bgg :=&FixedSizeTable {_bg :&InternalNode {}};if _ca :=_bgg .InitTree (codeTable );_ca !=nil {return nil ,_ca ;};return _bgg ,nil ;};func (_gb *EncodedTable )String ()string {return _gb ._e .String ()+"\u000a"};func _fef (_fab ,_aag int32 )string {var _efd int32 ;_ea :=make ([]rune ,_aag );for _fgc :=int32 (1);_fgc <=_aag ;_fgc ++{_efd =_fab >>uint (_aag -_fgc )&1;if _efd !=0{_ea [_fgc -1]='1';}else {_ea [_fgc -1]='0';};};return string (_ea );};func (_cf *EncodedTable )Decode (r _aa .StreamReader )(int64 ,error ){return _cf ._e .Decode (r )};func (_ga *Code )String ()string {var _ffb string ;if _ga ._cae !=-1{_ffb =_fef (_ga ._cae ,_ga ._bb );}else {_ffb ="\u003f";};return _c .Sprintf ("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_ffb ,_ga ._bb ,_ga ._afef ,_ga ._bdf );};func (_cbd *StandardTable )RootNode ()*InternalNode {return _cbd ._bc };func NewCode (prefixLength ,rangeLength ,rangeLow int32 ,isLowerRange bool )*Code {return &Code {_bb :prefixLength ,_afef :rangeLength ,_bdf :rangeLow ,_ceb :isLowerRange ,_cae :-1};};var _ Tabler =&EncodedTable {};func (_fdgd *ValueNode )String ()string {return _c .Sprintf ("\u0025\u0064\u002f%\u0064",_fdgd ._be ,_fdgd ._ae );};type EncodedTable struct{BasicTabler ;_e *InternalNode ;};func _cbga (_afa []*Code ){var _gfg int32 ;for _ ,_fce :=range _afa {_gfg =_fgca (_gfg ,_fce ._bb );};_cd :=make ([]int32 ,_gfg +1);for _ ,_bga :=range _afa {_cd [_bga ._bb ]++;};var _dfc int32 ;_gfa :=make ([]int32 ,len (_cd )+1);_cd [0]=0;for _db :=int32 (1);_db <=int32 (len (_cd ));_db ++{_gfa [_db ]=(_gfa [_db -1]+(_cd [_db -1]))<<1;_dfc =_gfa [_db ];for _ ,_bfa :=range _afa {if _bfa ._bb ==_db {_bfa ._cae =_dfc ;_dfc ++;};};};};func (_fdg *ValueNode )Decode (r _aa .StreamReader )(int64 ,error ){_afb ,_eed :=r .ReadBits (byte (_fdg ._be ));if _eed !=nil {return 0,_eed ;};if _fdg ._bec {_afb =-_afb ;};return int64 (_fdg ._ae )+int64 (_afb ),nil ;};func (_df *InternalNode )Decode (r _aa .StreamReader )(int64 ,error ){_fc ,_cfb :=r .ReadBit ();if _cfb !=nil {return 0,_cfb ;};if _fc ==0{return _df ._cabd .Decode (r );};return _df ._ge .Decode (r );};type Tabler interface{Decode (_afd _aa .StreamReader )(int64 ,error );InitTree (_fdd []*Code )error ;String ()string ;RootNode ()*InternalNode ;};func (_ffg *InternalNode )pad (_ffe *_b .Builder ){for _eb :=int32 (0);_eb < _ffg ._fbb ;_eb ++{_ffe .WriteString ("\u0020\u0020\u0020");};};var _ Node =&OutOfBandNode {};var _ada =make ([]Tabler ,len (_cbdd ));func (_dg *FixedSizeTable )InitTree (codeTable []*Code )error {_cbga (codeTable );for _ ,_bf :=range codeTable {_dgg :=_dg ._bg .append (_bf );if _dgg !=nil {return _dgg ;};};return nil ;};var _ Node =&InternalNode {};type InternalNode struct{_fbb int32 ;_cabd Node ;_ge Node ;};func _fgca (_aagf ,_ed int32 )int32 {if _aagf > _ed {return _aagf ;};return _ed ;};var _cbdd =[][][]int32 {{{1,4,0},{2,8,16},{3,16,272},{3,32,65808}},{{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{6,32,75},{6,-1,0}},{{8,8,-256},{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{8,32,-257,999},{7,32,75},{6,-1,0}},{{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{5,32,76}},{{7,8,-255},{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{7,32,-256,999},{6,32,76}},{{5,10,-2048},{4,9,-1024},{4,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{2,7,0},{3,7,128},{3,8,256},{4,9,512},{4,10,1024},{6,32,-2049,999},{6,32,2048}},{{4,9,-1024},{3,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{4,5,0},{5,5,32},{5,6,64},{4,7,128},{3,8,256},{3,9,512},{3,10,1024},{5,32,-1025,999},{5,32,2048}},{{8,3,-15},{9,1,-7},{8,1,-5},{9,0,-3},{7,0,-2},{4,0,-1},{2,1,0},{5,0,2},{6,0,3},{3,4,4},{6,1,20},{4,4,22},{4,5,38},{5,6,70},{5,7,134},{6,7,262},{7,8,390},{6,10,646},{9,32,-16,999},{9,32,1670},{2,-1,0}},{{8,4,-31},{9,2,-15},{8,2,-11},{9,1,-7},{7,1,-5},{4,1,-3},{3,1,-1},{3,1,1},{5,1,3},{6,1,5},{3,5,7},{6,2,39},{4,5,43},{4,6,75},{5,7,139},{5,8,267},{6,8,523},{7,9,779},{6,11,1291},{9,32,-32,999},{9,32,3339},{2,-1,0}},{{7,4,-21},{8,0,-5},{7,0,-4},{5,0,-3},{2,2,-2},{5,0,2},{6,0,3},{7,0,4},{8,0,5},{2,6,6},{5,5,70},{6,5,102},{6,6,134},{6,7,198},{6,8,326},{6,9,582},{6,10,1094},{7,11,2118},{8,32,-22,999},{8,32,4166},{2,-1,0}},{{1,0,1},{2,1,2},{4,0,4},{4,1,5},{5,1,7},{5,2,9},{6,2,13},{7,2,17},{7,3,21},{7,4,29},{7,5,45},{7,6,77},{7,32,141}},{{1,0,1},{2,0,2},{3,1,3},{5,0,5},{5,1,6},{6,1,8},{7,0,10},{7,1,11},{7,2,13},{7,3,17},{7,4,25},{8,5,41},{8,32,73}},{{1,0,1},{3,0,2},{4,0,3},{5,0,4},{4,1,5},{3,3,7},{6,1,15},{6,2,17},{6,3,21},{6,4,29},{6,5,45},{7,6,77},{7,32,141}},{{3,0,-2},{3,0,-1},{1,0,0},{3,0,1},{3,0,2}},{{7,4,-24},{6,2,-8},{5,1,-4},{4,0,-2},{3,0,-1},{1,0,0},{3,0,1},{4,0,2},{5,1,3},{6,2,5},{7,4,9},{7,32,-25,999},{7,32,25}}};func NewEncodedTable (table BasicTabler )(*EncodedTable ,error ){_ag :=&EncodedTable {_e :&InternalNode {},BasicTabler :table };if _ff :=_ag .parseTable ();_ff !=nil {return nil ,_ff ;};return _ag ,nil ;};func GetStandardTable (number int )(Tabler ,error ){if number <=0||number > len (_ada ){return nil ,_g .New ("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_ece :=_ada [number -1];if _ece ==nil {var _ad error ;_ece ,_ad =_cbg (_cbdd [number -1]);if _ad !=nil {return nil ,_ad ;};_ada [number -1]=_ece ;};return _ece ,nil ;};func (_ac *OutOfBandNode )String ()string {return _c .Sprintf ("\u0025\u0030\u00364\u0062",int64 (_f .MaxInt64 ));};func (_cab *FixedSizeTable )Decode (r _aa .StreamReader )(int64 ,error ){return _cab ._bg .Decode (r )};func (_cea *StandardTable )String ()string {return _cea ._bc .String ()+"\u000a"};type Code struct{_bb int32 ;_afef int32 ;_bdf int32 ;_ceb bool ;_cae int32 ;};func _bab (_ce *Code )*OutOfBandNode {return &OutOfBandNode {}};func (_ba *FixedSizeTable )RootNode ()*InternalNode {return _ba ._bg };func _egf (_afe int32 )*InternalNode {return &InternalNode {_fbb :_afe }};type FixedSizeTable struct{_bg *InternalNode };func (_fe *EncodedTable )RootNode ()*InternalNode {return _fe ._e };func (_dgd *InternalNode )append (_fg *Code )(_cag error ){if _fg ._bb ==0{return nil ;};_eg :=_fg ._bb -1-_dgd ._fbb ;if _eg < 0{return _g .New ("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064");};_ccd :=(_fg ._cae >>uint (_eg ))&0x1;if _eg ==0{if _fg ._afef ==-1{if _ccd ==1{if _dgd ._ge !=nil {return _c .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_fg );};_dgd ._ge =_bab (_fg );}else {if _dgd ._cabd !=nil {return _c .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_fg );};_dgd ._cabd =_bab (_fg );};}else {if _ccd ==1{if _dgd ._ge !=nil {return _c .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_fg );};_dgd ._ge =_dc (_fg );}else {if _dgd ._cabd !=nil {return _c .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_fg );};_dgd ._cabd =_dc (_fg );};};}else {if _ccd ==1{if _dgd ._ge ==nil {_dgd ._ge =_egf (_dgd ._fbb +1);};if _cag =_dgd ._ge .(*InternalNode ).append (_fg );_cag !=nil {return _cag ;};}else {if _dgd ._cabd ==nil {_dgd ._cabd =_egf (_dgd ._fbb +1);};if _cag =_dgd ._cabd .(*InternalNode ).append (_fg );_cag !=nil {return _cag ;};};};return nil ;};var _ Node =&ValueNode {};func _cbg (_gbd [][]int32 )(*StandardTable ,error ){var _eee []*Code ;for _gd :=0;_gd < len (_gbd );_gd ++{_fbbf :=_gbd [_gd ][0];_fgb :=_gbd [_gd ][1];_cbbd :=_gbd [_gd ][2];var _efcb bool ;if len (_gbd [_gd ])> 3{_efcb =true ;};_eee =append (_eee ,NewCode (_fbbf ,_fgb ,_cbbd ,_efcb ));};_dd :=&StandardTable {_bc :_egf (0)};if _gc :=_dd .InitTree (_eee );_gc !=nil {return nil ,_gc ;};return _dd ,nil ;};func _dc (_cbb *Code )*ValueNode {return &ValueNode {_be :_cbb ._afef ,_ae :_cbb ._bdf ,_bec :_cbb ._ceb }};func (_caf *FixedSizeTable )String ()string {return _caf ._bg .String ()+"\u000a"};func (_gbf *OutOfBandNode )Decode (r _aa .StreamReader )(int64 ,error ){return int64 (_f .MaxInt64 ),nil };type StandardTable struct{_bc *InternalNode };func (_ec *StandardTable )Decode (r _aa .StreamReader )(int64 ,error ){return _ec ._bc .Decode (r )};type ValueNode struct{_be int32 ;_ae int32 ;_bec bool ;};func (_cff *EncodedTable )parseTable ()error {var (_af []*Code ;_gg ,_feg ,_fa int32 ;_efg uint64 ;_fba error ;);_agc :=_cff .StreamReader ();_ee :=_cff .HtLow ();for _ee < _cff .HtHigh (){_efg ,_fba =_agc .ReadBits (byte (_cff .HtPS ()));if _fba !=nil {return _fba ;};_gg =int32 (_efg );_efg ,_fba =_agc .ReadBits (byte (_cff .HtRS ()));if _fba !=nil {return _fba ;};_feg =int32 (_efg );_af =append (_af ,NewCode (_gg ,_feg ,_fa ,false ));_ee +=1<<uint (_feg );};_efg ,_fba =_agc .ReadBits (byte (_cff .HtPS ()));if _fba !=nil {return _fba ;};_gg =int32 (_efg );_feg =32;_fa =_cff .HtLow ()-1;_af =append (_af ,NewCode (_gg ,_feg ,_fa ,true ));_efg ,_fba =_agc .ReadBits (byte (_cff .HtPS ()));if _fba !=nil {return _fba ;};_gg =int32 (_efg );_feg =32;_fa =_cff .HtHigh ();_af =append (_af ,NewCode (_gg ,_feg ,_fa ,false ));if _cff .HtOOB ()==1{_efg ,_fba =_agc .ReadBits (byte (_cff .HtPS ()));if _fba !=nil {return _fba ;};_gg =int32 (_efg );_af =append (_af ,NewCode (_gg ,-1,-1,false ));};if _fba =_cff .InitTree (_af );_fba !=nil {return _fba ;};return nil ;};