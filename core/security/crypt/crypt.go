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

package crypt ;import (_a "crypto/aes";_g "crypto/cipher";_c "crypto/md5";_ea "crypto/rand";_bac "crypto/rc4";_ba "fmt";_ed "github.com/unidoc/unipdf/v3/common";_ac "github.com/unidoc/unipdf/v3/core/security";_e "io";);func init (){_abd ("\u0041\u0045\u0053V\u0032",_gc )};

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2 (length int )Filter {_ff ,_gg :=_gce (FilterDict {Length :length });if _gg !=nil {_ed .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_gg );return filterV2 {_da :length };};return _ff ;};

// MakeKey implements Filter interface.
func (filterAESV3 )MakeKey (_ ,_ uint32 ,ekey []byte )([]byte ,error ){return ekey ,nil };

// Name implements Filter interface.
func (filterV2 )Name ()string {return "\u0056\u0032"};type filterAES struct{};

// HandlerVersion implements Filter interface.
func (filterAESV3 )HandlerVersion ()(V ,R int ){V ,R =5,6;return ;};func _ab (_d FilterDict )(Filter ,error ){if _d .Length ==256{_ed .Log .Debug ("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_d .Length );_d .Length /=8;};if _d .Length !=0&&_d .Length !=32{return nil ,_ba .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_d .Length );};return filterAESV3 {},nil ;};var (_acb =make (map[string ]filterFunc ););

// MakeKey implements Filter interface.
func (filterAESV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _ebc (objNum ,genNum ,ekey ,true );};

// KeyLength implements Filter interface.
func (filterAESV3 )KeyLength ()int {return 256/8};func init (){_abd ("\u0041\u0045\u0053V\u0033",_ab )};func _aadf (_cad string )(filterFunc ,error ){_ceg :=_acb [_cad ];if _ceg ==nil {return nil ,_ba .Errorf ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071",_cad );};return _ceg ,nil ;};

// DecryptBytes implements Filter interface.
func (filterV2 )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_gfdd ,_eab :=_bac .NewCipher (okey );if _eab !=nil {return nil ,_eab ;};_ed .Log .Trace ("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_gfdd .XORKeyStream (buf ,buf );_ed .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};func _gc (_gb FilterDict )(Filter ,error ){if _gb .Length ==128{_ed .Log .Debug ("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_gb .Length );_gb .Length /=8;};if _gb .Length !=0&&_gb .Length !=16{return nil ,_ba .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_gb .Length );};return filterAESV2 {},nil ;};

// EncryptBytes implements Filter interface.
func (filterV2 )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_ddg ,_gfd :=_bac .NewCipher (okey );if _gfd !=nil {return nil ,_gfd ;};_ed .Log .Trace ("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_ddg .XORKeyStream (buf ,buf );_ed .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};func (filterIdentity )EncryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };

// PDFVersion implements Filter interface.
func (filterAESV3 )PDFVersion ()[2]int {return [2]int {2,0}};func (filterAES )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_ga ,_ag :=_a .NewCipher (okey );if _ag !=nil {return nil ,_ag ;};if len (buf )< 16{_ed .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073",buf );return buf ,_ba .Errorf ("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029",len (buf ));};_edc :=buf [:16];buf =buf [16:];if len (buf )%16!=0{_ed .Log .Debug ("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (_edc ),_edc );_ed .Log .Debug ("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,_ba .Errorf ("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029",len (buf ));};_gba :=_g .NewCBCDecrypter (_ga ,_edc );_ed .Log .Trace ("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );_ed .Log .Trace ("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078",len (buf ),buf );_gba .CryptBlocks (buf ,buf );_ed .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );if len (buf )==0{_ed .Log .Trace ("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067");return buf ,nil ;};_ec :=int (buf [len (buf )-1]);if _ec > len (buf ){_ed .Log .Debug ("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029",_ec ,len (buf ));return buf ,_ba .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068");};buf =buf [:len (buf )-_ec ];return buf ,nil ;};

// PDFVersion implements Filter interface.
func (filterAESV2 )PDFVersion ()[2]int {return [2]int {1,5}};type filterAESV3 struct{filterAES };

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3 ()Filter {_be ,_gcf :=_ab (FilterDict {});if _gcf !=nil {_ed .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_gcf );return filterAESV3 {};};return _be ;};

// MakeKey implements Filter interface.
func (_cge filterV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _ebc (objNum ,genNum ,ekey ,false );};

// KeyLength implements Filter interface.
func (_ccf filterV2 )KeyLength ()int {return _ccf ._da };func init (){_abd ("\u0056\u0032",_gce )};

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity ()Filter {return filterIdentity {}};var _ Filter =filterV2 {};func _gce (_bcf FilterDict )(Filter ,error ){if _bcf .Length %8!=0{return nil ,_ba .Errorf ("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029",_bcf .Length );};if _bcf .Length < 5||_bcf .Length > 16{if _bcf .Length ==40||_bcf .Length ==64||_bcf .Length ==128{_ed .Log .Debug ("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029",_bcf .Length );_bcf .Length /=8;}else {return nil ,_ba .Errorf ("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029",_bcf .Length );};};return filterV2 {_da :_bcf .Length },nil ;};func (filterIdentity )DecryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };

// Name implements Filter interface.
func (filterAESV2 )Name ()string {return "\u0041\u0045\u0053V\u0032"};

// Name implements Filter interface.
func (filterAESV3 )Name ()string {return "\u0041\u0045\u0053V\u0033"};func _abd (_dbd string ,_aac filterFunc ){if _ ,_age :=_acb [_dbd ];_age {panic ("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064");};_acb [_dbd ]=_aac ;};type filterAESV2 struct{filterAES };func (filterIdentity )KeyLength ()int {return 0};

// HandlerVersion implements Filter interface.
func (filterAESV2 )HandlerVersion ()(V ,R int ){V ,R =4,4;return ;};

// PDFVersion implements Filter interface.
func (_bce filterV2 )PDFVersion ()[2]int {return [2]int {}};func (filterIdentity )Name ()string {return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"};

// HandlerVersion implements Filter interface.
func (_bae filterV2 )HandlerVersion ()(V ,R int ){V ,R =2,3;return ;};

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter (d FilterDict )(Filter ,error ){_fg ,_eac :=_aadf (d .CFM );if _eac !=nil {return nil ,_eac ;};_cb ,_eac :=_fg (d );if _eac !=nil {return nil ,_eac ;};return _cb ,nil ;};

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct{CFM string ;AuthEvent _ac .AuthEvent ;Length int ;};func (filterIdentity )PDFVersion ()[2]int {return [2]int {}};type filterIdentity struct{};var _ Filter =filterAESV3 {};

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2 ()Filter {_ad ,_ef :=_gc (FilterDict {});if _ef !=nil {_ed .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_ef );return filterAESV2 {};};return _ad ;};func _ebc (_cf ,_dc uint32 ,_gga []byte ,_db bool )([]byte ,error ){_aad :=make ([]byte ,len (_gga )+5);for _gbd :=0;_gbd < len (_gga );_gbd ++{_aad [_gbd ]=_gga [_gbd ];};for _cd :=0;_cd < 3;_cd ++{_ggf :=byte ((_cf >>uint32 (8*_cd ))&0xff);_aad [_cd +len (_gga )]=_ggf ;};for _ebb :=0;_ebb < 2;_ebb ++{_gd :=byte ((_dc >>uint32 (8*_ebb ))&0xff);_aad [_ebb +len (_gga )+3]=_gd ;};if _db {_aad =append (_aad ,0x73);_aad =append (_aad ,0x41);_aad =append (_aad ,0x6C);_aad =append (_aad ,0x54);};_bba :=_c .New ();_bba .Write (_aad );_gfc :=_bba .Sum (nil );if len (_gga )+5< 16{return _gfc [0:len (_gga )+5],nil ;};return _gfc ,nil ;};

// Filter is a common interface for crypt filter methods.
type Filter interface{

// Name returns a name of the filter that should be used in CFM field of Encrypt dictionary.
Name ()string ;

// KeyLength returns a length of the encryption key in bytes.
KeyLength ()int ;

// PDFVersion reports the minimal version of PDF document that introduced this filter.
PDFVersion ()[2]int ;

// HandlerVersion reports V and R parameters that should be used for this filter.
HandlerVersion ()(V ,R int );

// MakeKey generates a object encryption key based on file encryption key and object numbers.
// Used only for legacy filters - AESV3 doesn't change the key for each object.
MakeKey (_gcb ,_bfb uint32 ,_gcbb []byte )([]byte ,error );

// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and encrypt data in-place.
EncryptBytes (_de []byte ,_bacd []byte )([]byte ,error );

// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and decrypt data in-place.
DecryptBytes (_dae []byte ,_bbe []byte )([]byte ,error );};var _ Filter =filterAESV2 {};

// KeyLength implements Filter interface.
func (filterAESV2 )KeyLength ()int {return 128/8};func (filterIdentity )HandlerVersion ()(V ,R int ){return ;};func (filterAES )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_aa ,_cc :=_a .NewCipher (okey );if _cc !=nil {return nil ,_cc ;};_ed .Log .Trace ("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );const _bf =_a .BlockSize ;_f :=_bf -len (buf )%_bf ;for _ca :=0;_ca < _f ;_ca ++{buf =append (buf ,byte (_f ));};_ed .Log .Trace ("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073",len (buf ));_af :=make ([]byte ,_bf +len (buf ));_bb :=_af [:_bf ];if _ ,_fe :=_e .ReadFull (_ea .Reader ,_bb );_fe !=nil {return nil ,_fe ;};_gf :=_g .NewCBCEncrypter (_aa ,_bb );_gf .CryptBlocks (_af [_bf :],buf );buf =_af ;_ed .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,nil ;};type filterV2 struct{_da int };type filterFunc func (_ffe FilterDict )(Filter ,error );func (filterIdentity )MakeKey (objNum ,genNum uint32 ,fkey []byte )([]byte ,error ){return fkey ,nil };